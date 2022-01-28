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
    "github.com/gs2io/gs2-golang-sdk/money"
)

type WalletDomainCache struct {
    items map[string]money.Wallet
}

func NewWalletDomainCache() WalletDomainCache {
    return WalletDomainCache{
        items: map[string]money.Wallet{},
    }
}

func (p *WalletDomainCache) Update(
    item money.Wallet,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Slot),
    }, ":")
    p.items[keys] = item
}

func (p WalletDomainCache) Get(
    slot int32,
) money.Wallet {
    keys := strings.Join([]string{
        core.ToString(slot),
    }, ":")
    return p.items[keys]
}

func (p *WalletDomainCache) Delete(
    item money.Wallet,
) {
    keys := strings.Join([]string{
        core.ToString(*item.Slot),
    }, ":")
    delete(p.items, keys)
}
