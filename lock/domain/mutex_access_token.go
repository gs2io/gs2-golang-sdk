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

type MutexAccessTokenDomain struct {
    session *core.Gs2RestSession
    client lock.Gs2LockRestClient
    mutexCache cache.MutexDomainCache
    namespaceName string
    accessToken auth.AccessToken
    propertyId string
}

func NewMutexAccessTokenDomain(
    session *core.Gs2RestSession,
    mutexCache cache.MutexDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    propertyId string,
) MutexAccessTokenDomain {
    return MutexAccessTokenDomain {
        session: session,
        client: lock.Gs2LockRestClient{
            Session: session,
        },
        mutexCache: mutexCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        propertyId: propertyId,
    }
}

func (p MutexAccessTokenDomain) Lock(
    request lock.LockRequest,
) (*lock.LockResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.PropertyId = &p.propertyId
    r, err := p.client.Lock(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}

func (p MutexAccessTokenDomain) Unlock(
    request lock.UnlockRequest,
) (*lock.UnlockResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.PropertyId = &p.propertyId
    r, err := p.client.Unlock(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}

func (p MutexAccessTokenDomain) Load(
    request lock.GetMutexRequest,
) (*lock.GetMutexResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.PropertyId = &p.propertyId
    r, err := p.client.GetMutex(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.mutexCache.Update(*r.Item)
    return r, nil
}
