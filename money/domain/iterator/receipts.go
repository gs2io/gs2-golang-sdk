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
    "github.com/gs2io/gs2-golang-sdk/money"
    "github.com/gs2io/gs2-golang-sdk/money/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeReceiptsIterator struct {
    receiptCache cache.ReceiptDomainCache
    client money.Gs2MoneyRestClient
    namespaceName string
    userId string
    slot *int32
    begin *int64
    end *int64

    index int64
    pageToken *string
    last bool
    result []money.Receipt

    FetchSize *int32
}

func NewDescribeReceiptsIterator(
    receiptCache cache.ReceiptDomainCache,
    client money.Gs2MoneyRestClient,
    namespaceName string,
    userId string,
    slot *int32,
    begin *int64,
    end *int64,
) DescribeReceiptsIterator {
    it := DescribeReceiptsIterator{
        receiptCache: receiptCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,
        slot: slot,
        begin: begin,
        end: end,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]money.Receipt, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeReceiptsIterator) load() error {
    r, err := p.client.DescribeReceipts(
        &money.DescribeReceiptsRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            Slot: p.slot,
            Begin: p.begin,
            End: p.end,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.receiptCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeReceiptsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeReceiptsIterator) Next(

) (money.Receipt, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return money.Receipt{}, err
        }
    }
    if len(p.result) == 0 {
        return money.Receipt{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return money.Receipt{}, err
        }
    }
    return ret, nil
}
