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
    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeRateModelsIterator struct {
    rateModelCache cache.RateModelDomainCache
    client enhance.Gs2EnhanceRestClient
    namespaceName string

    index int64
    last bool
    result []enhance.RateModel

    FetchSize *int32
}

func NewDescribeRateModelsIterator(
    rateModelCache cache.RateModelDomainCache,
    client enhance.Gs2EnhanceRestClient,
    namespaceName string,
) DescribeRateModelsIterator {
    it := DescribeRateModelsIterator{
        rateModelCache: rateModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]enhance.RateModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeRateModelsIterator) load() error {
    r, err := p.client.DescribeRateModels(
        &enhance.DescribeRateModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.rateModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeRateModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeRateModelsIterator) Next(

) (enhance.RateModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return enhance.RateModel{}, err
        }
    }
    if len(p.result) == 0 {
        return enhance.RateModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return enhance.RateModel{}, err
        }
    }
    return ret, nil
}
