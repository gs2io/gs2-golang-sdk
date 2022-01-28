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
    "github.com/gs2io/gs2-golang-sdk/chat"
)

type MessageDomainCache struct {
    items map[string]chat.Message
}

func NewMessageDomainCache() MessageDomainCache {
    return MessageDomainCache{
        items: map[string]chat.Message{},
    }
}

func (p *MessageDomainCache) Update(
    item chat.Message,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    p.items[keys] = item
}

func (p MessageDomainCache) Get(
    messageName string,
) chat.Message {
    keys := strings.Join([]string{
        core.ToString(messageName),
    }, ":")
    return p.items[keys]
}

func (p *MessageDomainCache) Delete(
    item chat.Message,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Name),
    }, ":")
    delete(p.items, keys)
}
