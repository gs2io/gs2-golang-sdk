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

type RoomDomain struct {
    session *core.Gs2RestSession
    client realtime.Gs2RealtimeRestClient
    roomCache cache.RoomDomainCache
    namespaceName string
    roomName string
}

func NewRoomDomain(
    session *core.Gs2RestSession,
    roomCache cache.RoomDomainCache,
    namespaceName string,
    roomName string,
) RoomDomain {
    return RoomDomain {
        session: session,
        client: realtime.Gs2RealtimeRestClient{
            Session: session,
        },
        roomCache: roomCache,
        namespaceName: namespaceName,
        roomName: roomName,
    }
}

func (p RoomDomain) Load(
    request realtime.GetRoomRequest,
) (*realtime.GetRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RoomName = &p.roomName
    r, err := p.client.GetRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Update(*r.Item)
    return r, nil
}

func (p RoomDomain) Delete(
    request realtime.DeleteRoomRequest,
) (*realtime.DeleteRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RoomName = &p.roomName
    r, err := p.client.DeleteRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Delete(*r.Item)
    return r, nil
}
