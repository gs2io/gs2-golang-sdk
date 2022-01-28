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


type DescribeShowcaseMastersIterator struct {
    showcaseMasterCache cache.ShowcaseMasterDomainCache
    client showcase.Gs2ShowcaseRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []showcase.ShowcaseMaster

    FetchSize *int32
}

func NewDescribeShowcaseMastersIterator(
    showcaseMasterCache cache.ShowcaseMasterDomainCache,
    client showcase.Gs2ShowcaseRestClient,
    namespaceName string,
) DescribeShowcaseMastersIterator {
    it := DescribeShowcaseMastersIterator{
        showcaseMasterCache: showcaseMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]showcase.ShowcaseMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeShowcaseMastersIterator) load() error {
    r, err := p.client.DescribeShowcaseMasters(
        &showcase.DescribeShowcaseMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.showcaseMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeShowcaseMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeShowcaseMastersIterator) Next(

) (showcase.ShowcaseMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.ShowcaseMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return showcase.ShowcaseMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.ShowcaseMaster{}, err
        }
    }
    return ret, nil
}
