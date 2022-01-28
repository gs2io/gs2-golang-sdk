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

type ExperienceModelMasterDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    experienceModelMasterCache cache.ExperienceModelMasterDomainCache
    namespaceName string
    experienceName string
}

func NewExperienceModelMasterDomain(
    session *core.Gs2RestSession,
    experienceModelMasterCache cache.ExperienceModelMasterDomainCache,
    namespaceName string,
    experienceName string,
) ExperienceModelMasterDomain {
    return ExperienceModelMasterDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        experienceModelMasterCache: experienceModelMasterCache,
        namespaceName: namespaceName,
        experienceName: experienceName,
    }
}

func (p ExperienceModelMasterDomain) Load(
    request experience.GetExperienceModelMasterRequest,
) (*experience.GetExperienceModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ExperienceName = &p.experienceName
    r, err := p.client.GetExperienceModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.experienceModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p ExperienceModelMasterDomain) Update(
    request experience.UpdateExperienceModelMasterRequest,
) (*experience.UpdateExperienceModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ExperienceName = &p.experienceName
    r, err := p.client.UpdateExperienceModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.experienceModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p ExperienceModelMasterDomain) Delete(
    request experience.DeleteExperienceModelMasterRequest,
) (*experience.DeleteExperienceModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ExperienceName = &p.experienceName
    r, err := p.client.DeleteExperienceModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.experienceModelMasterCache.Delete(*r.Item)
    return r, nil
}
