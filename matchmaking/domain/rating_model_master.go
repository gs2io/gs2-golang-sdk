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

type RatingModelMasterDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    ratingModelMasterCache cache.RatingModelMasterDomainCache
    namespaceName string
    ratingName string
}

func NewRatingModelMasterDomain(
    session *core.Gs2RestSession,
    ratingModelMasterCache cache.RatingModelMasterDomainCache,
    namespaceName string,
    ratingName string,
) RatingModelMasterDomain {
    return RatingModelMasterDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        ratingModelMasterCache: ratingModelMasterCache,
        namespaceName: namespaceName,
        ratingName: ratingName,
    }
}

func (p RatingModelMasterDomain) Load(
    request matchmaking.GetRatingModelMasterRequest,
) (*matchmaking.GetRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RatingName = &p.ratingName
    r, err := p.client.GetRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.ratingModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p RatingModelMasterDomain) Update(
    request matchmaking.UpdateRatingModelMasterRequest,
) (*matchmaking.UpdateRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RatingName = &p.ratingName
    r, err := p.client.UpdateRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.ratingModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p RatingModelMasterDomain) Delete(
    request matchmaking.DeleteRatingModelMasterRequest,
) (*matchmaking.DeleteRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RatingName = &p.ratingName
    r, err := p.client.DeleteRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.ratingModelMasterCache.Delete(*r.Item)
    return r, nil
}
