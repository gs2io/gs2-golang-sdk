package domainmoney
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
    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/core"

    "github.com/gs2io/gs2-golang-sdk/money"
    "github.com/gs2io/gs2-golang-sdk/money/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/money/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type WalletAccessTokenDomain struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    walletCache cache.WalletDomainCache
    namespaceName string
    accessToken auth.AccessToken
    slot int32
}

func NewWalletAccessTokenDomain(
    session *core.Gs2RestSession,
    walletCache cache.WalletDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    slot int32,
) WalletAccessTokenDomain {
    return WalletAccessTokenDomain {
        session: session,
        client: money.Gs2MoneyRestClient{
            Session: session,
        },
        walletCache: walletCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        slot: slot,
    }
}

func (p WalletAccessTokenDomain) Load(
    request money.GetWalletRequest,
) (*money.GetWalletResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.Slot = &p.slot
    r, err := p.client.GetWallet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.walletCache.Update(*r.Item)
    return r, nil
}

func (p WalletAccessTokenDomain) Withdraw(
    request money.WithdrawRequest,
) (*money.WithdrawResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.Slot = &p.slot
    r, err := p.client.Withdraw(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.walletCache.Update(*r.Item)
    return r, nil
}
