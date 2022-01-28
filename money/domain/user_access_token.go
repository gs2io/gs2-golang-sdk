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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    namespaceName string
    accessToken auth.AccessToken
    walletCache cache.WalletDomainCache
    receiptCache cache.ReceiptDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: money.Gs2MoneyRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        walletCache: cache.NewWalletDomainCache(),
        receiptCache: cache.NewReceiptDomainCache(),
    }
}

func (p UserAccessTokenDomain) Wallets(
) iterator.DescribeWalletsIterator {
    return iterator.NewDescribeWalletsIterator(
        p.walletCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Wallet(
    slot int32,
) WalletAccessTokenDomain {
    return NewWalletAccessTokenDomain(
        p.session,
        p.walletCache,
        p.namespaceName,
        p.accessToken,
        slot,
    )
}

func (p UserAccessTokenDomain) Receipt(
    transactionId string,
) ReceiptAccessTokenDomain {
    return NewReceiptAccessTokenDomain(
        p.session,
        p.receiptCache,
        p.namespaceName,
        p.accessToken,
        transactionId,
    )
}