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


type DescribeQuestGroupModelsIterator struct {
    questGroupModelCache cache.QuestGroupModelDomainCache
    client quest.Gs2QuestRestClient
    namespaceName string

    index int64
    last bool
    result []quest.QuestGroupModel

    FetchSize *int32
}

func NewDescribeQuestGroupModelsIterator(
    questGroupModelCache cache.QuestGroupModelDomainCache,
    client quest.Gs2QuestRestClient,
    namespaceName string,
) DescribeQuestGroupModelsIterator {
    it := DescribeQuestGroupModelsIterator{
        questGroupModelCache: questGroupModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]quest.QuestGroupModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeQuestGroupModelsIterator) load() error {
    r, err := p.client.DescribeQuestGroupModels(
        &quest.DescribeQuestGroupModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.questGroupModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeQuestGroupModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeQuestGroupModelsIterator) Next(

) (quest.QuestGroupModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.QuestGroupModel{}, err
        }
    }
    if len(p.result) == 0 {
        return quest.QuestGroupModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.QuestGroupModel{}, err
        }
    }
    return ret, nil
}
