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


type DescribeRankingsByUserIdIterator struct {
    rankingCache cache.RankingDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string
    categoryName string
    userId *string

    index int64
    pageToken *string
    last bool
    result []ranking.Ranking

    FetchSize *int32
}

func NewDescribeRankingsByUserIdIterator(
    rankingCache cache.RankingDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
    categoryName string,
    userId *string,
) DescribeRankingsByUserIdIterator {
    it := DescribeRankingsByUserIdIterator{
        rankingCache: rankingCache,
        client: client,
        namespaceName: namespaceName,
        categoryName: categoryName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]ranking.Ranking, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeRankingsByUserIdIterator) load() error {
    r, err := p.client.DescribeRankingssByUserId(
        &ranking.DescribeRankingssByUserIdRequest{
            NamespaceName: &p.namespaceName,
            CategoryName: &p.categoryName,
            UserId: p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.rankingCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeRankingsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeRankingsByUserIdIterator) Next(

) (ranking.Ranking, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.Ranking{}, err
        }
    }
    if len(p.result) == 0 {
        return ranking.Ranking{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.Ranking{}, err
        }
    }
    return ret, nil
}
