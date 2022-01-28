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


type DescribeMissionTaskModelMastersIterator struct {
    missionTaskModelMasterCache cache.MissionTaskModelMasterDomainCache
    client mission.Gs2MissionRestClient
    namespaceName string
    missionGroupName string

    index int64
    pageToken *string
    last bool
    result []mission.MissionTaskModelMaster

    FetchSize *int32
}

func NewDescribeMissionTaskModelMastersIterator(
    missionTaskModelMasterCache cache.MissionTaskModelMasterDomainCache,
    client mission.Gs2MissionRestClient,
    namespaceName string,
    missionGroupName string,
) DescribeMissionTaskModelMastersIterator {
    it := DescribeMissionTaskModelMastersIterator{
        missionTaskModelMasterCache: missionTaskModelMasterCache,
        client: client,
        namespaceName: namespaceName,
        missionGroupName: missionGroupName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]mission.MissionTaskModelMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeMissionTaskModelMastersIterator) load() error {
    r, err := p.client.DescribeMissionTaskModelMasters(
        &mission.DescribeMissionTaskModelMastersRequest{
            NamespaceName: &p.namespaceName,
            MissionGroupName: &p.missionGroupName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.missionTaskModelMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeMissionTaskModelMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeMissionTaskModelMastersIterator) Next(

) (mission.MissionTaskModelMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.MissionTaskModelMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return mission.MissionTaskModelMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.MissionTaskModelMaster{}, err
        }
    }
    return ret, nil
}
