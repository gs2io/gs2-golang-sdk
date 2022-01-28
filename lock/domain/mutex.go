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

type MutexDomain struct {
    session *core.Gs2RestSession
    client lock.Gs2LockRestClient
    mutexCache cache.MutexDomainCache
    namespaceName string
    userId string
    propertyId string
}

func NewMutexDomain(
    session *core.Gs2RestSession,
    mutexCache cache.MutexDomainCache,
    namespaceName string,
    userId string,
    propertyId string,
) MutexDomain {
    return MutexDomain {
        session: session,
        client: lock.Gs2LockRestClient{
            Session: session,
        },
        mutexCache: mutexCache,
        namespaceName: namespaceName,
        userId: userId,
        propertyId: propertyId,
    }
}

func (p MutexDomain) Lock(
    request lock.LockByUserIdRequest,
) (*lock.LockByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PropertyId = &p.propertyId
    r, err := p.client.LockByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}

func (p MutexDomain) Unlock(
    request lock.UnlockByUserIdRequest,
) (*lock.UnlockByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PropertyId = &p.propertyId
    r, err := p.client.UnlockByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}

func (p MutexDomain) Load(
    request lock.GetMutexByUserIdRequest,
) (*lock.GetMutexByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PropertyId = &p.propertyId
    r, err := p.client.GetMutexByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}

func (p MutexDomain) Delete(
    request lock.DeleteMutexByUserIdRequest,
) (*lock.DeleteMutexByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PropertyId = &p.propertyId
    r, err := p.client.DeleteMutexByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Delete(*r.Item)
    return r, nil
}
