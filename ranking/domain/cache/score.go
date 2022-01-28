package cache
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
    "strings"

    "github.com/gs2io/gs2-golang-sdk/core"
    "github.com/gs2io/gs2-golang-sdk/ranking"
)

type ScoreDomainCache struct {
    items map[string]ranking.Score
}

func NewScoreDomainCache() ScoreDomainCache {
    return ScoreDomainCache{
        items: map[string]ranking.Score{},
    }
}

func (p *ScoreDomainCache) Update(
    item ranking.Score,
) {
    keys := strings.Join([]string{
        core.ToString(*item.CategoryName),
        core.ToString(*item.ScorerUserId),
        core.ToString(*item.UniqueId),
    }, ":")
    p.items[keys] = item
}

func (p ScoreDomainCache) Get(
    categoryName string,
    scorerUserId string,
    uniqueId string,
) ranking.Score {
    keys := strings.Join([]string{
        core.ToString(categoryName),
        core.ToString(scorerUserId),
        core.ToString(uniqueId),
    }, ":")
    return p.items[keys]
}

func (p *ScoreDomainCache) Delete(
    item ranking.Score,
) {
    keys := strings.Join([]string{
        core.ToString(*item.CategoryName),
        core.ToString(*item.ScorerUserId),
        core.ToString(*item.UniqueId),
    }, ":")
    delete(p.items, keys)
}
