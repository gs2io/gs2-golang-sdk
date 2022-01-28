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

type ReceiptDomain struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    receiptCache cache.ReceiptDomainCache
    namespaceName string
    userId string
    transactionId string
}

func NewReceiptDomain(
    session *core.Gs2RestSession,
    receiptCache cache.ReceiptDomainCache,
    namespaceName string,
    userId string,
    transactionId string,
) ReceiptDomain {
    return ReceiptDomain {
        session: session,
        client: money.Gs2MoneyRestClient{
            Session: session,
        },
        receiptCache: receiptCache,
        namespaceName: namespaceName,
        userId: userId,
        transactionId: transactionId,
    }
}

func (p ReceiptDomain) GetByUserIdAndTransactionId(
    request money.GetByUserIdAndTransactionIdRequest,
) (*money.GetByUserIdAndTransactionIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.TransactionId = &p.transactionId
    r, err := p.client.GetByUserIdAndTransactionId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.receiptCache.Update(*r.Item)
    return r, nil
}