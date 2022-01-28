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
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Inventory struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Inventory(
    session *core.Gs2RestSession,
) Gs2Inventory {
    return Gs2Inventory {
        session: session,
        client: inventory.Gs2InventoryRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Inventory) CreateNamespace(
    request inventory.CreateNamespaceRequest,
) (*inventory.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Inventory) AddCapacityByStampSheet(
    request inventory.AddCapacityByStampSheetRequest,
) (*inventory.AddCapacityByStampSheetResult, error) {
    r, err := p.client.AddCapacityByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) SetCapacityByStampSheet(
    request inventory.SetCapacityByStampSheetRequest,
) (*inventory.SetCapacityByStampSheetResult, error) {
    r, err := p.client.SetCapacityByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) AcquireItemSetByStampSheet(
    request inventory.AcquireItemSetByStampSheetRequest,
) (*inventory.AcquireItemSetByStampSheetResult, error) {
    r, err := p.client.AcquireItemSetByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) ConsumeItemSetByStampTask(
    request inventory.ConsumeItemSetByStampTaskRequest,
) (*inventory.ConsumeItemSetByStampTaskResult, error) {
    r, err := p.client.ConsumeItemSetByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) AddReferenceOfItemSetByStampSheet(
    request inventory.AddReferenceOfItemSetByStampSheetRequest,
) (*inventory.AddReferenceOfItemSetByStampSheetResult, error) {
    r, err := p.client.AddReferenceOfItemSetByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) DeleteReferenceOfItemSetByStampSheet(
    request inventory.DeleteReferenceOfItemSetByStampSheetRequest,
) (*inventory.DeleteReferenceOfItemSetByStampSheetResult, error) {
    r, err := p.client.DeleteReferenceOfItemSetByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) VerifyReferenceOfByStampTask(
    request inventory.VerifyReferenceOfByStampTaskRequest,
) (*inventory.VerifyReferenceOfByStampTaskResult, error) {
    r, err := p.client.VerifyReferenceOfByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Inventory) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Inventory) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
