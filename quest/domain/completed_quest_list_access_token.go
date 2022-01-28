package domainquest
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

    "github.com/gs2io/gs2-golang-sdk/quest"
    "github.com/gs2io/gs2-golang-sdk/quest/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/quest/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CompletedQuestListAccessTokenDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    completedQuestListCache cache.CompletedQuestListDomainCache
    namespaceName string
    accessToken auth.AccessToken
    questGroupName string
}

func NewCompletedQuestListAccessTokenDomain(
    session *core.Gs2RestSession,
    completedQuestListCache cache.CompletedQuestListDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    questGroupName string,
) CompletedQuestListAccessTokenDomain {
    return CompletedQuestListAccessTokenDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        completedQuestListCache: completedQuestListCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        questGroupName: questGroupName,
    }
}

func (p CompletedQuestListAccessTokenDomain) Load(
    request quest.GetCompletedQuestListRequest,
) (*quest.GetCompletedQuestListResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.GetCompletedQuestList(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completedQuestListCache.Update(*r.Item)
    return r, nil
}
