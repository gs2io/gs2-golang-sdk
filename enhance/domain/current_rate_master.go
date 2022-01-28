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

type CurrentRateMasterDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    namespaceName string
}

func NewCurrentRateMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentRateMasterDomain {
    return CurrentRateMasterDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentRateMasterDomain) ExportMaster(
    request enhance.ExportMasterRequest,
) (*enhance.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRateMasterDomain) Load(
    request enhance.GetCurrentRateMasterRequest,
) (*enhance.GetCurrentRateMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentRateMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRateMasterDomain) Update(
    request enhance.UpdateCurrentRateMasterRequest,
) (*enhance.UpdateCurrentRateMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRateMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentRateMasterDomain) UpdateFromGitHub(
    request enhance.UpdateCurrentRateMasterFromGitHubRequest,
) (*enhance.UpdateCurrentRateMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentRateMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
