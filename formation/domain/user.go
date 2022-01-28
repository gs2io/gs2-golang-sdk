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

type UserDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    namespaceName string
    userId string
    moldCache cache.MoldDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        moldCache: cache.NewMoldDomainCache(),
    }
}

func (p UserDomain) Molds(
) iterator.DescribeMoldsByUserIdIterator {
    return iterator.NewDescribeMoldsByUserIdIterator(
        p.moldCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Mold(
    moldName string,
) MoldDomain {
    return NewMoldDomain(
        p.session,
        p.moldCache,
        p.namespaceName,
        p.userId,
        moldName,
    )
}
