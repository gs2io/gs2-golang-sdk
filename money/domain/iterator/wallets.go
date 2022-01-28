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


type DescribeWalletsIterator struct {
    walletCache cache.WalletDomainCache
    client money.Gs2MoneyRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []money.Wallet

    FetchSize *int32
}

func NewDescribeWalletsIterator(
    walletCache cache.WalletDomainCache,
    client money.Gs2MoneyRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeWalletsIterator {
    it := DescribeWalletsIterator{
        walletCache: walletCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]money.Wallet, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeWalletsIterator) load() error {
    r, err := p.client.DescribeWallets(
        &money.DescribeWalletsRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.walletCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeWalletsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeWalletsIterator) Next(

) (money.Wallet, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return money.Wallet{}, err
        }
    }
    if len(p.result) == 0 {
        return money.Wallet{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return money.Wallet{}, err
        }
    }
    return ret, nil
}
