package domaindistributor
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

    "github.com/gs2io/gs2-golang-sdk/distributor"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentDistributorMasterDomain struct {
    session *core.Gs2RestSession
    client distributor.Gs2DistributorRestClient
    namespaceName string
}

func NewCurrentDistributorMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentDistributorMasterDomain {
    return CurrentDistributorMasterDomain {
        session: session,
        client: distributor.Gs2DistributorRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentDistributorMasterDomain) ExportMaster(
    request distributor.ExportMasterRequest,
) (*distributor.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentDistributorMasterDomain) Load(
    request distributor.GetCurrentDistributorMasterRequest,
) (*distributor.GetCurrentDistributorMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentDistributorMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentDistributorMasterDomain) Update(
    request distributor.UpdateCurrentDistributorMasterRequest,
) (*distributor.UpdateCurrentDistributorMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentDistributorMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentDistributorMasterDomain) UpdateFromGitHub(
    request distributor.UpdateCurrentDistributorMasterFromGitHubRequest,
) (*distributor.UpdateCurrentDistributorMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentDistributorMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
