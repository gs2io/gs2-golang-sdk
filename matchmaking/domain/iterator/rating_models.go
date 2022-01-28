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
    "github.com/gs2io/gs2-golang-sdk/matchmaking"
    "github.com/gs2io/gs2-golang-sdk/matchmaking/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeRatingModelsIterator struct {
    ratingModelCache cache.RatingModelDomainCache
    client matchmaking.Gs2MatchmakingRestClient
    namespaceName string

    index int64
    last bool
    result []matchmaking.RatingModel

    FetchSize *int32
}

func NewDescribeRatingModelsIterator(
    ratingModelCache cache.RatingModelDomainCache,
    client matchmaking.Gs2MatchmakingRestClient,
    namespaceName string,
) DescribeRatingModelsIterator {
    it := DescribeRatingModelsIterator{
        ratingModelCache: ratingModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]matchmaking.RatingModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeRatingModelsIterator) load() error {
    r, err := p.client.DescribeRatingModels(
        &matchmaking.DescribeRatingModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.ratingModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeRatingModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeRatingModelsIterator) Next(

) (matchmaking.RatingModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return matchmaking.RatingModel{}, err
        }
    }
    if len(p.result) == 0 {
        return matchmaking.RatingModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return matchmaking.RatingModel{}, err
        }
    }
    return ret, nil
}
