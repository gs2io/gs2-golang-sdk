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
    "github.com/gs2io/gs2-golang-sdk/exchange"
    "github.com/gs2io/gs2-golang-sdk/exchange/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeAwaitsByUserIdIterator struct {
    awaitCache cache.AwaitDomainCache
    client exchange.Gs2ExchangeRestClient
    namespaceName string
    userId string
    rateName *string

    index int64
    pageToken *string
    last bool
    result []exchange.Await

    FetchSize *int32
}

func NewDescribeAwaitsByUserIdIterator(
    awaitCache cache.AwaitDomainCache,
    client exchange.Gs2ExchangeRestClient,
    namespaceName string,
    userId string,
    rateName *string,
) DescribeAwaitsByUserIdIterator {
    it := DescribeAwaitsByUserIdIterator{
        awaitCache: awaitCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,
        rateName: rateName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]exchange.Await, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeAwaitsByUserIdIterator) load() error {
    r, err := p.client.DescribeAwaitsByUserId(
        &exchange.DescribeAwaitsByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            RateName: p.rateName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.awaitCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeAwaitsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeAwaitsByUserIdIterator) Next(

) (exchange.Await, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return exchange.Await{}, err
        }
    }
    if len(p.result) == 0 {
        return exchange.Await{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return exchange.Await{}, err
        }
    }
    return ret, nil
}
