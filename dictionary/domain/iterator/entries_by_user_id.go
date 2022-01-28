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
    "github.com/gs2io/gs2-golang-sdk/dictionary"
    "github.com/gs2io/gs2-golang-sdk/dictionary/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeEntriesByUserIdIterator struct {
    entryCache cache.EntryDomainCache
    client dictionary.Gs2DictionaryRestClient
    namespaceName string
    userId string

    index int64
    pageToken *string
    last bool
    result []dictionary.Entry

    FetchSize *int32
}

func NewDescribeEntriesByUserIdIterator(
    entryCache cache.EntryDomainCache,
    client dictionary.Gs2DictionaryRestClient,
    namespaceName string,
    userId string,
) DescribeEntriesByUserIdIterator {
    it := DescribeEntriesByUserIdIterator{
        entryCache: entryCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]dictionary.Entry, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeEntriesByUserIdIterator) load() error {
    r, err := p.client.DescribeEntriesByUserId(
        &dictionary.DescribeEntriesByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.entryCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeEntriesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeEntriesByUserIdIterator) Next(

) (dictionary.Entry, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return dictionary.Entry{}, err
        }
    }
    if len(p.result) == 0 {
        return dictionary.Entry{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return dictionary.Entry{}, err
        }
    }
    return ret, nil
}
