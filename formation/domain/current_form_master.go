package domainformation
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

    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentFormMasterDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    namespaceName string
}

func NewCurrentFormMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentFormMasterDomain {
    return CurrentFormMasterDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentFormMasterDomain) ExportMaster(
    request formation.ExportMasterRequest,
) (*formation.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentFormMasterDomain) Load(
    request formation.GetCurrentFormMasterRequest,
) (*formation.GetCurrentFormMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentFormMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentFormMasterDomain) Update(
    request formation.UpdateCurrentFormMasterRequest,
) (*formation.UpdateCurrentFormMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentFormMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentFormMasterDomain) UpdateFromGitHub(
    request formation.UpdateCurrentFormMasterFromGitHubRequest,
) (*formation.UpdateCurrentFormMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentFormMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
