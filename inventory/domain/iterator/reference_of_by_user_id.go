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


type DescribeReferenceOfByUserIdIterator struct {
    client inventory.Gs2InventoryRestClient
    namespaceName string
    inventoryName string
    userId string
    itemName string
    itemSetName string

    index int64
    last bool
    result []string

    FetchSize *int32
}

func NewDescribeReferenceOfByUserIdIterator(
    client inventory.Gs2InventoryRestClient,
    namespaceName string,
    inventoryName string,
    userId string,
    itemName string,
    itemSetName string,
) DescribeReferenceOfByUserIdIterator {
    it := DescribeReferenceOfByUserIdIterator{
        client: client,
        namespaceName: namespaceName,
        inventoryName: inventoryName,
        userId: userId,
        itemName: itemName,
        itemSetName: itemSetName,

        index: 0,
        last: false,
        result: make([]string, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeReferenceOfByUserIdIterator) load() error {
    r, err := p.client.DescribeReferenceOfByUserId(
        &inventory.DescribeReferenceOfByUserIdRequest{
            NamespaceName: &p.namespaceName,
            InventoryName: &p.inventoryName,
            UserId: &p.userId,
            ItemName: &p.itemName,
            ItemSetName: &p.itemSetName,
        },
    )
    if err != nil {
        return err
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeReferenceOfByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeReferenceOfByUserIdIterator) Next(

) (string, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return "", err
        }
    }
    if len(p.result) == 0 {
        return "", errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return "", err
        }
    }
    return ret, nil
}
