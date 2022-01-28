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


type DescribeLimitModelsIterator struct {
    limitModelCache cache.LimitModelDomainCache
    client limit.Gs2LimitRestClient
    namespaceName string

    index int64
    last bool
    result []limit.LimitModel

    FetchSize *int32
}

func NewDescribeLimitModelsIterator(
    limitModelCache cache.LimitModelDomainCache,
    client limit.Gs2LimitRestClient,
    namespaceName string,
) DescribeLimitModelsIterator {
    it := DescribeLimitModelsIterator{
        limitModelCache: limitModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]limit.LimitModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeLimitModelsIterator) load() error {
    r, err := p.client.DescribeLimitModels(
        &limit.DescribeLimitModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.limitModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeLimitModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeLimitModelsIterator) Next(

) (limit.LimitModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return limit.LimitModel{}, err
        }
    }
    if len(p.result) == 0 {
        return limit.LimitModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return limit.LimitModel{}, err
        }
    }
    return ret, nil
}
