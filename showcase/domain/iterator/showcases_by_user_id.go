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
    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeShowcasesByUserIdIterator struct {
    showcaseCache cache.ShowcaseDomainCache
    client showcase.Gs2ShowcaseRestClient
    namespaceName string
    userId string

    index int64
    last bool
    result []showcase.Showcase

    FetchSize *int32
}

func NewDescribeShowcasesByUserIdIterator(
    showcaseCache cache.ShowcaseDomainCache,
    client showcase.Gs2ShowcaseRestClient,
    namespaceName string,
    userId string,
) DescribeShowcasesByUserIdIterator {
    it := DescribeShowcasesByUserIdIterator{
        showcaseCache: showcaseCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        last: false,
        result: make([]showcase.Showcase, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeShowcasesByUserIdIterator) load() error {
    r, err := p.client.DescribeShowcasesByUserId(
        &showcase.DescribeShowcasesByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.showcaseCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeShowcasesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeShowcasesByUserIdIterator) Next(

) (showcase.Showcase, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.Showcase{}, err
        }
    }
    if len(p.result) == 0 {
        return showcase.Showcase{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.Showcase{}, err
        }
    }
    return ret, nil
}
