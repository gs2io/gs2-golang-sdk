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

type UserDomain struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    namespaceName string
    userId string
    walletCache cache.WalletDomainCache
    receiptCache cache.ReceiptDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: money.Gs2MoneyRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        walletCache: cache.NewWalletDomainCache(),
        receiptCache: cache.NewReceiptDomainCache(),
    }
}

func (p UserDomain) Wallets(
) iterator.DescribeWalletsByUserIdIterator {
    return iterator.NewDescribeWalletsByUserIdIterator(
        p.walletCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Receipts(
    slot *int32,
    begin *int64,
    end *int64,
) iterator.DescribeReceiptsIterator {
    return iterator.NewDescribeReceiptsIterator(
        p.receiptCache,
        p.client,
        p.namespaceName,
        p.userId,
        slot,
        begin,
        end,
    )
}

func (p UserDomain) Wallet(
    slot int32,
) WalletDomain {
    return NewWalletDomain(
        p.session,
        p.walletCache,
        p.namespaceName,
        p.userId,
        slot,
    )
}

func (p UserDomain) Receipt(
    transactionId string,
) ReceiptDomain {
    return NewReceiptDomain(
        p.session,
        p.receiptCache,
        p.namespaceName,
        p.userId,
        transactionId,
    )
}