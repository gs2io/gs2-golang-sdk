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


type DescribeSubscribesByCategoryNameAndUserIdIterator struct {
    subscribeUserCache cache.SubscribeUserDomainCache
    client ranking.Gs2RankingRestClient
    namespaceName string
    categoryName string
    userId string

    index int64
    last bool
    result []ranking.SubscribeUser

    FetchSize *int32
}

func NewDescribeSubscribesByCategoryNameAndUserIdIterator(
    subscribeUserCache cache.SubscribeUserDomainCache,
    client ranking.Gs2RankingRestClient,
    namespaceName string,
    categoryName string,
    userId string,
) DescribeSubscribesByCategoryNameAndUserIdIterator {
    it := DescribeSubscribesByCategoryNameAndUserIdIterator{
        subscribeUserCache: subscribeUserCache,
        client: client,
        namespaceName: namespaceName,
        categoryName: categoryName,
        userId: userId,

        index: 0,
        last: false,
        result: make([]ranking.SubscribeUser, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeSubscribesByCategoryNameAndUserIdIterator) load() error {
    r, err := p.client.DescribeSubscribesByCategoryNameAndUserId(
        &ranking.DescribeSubscribesByCategoryNameAndUserIdRequest{
            NamespaceName: &p.namespaceName,
            CategoryName: &p.categoryName,
            UserId: &p.userId,
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

func (p *DescribeSubscribesByCategoryNameAndUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeSubscribesByCategoryNameAndUserIdIterator) Next(

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
