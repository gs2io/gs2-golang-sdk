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
    "github.com/gs2io/gs2-golang-sdk/account"
    "github.com/gs2io/gs2-golang-sdk/account/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeTakeOversIterator struct {
    takeOverCache cache.TakeOverDomainCache
    client account.Gs2AccountRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []account.TakeOver

    FetchSize *int32
}

func NewDescribeTakeOversIterator(
    takeOverCache cache.TakeOverDomainCache,
    client account.Gs2AccountRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeTakeOversIterator {
    it := DescribeTakeOversIterator{
        takeOverCache: takeOverCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]account.TakeOver, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeTakeOversIterator) load() error {
    r, err := p.client.DescribeTakeOvers(
        &account.DescribeTakeOversRequest{
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
        p.takeOverCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeTakeOversIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeTakeOversIterator) Next(

) (account.TakeOver, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return account.TakeOver{}, err
        }
    }
    if len(p.result) == 0 {
        return account.TakeOver{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return account.TakeOver{}, err
        }
    }
    return ret, nil
}
