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

type FormModelMasterDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    formModelMasterCache cache.FormModelMasterDomainCache
    namespaceName string
    formModelName string
}

func NewFormModelMasterDomain(
    session *core.Gs2RestSession,
    formModelMasterCache cache.FormModelMasterDomainCache,
    namespaceName string,
    formModelName string,
) FormModelMasterDomain {
    return FormModelMasterDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        formModelMasterCache: formModelMasterCache,
        namespaceName: namespaceName,
        formModelName: formModelName,
    }
}

func (p FormModelMasterDomain) Load(
    request formation.GetFormModelMasterRequest,
) (*formation.GetFormModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.FormModelName = &p.formModelName
    r, err := p.client.GetFormModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p FormModelMasterDomain) Update(
    request formation.UpdateFormModelMasterRequest,
) (*formation.UpdateFormModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.FormModelName = &p.formModelName
    r, err := p.client.UpdateFormModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p FormModelMasterDomain) Delete(
    request formation.DeleteFormModelMasterRequest,
) (*formation.DeleteFormModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.FormModelName = &p.formModelName
    r, err := p.client.DeleteFormModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formModelMasterCache.Delete(*r.Item)
    return r, nil
}
