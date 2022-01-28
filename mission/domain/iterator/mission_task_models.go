package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/mission"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeMissionTaskModelsIterator struct {
    missionTaskModelCache cache.MissionTaskModelDomainCache
    client mission.Gs2MissionRestClient
    namespaceName string
    missionGroupName string

    index int64
    last bool
    result []mission.MissionTaskModel

    FetchSize *int32
}

func NewDescribeMissionTaskModelsIterator(
    missionTaskModelCache cache.MissionTaskModelDomainCache,
    client mission.Gs2MissionRestClient,
    namespaceName string,
    missionGroupName string,
) DescribeMissionTaskModelsIterator {
    it := DescribeMissionTaskModelsIterator{
        missionTaskModelCache: missionTaskModelCache,
        client: client,
        namespaceName: namespaceName,
        missionGroupName: missionGroupName,

        index: 0,
        last: false,
        result: make([]mission.MissionTaskModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeMissionTaskModelsIterator) load() error {
    r, err := p.client.DescribeMissionTaskModels(
        &mission.DescribeMissionTaskModelsRequest{
            NamespaceName: &p.namespaceName,
            MissionGroupName: &p.missionGroupName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.missionTaskModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeMissionTaskModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeMissionTaskModelsIterator) Next(

) (mission.MissionTaskModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.MissionTaskModel{}, err
        }
    }
    if len(p.result) == 0 {
        return mission.MissionTaskModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.MissionTaskModel{}, err
        }
    }
    return ret, nil
}
