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

type AcceptVersionDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    acceptVersionCache cache.AcceptVersionDomainCache
    namespaceName string
    userId string
    versionName string
}

func NewAcceptVersionDomain(
    session *core.Gs2RestSession,
    acceptVersionCache cache.AcceptVersionDomainCache,
    namespaceName string,
    userId string,
    versionName string,
) AcceptVersionDomain {
    return AcceptVersionDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        acceptVersionCache: acceptVersionCache,
        namespaceName: namespaceName,
        userId: userId,
        versionName: versionName,
    }
}

func (p AcceptVersionDomain) Accept(
    request version.AcceptByUserIdRequest,
) (*version.AcceptByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.VersionName = &p.versionName
    r, err := p.client.AcceptByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.acceptVersionCache.Update(*r.Item)
    return r, nil
}

func (p AcceptVersionDomain) Load(
    request version.GetAcceptVersionByUserIdRequest,
) (*version.GetAcceptVersionByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.VersionName = &p.versionName
    r, err := p.client.GetAcceptVersionByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.acceptVersionCache.Update(*r.Item)
    return r, nil
}

func (p AcceptVersionDomain) Delete(
    request version.DeleteAcceptVersionByUserIdRequest,
) (*version.DeleteAcceptVersionByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.VersionName = &p.versionName
    r, err := p.client.DeleteAcceptVersionByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
