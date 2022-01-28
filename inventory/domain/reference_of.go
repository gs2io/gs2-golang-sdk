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

type ReferenceOfDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceName string
    userId string
    inventoryName string
    itemName string
    itemSetName string
}

func NewReferenceOfDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
    inventoryName string,
    itemName string,
    itemSetName string,
) ReferenceOfDomain {
    return ReferenceOfDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        inventoryName: inventoryName,
        itemName: itemName,
        itemSetName: itemSetName,
    }
}

func (p ReferenceOfDomain) Load(
    request inventory.GetReferenceOfByUserIdRequest,
) (*inventory.GetReferenceOfByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetReferenceOfByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfDomain) Verify(
    request inventory.VerifyReferenceOfByUserIdRequest,
) (*inventory.VerifyReferenceOfByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.VerifyReferenceOfByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfDomain) Add(
    request inventory.AddReferenceOfByUserIdRequest,
) (*inventory.AddReferenceOfByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.AddReferenceOfByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfDomain) Delete(
    request inventory.DeleteReferenceOfByUserIdRequest,
) (*inventory.DeleteReferenceOfByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.DeleteReferenceOfByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ReferenceOfDomain) List(
) iterator.DescribeReferenceOfByUserIdIterator {
    return iterator.NewDescribeReferenceOfByUserIdIterator(
        p.client,
        p.namespaceName,
        p.inventoryName,
        p.userId,
        p.itemName,
        p.itemSetName,
    )
}
