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
    "github.com/gs2io/gs2-golang-sdk/lock"
    "github.com/gs2io/gs2-golang-sdk/lock/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeMutexesByUserIdIterator struct {
    mutexCache cache.MutexDomainCache
    client lock.Gs2LockRestClient
    namespaceName string
    userId string

    index int64
    pageToken *string
    last bool
    result []lock.Mutex

    FetchSize *int32
}

func NewDescribeMutexesByUserIdIterator(
    mutexCache cache.MutexDomainCache,
    client lock.Gs2LockRestClient,
    namespaceName string,
    userId string,
) DescribeMutexesByUserIdIterator {
    it := DescribeMutexesByUserIdIterator{
        mutexCache: mutexCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]lock.Mutex, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeMutexesByUserIdIterator) load() error {
    r, err := p.client.DescribeMutexesByUserId(
        &lock.DescribeMutexesByUserIdRequest{
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
        p.mutexCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeMutexesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeMutexesByUserIdIterator) Next(

) (lock.Mutex, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lock.Mutex{}, err
        }
    }
    if len(p.result) == 0 {
        return lock.Mutex{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return lock.Mutex{}, err
        }
    }
    return ret, nil
}
