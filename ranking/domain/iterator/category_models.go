package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/ranking"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeCategoryModelsIterator struct {
    categoryModelCache cache.CategoryModelDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string

    index int64
    last bool
    result []ranking.CategoryModel

    FetchSize *int32
}

func NewDescribeCategoryModelsIterator(
    categoryModelCache cache.CategoryModelDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
) DescribeCategoryModelsIterator {
    it := DescribeCategoryModelsIterator{
        categoryModelCache: categoryModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]ranking.CategoryModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCategoryModelsIterator) load() error {
    r, err := p.client.DescribeCategoryModels(
        &ranking.DescribeCategoryModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.categoryModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeCategoryModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCategoryModelsIterator) Next(

) (ranking.CategoryModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.CategoryModel{}, err
        }
    }
    if len(p.result) == 0 {
        return ranking.CategoryModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.CategoryModel{}, err
        }
    }
    return ret, nil
}
