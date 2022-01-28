package domainformation
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

    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    namespaceName string
    accessToken auth.AccessToken
    moldCache cache.MoldDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        moldCache: cache.NewMoldDomainCache(),
    }
}

func (p UserAccessTokenDomain) Molds(
) iterator.DescribeMoldsIterator {
    return iterator.NewDescribeMoldsIterator(
        p.moldCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Mold(
    moldName string,
) MoldAccessTokenDomain {
    return NewMoldAccessTokenDomain(
        p.session,
        p.moldCache,
        p.namespaceName,
        p.accessToken,
        moldName,
    )
}
