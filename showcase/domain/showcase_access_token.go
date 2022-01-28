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

type ShowcaseAccessTokenDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    showcaseCache cache.ShowcaseDomainCache
    namespaceName string
    accessToken auth.AccessToken
    showcaseName string
}

func NewShowcaseAccessTokenDomain(
    session *core.Gs2RestSession,
    showcaseCache cache.ShowcaseDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    showcaseName string,
) ShowcaseAccessTokenDomain {
    return ShowcaseAccessTokenDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        showcaseCache: showcaseCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        showcaseName: showcaseName,
    }
}

func (p ShowcaseAccessTokenDomain) Load(
    request showcase.GetShowcaseRequest,
) (*showcase.GetShowcaseResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.ShowcaseName = &p.showcaseName
    r, err := p.client.GetShowcase(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.showcaseCache.Update(*r.Item)
    return r, nil
}

func (p ShowcaseAccessTokenDomain) Buy(
    request showcase.BuyRequest,
) (*showcase.BuyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.ShowcaseName = &p.showcaseName
    r, err := p.client.Buy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
