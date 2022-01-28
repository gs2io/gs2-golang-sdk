package domainmission
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

    "github.com/gs2io/gs2-golang-sdk/mission"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CompleteDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    completeCache cache.CompleteDomainCache
    namespaceName string
    userId string
    missionGroupName string
}

func NewCompleteDomain(
    session *core.Gs2RestSession,
    completeCache cache.CompleteDomainCache,
    namespaceName string,
    userId string,
    missionGroupName string,
) CompleteDomain {
    return CompleteDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        completeCache: completeCache,
        namespaceName: namespaceName,
        userId: userId,
        missionGroupName: missionGroupName,
    }
}

func (p CompleteDomain) Complete(
    request mission.CompleteByUserIdRequest,
) (*mission.CompleteByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.CompleteByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CompleteDomain) Receive(
    request mission.ReceiveByUserIdRequest,
) (*mission.ReceiveByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.ReceiveByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completeCache.Update(*r.Item)
    return r, nil
}

func (p CompleteDomain) Load(
    request mission.GetCompleteByUserIdRequest,
) (*mission.GetCompleteByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.GetCompleteByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completeCache.Update(*r.Item)
    return r, nil
}

func (p CompleteDomain) Delete(
    request mission.DeleteCompleteByUserIdRequest,
) (*mission.DeleteCompleteByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.DeleteCompleteByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completeCache.Delete(*r.Item)
    return r, nil
}
