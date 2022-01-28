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


type DescribeCountersIterator struct {
    counterCache cache.CounterDomainCache
    client mission.Gs2MissionRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []mission.Counter

    FetchSize *int32
}

func NewDescribeCountersIterator(
    counterCache cache.CounterDomainCache,
    client mission.Gs2MissionRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeCountersIterator {
    it := DescribeCountersIterator{
        counterCache: counterCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]mission.Counter, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCountersIterator) load() error {
    r, err := p.client.DescribeCounters(
        &mission.DescribeCountersRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.counterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeCountersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCountersIterator) Next(

) (mission.Counter, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.Counter{}, err
        }
    }
    if len(p.result) == 0 {
        return mission.Counter{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return mission.Counter{}, err
        }
    }
    return ret, nil
}
