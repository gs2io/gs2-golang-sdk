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

type MaxStaminaTableMasterDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    maxStaminaTableMasterCache cache.MaxStaminaTableMasterDomainCache
    namespaceName string
    maxStaminaTableName string
}

func NewMaxStaminaTableMasterDomain(
    session *core.Gs2RestSession,
    maxStaminaTableMasterCache cache.MaxStaminaTableMasterDomainCache,
    namespaceName string,
    maxStaminaTableName string,
) MaxStaminaTableMasterDomain {
    return MaxStaminaTableMasterDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        maxStaminaTableMasterCache: maxStaminaTableMasterCache,
        namespaceName: namespaceName,
        maxStaminaTableName: maxStaminaTableName,
    }
}

func (p MaxStaminaTableMasterDomain) Load(
    request stamina.GetMaxStaminaTableMasterRequest,
) (*stamina.GetMaxStaminaTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MaxStaminaTableName = &p.maxStaminaTableName
    r, err := p.client.GetMaxStaminaTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.maxStaminaTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p MaxStaminaTableMasterDomain) Update(
    request stamina.UpdateMaxStaminaTableMasterRequest,
) (*stamina.UpdateMaxStaminaTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MaxStaminaTableName = &p.maxStaminaTableName
    r, err := p.client.UpdateMaxStaminaTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.maxStaminaTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p MaxStaminaTableMasterDomain) Delete(
    request stamina.DeleteMaxStaminaTableMasterRequest,
) (*stamina.DeleteMaxStaminaTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MaxStaminaTableName = &p.maxStaminaTableName
    r, err := p.client.DeleteMaxStaminaTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.maxStaminaTableMasterCache.Delete(*r.Item)
    return r, nil
}
