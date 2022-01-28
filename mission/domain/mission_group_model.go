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

type MissionGroupModelDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    missionGroupModelCache cache.MissionGroupModelDomainCache
    namespaceName string
    missionGroupName string
    missionTaskModelCache cache.MissionTaskModelDomainCache
}

func NewMissionGroupModelDomain(
    session *core.Gs2RestSession,
    missionGroupModelCache cache.MissionGroupModelDomainCache,
    namespaceName string,
    missionGroupName string,
) MissionGroupModelDomain {
    return MissionGroupModelDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        missionGroupModelCache: missionGroupModelCache,
        namespaceName: namespaceName,
        missionGroupName: missionGroupName,
        missionTaskModelCache: cache.NewMissionTaskModelDomainCache(),
    }
}

func (p MissionGroupModelDomain) Load(
    request mission.GetMissionGroupModelRequest,
) (*mission.GetMissionGroupModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.GetMissionGroupModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.missionGroupModelCache.Update(*r.Item)
    return r, nil
}

func (p MissionGroupModelDomain) MissionTaskModels(
) iterator.DescribeMissionTaskModelsIterator {
    return iterator.NewDescribeMissionTaskModelsIterator(
        p.missionTaskModelCache,
        p.client,
        p.namespaceName,
        p.missionGroupName,
    )
}

func (p MissionGroupModelDomain) MissionTaskModel(
    missionTaskName string,
) MissionTaskModelDomain {
    return NewMissionTaskModelDomain(
        p.session,
        p.missionTaskModelCache,
        p.namespaceName,
        p.missionGroupName,
        missionTaskName,
    )
}
