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


type DescribeProbabilitiesByUserIdIterator struct {
    probabilityCache cache.ProbabilityDomainCache
    client lottery.Gs2LotteryRestClient
    namespaceName string
    lotteryName string
    userId string

    index int64
    last bool
    result []lottery.Probability

    FetchSize *int32
}

func NewDescribeProbabilitiesByUserIdIterator(
    probabilityCache cache.ProbabilityDomainCache,
    client lottery.Gs2LotteryRestClient,
    namespaceName string,
    lotteryName string,
    userId string,
) DescribeProbabilitiesByUserIdIterator {
    it := DescribeProbabilitiesByUserIdIterator{
        probabilityCache: probabilityCache,
        client: client,
        namespaceName: namespaceName,
        lotteryName: lotteryName,
        userId: userId,

        index: 0,
        last: false,
        result: make([]lottery.Probability, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeProbabilitiesByUserIdIterator) load() error {
    r, err := p.client.DescribeProbabilitiesByUserId(
        &lottery.DescribeProbabilitiesByUserIdRequest{
            NamespaceName: &p.namespaceName,
            LotteryName: &p.lotteryName,
            UserId: &p.userId,
        },
    )
    if err != nil {
        return err
    }
    p.probabilityCache.Update(r.Items)
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeProbabilitiesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeProbabilitiesByUserIdIterator) Next(

) (lottery.Probability, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.Probability{}, err
        }
    }
    if len(p.result) == 0 {
        return lottery.Probability{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.Probability{}, err
        }
    }
    return ret, nil
}
