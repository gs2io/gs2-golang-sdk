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

type MissionGroupModelMasterDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    missionGroupModelMasterCache cache.MissionGroupModelMasterDomainCache
    namespaceName string
    missionGroupName string
    missionTaskModelMasterCache cache.MissionTaskModelMasterDomainCache
}

func NewMissionGroupModelMasterDomain(
    session *core.Gs2RestSession,
    missionGroupModelMasterCache cache.MissionGroupModelMasterDomainCache,
    namespaceName string,
    missionGroupName string,
) MissionGroupModelMasterDomain {
    return MissionGroupModelMasterDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        missionGroupModelMasterCache: missionGroupModelMasterCache,
        namespaceName: namespaceName,
        missionGroupName: missionGroupName,
        missionTaskModelMasterCache: cache.NewMissionTaskModelMasterDomainCache(),
    }
}

func (p MissionGroupModelMasterDomain) Load(
    request mission.GetMissionGroupModelMasterRequest,
) (*mission.GetMissionGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.GetMissionGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionGroupModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MissionGroupModelMasterDomain) Update(
    request mission.UpdateMissionGroupModelMasterRequest,
) (*mission.UpdateMissionGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.UpdateMissionGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionGroupModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MissionGroupModelMasterDomain) Delete(
    request mission.DeleteMissionGroupModelMasterRequest,
) (*mission.DeleteMissionGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.DeleteMissionGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionGroupModelMasterCache.Delete(*r.Item)
    return r, nil
}

func (p MissionGroupModelMasterDomain) CreateMissionTaskModelMaster(
    request mission.CreateMissionTaskModelMasterRequest,
) (*mission.CreateMissionTaskModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.CreateMissionTaskModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p MissionGroupModelMasterDomain) MissionTaskModelMasters(
) iterator.DescribeMissionTaskModelMastersIterator {
    return iterator.NewDescribeMissionTaskModelMastersIterator(
        p.missionTaskModelMasterCache,
        p.client,
        p.namespaceName,
        p.missionGroupName,
    )
}

func (p MissionGroupModelMasterDomain) MissionTaskModelMaster(
    missionTaskName string,
) MissionTaskModelMasterDomain {
    return NewMissionTaskModelMasterDomain(
        p.session,
        p.missionTaskModelMasterCache,
        p.namespaceName,
        p.missionGroupName,
        missionTaskName,
    )
}
