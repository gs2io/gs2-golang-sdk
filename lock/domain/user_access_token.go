package domainlock
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

    "github.com/gs2io/gs2-golang-sdk/lock"
    "github.com/gs2io/gs2-golang-sdk/lock/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/lock/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client lock.Gs2LockRestClient
    namespaceName string
    accessToken auth.AccessToken
    mutexCache cache.MutexDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: lock.Gs2LockRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        mutexCache: cache.NewMutexDomainCache(),
    }
}

func (p UserAccessTokenDomain) Mutexes(
) iterator.DescribeMutexesIterator {
    return iterator.NewDescribeMutexesIterator(
        p.mutexCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Mutex(
    propertyId string,
) MutexAccessTokenDomain {
    return NewMutexAccessTokenDomain(
        p.session,
        p.mutexCache,
        p.namespaceName,
        p.accessToken,
        propertyId,
    )
}
