package domaininbox
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

    "github.com/gs2io/gs2-golang-sdk/inbox"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentMessageMasterDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    namespaceName string
}

func NewCurrentMessageMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentMessageMasterDomain {
    return CurrentMessageMasterDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentMessageMasterDomain) ExportMaster(
    request inbox.ExportMasterRequest,
) (*inbox.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMessageMasterDomain) Load(
    request inbox.GetCurrentMessageMasterRequest,
) (*inbox.GetCurrentMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMessageMasterDomain) Update(
    request inbox.UpdateCurrentMessageMasterRequest,
) (*inbox.UpdateCurrentMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentMessageMasterDomain) UpdateFromGitHub(
    request inbox.UpdateCurrentMessageMasterFromGitHubRequest,
) (*inbox.UpdateCurrentMessageMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentMessageMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
