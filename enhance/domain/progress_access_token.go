package domainenhance
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

    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type ProgressAccessTokenDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    progressCache cache.ProgressDomainCache
    namespaceName string
    accessToken auth.AccessToken
}

func NewProgressAccessTokenDomain(
    session *core.Gs2RestSession,
    progressCache cache.ProgressDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
) ProgressAccessTokenDomain {
    return ProgressAccessTokenDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        progressCache: progressCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p ProgressAccessTokenDomain) Load(
    request enhance.GetProgressRequest,
) (*enhance.GetProgressResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.GetProgress(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.progressCache.Update(*r.Item)
    return r, nil
}