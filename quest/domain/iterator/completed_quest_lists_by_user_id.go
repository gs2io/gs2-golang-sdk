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


type DescribeCompletedQuestListsByUserIdIterator struct {
    completedQuestListCache cache.CompletedQuestListDomainCache
    client quest.Gs2QuestRestClient
    namespaceName string
    userId string

    index int64
    pageToken *string
    last bool
    result []quest.CompletedQuestList

    FetchSize *int32
}

func NewDescribeCompletedQuestListsByUserIdIterator(
    completedQuestListCache cache.CompletedQuestListDomainCache,
    client quest.Gs2QuestRestClient,
    namespaceName string,
    userId string,
) DescribeCompletedQuestListsByUserIdIterator {
    it := DescribeCompletedQuestListsByUserIdIterator{
        completedQuestListCache: completedQuestListCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]quest.CompletedQuestList, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCompletedQuestListsByUserIdIterator) load() error {
    r, err := p.client.DescribeCompletedQuestListsByUserId(
        &quest.DescribeCompletedQuestListsByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.completedQuestListCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeCompletedQuestListsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCompletedQuestListsByUserIdIterator) Next(

) (quest.CompletedQuestList, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.CompletedQuestList{}, err
        }
    }
    if len(p.result) == 0 {
        return quest.CompletedQuestList{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return quest.CompletedQuestList{}, err
        }
    }
    return ret, nil
}
