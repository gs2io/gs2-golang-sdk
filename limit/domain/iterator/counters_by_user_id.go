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
    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeCountersByUserIdIterator struct {
    counterCache cache.CounterDomainCache
    client limit.Gs2LimitRestClient
    namespaceName string
    userId string
    limitName *string

    index int64
    pageToken *string
    last bool
    result []limit.Counter

    FetchSize *int32
}

func NewDescribeCountersByUserIdIterator(
    counterCache cache.CounterDomainCache,
    client limit.Gs2LimitRestClient,
    namespaceName string,
    userId string,
    limitName *string,
) DescribeCountersByUserIdIterator {
    it := DescribeCountersByUserIdIterator{
        counterCache: counterCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,
        limitName: limitName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]limit.Counter, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCountersByUserIdIterator) load() error {
    r, err := p.client.DescribeCountersByUserId(
        &limit.DescribeCountersByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            LimitName: p.limitName,
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

func (p *DescribeCountersByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCountersByUserIdIterator) Next(

) (limit.Counter, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return limit.Counter{}, err
        }
    }
    if len(p.result) == 0 {
        return limit.Counter{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return limit.Counter{}, err
        }
    }
    return ret, nil
}
