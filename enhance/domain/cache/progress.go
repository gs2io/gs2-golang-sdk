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
    "github.com/gs2io/gs2-golang-sdk/enhance"
)

type ProgressDomainCache struct {
    items map[string]enhance.Progress
}

func NewProgressDomainCache() ProgressDomainCache {
    return ProgressDomainCache{
        items: map[string]enhance.Progress{},
    }
}

func (p *ProgressDomainCache) Update(
    item enhance.Progress,
) {
    keys := strings.Join([]string{
        core.ToString(*item.ProgressId),
    }, ":")
    p.items[keys] = item
}

func (p ProgressDomainCache) Get(
    progressId string,
) enhance.Progress {
    keys := strings.Join([]string{
        core.ToString(progressId),
    }, ":")
    return p.items[keys]
}

func (p *ProgressDomainCache) Delete(
    item enhance.Progress,
) {
    keys := strings.Join([]string{
        core.ToString(*item.ProgressId),
    }, ":")
    delete(p.items, keys)
}
