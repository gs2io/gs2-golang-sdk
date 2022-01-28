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


type DescribeStatusesByUserIdIterator struct {
    statusCache cache.StatusDomainCache
    client experience.Gs2ExperienceRestClient
    namespaceName string
    experienceName *string
    userId string

    index int64
    pageToken *string
    last bool
    result []experience.Status

    FetchSize *int32
}

func NewDescribeStatusesByUserIdIterator(
    statusCache cache.StatusDomainCache,
    client experience.Gs2ExperienceRestClient,
    namespaceName string,
    experienceName *string,
    userId string,
) DescribeStatusesByUserIdIterator {
    it := DescribeStatusesByUserIdIterator{
        statusCache: statusCache,
        client: client,
        namespaceName: namespaceName,
        experienceName: experienceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]experience.Status, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeStatusesByUserIdIterator) load() error {
    r, err := p.client.DescribeStatusesByUserId(
        &experience.DescribeStatusesByUserIdRequest{
            NamespaceName: &p.namespaceName,
            ExperienceName: p.experienceName,
            UserId: &p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.statusCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeStatusesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeStatusesByUserIdIterator) Next(

) (experience.Status, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.Status{}, err
        }
    }
    if len(p.result) == 0 {
        return experience.Status{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.Status{}, err
        }
    }
    return ret, nil
}
