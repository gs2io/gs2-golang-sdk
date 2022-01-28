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
    "github.com/gs2io/gs2-golang-sdk/lottery"
    "github.com/gs2io/gs2-golang-sdk/lottery/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeLotteryModelsIterator struct {
    lotteryModelCache cache.LotteryModelDomainCache
    client lottery.Gs2LotteryRestClient
    namespaceName string

    index int64
    last bool
    result []lottery.LotteryModel

    FetchSize *int32
}

func NewDescribeLotteryModelsIterator(
    lotteryModelCache cache.LotteryModelDomainCache,
    client lottery.Gs2LotteryRestClient,
    namespaceName string,
) DescribeLotteryModelsIterator {
    it := DescribeLotteryModelsIterator{
        lotteryModelCache: lotteryModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]lottery.LotteryModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeLotteryModelsIterator) load() error {
    r, err := p.client.DescribeLotteryModels(
        &lottery.DescribeLotteryModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.lotteryModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeLotteryModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeLotteryModelsIterator) Next(

) (lottery.LotteryModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.LotteryModel{}, err
        }
    }
    if len(p.result) == 0 {
        return lottery.LotteryModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.LotteryModel{}, err
        }
    }
    return ret, nil
}