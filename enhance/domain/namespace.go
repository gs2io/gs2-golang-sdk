package domainenhance
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

    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    rateModelCache cache.RateModelDomainCache
    rateModelMasterCache cache.RateModelMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        rateModelCache: cache.NewRateModelDomainCache(),
        rateModelMasterCache: cache.NewRateModelMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request enhance.GetNamespaceStatusRequest,
) (*enhance.GetNamespaceStatusResult, error) {
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
    request enhance.GetNamespaceRequest,
) (*enhance.GetNamespaceResult, error) {
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
    request enhance.UpdateNamespaceRequest,
) (*enhance.UpdateNamespaceResult, error) {
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
    request enhance.DeleteNamespaceRequest,
) (*enhance.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateRateModelMaster(
    request enhance.CreateRateModelMasterRequest,
) (*enhance.CreateRateModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateRateModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) RateModels(
) iterator.DescribeRateModelsIterator {
    return iterator.NewDescribeRateModelsIterator(
        p.rateModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RateModelMasters(
) iterator.DescribeRateModelMastersIterator {
    return iterator.NewDescribeRateModelMastersIterator(
        p.rateModelMasterCache,
        p.client,
        p.namespaceName,
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

func (p NamespaceDomain) CurrentRateMaster(
) CurrentRateMasterDomain {
    return NewCurrentRateMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RateModel(
    rateName string,
) RateModelDomain {
    return NewRateModelDomain(
        p.session,
        p.rateModelCache,
        p.namespaceName,
        rateName,
    )
}

func (p NamespaceDomain) RateModelMaster(
    rateName string,
) RateModelMasterDomain {
    return NewRateModelMasterDomain(
        p.session,
        p.rateModelMasterCache,
        p.namespaceName,
        rateName,
    )
}
