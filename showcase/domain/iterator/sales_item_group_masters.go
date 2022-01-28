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
    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeSalesItemGroupMastersIterator struct {
    salesItemGroupMasterCache cache.SalesItemGroupMasterDomainCache
    client showcase.Gs2ShowcaseRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []showcase.SalesItemGroupMaster

    FetchSize *int32
}

func NewDescribeSalesItemGroupMastersIterator(
    salesItemGroupMasterCache cache.SalesItemGroupMasterDomainCache,
    client showcase.Gs2ShowcaseRestClient,
    namespaceName string,
) DescribeSalesItemGroupMastersIterator {
    it := DescribeSalesItemGroupMastersIterator{
        salesItemGroupMasterCache: salesItemGroupMasterCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]showcase.SalesItemGroupMaster, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeSalesItemGroupMastersIterator) load() error {
    r, err := p.client.DescribeSalesItemGroupMasters(
        &showcase.DescribeSalesItemGroupMastersRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.salesItemGroupMasterCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeSalesItemGroupMastersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeSalesItemGroupMastersIterator) Next(

) (showcase.SalesItemGroupMaster, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.SalesItemGroupMaster{}, err
        }
    }
    if len(p.result) == 0 {
        return showcase.SalesItemGroupMaster{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return showcase.SalesItemGroupMaster{}, err
        }
    }
    return ret, nil
}