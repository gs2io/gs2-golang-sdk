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

type CompleteAccessTokenDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    completeCache cache.CompleteDomainCache
    namespaceName string
    accessToken auth.AccessToken
    missionGroupName string
}

func NewCompleteAccessTokenDomain(
    session *core.Gs2RestSession,
    completeCache cache.CompleteDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    missionGroupName string,
) CompleteAccessTokenDomain {
    return CompleteAccessTokenDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        completeCache: completeCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        missionGroupName: missionGroupName,
    }
}

func (p CompleteAccessTokenDomain) Complete(
    request mission.CompleteRequest,
) (*mission.CompleteResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.Complete(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CompleteAccessTokenDomain) Load(
    request mission.GetCompleteRequest,
) (*mission.GetCompleteResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MissionGroupName = &p.missionGroupName
    r, err := p.client.GetComplete(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completeCache.Update(*r.Item)
    return r, nil
}
