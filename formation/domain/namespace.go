package domainformation
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

    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    formModelMasterCache cache.FormModelMasterDomainCache
    moldModelCache cache.MoldModelDomainCache
    moldModelMasterCache cache.MoldModelMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        formModelMasterCache: cache.NewFormModelMasterDomainCache(),
        moldModelCache: cache.NewMoldModelDomainCache(),
        moldModelMasterCache: cache.NewMoldModelMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request formation.GetNamespaceStatusRequest,
) (*formation.GetNamespaceStatusResult, error) {
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
    request formation.GetNamespaceRequest,
) (*formation.GetNamespaceResult, error) {
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
    request formation.UpdateNamespaceRequest,
) (*formation.UpdateNamespaceResult, error) {
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
    request formation.DeleteNamespaceRequest,
) (*formation.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateFormModelMaster(
    request formation.CreateFormModelMasterRequest,
) (*formation.CreateFormModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateFormModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateMoldModelMaster(
    request formation.CreateMoldModelMasterRequest,
) (*formation.CreateMoldModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateMoldModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) FormModelMasters(
) iterator.DescribeFormModelMastersIterator {
    return iterator.NewDescribeFormModelMastersIterator(
        p.formModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MoldModels(
) iterator.DescribeMoldModelsIterator {
    return iterator.NewDescribeMoldModelsIterator(
        p.moldModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MoldModelMasters(
) iterator.DescribeMoldModelMastersIterator {
    return iterator.NewDescribeMoldModelMastersIterator(
        p.moldModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentFormMaster(
) CurrentFormMasterDomain {
    return NewCurrentFormMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) MoldModel(
    moldName string,
) MoldModelDomain {
    return NewMoldModelDomain(
        p.session,
        p.moldModelCache,
        p.namespaceName,
        moldName,
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

func (p NamespaceDomain) FormModelMaster(
    formModelName string,
) FormModelMasterDomain {
    return NewFormModelMasterDomain(
        p.session,
        p.formModelMasterCache,
        p.namespaceName,
        formModelName,
    )
}

func (p NamespaceDomain) MoldModelMaster(
    moldName string,
) MoldModelMasterDomain {
    return NewMoldModelMasterDomain(
        p.session,
        p.moldModelMasterCache,
        p.namespaceName,
        moldName,
    )
}
