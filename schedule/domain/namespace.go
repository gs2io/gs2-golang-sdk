package domainschedule
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

    "github.com/gs2io/gs2-golang-sdk/schedule"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    eventMasterCache cache.EventMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        eventMasterCache: cache.NewEventMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request schedule.GetNamespaceStatusRequest,
) (*schedule.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request schedule.GetNamespaceRequest,
) (*schedule.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request schedule.UpdateNamespaceRequest,
) (*schedule.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request schedule.DeleteNamespaceRequest,
) (*schedule.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateEventMaster(
    request schedule.CreateEventMasterRequest,
) (*schedule.CreateEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) EventMasters(
) iterator.DescribeEventMastersIterator {
    return iterator.NewDescribeEventMastersIterator(
        p.eventMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) CurrentEventMaster(
) CurrentEventMasterDomain {
    return NewCurrentEventMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) EventMaster(
    eventName string,
) EventMasterDomain {
    return NewEventMasterDomain(
        p.session,
        p.eventMasterCache,
        p.namespaceName,
        eventName,
    )
}
