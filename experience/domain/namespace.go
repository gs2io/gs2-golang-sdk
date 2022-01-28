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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    experienceModelMasterCache cache.ExperienceModelMasterDomainCache
    experienceModelCache cache.ExperienceModelDomainCache
    thresholdMasterCache cache.ThresholdMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        experienceModelMasterCache: cache.NewExperienceModelMasterDomainCache(),
        experienceModelCache: cache.NewExperienceModelDomainCache(),
        thresholdMasterCache: cache.NewThresholdMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request experience.GetNamespaceStatusRequest,
) (*experience.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request experience.GetNamespaceRequest,
) (*experience.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request experience.UpdateNamespaceRequest,
) (*experience.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request experience.DeleteNamespaceRequest,
) (*experience.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateExperienceModelMaster(
    request experience.CreateExperienceModelMasterRequest,
) (*experience.CreateExperienceModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateExperienceModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateThresholdMaster(
    request experience.CreateThresholdMasterRequest,
) (*experience.CreateThresholdMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateThresholdMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) ExperienceModelMasters(
) iterator.DescribeExperienceModelMastersIterator {
    return iterator.NewDescribeExperienceModelMastersIterator(
        p.experienceModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) ExperienceModels(
) iterator.DescribeExperienceModelsIterator {
    return iterator.NewDescribeExperienceModelsIterator(
        p.experienceModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) ThresholdMasters(
) iterator.DescribeThresholdMastersIterator {
    return iterator.NewDescribeThresholdMastersIterator(
        p.thresholdMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentExperienceMaster(
) CurrentExperienceMasterDomain {
    return NewCurrentExperienceMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) ExperienceModel(
    experienceName string,
) ExperienceModelDomain {
    return NewExperienceModelDomain(
        p.session,
        p.experienceModelCache,
        p.namespaceName,
        experienceName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) ThresholdMaster(
    thresholdName string,
) ThresholdMasterDomain {
    return NewThresholdMasterDomain(
        p.session,
        p.thresholdMasterCache,
        p.namespaceName,
        thresholdName,
    )
}

func (p NamespaceDomain) ExperienceModelMaster(
    experienceName string,
) ExperienceModelMasterDomain {
    return NewExperienceModelMasterDomain(
        p.session,
        p.experienceModelMasterCache,
        p.namespaceName,
        experienceName,
    )
}
