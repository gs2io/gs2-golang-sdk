package domainranking
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

    "github.com/gs2io/gs2-golang-sdk/ranking"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    categoryModelCache cache.CategoryModelDomainCache
    categoryModelMasterCache cache.CategoryModelMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        categoryModelCache: cache.NewCategoryModelDomainCache(),
        categoryModelMasterCache: cache.NewCategoryModelMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request ranking.GetNamespaceStatusRequest,
) (*ranking.GetNamespaceStatusResult, error) {
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
    request ranking.GetNamespaceRequest,
) (*ranking.GetNamespaceResult, error) {
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
    request ranking.UpdateNamespaceRequest,
) (*ranking.UpdateNamespaceResult, error) {
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
    request ranking.DeleteNamespaceRequest,
) (*ranking.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateCategoryModelMaster(
    request ranking.CreateCategoryModelMasterRequest,
) (*ranking.CreateCategoryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateCategoryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CategoryModels(
) iterator.DescribeCategoryModelsIterator {
    return iterator.NewDescribeCategoryModelsIterator(
        p.categoryModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CategoryModelMasters(
) iterator.DescribeCategoryModelMastersIterator {
    return iterator.NewDescribeCategoryModelMastersIterator(
        p.categoryModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentRankingMaster(
) CurrentRankingMasterDomain {
    return NewCurrentRankingMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CategoryModel(
    categoryName string,
) CategoryModelDomain {
    return NewCategoryModelDomain(
        p.session,
        p.categoryModelCache,
        p.namespaceName,
        categoryName,
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

func (p NamespaceDomain) CategoryModelMaster(
    categoryName string,
) CategoryModelMasterDomain {
    return NewCategoryModelMasterDomain(
        p.session,
        p.categoryModelMasterCache,
        p.namespaceName,
        categoryName,
    )
}
