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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    ratingModelMasterCache cache.RatingModelMasterDomainCache
    ratingModelCache cache.RatingModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        ratingModelMasterCache: cache.NewRatingModelMasterDomainCache(),
        ratingModelCache: cache.NewRatingModelDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request matchmaking.GetNamespaceStatusRequest,
) (*matchmaking.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request matchmaking.GetNamespaceRequest,
) (*matchmaking.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request matchmaking.UpdateNamespaceRequest,
) (*matchmaking.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request matchmaking.DeleteNamespaceRequest,
) (*matchmaking.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateRatingModelMaster(
    request matchmaking.CreateRatingModelMasterRequest,
) (*matchmaking.CreateRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) RatingModelMasters(
) iterator.DescribeRatingModelMastersIterator {
    return iterator.NewDescribeRatingModelMastersIterator(
        p.ratingModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RatingModels(
) iterator.DescribeRatingModelsIterator {
    return iterator.NewDescribeRatingModelsIterator(
        p.ratingModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) CurrentRatingModelMaster(
) CurrentRatingModelMasterDomain {
    return NewCurrentRatingModelMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RatingModel(
    ratingName string,
) RatingModelDomain {
    return NewRatingModelDomain(
        p.session,
        p.ratingModelCache,
        p.namespaceName,
        ratingName,
    )
}

func (p NamespaceDomain) RatingModelMaster(
    ratingName string,
) RatingModelMasterDomain {
    return NewRatingModelMasterDomain(
        p.session,
        p.ratingModelMasterCache,
        p.namespaceName,
        ratingName,
    )
}
