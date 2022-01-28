package domainranking
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

    "github.com/gs2io/gs2-golang-sdk/ranking"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type SubscribeDomain struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    subscribeUserCache cache.SubscribeUserDomainCache
    namespaceName string
    userId string
    categoryName string
}

func NewSubscribeDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
    categoryName string,
) SubscribeDomain {
    return SubscribeDomain {
        session: session,
        client: ranking.Gs2RankingRestClient{
            Session: session,
        },
        subscribeUserCache: cache.NewSubscribeUserDomainCache(),
        namespaceName: namespaceName,
        userId: userId,
        categoryName: categoryName,
    }
}

func (p SubscribeDomain) Subscribe(
    request ranking.SubscribeByUserIdRequest,
) (*ranking.SubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.CategoryName = &p.categoryName
    r, err := p.client.SubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p SubscribeDomain) Load(
    request ranking.GetSubscribeByUserIdRequest,
) (*ranking.GetSubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.CategoryName = &p.categoryName
    r, err := p.client.GetSubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p SubscribeDomain) Unsubscribe(
    request ranking.UnsubscribeByUserIdRequest,
) (*ranking.UnsubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.CategoryName = &p.categoryName
    r, err := p.client.UnsubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
