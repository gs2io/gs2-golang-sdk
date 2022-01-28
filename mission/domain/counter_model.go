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

type CounterModelDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    counterModelCache cache.CounterModelDomainCache
    namespaceName string
    counterName string
}

func NewCounterModelDomain(
    session *core.Gs2RestSession,
    counterModelCache cache.CounterModelDomainCache,
    namespaceName string,
    counterName string,
) CounterModelDomain {
    return CounterModelDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        counterModelCache: counterModelCache,
        namespaceName: namespaceName,
        counterName: counterName,
    }
}

func (p CounterModelDomain) Load(
    request mission.GetCounterModelRequest,
) (*mission.GetCounterModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.CounterName = &p.counterName
    r, err := p.client.GetCounterModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterModelCache.Update(*r.Item)
    return r, nil
}
