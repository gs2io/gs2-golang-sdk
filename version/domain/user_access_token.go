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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    namespaceName string
    accessToken auth.AccessToken
    acceptVersionCache cache.AcceptVersionDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        acceptVersionCache: cache.NewAcceptVersionDomainCache(),
    }
}

func (p UserAccessTokenDomain) CheckVersion(
    request version.CheckVersionRequest,
) (*version.CheckVersionResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.CheckVersion(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) AcceptVersions(
) iterator.DescribeAcceptVersionsIterator {
    return iterator.NewDescribeAcceptVersionsIterator(
        p.acceptVersionCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) AcceptVersion(
    versionName string,
) AcceptVersionAccessTokenDomain {
    return NewAcceptVersionAccessTokenDomain(
        p.session,
        p.acceptVersionCache,
        p.namespaceName,
        p.accessToken,
        versionName,
    )
}
