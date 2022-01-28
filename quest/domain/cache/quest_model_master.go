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
    "github.com/gs2io/gs2-golang-sdk/quest"
)

type QuestModelMasterDomainCache struct {
    items map[string]quest.QuestModelMaster
}

func NewQuestModelMasterDomainCache() QuestModelMasterDomainCache {
    return QuestModelMasterDomainCache{
        items: map[string]quest.QuestModelMaster{},
    }
}

func (p *QuestModelMasterDomainCache) Update(
    item quest.QuestModelMaster,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    p.items[keys] = item
}

func (p QuestModelMasterDomainCache) Get(
    questName string,
) quest.QuestModelMaster {
    keys := strings.Join([]string{
        core.ToString(questName),
    }, ":")
    return p.items[keys]
}

func (p *QuestModelMasterDomainCache) Delete(
    item quest.QuestModelMaster,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    delete(p.items, keys)
}
