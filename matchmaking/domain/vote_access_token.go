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

type VoteAccessTokenDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceName string
    accessToken auth.AccessToken
    ratingName string
    gatheringName string
}

func NewVoteAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
    ratingName string,
    gatheringName string,
) VoteAccessTokenDomain {
    return VoteAccessTokenDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        ratingName: ratingName,
        gatheringName: gatheringName,
    }
}

func (p VoteAccessTokenDomain) GetBallot(
    request matchmaking.GetBallotRequest,
) (*matchmaking.GetBallotResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RatingName = &p.ratingName
    request.GatheringName = &p.gatheringName
    r, err := p.client.GetBallot(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
