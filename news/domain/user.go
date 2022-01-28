package domainnews
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

    "github.com/gs2io/gs2-golang-sdk/news"
    "github.com/gs2io/gs2-golang-sdk/news/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/news/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client news.Gs2NewsRestClient
    namespaceName string
    userId string
    newsCache cache.NewsDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: news.Gs2NewsRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        newsCache: cache.NewNewsDomainCache(),
    }
}

func (p UserDomain) WantGrant(
    request news.WantGrantByUserIdRequest,
) (*news.WantGrantByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.WantGrantByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Newses(
) iterator.DescribeNewsByUserIdIterator {
    return iterator.NewDescribeNewsByUserIdIterator(
        p.newsCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}
