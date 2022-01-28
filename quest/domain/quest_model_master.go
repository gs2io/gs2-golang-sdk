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

type QuestModelMasterDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    questModelMasterCache cache.QuestModelMasterDomainCache
    namespaceName string
    questGroupName string
    questName string
}

func NewQuestModelMasterDomain(
    session *core.Gs2RestSession,
    questModelMasterCache cache.QuestModelMasterDomainCache,
    namespaceName string,
    questGroupName string,
    questName string,
) QuestModelMasterDomain {
    return QuestModelMasterDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        questModelMasterCache: questModelMasterCache,
        namespaceName: namespaceName,
        questGroupName: questGroupName,
        questName: questName,
    }
}

func (p QuestModelMasterDomain) Load(
    request quest.GetQuestModelMasterRequest,
) (*quest.GetQuestModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    request.QuestName = &p.questName
    r, err := p.client.GetQuestModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p QuestModelMasterDomain) Update(
    request quest.UpdateQuestModelMasterRequest,
) (*quest.UpdateQuestModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    request.QuestName = &p.questName
    r, err := p.client.UpdateQuestModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p QuestModelMasterDomain) Delete(
    request quest.DeleteQuestModelMasterRequest,
) (*quest.DeleteQuestModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    request.QuestName = &p.questName
    r, err := p.client.DeleteQuestModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questModelMasterCache.Delete(*r.Item)
    return r, nil
}
