package domaindistributor
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

    "github.com/gs2io/gs2-golang-sdk/distributor"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client distributor.Gs2DistributorRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    distributorModelMasterCache cache.DistributorModelMasterDomainCache
    distributorModelCache cache.DistributorModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: distributor.Gs2DistributorRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        distributorModelMasterCache: cache.NewDistributorModelMasterDomainCache(),
        distributorModelCache: cache.NewDistributorModelDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request distributor.GetNamespaceStatusRequest,
) (*distributor.GetNamespaceStatusResult, error) {
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
    request distributor.GetNamespaceRequest,
) (*distributor.GetNamespaceResult, error) {
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
    request distributor.UpdateNamespaceRequest,
) (*distributor.UpdateNamespaceResult, error) {
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
    request distributor.DeleteNamespaceRequest,
) (*distributor.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateDistributorModelMaster(
    request distributor.CreateDistributorModelMasterRequest,
) (*distributor.CreateDistributorModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateDistributorModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Distribute(
    request distributor.DistributeRequest,
) (*distributor.DistributeResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.Distribute(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) RunStampTask(
    request distributor.RunStampTaskRequest,
) (*distributor.RunStampTaskResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.RunStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) RunStampSheet(
    request distributor.RunStampSheetRequest,
) (*distributor.RunStampSheetResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.RunStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) RunStampSheetExpress(
    request distributor.RunStampSheetExpressRequest,
) (*distributor.RunStampSheetExpressResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.RunStampSheetExpress(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) DistributorModelMasters(
) iterator.DescribeDistributorModelMastersIterator {
    return iterator.NewDescribeDistributorModelMastersIterator(
        p.distributorModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) DistributorModels(
) iterator.DescribeDistributorModelsIterator {
    return iterator.NewDescribeDistributorModelsIterator(
        p.distributorModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentDistributorMaster(
) CurrentDistributorMasterDomain {
    return NewCurrentDistributorMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) DistributorModel(
    distributorName string,
) DistributorModelDomain {
    return NewDistributorModelDomain(
        p.session,
        p.distributorModelCache,
        p.namespaceName,
        distributorName,
    )
}

func (p NamespaceDomain) DistributorModelMaster(
    distributorName string,
) DistributorModelMasterDomain {
    return NewDistributorModelMasterDomain(
        p.session,
        p.distributorModelMasterCache,
        p.namespaceName,
        distributorName,
    )
}
