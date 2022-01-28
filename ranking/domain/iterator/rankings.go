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


type DescribeRankingsIterator struct {
    rankingCache cache.RankingDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string
    categoryName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []ranking.Ranking

    FetchSize *int32
}

func NewDescribeRankingsIterator(
    rankingCache cache.RankingDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
    categoryName string,
    accessToken auth.AccessToken,
) DescribeRankingsIterator {
    it := DescribeRankingsIterator{
        rankingCache: rankingCache,
        client: client,
        namespaceName: namespaceName,
        categoryName: categoryName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]ranking.Ranking, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeRankingsIterator) load() error {
    r, err := p.client.DescribeRankings(
        &ranking.DescribeRankingsRequest{
            NamespaceName: &p.namespaceName,
            CategoryName: &p.categoryName,
            AccessToken: p.accessToken.Token,
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

func (p *DescribeRankingsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeRankingsIterator) Next(

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
