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
    "github.com/gs2io/gs2-golang-sdk/key"
    "github.com/gs2io/gs2-golang-sdk/key/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeKeysIterator struct {
    keyCache cache.KeyDomainCache
    client key.Gs2KeyRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []key.Key

    FetchSize *int32
}

func NewDescribeKeysIterator(
    keyCache cache.KeyDomainCache,
    client key.Gs2KeyRestClient,
    namespaceName string,
) DescribeKeysIterator {
    it := DescribeKeysIterator{
        keyCache: keyCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]key.Key, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeKeysIterator) load() error {
    r, err := p.client.DescribeKeys(
        &key.DescribeKeysRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.keyCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeKeysIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeKeysIterator) Next(

) (key.Key, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return key.Key{}, err
        }
    }
    if len(p.result) == 0 {
        return key.Key{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return key.Key{}, err
        }
    }
    return ret, nil
}
