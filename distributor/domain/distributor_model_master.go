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

type DistributorModelMasterDomain struct {
    session *core.Gs2RestSession
    client distributor.Gs2DistributorRestClient
    distributorModelMasterCache cache.DistributorModelMasterDomainCache
    namespaceName string
    distributorName string
}

func NewDistributorModelMasterDomain(
    session *core.Gs2RestSession,
    distributorModelMasterCache cache.DistributorModelMasterDomainCache,
    namespaceName string,
    distributorName string,
) DistributorModelMasterDomain {
    return DistributorModelMasterDomain {
        session: session,
        client: distributor.Gs2DistributorRestClient{
            Session: session,
        },
        distributorModelMasterCache: distributorModelMasterCache,
        namespaceName: namespaceName,
        distributorName: distributorName,
    }
}

func (p DistributorModelMasterDomain) Load(
    request distributor.GetDistributorModelMasterRequest,
) (*distributor.GetDistributorModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.DistributorName = &p.distributorName
    r, err := p.client.GetDistributorModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.distributorModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p DistributorModelMasterDomain) Update(
    request distributor.UpdateDistributorModelMasterRequest,
) (*distributor.UpdateDistributorModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.DistributorName = &p.distributorName
    r, err := p.client.UpdateDistributorModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.distributorModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p DistributorModelMasterDomain) Delete(
    request distributor.DeleteDistributorModelMasterRequest,
) (*distributor.DeleteDistributorModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.DistributorName = &p.distributorName
    r, err := p.client.DeleteDistributorModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.distributorModelMasterCache.Delete(*r.Item)
    return r, nil
}
