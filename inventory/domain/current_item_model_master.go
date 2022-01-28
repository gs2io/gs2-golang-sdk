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

type CurrentItemModelMasterDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceName string
}

func NewCurrentItemModelMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentItemModelMasterDomain {
    return CurrentItemModelMasterDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentItemModelMasterDomain) ExportMaster(
    request inventory.ExportMasterRequest,
) (*inventory.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentItemModelMasterDomain) Load(
    request inventory.GetCurrentItemModelMasterRequest,
) (*inventory.GetCurrentItemModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentItemModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentItemModelMasterDomain) Update(
    request inventory.UpdateCurrentItemModelMasterRequest,
) (*inventory.UpdateCurrentItemModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentItemModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentItemModelMasterDomain) UpdateFromGitHub(
    request inventory.UpdateCurrentItemModelMasterFromGitHubRequest,
) (*inventory.UpdateCurrentItemModelMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentItemModelMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
