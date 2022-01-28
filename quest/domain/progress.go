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

type ProgressDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    progressCache cache.ProgressDomainCache
    namespaceName string
    userId string
}

func NewProgressDomain(
    session *core.Gs2RestSession,
    progressCache cache.ProgressDomainCache,
    namespaceName string,
    userId string,
) ProgressDomain {
    return ProgressDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        progressCache: progressCache,
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p ProgressDomain) Load(
    request quest.GetProgressByUserIdRequest,
) (*quest.GetProgressByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetProgressByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.progressCache.Update(*r.Item)
    return r, nil
}

func (p ProgressDomain) Start(
    request quest.StartByUserIdRequest,
) (*quest.StartByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.StartByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ProgressDomain) End(
    request quest.EndByUserIdRequest,
) (*quest.EndByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.EndByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.progressCache.Update(*r.Item)
    return r, nil
}

func (p ProgressDomain) Delete(
    request quest.DeleteProgressByUserIdRequest,
) (*quest.DeleteProgressByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteProgressByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.progressCache.Delete(*r.Item)
    return r, nil
}

func (p ProgressDomain) List(
) iterator.DescribeProgressesByUserIdIterator {
    return iterator.NewDescribeProgressesByUserIdIterator(
        p.progressCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}
