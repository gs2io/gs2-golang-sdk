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
    "github.com/gs2io/gs2-golang-sdk/lottery"
)

type ProbabilityDomainCache struct {
    items []lottery.Probability
}

func NewProbabilityDomainCache() ProbabilityDomainCache {
    return ProbabilityDomainCache{
        items: make([]lottery.Probability, 0),
    }
}

func (p *ProbabilityDomainCache) Update(
    items []lottery.Probability,
) {
    p.items = items
}

func (p ProbabilityDomainCache) Get(
) []lottery.Probability {
    return p.items
}

func (p *ProbabilityDomainCache) Delete(
) {
    p.items = make([]lottery.Probability, 0)
}
