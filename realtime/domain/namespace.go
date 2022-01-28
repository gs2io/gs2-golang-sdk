package domainrealtime
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

    "github.com/gs2io/gs2-golang-sdk/realtime"
    "github.com/gs2io/gs2-golang-sdk/realtime/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/realtime/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client realtime.Gs2RealtimeRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    roomCache cache.RoomDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: realtime.Gs2RealtimeRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        roomCache: cache.NewRoomDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request realtime.GetNamespaceStatusRequest,
) (*realtime.GetNamespaceStatusResult, error) {
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
    request realtime.GetNamespaceRequest,
) (*realtime.GetNamespaceResult, error) {
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
    request realtime.UpdateNamespaceRequest,
) (*realtime.UpdateNamespaceResult, error) {
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
    request realtime.DeleteNamespaceRequest,
) (*realtime.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) Rooms(
) iterator.DescribeRoomsIterator {
    return iterator.NewDescribeRoomsIterator(
        p.roomCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) Room(
    roomName string,
) RoomDomain {
    return NewRoomDomain(
        p.session,
        p.roomCache,
        p.namespaceName,
        roomName,
    )
}
