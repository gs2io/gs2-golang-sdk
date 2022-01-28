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

type VersionModelMasterDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    versionModelMasterCache cache.VersionModelMasterDomainCache
    namespaceName string
    versionName string
}

func NewVersionModelMasterDomain(
    session *core.Gs2RestSession,
    versionModelMasterCache cache.VersionModelMasterDomainCache,
    namespaceName string,
    versionName string,
) VersionModelMasterDomain {
    return VersionModelMasterDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        versionModelMasterCache: versionModelMasterCache,
        namespaceName: namespaceName,
        versionName: versionName,
    }
}

func (p VersionModelMasterDomain) Load(
    request version.GetVersionModelMasterRequest,
) (*version.GetVersionModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.VersionName = &p.versionName
    r, err := p.client.GetVersionModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.versionModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p VersionModelMasterDomain) Update(
    request version.UpdateVersionModelMasterRequest,
) (*version.UpdateVersionModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.VersionName = &p.versionName
    r, err := p.client.UpdateVersionModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.versionModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p VersionModelMasterDomain) Delete(
    request version.DeleteVersionModelMasterRequest,
) (*version.DeleteVersionModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.VersionName = &p.versionName
    r, err := p.client.DeleteVersionModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.versionModelMasterCache.Delete(*r.Item)
    return r, nil
}
