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

type MissionTaskModelMasterDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    missionTaskModelMasterCache cache.MissionTaskModelMasterDomainCache
    namespaceName string
    missionGroupName string
    missionTaskName string
}

func NewMissionTaskModelMasterDomain(
    session *core.Gs2RestSession,
    missionTaskModelMasterCache cache.MissionTaskModelMasterDomainCache,
    namespaceName string,
    missionGroupName string,
    missionTaskName string,
) MissionTaskModelMasterDomain {
    return MissionTaskModelMasterDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        missionTaskModelMasterCache: missionTaskModelMasterCache,
        namespaceName: namespaceName,
        missionGroupName: missionGroupName,
        missionTaskName: missionTaskName,
    }
}

func (p MissionTaskModelMasterDomain) Load(
    request mission.GetMissionTaskModelMasterRequest,
) (*mission.GetMissionTaskModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    request.MissionTaskName = &p.missionTaskName
    r, err := p.client.GetMissionTaskModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionTaskModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MissionTaskModelMasterDomain) Update(
    request mission.UpdateMissionTaskModelMasterRequest,
) (*mission.UpdateMissionTaskModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    request.MissionTaskName = &p.missionTaskName
    r, err := p.client.UpdateMissionTaskModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionTaskModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MissionTaskModelMasterDomain) Delete(
    request mission.DeleteMissionTaskModelMasterRequest,
) (*mission.DeleteMissionTaskModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    request.MissionTaskName = &p.missionTaskName
    r, err := p.client.DeleteMissionTaskModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionTaskModelMasterCache.Delete(*r.Item)
    return r, nil
}
