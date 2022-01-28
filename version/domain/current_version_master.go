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

type CurrentVersionMasterDomain struct {
    session *core.Gs2RestSession
    client version.Gs2VersionRestClient
    namespaceName string
}

func NewCurrentVersionMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentVersionMasterDomain {
    return CurrentVersionMasterDomain {
        session: session,
        client: version.Gs2VersionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentVersionMasterDomain) ExportMaster(
    request version.ExportMasterRequest,
) (*version.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentVersionMasterDomain) Load(
    request version.GetCurrentVersionMasterRequest,
) (*version.GetCurrentVersionMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentVersionMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentVersionMasterDomain) Update(
    request version.UpdateCurrentVersionMasterRequest,
) (*version.UpdateCurrentVersionMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentVersionMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentVersionMasterDomain) UpdateFromGitHub(
    request version.UpdateCurrentVersionMasterFromGitHubRequest,
) (*version.UpdateCurrentVersionMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentVersionMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
