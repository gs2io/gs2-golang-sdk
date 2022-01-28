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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    staminaModelMasterCache cache.StaminaModelMasterDomainCache
    maxStaminaTableMasterCache cache.MaxStaminaTableMasterDomainCache
    recoverIntervalTableMasterCache cache.RecoverIntervalTableMasterDomainCache
    recoverValueTableMasterCache cache.RecoverValueTableMasterDomainCache
    staminaModelCache cache.StaminaModelDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        staminaModelMasterCache: cache.NewStaminaModelMasterDomainCache(),
        maxStaminaTableMasterCache: cache.NewMaxStaminaTableMasterDomainCache(),
        recoverIntervalTableMasterCache: cache.NewRecoverIntervalTableMasterDomainCache(),
        recoverValueTableMasterCache: cache.NewRecoverValueTableMasterDomainCache(),
        staminaModelCache: cache.NewStaminaModelDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request stamina.GetNamespaceStatusRequest,
) (*stamina.GetNamespaceStatusResult, error) {
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
    request stamina.GetNamespaceRequest,
) (*stamina.GetNamespaceResult, error) {
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
    request stamina.UpdateNamespaceRequest,
) (*stamina.UpdateNamespaceResult, error) {
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
    request stamina.DeleteNamespaceRequest,
) (*stamina.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateStaminaModelMaster(
    request stamina.CreateStaminaModelMasterRequest,
) (*stamina.CreateStaminaModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateStaminaModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateMaxStaminaTableMaster(
    request stamina.CreateMaxStaminaTableMasterRequest,
) (*stamina.CreateMaxStaminaTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateMaxStaminaTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateRecoverIntervalTableMaster(
    request stamina.CreateRecoverIntervalTableMasterRequest,
) (*stamina.CreateRecoverIntervalTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateRecoverIntervalTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateRecoverValueTableMaster(
    request stamina.CreateRecoverValueTableMasterRequest,
) (*stamina.CreateRecoverValueTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateRecoverValueTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) StaminaModelMasters(
) iterator.DescribeStaminaModelMastersIterator {
    return iterator.NewDescribeStaminaModelMastersIterator(
        p.staminaModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MaxStaminaTableMasters(
) iterator.DescribeMaxStaminaTableMastersIterator {
    return iterator.NewDescribeMaxStaminaTableMastersIterator(
        p.maxStaminaTableMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RecoverIntervalTableMasters(
) iterator.DescribeRecoverIntervalTableMastersIterator {
    return iterator.NewDescribeRecoverIntervalTableMastersIterator(
        p.recoverIntervalTableMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) RecoverValueTableMasters(
) iterator.DescribeRecoverValueTableMastersIterator {
    return iterator.NewDescribeRecoverValueTableMastersIterator(
        p.recoverValueTableMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) StaminaModels(
) iterator.DescribeStaminaModelsIterator {
    return iterator.NewDescribeStaminaModelsIterator(
        p.staminaModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentStaminaMaster(
) CurrentStaminaMasterDomain {
    return NewCurrentStaminaMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) StaminaModel(
    staminaName string,
) StaminaModelDomain {
    return NewStaminaModelDomain(
        p.session,
        p.staminaModelCache,
        p.namespaceName,
        staminaName,
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

func (p NamespaceDomain) RecoverIntervalTableMaster(
    recoverIntervalTableName string,
) RecoverIntervalTableMasterDomain {
    return NewRecoverIntervalTableMasterDomain(
        p.session,
        p.recoverIntervalTableMasterCache,
        p.namespaceName,
        recoverIntervalTableName,
    )
}

func (p NamespaceDomain) MaxStaminaTableMaster(
    maxStaminaTableName string,
) MaxStaminaTableMasterDomain {
    return NewMaxStaminaTableMasterDomain(
        p.session,
        p.maxStaminaTableMasterCache,
        p.namespaceName,
        maxStaminaTableName,
    )
}

func (p NamespaceDomain) RecoverValueTableMaster(
    recoverValueTableName string,
) RecoverValueTableMasterDomain {
    return NewRecoverValueTableMasterDomain(
        p.session,
        p.recoverValueTableMasterCache,
        p.namespaceName,
        recoverValueTableName,
    )
}

func (p NamespaceDomain) StaminaModelMaster(
    staminaName string,
) StaminaModelMasterDomain {
    return NewStaminaModelMasterDomain(
        p.session,
        p.staminaModelMasterCache,
        p.namespaceName,
        staminaName,
    )
}
