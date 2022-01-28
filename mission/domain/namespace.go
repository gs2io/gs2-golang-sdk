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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    counterModelMasterCache cache.CounterModelMasterDomainCache
    missionGroupModelMasterCache cache.MissionGroupModelMasterDomainCache
    counterModelCache cache.CounterModelDomainCache
    missionGroupModelCache cache.MissionGroupModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        counterModelMasterCache: cache.NewCounterModelMasterDomainCache(),
        missionGroupModelMasterCache: cache.NewMissionGroupModelMasterDomainCache(),
        counterModelCache: cache.NewCounterModelDomainCache(),
        missionGroupModelCache: cache.NewMissionGroupModelDomainCache(),
    }
}

func (p NamespaceDomain) CreateCounterModelMaster(
    request mission.CreateCounterModelMasterRequest,
) (*mission.CreateCounterModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateCounterModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateMissionGroupModelMaster(
    request mission.CreateMissionGroupModelMasterRequest,
) (*mission.CreateMissionGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateMissionGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) GetStatus(
    request mission.GetNamespaceStatusRequest,
) (*mission.GetNamespaceStatusResult, error) {
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
    request mission.GetNamespaceRequest,
) (*mission.GetNamespaceResult, error) {
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
    request mission.UpdateNamespaceRequest,
) (*mission.UpdateNamespaceResult, error) {
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
    request mission.DeleteNamespaceRequest,
) (*mission.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CounterModelMasters(
) iterator.DescribeCounterModelMastersIterator {
    return iterator.NewDescribeCounterModelMastersIterator(
        p.counterModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MissionGroupModelMasters(
) iterator.DescribeMissionGroupModelMastersIterator {
    return iterator.NewDescribeMissionGroupModelMastersIterator(
        p.missionGroupModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CounterModels(
) iterator.DescribeCounterModelsIterator {
    return iterator.NewDescribeCounterModelsIterator(
        p.counterModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MissionGroupModels(
) iterator.DescribeMissionGroupModelsIterator {
    return iterator.NewDescribeMissionGroupModelsIterator(
        p.missionGroupModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentMissionMaster(
) CurrentMissionMasterDomain {
    return NewCurrentMissionMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MissionGroupModel(
    missionGroupName string,
) MissionGroupModelDomain {
    return NewMissionGroupModelDomain(
        p.session,
        p.missionGroupModelCache,
        p.namespaceName,
        missionGroupName,
    )
}

func (p NamespaceDomain) CounterModel(
    counterName string,
) CounterModelDomain {
    return NewCounterModelDomain(
        p.session,
        p.counterModelCache,
        p.namespaceName,
        counterName,
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

func (p NamespaceDomain) MissionGroupModelMaster(
    missionGroupName string,
) MissionGroupModelMasterDomain {
    return NewMissionGroupModelMasterDomain(
        p.session,
        p.missionGroupModelMasterCache,
        p.namespaceName,
        missionGroupName,
    )
}

func (p NamespaceDomain) CounterModelMaster(
    counterName string,
) CounterModelMasterDomain {
    return NewCounterModelMasterDomain(
        p.session,
        p.counterModelMasterCache,
        p.namespaceName,
        counterName,
    )
}
