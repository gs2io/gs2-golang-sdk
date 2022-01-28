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

type ThresholdMasterDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    thresholdMasterCache cache.ThresholdMasterDomainCache
    namespaceName string
    thresholdName string
}

func NewThresholdMasterDomain(
    session *core.Gs2RestSession,
    thresholdMasterCache cache.ThresholdMasterDomainCache,
    namespaceName string,
    thresholdName string,
) ThresholdMasterDomain {
    return ThresholdMasterDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        thresholdMasterCache: thresholdMasterCache,
        namespaceName: namespaceName,
        thresholdName: thresholdName,
    }
}

func (p ThresholdMasterDomain) Load(
    request experience.GetThresholdMasterRequest,
) (*experience.GetThresholdMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ThresholdName = &p.thresholdName
    r, err := p.client.GetThresholdMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.thresholdMasterCache.Update(*r.Item)
    return r, nil
}

func (p ThresholdMasterDomain) Update(
    request experience.UpdateThresholdMasterRequest,
) (*experience.UpdateThresholdMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ThresholdName = &p.thresholdName
    r, err := p.client.UpdateThresholdMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.thresholdMasterCache.Update(*r.Item)
    return r, nil
}

func (p ThresholdMasterDomain) Delete(
    request experience.DeleteThresholdMasterRequest,
) (*experience.DeleteThresholdMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ThresholdName = &p.thresholdName
    r, err := p.client.DeleteThresholdMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.thresholdMasterCache.Delete(*r.Item)
    return r, nil
}
