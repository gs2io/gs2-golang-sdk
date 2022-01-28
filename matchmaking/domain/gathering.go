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

type GatheringDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    gatheringCache cache.GatheringDomainCache
    namespaceName string
    userId string
    gatheringName string
}

func NewGatheringDomain(
    session *core.Gs2RestSession,
    gatheringCache cache.GatheringDomainCache,
    namespaceName string,
    userId string,
    gatheringName string,
) GatheringDomain {
    return GatheringDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        gatheringCache: gatheringCache,
        namespaceName: namespaceName,
        userId: userId,
        gatheringName: gatheringName,
    }
}

func (p GatheringDomain) Update(
    request matchmaking.UpdateGatheringByUserIdRequest,
) (*matchmaking.UpdateGatheringByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.GatheringName = &p.gatheringName
    r, err := p.client.UpdateGatheringByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gatheringCache.Update(*r.Item)
    return r, nil
}

func (p GatheringDomain) Load(
    request matchmaking.GetGatheringRequest,
) (*matchmaking.GetGatheringResult, error) {
    request.NamespaceName = &p.namespaceName
    request.GatheringName = &p.gatheringName
    r, err := p.client.GetGathering(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gatheringCache.Update(*r.Item)
    return r, nil
}

func (p GatheringDomain) CancelMatchmaking(
    request matchmaking.CancelMatchmakingByUserIdRequest,
) (*matchmaking.CancelMatchmakingByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.GatheringName = &p.gatheringName
    r, err := p.client.CancelMatchmakingByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gatheringCache.Update(*r.Item)
    return r, nil
}
