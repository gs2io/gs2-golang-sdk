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
    "github.com/gs2io/gs2-golang-sdk/ranking"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeSubscribesByCategoryNameIterator struct {
    subscribeUserCache cache.SubscribeUserDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string
    categoryName string
    accessToken auth.AccessToken

    index int64
    last bool
    result []ranking.SubscribeUser

    FetchSize *int32
}

func NewDescribeSubscribesByCategoryNameIterator(
    subscribeUserCache cache.SubscribeUserDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
    categoryName string,
    accessToken auth.AccessToken,
) DescribeSubscribesByCategoryNameIterator {
    it := DescribeSubscribesByCategoryNameIterator{
        subscribeUserCache: subscribeUserCache,
        client: client,
        namespaceName: namespaceName,
        categoryName: categoryName,
        accessToken: accessToken,

        index: 0,
        last: false,
        result: make([]ranking.SubscribeUser, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeSubscribesByCategoryNameIterator) load() error {
    r, err := p.client.DescribeSubscribesByCategoryName(
        &ranking.DescribeSubscribesByCategoryNameRequest{
            NamespaceName: &p.namespaceName,
            CategoryName: &p.categoryName,
            AccessToken: p.accessToken.Token,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.subscribeUserCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeSubscribesByCategoryNameIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeSubscribesByCategoryNameIterator) Next(

) (ranking.SubscribeUser, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.SubscribeUser{}, err
        }
    }
    if len(p.result) == 0 {
        return ranking.SubscribeUser{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return ranking.SubscribeUser{}, err
        }
    }
    return ret, nil
}
