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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    namespaceName string
    accessToken auth.AccessToken
    progressCache cache.ProgressDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        progressCache: cache.NewProgressDomainCache(),
    }
}

func (p UserAccessTokenDomain) DirectEnhance(
    request enhance.DirectEnhanceRequest,
) (*enhance.DirectEnhanceResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.DirectEnhance(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Start(
    request enhance.StartRequest,
) (*enhance.StartResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.Start(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) End(
    request enhance.EndRequest,
) (*enhance.EndResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.End(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) DeleteProgress(
    request enhance.DeleteProgressRequest,
) (*enhance.DeleteProgressResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.DeleteProgress(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Progress(
) ProgressAccessTokenDomain {
    return NewProgressAccessTokenDomain(
        p.session,
        p.progressCache,
        p.namespaceName,
        p.accessToken,
    )
}
