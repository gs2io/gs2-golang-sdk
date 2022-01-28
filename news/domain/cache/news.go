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
    "github.com/gs2io/gs2-golang-sdk/news"
)

type NewsDomainCache struct {
    items []news.News
}

func NewNewsDomainCache() NewsDomainCache {
    return NewsDomainCache{
        items: make([]news.News, 0),
    }
}

func (p *NewsDomainCache) Update(
    items []news.News,
) {
    p.items = items
}

func (p NewsDomainCache) Get(
) []news.News {
    return p.items
}

func (p *NewsDomainCache) Delete(
) {
    p.items = make([]news.News, 0)
}
