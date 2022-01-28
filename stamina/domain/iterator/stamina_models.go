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


type DescribeStaminaModelsIterator struct {
    staminaModelCache cache.StaminaModelDomainCache
    client stamina.Gs2StaminaRestClient
    namespaceName string

    index int64
    last bool
    result []stamina.StaminaModel

    FetchSize *int32
}

func NewDescribeStaminaModelsIterator(
    staminaModelCache cache.StaminaModelDomainCache,
    client stamina.Gs2StaminaRestClient,
    namespaceName string,
) DescribeStaminaModelsIterator {
    it := DescribeStaminaModelsIterator{
        staminaModelCache: staminaModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]stamina.StaminaModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeStaminaModelsIterator) load() error {
    r, err := p.client.DescribeStaminaModels(
        &stamina.DescribeStaminaModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.staminaModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeStaminaModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeStaminaModelsIterator) Next(

) (stamina.StaminaModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return stamina.StaminaModel{}, err
        }
    }
    if len(p.result) == 0 {
        return stamina.StaminaModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return stamina.StaminaModel{}, err
        }
    }
    return ret, nil
}
