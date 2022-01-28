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

type StatusDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    statusCache cache.StatusDomainCache
    namespaceName string
    userId string
    experienceName string
    propertyId string
}

func NewStatusDomain(
    session *core.Gs2RestSession,
    statusCache cache.StatusDomainCache,
    namespaceName string,
    userId string,
    experienceName string,
    propertyId string,
) StatusDomain {
    return StatusDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        statusCache: statusCache,
        namespaceName: namespaceName,
        userId: userId,
        experienceName: experienceName,
        propertyId: propertyId,
    }
}

func (p StatusDomain) Load(
    request experience.GetStatusByUserIdRequest,
) (*experience.GetStatusByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.GetStatusByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) GetWithSignature(
    request experience.GetStatusWithSignatureByUserIdRequest,
) (*experience.GetStatusWithSignatureByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.GetStatusWithSignatureByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) AddExperience(
    request experience.AddExperienceByUserIdRequest,
) (*experience.AddExperienceByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.AddExperienceByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) SetExperience(
    request experience.SetExperienceByUserIdRequest,
) (*experience.SetExperienceByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.SetExperienceByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) AddRankCap(
    request experience.AddRankCapByUserIdRequest,
) (*experience.AddRankCapByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.AddRankCapByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) SetRankCap(
    request experience.SetRankCapByUserIdRequest,
) (*experience.SetRankCapByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.SetRankCapByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusDomain) Delete(
    request experience.DeleteStatusByUserIdRequest,
) (*experience.DeleteStatusByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.DeleteStatusByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Delete(*r.Item)
    return r, nil
}
