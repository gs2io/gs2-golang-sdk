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
    "github.com/gs2io/gs2-golang-sdk/inventory"
    "github.com/gs2io/gs2-golang-sdk/inventory/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeItemSetsByUserIdIterator struct {
    itemSetCache cache.ItemSetDomainCache
    client inventory.Gs2InventoryRestClient
    namespaceName string
    inventoryName string
    userId string

    index int64
    pageToken *string
    last bool
    result []inventory.ItemSet

    FetchSize *int32
}

func NewDescribeItemSetsByUserIdIterator(
    itemSetCache cache.ItemSetDomainCache,
    client inventory.Gs2InventoryRestClient,
    namespaceName string,
    inventoryName string,
    userId string,
) DescribeItemSetsByUserIdIterator {
    it := DescribeItemSetsByUserIdIterator{
        itemSetCache: itemSetCache,
        client: client,
        namespaceName: namespaceName,
        inventoryName: inventoryName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]inventory.ItemSet, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeItemSetsByUserIdIterator) load() error {
    r, err := p.client.DescribeItemSetsByUserId(
        &inventory.DescribeItemSetsByUserIdRequest{
            NamespaceName: &p.namespaceName,
            InventoryName: &p.inventoryName,
            UserId: &p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.itemSetCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeItemSetsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeItemSetsByUserIdIterator) Next(

) (inventory.ItemSet, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return inventory.ItemSet{}, err
        }
    }
    if len(p.result) == 0 {
        return inventory.ItemSet{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return inventory.ItemSet{}, err
        }
    }
    return ret, nil
}
