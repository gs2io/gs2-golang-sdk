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

type CurrentRankingMasterDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    namespaceName string
}

func NewCurrentRankingMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentRankingMasterDomain {
    return CurrentRankingMasterDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentRankingMasterDomain) ExportMaster(
    request ranking.ExportMasterRequest,
) (*ranking.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRankingMasterDomain) Load(
    request ranking.GetCurrentRankingMasterRequest,
) (*ranking.GetCurrentRankingMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentRankingMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRankingMasterDomain) Update(
    request ranking.UpdateCurrentRankingMasterRequest,
) (*ranking.UpdateCurrentRankingMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRankingMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRankingMasterDomain) UpdateFromGitHub(
    request ranking.UpdateCurrentRankingMasterFromGitHubRequest,
) (*ranking.UpdateCurrentRankingMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRankingMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
