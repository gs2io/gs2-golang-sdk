package domainshowcase
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

    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentShowcaseMasterDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    namespaceName string
}

func NewCurrentShowcaseMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentShowcaseMasterDomain {
    return CurrentShowcaseMasterDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentShowcaseMasterDomain) ExportMaster(
    request showcase.ExportMasterRequest,
) (*showcase.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentShowcaseMasterDomain) Load(
    request showcase.GetCurrentShowcaseMasterRequest,
) (*showcase.GetCurrentShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentShowcaseMasterDomain) Update(
    request showcase.UpdateCurrentShowcaseMasterRequest,
) (*showcase.UpdateCurrentShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentShowcaseMasterDomain) UpdateFromGitHub(
    request showcase.UpdateCurrentShowcaseMasterFromGitHubRequest,
) (*showcase.UpdateCurrentShowcaseMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentShowcaseMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
