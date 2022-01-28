package domainmatchmaking
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

    "github.com/gs2io/gs2-golang-sdk/matchmaking"
    "github.com/gs2io/gs2-golang-sdk/matchmaking/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/matchmaking/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceName string
    userId string
    gatheringCache cache.GatheringDomainCache
    ratingCache cache.RatingDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        gatheringCache: cache.NewGatheringDomainCache(),
        ratingCache: cache.NewRatingDomainCache(),
    }
}

func (p UserDomain) CreateGathering(
    request matchmaking.CreateGatheringByUserIdRequest,
) (*matchmaking.CreateGatheringByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CreateGatheringByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) DoMatchmaking(
    request matchmaking.DoMatchmakingByUserIdRequest,
) (*matchmaking.DoMatchmakingByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DoMatchmakingByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Gatherings(
) iterator.DescribeGatheringsIterator {
    return iterator.NewDescribeGatheringsIterator(
        p.gatheringCache,
        p.client,
        p.namespaceName,
    )
}

func (p UserDomain) Ratings(
) iterator.DescribeRatingsByUserIdIterator {
    return iterator.NewDescribeRatingsByUserIdIterator(
        p.ratingCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Rating(
    ratingName string,
) RatingDomain {
    return NewRatingDomain(
        p.session,
        p.ratingCache,
        p.namespaceName,
        p.userId,
        ratingName,
    )
}

func (p UserDomain) Gathering(
    gatheringName string,
) GatheringDomain {
    return NewGatheringDomain(
        p.session,
        p.gatheringCache,
        p.namespaceName,
        p.userId,
        gatheringName,
    )
}

func (p UserDomain) Vote(
    ratingName string,
    gatheringName string,
) VoteDomain {
    return NewVoteDomain(
        p.session,
        p.namespaceName,
        p.userId,
        ratingName,
        gatheringName,
    )
}
