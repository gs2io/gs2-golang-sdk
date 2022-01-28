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
    "github.com/gs2io/gs2-golang-sdk/schedule"
)

type EventMasterDomainCache struct {
    items map[string]schedule.EventMaster
}

func NewEventMasterDomainCache() EventMasterDomainCache {
    return EventMasterDomainCache{
        items: map[string]schedule.EventMaster{},
    }
}

func (p *EventMasterDomainCache) Update(
    item schedule.EventMaster,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    p.items[keys] = item
}

func (p EventMasterDomainCache) Get(
    eventName string,
) schedule.EventMaster {
    keys := strings.Join([]string{
        core.ToString(eventName),
    }, ":")
    return p.items[keys]
}

func (p *EventMasterDomainCache) Delete(
    item schedule.EventMaster,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    delete(p.items, keys)
}
