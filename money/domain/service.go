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
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Money struct {
    session *core.Gs2RestSession
    client money.Gs2MoneyRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Money(
    session *core.Gs2RestSession,
) Gs2Money {
    return Gs2Money {
        session: session,
        client: money.Gs2MoneyRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Money) CreateNamespace(
    request money.CreateNamespaceRequest,
) (*money.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Money) DepositByStampSheet(
    request money.DepositByStampSheetRequest,
) (*money.DepositByStampSheetResult, error) {
    r, err := p.client.DepositByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Money) WithdrawByStampTask(
    request money.WithdrawByStampTaskRequest,
) (*money.WithdrawByStampTaskResult, error) {
    r, err := p.client.WithdrawByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Money) RecordReceipt(
    request money.RecordReceiptRequest,
) (*money.RecordReceiptResult, error) {
    r, err := p.client.RecordReceipt(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Money) RecordReceiptByStampTask(
    request money.RecordReceiptByStampTaskRequest,
) (*money.RecordReceiptByStampTaskResult, error) {
    r, err := p.client.RecordReceiptByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Money) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Money) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
