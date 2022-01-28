package domainexperience
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

    "github.com/gs2io/gs2-golang-sdk/experience"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentExperienceMasterDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    namespaceName string
}

func NewCurrentExperienceMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentExperienceMasterDomain {
    return CurrentExperienceMasterDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentExperienceMasterDomain) ExportMaster(
    request experience.ExportMasterRequest,
) (*experience.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentExperienceMasterDomain) Load(
    request experience.GetCurrentExperienceMasterRequest,
) (*experience.GetCurrentExperienceMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentExperienceMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentExperienceMasterDomain) Update(
    request experience.UpdateCurrentExperienceMasterRequest,
) (*experience.UpdateCurrentExperienceMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentExperienceMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentExperienceMasterDomain) UpdateFromGitHub(
    request experience.UpdateCurrentExperienceMasterFromGitHubRequest,
) (*experience.UpdateCurrentExperienceMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentExperienceMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
