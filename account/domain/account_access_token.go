package domainaccount
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

    "github.com/gs2io/gs2-golang-sdk/account"
    "github.com/gs2io/gs2-golang-sdk/account/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/account/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type AccountAccessTokenDomain struct {
    session *core.Gs2RestSession
    client account.Gs2AccountRestClient
    accountCache cache.AccountDomainCache
    namespaceName string
    accessToken auth.AccessToken
    takeOverCache cache.TakeOverDomainCache
}

func NewAccountAccessTokenDomain(
    session *core.Gs2RestSession,
    accountCache cache.AccountDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
) AccountAccessTokenDomain {
    return AccountAccessTokenDomain {
        session: session,
        client: account.Gs2AccountRestClient{
            Session: session,
        },
        accountCache: accountCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        takeOverCache: cache.NewTakeOverDomainCache(),
    }
}

func (p AccountAccessTokenDomain) CreateTakeOver(
    request account.CreateTakeOverRequest,
) (*account.CreateTakeOverResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.CreateTakeOver(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p AccountAccessTokenDomain) TakeOvers(
) iterator.DescribeTakeOversIterator {
    return iterator.NewDescribeTakeOversIterator(
        p.takeOverCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p AccountAccessTokenDomain) TakeOver(
    Type int32,
) TakeOverAccessTokenDomain {
    return NewTakeOverAccessTokenDomain(
        p.session,
        p.takeOverCache,
        p.namespaceName,
        p.accessToken,
        Type,
    )
}
