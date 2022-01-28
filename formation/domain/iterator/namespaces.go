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
    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeNamespacesIterator struct {
    namespaceCache cache.NamespaceDomainCache
    client formation.Gs2FormationRestClient

    index int64
    pageToken *string
    last bool
    result []formation.Namespace

    FetchSize *int32
}

func NewDescribeNamespacesIterator(
    namespaceCache cache.NamespaceDomainCache,
    client formation.Gs2FormationRestClient,
) DescribeNamespacesIterator {
    it := DescribeNamespacesIterator{
        namespaceCache: namespaceCache,
        client: client,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]formation.Namespace, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeNamespacesIterator) load() error {
    r, err := p.client.DescribeNamespaces(
        &formation.DescribeNamespacesRequest{
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.namespaceCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeNamespacesIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeNamespacesIterator) Next(

) (formation.Namespace, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Namespace{}, err
        }
    }
    if len(p.result) == 0 {
        return formation.Namespace{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Namespace{}, err
        }
    }
    return ret, nil
}
