package domainmission
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

    "github.com/gs2io/gs2-golang-sdk/mission"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentMissionMasterDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    namespaceName string
}

func NewCurrentMissionMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentMissionMasterDomain {
    return CurrentMissionMasterDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentMissionMasterDomain) ExportMaster(
    request mission.ExportMasterRequest,
) (*mission.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMissionMasterDomain) Load(
    request mission.GetCurrentMissionMasterRequest,
) (*mission.GetCurrentMissionMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentMissionMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMissionMasterDomain) Update(
    request mission.UpdateCurrentMissionMasterRequest,
) (*mission.UpdateCurrentMissionMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentMissionMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMissionMasterDomain) UpdateFromGitHub(
    request mission.UpdateCurrentMissionMasterFromGitHubRequest,
) (*mission.UpdateCurrentMissionMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentMissionMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
