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

type QuestModelDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    questModelCache cache.QuestModelDomainCache
    namespaceName string
    questGroupName string
    questName string
}

func NewQuestModelDomain(
    session *core.Gs2RestSession,
    questModelCache cache.QuestModelDomainCache,
    namespaceName string,
    questGroupName string,
    questName string,
) QuestModelDomain {
    return QuestModelDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        questModelCache: questModelCache,
        namespaceName: namespaceName,
        questGroupName: questGroupName,
        questName: questName,
    }
}

func (p QuestModelDomain) Load(
    request quest.GetQuestModelRequest,
) (*quest.GetQuestModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.QuestGroupName = &p.questGroupName
    request.QuestName = &p.questName
    r, err := p.client.GetQuestModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.questModelCache.Update(*r.Item)
    return r, nil
}
