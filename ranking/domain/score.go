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

type ScoreDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    scoreCache cache.ScoreDomainCache
    namespaceName string
    userId string
    categoryName string
    scorerUserId string
    uniqueId string
}

func NewScoreDomain(
    session *core.Gs2RestSession,
    scoreCache cache.ScoreDomainCache,
    namespaceName string,
    userId string,
    categoryName string,
    scorerUserId string,
    uniqueId string,
) ScoreDomain {
    return ScoreDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        scoreCache: scoreCache,
        namespaceName: namespaceName,
        userId: userId,
        categoryName: categoryName,
        scorerUserId: scorerUserId,
        uniqueId: uniqueId,
    }
}

func (p ScoreDomain) Load(
    request ranking.GetScoreByUserIdRequest,
) (*ranking.GetScoreByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.CategoryName = &p.categoryName
    request.ScorerUserId = &p.scorerUserId
    request.UniqueId = &p.uniqueId
    r, err := p.client.GetScoreByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.scoreCache.Update(*r.Item)
    return r, nil
}
