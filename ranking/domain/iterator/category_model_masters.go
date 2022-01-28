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


type DescribeCategoryModelMastersIterator struct {
    categoryModelMasterCache cache.CategoryModelMasterDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []ranking.CategoryModelMaster

    FetchSize *int32
}

func NewDescribeCategoryModelMastersIterator(
    categoryModelMasterCache cache.CategoryModelMasterDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
) DescribeCategoryModelMastersIterator {
    it := DescribeCategoryModelMastersIterator{
        categoryModelMasterCache: categoryModelMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]ranking.CategoryModelMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCategoryModelMastersIterator) load() error {
    r, err := p.client.DescribeCategoryModelMasters(
        &ranking.DescribeCategoryModelMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.categoryModelMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeCategoryModelMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCategoryModelMastersIterator) Next(

) (ranking.CategoryModelMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.CategoryModelMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return ranking.CategoryModelMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.CategoryModelMaster{}, err
        }
    }
    return ret, nil
}
