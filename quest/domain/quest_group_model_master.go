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

type QuestGroupModelMasterDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    questGroupModelMasterCache cache.QuestGroupModelMasterDomainCache
    namespaceName string
    questGroupName string
    questModelMasterCache cache.QuestModelMasterDomainCache
}

func NewQuestGroupModelMasterDomain(
    session *core.Gs2RestSession,
    questGroupModelMasterCache cache.QuestGroupModelMasterDomainCache,
    namespaceName string,
    questGroupName string,
) QuestGroupModelMasterDomain {
    return QuestGroupModelMasterDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        questGroupModelMasterCache: questGroupModelMasterCache,
        namespaceName: namespaceName,
        questGroupName: questGroupName,
        questModelMasterCache: cache.NewQuestModelMasterDomainCache(),
    }
}

func (p QuestGroupModelMasterDomain) Load(
    request quest.GetQuestGroupModelMasterRequest,
) (*quest.GetQuestGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.GetQuestGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questGroupModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p QuestGroupModelMasterDomain) Update(
    request quest.UpdateQuestGroupModelMasterRequest,
) (*quest.UpdateQuestGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.UpdateQuestGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questGroupModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p QuestGroupModelMasterDomain) Delete(
    request quest.DeleteQuestGroupModelMasterRequest,
) (*quest.DeleteQuestGroupModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.DeleteQuestGroupModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questGroupModelMasterCache.Delete(*r.Item)
    return r, nil
}

func (p QuestGroupModelMasterDomain) CreateQuestModelMaster(
    request quest.CreateQuestModelMasterRequest,
) (*quest.CreateQuestModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.CreateQuestModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p QuestGroupModelMasterDomain) QuestModelMasters(
) iterator.DescribeQuestModelMastersIterator {
    return iterator.NewDescribeQuestModelMastersIterator(
        p.questModelMasterCache,
        p.client,
        p.namespaceName,
        p.questGroupName,
    )
}

func (p QuestGroupModelMasterDomain) QuestModelMaster(
    questName string,
) QuestModelMasterDomain {
    return NewQuestModelMasterDomain(
        p.session,
        p.questModelMasterCache,
        p.namespaceName,
        p.questGroupName,
        questName,
    )
}
