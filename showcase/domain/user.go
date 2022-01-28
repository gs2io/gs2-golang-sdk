package domainshowcase
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

    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    namespaceName string
    userId string
    showcaseCache cache.ShowcaseDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        showcaseCache: cache.NewShowcaseDomainCache(),
    }
}

func (p UserDomain) Showcases(
) iterator.DescribeShowcasesByUserIdIterator {
    return iterator.NewDescribeShowcasesByUserIdIterator(
        p.showcaseCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Showcase(
    showcaseName string,
) ShowcaseDomain {
    return NewShowcaseDomain(
        p.session,
        p.showcaseCache,
        p.namespaceName,
        p.userId,
        showcaseName,
    )
}
