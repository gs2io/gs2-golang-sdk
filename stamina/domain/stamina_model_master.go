package domainstamina
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

    "github.com/gs2io/gs2-golang-sdk/stamina"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type StaminaModelMasterDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    staminaModelMasterCache cache.StaminaModelMasterDomainCache
    namespaceName string
    staminaName string
}

func NewStaminaModelMasterDomain(
    session *core.Gs2RestSession,
    staminaModelMasterCache cache.StaminaModelMasterDomainCache,
    namespaceName string,
    staminaName string,
) StaminaModelMasterDomain {
    return StaminaModelMasterDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        staminaModelMasterCache: staminaModelMasterCache,
        namespaceName: namespaceName,
        staminaName: staminaName,
    }
}

func (p StaminaModelMasterDomain) Load(
    request stamina.GetStaminaModelMasterRequest,
) (*stamina.GetStaminaModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.StaminaName = &p.staminaName
    r, err := p.client.GetStaminaModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p StaminaModelMasterDomain) Update(
    request stamina.UpdateStaminaModelMasterRequest,
) (*stamina.UpdateStaminaModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.StaminaName = &p.staminaName
    r, err := p.client.UpdateStaminaModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p StaminaModelMasterDomain) Delete(
    request stamina.DeleteStaminaModelMasterRequest,
) (*stamina.DeleteStaminaModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.StaminaName = &p.staminaName
    r, err := p.client.DeleteStaminaModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaModelMasterCache.Delete(*r.Item)
    return r, nil
}
