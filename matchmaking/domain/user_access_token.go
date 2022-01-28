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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceName string
    accessToken auth.AccessToken
    gatheringCache cache.GatheringDomainCache
    ratingCache cache.RatingDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        gatheringCache: cache.NewGatheringDomainCache(),
        ratingCache: cache.NewRatingDomainCache(),
    }
}

func (p UserAccessTokenDomain) CreateGathering(
    request matchmaking.CreateGatheringRequest,
) (*matchmaking.CreateGatheringResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.CreateGathering(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) DoMatchmaking(
    request matchmaking.DoMatchmakingRequest,
) (*matchmaking.DoMatchmakingResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.DoMatchmaking(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Ratings(
) iterator.DescribeRatingsIterator {
    return iterator.NewDescribeRatingsIterator(
        p.ratingCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Rating(
    ratingName string,
) RatingAccessTokenDomain {
    return NewRatingAccessTokenDomain(
        p.session,
        p.ratingCache,
        p.namespaceName,
        p.accessToken,
        ratingName,
    )
}

func (p UserAccessTokenDomain) Gathering(
    gatheringName string,
) GatheringAccessTokenDomain {
    return NewGatheringAccessTokenDomain(
        p.session,
        p.gatheringCache,
        p.namespaceName,
        p.accessToken,
        gatheringName,
    )
}

func (p UserAccessTokenDomain) Vote(
    ratingName string,
    gatheringName string,
) VoteAccessTokenDomain {
    return NewVoteAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
        ratingName,
        gatheringName,
    )
}
