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

type CurrentRatingModelMasterDomain struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceName string
}

func NewCurrentRatingModelMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentRatingModelMasterDomain {
    return CurrentRatingModelMasterDomain {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentRatingModelMasterDomain) ExportMaster(
    request matchmaking.ExportMasterRequest,
) (*matchmaking.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRatingModelMasterDomain) Load(
    request matchmaking.GetCurrentRatingModelMasterRequest,
) (*matchmaking.GetCurrentRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRatingModelMasterDomain) Update(
    request matchmaking.UpdateCurrentRatingModelMasterRequest,
) (*matchmaking.UpdateCurrentRatingModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRatingModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRatingModelMasterDomain) UpdateFromGitHub(
    request matchmaking.UpdateCurrentRatingModelMasterFromGitHubRequest,
) (*matchmaking.UpdateCurrentRatingModelMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRatingModelMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
