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

type UserDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    namespaceName string
    userId string
    acceptVersionCache cache.AcceptVersionDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        acceptVersionCache: cache.NewAcceptVersionDomainCache(),
    }
}

func (p UserDomain) CheckVersion(
    request version.CheckVersionByUserIdRequest,
) (*version.CheckVersionByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CheckVersionByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) AcceptVersions(
) iterator.DescribeAcceptVersionsByUserIdIterator {
    return iterator.NewDescribeAcceptVersionsByUserIdIterator(
        p.acceptVersionCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) AcceptVersion(
    versionName string,
) AcceptVersionDomain {
    return NewAcceptVersionDomain(
        p.session,
        p.acceptVersionCache,
        p.namespaceName,
        p.userId,
        versionName,
    )
}
