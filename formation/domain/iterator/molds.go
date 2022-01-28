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


type DescribeMoldsIterator struct {
    moldCache cache.MoldDomainCache
    client formation.Gs2FormationRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []formation.Mold

    FetchSize *int32
}

func NewDescribeMoldsIterator(
    moldCache cache.MoldDomainCache,
    client formation.Gs2FormationRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeMoldsIterator {
    it := DescribeMoldsIterator{
        moldCache: moldCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]formation.Mold, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeMoldsIterator) load() error {
    r, err := p.client.DescribeMolds(
        &formation.DescribeMoldsRequest{
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
        p.moldCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeMoldsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeMoldsIterator) Next(

) (formation.Mold, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Mold{}, err
        }
    }
    if len(p.result) == 0 {
        return formation.Mold{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Mold{}, err
        }
    }
    return ret, nil
}