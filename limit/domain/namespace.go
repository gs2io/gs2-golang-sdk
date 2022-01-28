package domainlimit
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

    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    limitModelMasterCache cache.LimitModelMasterDomainCache
    limitModelCache cache.LimitModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        limitModelMasterCache: cache.NewLimitModelMasterDomainCache(),
        limitModelCache: cache.NewLimitModelDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request limit.GetNamespaceStatusRequest,
) (*limit.GetNamespaceStatusResult, error) {
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
    request limit.GetNamespaceRequest,
) (*limit.GetNamespaceResult, error) {
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
    request limit.UpdateNamespaceRequest,
) (*limit.UpdateNamespaceResult, error) {
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
    request limit.DeleteNamespaceRequest,
) (*limit.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateLimitModelMaster(
    request limit.CreateLimitModelMasterRequest,
) (*limit.CreateLimitModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateLimitModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) LimitModelMasters(
) iterator.DescribeLimitModelMastersIterator {
    return iterator.NewDescribeLimitModelMastersIterator(
        p.limitModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) LimitModels(
) iterator.DescribeLimitModelsIterator {
    return iterator.NewDescribeLimitModelsIterator(
        p.limitModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentLimitMaster(
) CurrentLimitMasterDomain {
    return NewCurrentLimitMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) LimitModel(
    limitName string,
) LimitModelDomain {
    return NewLimitModelDomain(
        p.session,
        p.limitModelCache,
        p.namespaceName,
        limitName,
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

func (p NamespaceDomain) LimitModelMaster(
    limitName string,
) LimitModelMasterDomain {
    return NewLimitModelMasterDomain(
        p.session,
        p.limitModelMasterCache,
        p.namespaceName,
        limitName,
    )
}
