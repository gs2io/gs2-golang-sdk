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

type UserDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    namespaceName string
    userId string
    scoreCache cache.ScoreDomainCache
    rankingCache cache.RankingDomainCache
    subscribeUserCache cache.SubscribeUserDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        scoreCache: cache.NewScoreDomainCache(),
        rankingCache: cache.NewRankingDomainCache(),
        subscribeUserCache: cache.NewSubscribeUserDomainCache(),
    }
}

func (p UserDomain) PutScore(
    request ranking.PutScoreByUserIdRequest,
) (*ranking.PutScoreByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PutScoreByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) SubscribeUsers(
    categoryName string,
) iterator.DescribeSubscribesByCategoryNameAndUserIdIterator {
    return iterator.NewDescribeSubscribesByCategoryNameAndUserIdIterator(
        p.subscribeUserCache,
        p.client,
        p.namespaceName,
        categoryName,
        p.userId,
    )
}

func (p UserDomain) Scores(
    categoryName string,
    scorerUserId string,
) iterator.DescribeScoresByUserIdIterator {
    return iterator.NewDescribeScoresByUserIdIterator(
        p.scoreCache,
        p.client,
        p.namespaceName,
        categoryName,
        p.userId,
        scorerUserId,
    )
}

func (p UserDomain) Rankings(
    categoryName string,
) iterator.DescribeRankingsByUserIdIterator {
    return iterator.NewDescribeRankingsByUserIdIterator(
        p.rankingCache,
        p.client,
        p.namespaceName,
        categoryName,
        &p.userId,
    )
}

func (p UserDomain) NearRankings(
    categoryName string,
    score int64,
) iterator.DescribeNearRankingsIterator {
    return iterator.NewDescribeNearRankingsIterator(
        p.rankingCache,
        p.client,
        p.namespaceName,
        categoryName,
        score,
    )
}

func (p UserDomain) Ranking(
    categoryName *string,
) RankingDomain {
    return NewRankingDomain(
        p.session,
        p.rankingCache,
        p.namespaceName,
        p.userId,
        categoryName,
    )
}

func (p UserDomain) Score(
    categoryName string,
    scorerUserId string,
    uniqueId string,
) ScoreDomain {
    return NewScoreDomain(
        p.session,
        p.scoreCache,
        p.namespaceName,
        p.userId,
        categoryName,
        scorerUserId,
        uniqueId,
    )
}

func (p UserDomain) Subscribe(
    categoryName string,
) SubscribeDomain {
    return NewSubscribeDomain(
        p.session,
        p.namespaceName,
        p.userId,
        categoryName,
    )
}
