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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    namespaceName string
    accessToken auth.AccessToken
    scoreCache cache.ScoreDomainCache
    rankingCache cache.RankingDomainCache
    subscribeUserCache cache.SubscribeUserDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        scoreCache: cache.NewScoreDomainCache(),
        rankingCache: cache.NewRankingDomainCache(),
        subscribeUserCache: cache.NewSubscribeUserDomainCache(),
    }
}

func (p UserAccessTokenDomain) PutScore(
    request ranking.PutScoreRequest,
) (*ranking.PutScoreResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PutScore(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) SubscribeUsers(
    categoryName string,
) iterator.DescribeSubscribesByCategoryNameIterator {
    return iterator.NewDescribeSubscribesByCategoryNameIterator(
        p.subscribeUserCache,
        p.client,
        p.namespaceName,
        categoryName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Scores(
    categoryName string,
    scorerUserId string,
) iterator.DescribeScoresIterator {
    return iterator.NewDescribeScoresIterator(
        p.scoreCache,
        p.client,
        p.namespaceName,
        categoryName,
        p.accessToken,
        scorerUserId,
    )
}

func (p UserAccessTokenDomain) Rankings(
    categoryName string,
) iterator.DescribeRankingsIterator {
    return iterator.NewDescribeRankingsIterator(
        p.rankingCache,
        p.client,
        p.namespaceName,
        categoryName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Ranking(
    categoryName *string,
) RankingAccessTokenDomain {
    return NewRankingAccessTokenDomain(
        p.session,
        p.rankingCache,
        p.namespaceName,
        p.accessToken,
        categoryName,
    )
}

func (p UserAccessTokenDomain) Score(
    categoryName string,
    scorerUserId string,
    uniqueId string,
) ScoreAccessTokenDomain {
    return NewScoreAccessTokenDomain(
        p.session,
        p.scoreCache,
        p.namespaceName,
        p.accessToken,
        categoryName,
        scorerUserId,
        uniqueId,
    )
}

func (p UserAccessTokenDomain) Subscribe(
    categoryName string,
) SubscribeAccessTokenDomain {
    return NewSubscribeAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
        categoryName,
    )
}
