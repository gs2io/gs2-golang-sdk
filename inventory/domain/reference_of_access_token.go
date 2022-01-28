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

type ReferenceOfAccessTokenDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceName string
    accessToken auth.AccessToken
    inventoryName string
    itemName string
    itemSetName string
}

func NewReferenceOfAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
    inventoryName string,
    itemName string,
    itemSetName string,
) ReferenceOfAccessTokenDomain {
    return ReferenceOfAccessTokenDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        inventoryName: inventoryName,
        itemName: itemName,
        itemSetName: itemSetName,
    }
}

func (p ReferenceOfAccessTokenDomain) Load(
    request inventory.GetReferenceOfRequest,
) (*inventory.GetReferenceOfResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetReferenceOf(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfAccessTokenDomain) Verify(
    request inventory.VerifyReferenceOfRequest,
) (*inventory.VerifyReferenceOfResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.VerifyReferenceOf(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfAccessTokenDomain) Add(
    request inventory.AddReferenceOfRequest,
) (*inventory.AddReferenceOfResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.AddReferenceOf(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfAccessTokenDomain) Delete(
    request inventory.DeleteReferenceOfRequest,
) (*inventory.DeleteReferenceOfResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.DeleteReferenceOf(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfAccessTokenDomain) List(
) iterator.DescribeReferenceOfIterator {
    return iterator.NewDescribeReferenceOfIterator(
        p.client,
        p.namespaceName,
        p.inventoryName,
        p.accessToken,
        p.itemName,
        p.itemSetName,
    )
}
