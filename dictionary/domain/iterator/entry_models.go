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


type DescribeEntryModelsIterator struct {
    entryModelCache cache.EntryModelDomainCache
    client dictionary.Gs2DictionaryRestClient
    namespaceName string

    index int64
    last bool
    result []dictionary.EntryModel

    FetchSize *int32
}

func NewDescribeEntryModelsIterator(
    entryModelCache cache.EntryModelDomainCache,
    client dictionary.Gs2DictionaryRestClient,
    namespaceName string,
) DescribeEntryModelsIterator {
    it := DescribeEntryModelsIterator{
        entryModelCache: entryModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]dictionary.EntryModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeEntryModelsIterator) load() error {
    r, err := p.client.DescribeEntryModels(
        &dictionary.DescribeEntryModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.entryModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeEntryModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeEntryModelsIterator) Next(

) (dictionary.EntryModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return dictionary.EntryModel{}, err
        }
    }
    if len(p.result) == 0 {
        return dictionary.EntryModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return dictionary.EntryModel{}, err
        }
    }
    return ret, nil
}
