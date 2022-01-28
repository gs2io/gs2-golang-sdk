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

type QuestGroupModelDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    questGroupModelCache cache.QuestGroupModelDomainCache
    namespaceName string
    questGroupName string
    questModelCache cache.QuestModelDomainCache
}

func NewQuestGroupModelDomain(
    session *core.Gs2RestSession,
    questGroupModelCache cache.QuestGroupModelDomainCache,
    namespaceName string,
    questGroupName string,
) QuestGroupModelDomain {
    return QuestGroupModelDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        questGroupModelCache: questGroupModelCache,
        namespaceName: namespaceName,
        questGroupName: questGroupName,
        questModelCache: cache.NewQuestModelDomainCache(),
    }
}

func (p QuestGroupModelDomain) Load(
    request quest.GetQuestGroupModelRequest,
) (*quest.GetQuestGroupModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    r, err := p.client.GetQuestGroupModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questGroupModelCache.Update(*r.Item)
    return r, nil
}

func (p QuestGroupModelDomain) QuestModels(
) iterator.DescribeQuestModelsIterator {
    return iterator.NewDescribeQuestModelsIterator(
        p.questModelCache,
        p.client,
        p.namespaceName,
        p.questGroupName,
    )
}

func (p QuestGroupModelDomain) QuestModel(
    questName string,
) QuestModelDomain {
    return NewQuestModelDomain(
        p.session,
        p.questModelCache,
        p.namespaceName,
        p.questGroupName,
        questName,
    )
}
