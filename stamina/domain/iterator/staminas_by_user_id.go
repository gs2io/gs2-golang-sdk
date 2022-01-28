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
    "github.com/gs2io/gs2-golang-sdk/stamina"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeStaminasByUserIdIterator struct {
    staminaCache cache.StaminaDomainCache
    client stamina.Gs2StaminaRestClient
    namespaceName string
    userId string

    index int64
    pageToken *string
    last bool
    result []stamina.Stamina

    FetchSize *int32
}

func NewDescribeStaminasByUserIdIterator(
    staminaCache cache.StaminaDomainCache,
    client stamina.Gs2StaminaRestClient,
    namespaceName string,
    userId string,
) DescribeStaminasByUserIdIterator {
    it := DescribeStaminasByUserIdIterator{
        staminaCache: staminaCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]stamina.Stamina, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeStaminasByUserIdIterator) load() error {
    r, err := p.client.DescribeStaminasByUserId(
        &stamina.DescribeStaminasByUserIdRequest{
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
        p.staminaCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeStaminasByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeStaminasByUserIdIterator) Next(

) (stamina.Stamina, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return stamina.Stamina{}, err
        }
    }
    if len(p.result) == 0 {
        return stamina.Stamina{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return stamina.Stamina{}, err
        }
    }
    return ret, nil
}
