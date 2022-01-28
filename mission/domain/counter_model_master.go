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

type CounterModelMasterDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    counterModelMasterCache cache.CounterModelMasterDomainCache
    namespaceName string
    counterName string
}

func NewCounterModelMasterDomain(
    session *core.Gs2RestSession,
    counterModelMasterCache cache.CounterModelMasterDomainCache,
    namespaceName string,
    counterName string,
) CounterModelMasterDomain {
    return CounterModelMasterDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        counterModelMasterCache: counterModelMasterCache,
        namespaceName: namespaceName,
        counterName: counterName,
    }
}

func (p CounterModelMasterDomain) Load(
    request mission.GetCounterModelMasterRequest,
) (*mission.GetCounterModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.CounterName = &p.counterName
    r, err := p.client.GetCounterModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p CounterModelMasterDomain) Update(
    request mission.UpdateCounterModelMasterRequest,
) (*mission.UpdateCounterModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.CounterName = &p.counterName
    r, err := p.client.UpdateCounterModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p CounterModelMasterDomain) Delete(
    request mission.DeleteCounterModelMasterRequest,
) (*mission.DeleteCounterModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.CounterName = &p.counterName
    r, err := p.client.DeleteCounterModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterModelMasterCache.Delete(*r.Item)
    return r, nil
}
