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

type CompletedQuestListDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    completedQuestListCache cache.CompletedQuestListDomainCache
    namespaceName string
    userId string
    questGroupName string
}

func NewCompletedQuestListDomain(
    session *core.Gs2RestSession,
    completedQuestListCache cache.CompletedQuestListDomainCache,
    namespaceName string,
    userId string,
    questGroupName string,
) CompletedQuestListDomain {
    return CompletedQuestListDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        completedQuestListCache: completedQuestListCache,
        namespaceName: namespaceName,
        userId: userId,
        questGroupName: questGroupName,
    }
}

func (p CompletedQuestListDomain) Load(
    request quest.GetCompletedQuestListByUserIdRequest,
) (*quest.GetCompletedQuestListByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.GetCompletedQuestListByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completedQuestListCache.Update(*r.Item)
    return r, nil
}

func (p CompletedQuestListDomain) Delete(
    request quest.DeleteCompletedQuestListByUserIdRequest,
) (*quest.DeleteCompletedQuestListByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.DeleteCompletedQuestListByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.completedQuestListCache.Delete(*r.Item)
    return r, nil
}
