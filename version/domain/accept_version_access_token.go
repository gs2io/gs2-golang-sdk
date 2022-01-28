package domainversion
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

    "github.com/gs2io/gs2-golang-sdk/version"
    "github.com/gs2io/gs2-golang-sdk/version/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/version/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type AcceptVersionAccessTokenDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    acceptVersionCache cache.AcceptVersionDomainCache
    namespaceName string
    accessToken auth.AccessToken
    versionName string
}

func NewAcceptVersionAccessTokenDomain(
    session *core.Gs2RestSession,
    acceptVersionCache cache.AcceptVersionDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    versionName string,
) AcceptVersionAccessTokenDomain {
    return AcceptVersionAccessTokenDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        acceptVersionCache: acceptVersionCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        versionName: versionName,
    }
}

func (p AcceptVersionAccessTokenDomain) Accept(
    request version.AcceptRequest,
) (*version.AcceptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.VersionName = &p.versionName
    r, err := p.client.Accept(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.acceptVersionCache.Update(*r.Item)
    return r, nil
}

func (p AcceptVersionAccessTokenDomain) Load(
    request version.GetAcceptVersionRequest,
) (*version.GetAcceptVersionResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.VersionName = &p.versionName
    r, err := p.client.GetAcceptVersion(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.acceptVersionCache.Update(*r.Item)
    return r, nil
}

func (p AcceptVersionAccessTokenDomain) Delete(
    request version.DeleteAcceptVersionRequest,
) (*version.DeleteAcceptVersionResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.VersionName = &p.versionName
    r, err := p.client.DeleteAcceptVersion(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
