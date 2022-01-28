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
    "github.com/gs2io/gs2-golang-sdk/experience"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeThresholdMastersIterator struct {
    thresholdMasterCache cache.ThresholdMasterDomainCache
    client experience.Gs2ExperienceRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []experience.ThresholdMaster

    FetchSize *int32
}

func NewDescribeThresholdMastersIterator(
    thresholdMasterCache cache.ThresholdMasterDomainCache,
    client experience.Gs2ExperienceRestClient,
    namespaceName string,
) DescribeThresholdMastersIterator {
    it := DescribeThresholdMastersIterator{
        thresholdMasterCache: thresholdMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]experience.ThresholdMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeThresholdMastersIterator) load() error {
    r, err := p.client.DescribeThresholdMasters(
        &experience.DescribeThresholdMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.thresholdMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeThresholdMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeThresholdMastersIterator) Next(

) (experience.ThresholdMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.ThresholdMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return experience.ThresholdMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.ThresholdMaster{}, err
        }
    }
    return ret, nil
}
