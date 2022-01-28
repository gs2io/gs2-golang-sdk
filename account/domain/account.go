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

type AccountDomain struct {
    session *core.Gs2RestSession
    client account.Gs2AccountRestClient
    accountCache cache.AccountDomainCache
    namespaceName string
    userId string
    takeOverCache cache.TakeOverDomainCache
}

func NewAccountDomain(
    session *core.Gs2RestSession,
    accountCache cache.AccountDomainCache,
    namespaceName string,
    userId string,
) AccountDomain {
    return AccountDomain {
        session: session,
        client: account.Gs2AccountRestClient{
            Session: session,
        },
        accountCache: accountCache,
        namespaceName: namespaceName,
        userId: userId,
        takeOverCache: cache.NewTakeOverDomainCache(),
    }
}

func (p AccountDomain) UpdateTimeOffset(
    request account.UpdateTimeOffsetRequest,
) (*account.UpdateTimeOffsetResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.UpdateTimeOffset(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.accountCache.Update(*r.Item)
    return r, nil
}

func (p AccountDomain) Load(
    request account.GetAccountRequest,
) (*account.GetAccountResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetAccount(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.accountCache.Update(*r.Item)
    return r, nil
}

func (p AccountDomain) Delete(
    request account.DeleteAccountRequest,
) (*account.DeleteAccountResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteAccount(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.accountCache.Delete(*r.Item)
    return r, nil
}

func (p AccountDomain) Authentication(
    request account.AuthenticationRequest,
) (*account.AuthenticationResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.Authentication(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.accountCache.Update(*r.Item)
    return r, nil
}

func (p AccountDomain) CreateTakeOver(
    request account.CreateTakeOverByUserIdRequest,
) (*account.CreateTakeOverByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CreateTakeOverByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p AccountDomain) TakeOvers(
) iterator.DescribeTakeOversByUserIdIterator {
    return iterator.NewDescribeTakeOversByUserIdIterator(
        p.takeOverCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p AccountDomain) TakeOver(
    Type int32,
) TakeOverDomain {
    return NewTakeOverDomain(
        p.session,
        p.takeOverCache,
        p.namespaceName,
        p.userId,
        Type,
    )
}
