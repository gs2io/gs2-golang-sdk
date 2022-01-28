package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/quest"
    "github.com/gs2io/gs2-golang-sdk/quest/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeQuestModelsIterator struct {
    questModelCache cache.QuestModelDomainCache
    client quest.Gs2QuestRestClient
    namespaceName string
    questGroupName string

    index int64
    last bool
    result []quest.QuestModel

    FetchSize *int32
}

func NewDescribeQuestModelsIterator(
    questModelCache cache.QuestModelDomainCache,
    client quest.Gs2QuestRestClient,
    namespaceName string,
    questGroupName string,
) DescribeQuestModelsIterator {
    it := DescribeQuestModelsIterator{
        questModelCache: questModelCache,
        client: client,
        namespaceName: namespaceName,
        questGroupName: questGroupName,

        index: 0,
        last: false,
        result: make([]quest.QuestModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeQuestModelsIterator) load() error {
    r, err := p.client.DescribeQuestModels(
        &quest.DescribeQuestModelsRequest{
            NamespaceName: &p.namespaceName,
            QuestGroupName: &p.questGroupName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.questModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeQuestModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeQuestModelsIterator) Next(

) (quest.QuestModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.QuestModel{}, err
        }
    }
    if len(p.result) == 0 {
        return quest.QuestModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.QuestModel{}, err
        }
    }
    return ret, nil
}
