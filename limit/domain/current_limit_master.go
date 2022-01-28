package domainlimit
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

    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentLimitMasterDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    namespaceName string
}

func NewCurrentLimitMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentLimitMasterDomain {
    return CurrentLimitMasterDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentLimitMasterDomain) ExportMaster(
    request limit.ExportMasterRequest,
) (*limit.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLimitMasterDomain) Load(
    request limit.GetCurrentLimitMasterRequest,
) (*limit.GetCurrentLimitMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentLimitMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLimitMasterDomain) Update(
    request limit.UpdateCurrentLimitMasterRequest,
) (*limit.UpdateCurrentLimitMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentLimitMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLimitMasterDomain) UpdateFromGitHub(
    request limit.UpdateCurrentLimitMasterFromGitHubRequest,
) (*limit.UpdateCurrentLimitMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentLimitMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
