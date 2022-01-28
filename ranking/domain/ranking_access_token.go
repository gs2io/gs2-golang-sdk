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

type RankingAccessTokenDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    rankingCache cache.RankingDomainCache
    namespaceName string
    accessToken auth.AccessToken
    categoryName *string
}

func NewRankingAccessTokenDomain(
    session *core.Gs2RestSession,
    rankingCache cache.RankingDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    categoryName *string,
) RankingAccessTokenDomain {
    return RankingAccessTokenDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        rankingCache: rankingCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        categoryName: categoryName,
    }
}

func (p RankingAccessTokenDomain) Load(
    request ranking.GetRankingRequest,
) (*ranking.GetRankingResult, error) {
    request.CategoryName = p.categoryName
    r, err := p.client.GetRanking(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.rankingCache.Update(*r.Item)
    return r, nil
}
