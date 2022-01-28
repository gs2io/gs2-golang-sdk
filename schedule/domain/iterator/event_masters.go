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
    "github.com/gs2io/gs2-golang-sdk/schedule"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeEventMastersIterator struct {
    eventMasterCache cache.EventMasterDomainCache
    client schedule.Gs2ScheduleRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []schedule.EventMaster

    FetchSize *int32
}

func NewDescribeEventMastersIterator(
    eventMasterCache cache.EventMasterDomainCache,
    client schedule.Gs2ScheduleRestClient,
    namespaceName string,
) DescribeEventMastersIterator {
    it := DescribeEventMastersIterator{
        eventMasterCache: eventMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]schedule.EventMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeEventMastersIterator) load() error {
    r, err := p.client.DescribeEventMasters(
        &schedule.DescribeEventMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.eventMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeEventMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeEventMastersIterator) Next(

) (schedule.EventMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return schedule.EventMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return schedule.EventMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return schedule.EventMaster{}, err
        }
    }
    return ret, nil
}
