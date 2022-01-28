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

type WalletDomain struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    walletCache cache.WalletDomainCache
    namespaceName string
    userId string
    slot int32
}

func NewWalletDomain(
    session *core.Gs2RestSession,
    walletCache cache.WalletDomainCache,
    namespaceName string,
    userId string,
    slot int32,
) WalletDomain {
    return WalletDomain {
        session: session,
        client: money.Gs2MoneyRestClient{
            Session: session,
        },
        walletCache: walletCache,
        namespaceName: namespaceName,
        userId: userId,
        slot: slot,
    }
}

func (p WalletDomain) Load(
    request money.GetWalletByUserIdRequest,
) (*money.GetWalletByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.Slot = &p.slot
    r, err := p.client.GetWalletByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.walletCache.Update(*r.Item)
    return r, nil
}

func (p WalletDomain) Deposit(
    request money.DepositByUserIdRequest,
) (*money.DepositByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.Slot = &p.slot
    r, err := p.client.DepositByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.walletCache.Update(*r.Item)
    return r, nil
}

func (p WalletDomain) Withdraw(
    request money.WithdrawByUserIdRequest,
) (*money.WithdrawByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.Slot = &p.slot
    r, err := p.client.WithdrawByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.walletCache.Update(*r.Item)
    return r, nil
}
