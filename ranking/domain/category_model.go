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

type CategoryModelDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    categoryModelCache cache.CategoryModelDomainCache
    namespaceName string
    categoryName string
}

func NewCategoryModelDomain(
    session *core.Gs2RestSession,
    categoryModelCache cache.CategoryModelDomainCache,
    namespaceName string,
    categoryName string,
) CategoryModelDomain {
    return CategoryModelDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        categoryModelCache: categoryModelCache,
        namespaceName: namespaceName,
        categoryName: categoryName,
    }
}

func (p CategoryModelDomain) Load(
    request ranking.GetCategoryModelRequest,
) (*ranking.GetCategoryModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.CategoryName = &p.categoryName
    r, err := p.client.GetCategoryModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.categoryModelCache.Update(*r.Item)
    return r, nil
}
