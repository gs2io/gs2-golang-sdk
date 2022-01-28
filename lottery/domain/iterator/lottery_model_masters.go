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


type DescribeLotteryModelMastersIterator struct {
    lotteryModelMasterCache cache.LotteryModelMasterDomainCache
    client lottery.Gs2LotteryRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []lottery.LotteryModelMaster

    FetchSize *int32
}

func NewDescribeLotteryModelMastersIterator(
    lotteryModelMasterCache cache.LotteryModelMasterDomainCache,
    client lottery.Gs2LotteryRestClient,
    namespaceName string,
) DescribeLotteryModelMastersIterator {
    it := DescribeLotteryModelMastersIterator{
        lotteryModelMasterCache: lotteryModelMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]lottery.LotteryModelMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeLotteryModelMastersIterator) load() error {
    r, err := p.client.DescribeLotteryModelMasters(
        &lottery.DescribeLotteryModelMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.lotteryModelMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeLotteryModelMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeLotteryModelMastersIterator) Next(

) (lottery.LotteryModelMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.LotteryModelMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return lottery.LotteryModelMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lottery.LotteryModelMaster{}, err
        }
    }
    return ret, nil
}
