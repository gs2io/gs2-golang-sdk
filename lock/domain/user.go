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

type UserDomain struct {
    session *core.Gs2RestSession
    client lock.Gs2LockRestClient
    namespaceName string
    userId string
    mutexCache cache.MutexDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: lock.Gs2LockRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        mutexCache: cache.NewMutexDomainCache(),
    }
}

func (p UserDomain) Mutexes(
) iterator.DescribeMutexesByUserIdIterator {
    return iterator.NewDescribeMutexesByUserIdIterator(
        p.mutexCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Mutex(
    propertyId string,
) MutexDomain {
    return NewMutexDomain(
        p.session,
        p.mutexCache,
        p.namespaceName,
        p.userId,
        propertyId,
    )
}
