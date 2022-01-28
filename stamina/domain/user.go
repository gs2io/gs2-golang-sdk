package domainstamina
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

    "github.com/gs2io/gs2-golang-sdk/stamina"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    namespaceName string
    userId string
    staminaCache cache.StaminaDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        staminaCache: cache.NewStaminaDomainCache(),
    }
}

func (p UserDomain) Staminas(
) iterator.DescribeStaminasByUserIdIterator {
    return iterator.NewDescribeStaminasByUserIdIterator(
        p.staminaCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Stamina(
    staminaName string,
) StaminaDomain {
    return NewStaminaDomain(
        p.session,
        p.staminaCache,
        p.namespaceName,
        p.userId,
        staminaName,
    )
}
