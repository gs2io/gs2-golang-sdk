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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    inventoryModelMasterCache cache.InventoryModelMasterDomainCache
    inventoryModelCache cache.InventoryModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        inventoryModelMasterCache: cache.NewInventoryModelMasterDomainCache(),
        inventoryModelCache: cache.NewInventoryModelDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request inventory.GetNamespaceStatusRequest,
) (*inventory.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request inventory.GetNamespaceRequest,
) (*inventory.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request inventory.UpdateNamespaceRequest,
) (*inventory.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request inventory.DeleteNamespaceRequest,
) (*inventory.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateInventoryModelMaster(
    request inventory.CreateInventoryModelMasterRequest,
) (*inventory.CreateInventoryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateInventoryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) InventoryModelMasters(
) iterator.DescribeInventoryModelMastersIterator {
    return iterator.NewDescribeInventoryModelMastersIterator(
        p.inventoryModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) InventoryModels(
) iterator.DescribeInventoryModelsIterator {
    return iterator.NewDescribeInventoryModelsIterator(
        p.inventoryModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentItemModelMaster(
) CurrentItemModelMasterDomain {
    return NewCurrentItemModelMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) InventoryModel(
    inventoryName string,
) InventoryModelDomain {
    return NewInventoryModelDomain(
        p.session,
        p.inventoryModelCache,
        p.namespaceName,
        inventoryName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) InventoryModelMaster(
    inventoryName string,
) InventoryModelMasterDomain {
    return NewInventoryModelMasterDomain(
        p.session,
        p.inventoryModelMasterCache,
        p.namespaceName,
        inventoryName,
    )
}
