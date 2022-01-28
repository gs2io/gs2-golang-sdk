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

type TakeOverAccessTokenDomain struct {
    session *core.Gs2RestSession
    client account.Gs2AccountRestClient
    takeOverCache cache.TakeOverDomainCache
    namespaceName string
    accessToken auth.AccessToken
    Type int32
}

func NewTakeOverAccessTokenDomain(
    session *core.Gs2RestSession,
    takeOverCache cache.TakeOverDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    Type int32,
) TakeOverAccessTokenDomain {
    return TakeOverAccessTokenDomain {
        session: session,
        client: account.Gs2AccountRestClient{
            Session: session,
        },
        takeOverCache: takeOverCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        Type: Type,
    }
}

func (p TakeOverAccessTokenDomain) Load(
    request account.GetTakeOverRequest,
) (*account.GetTakeOverResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.Type = &p.Type
    r, err := p.client.GetTakeOver(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.takeOverCache.Update(*r.Item)
    return r, nil
}

func (p TakeOverAccessTokenDomain) Update(
    request account.UpdateTakeOverRequest,
) (*account.UpdateTakeOverResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.Type = &p.Type
    r, err := p.client.UpdateTakeOver(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.takeOverCache.Update(*r.Item)
    return r, nil
}

func (p TakeOverAccessTokenDomain) Delete(
    request account.DeleteTakeOverRequest,
) (*account.DeleteTakeOverResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.Type = &p.Type
    r, err := p.client.DeleteTakeOver(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.takeOverCache.Delete(*r.Item)
    return r, nil
}
