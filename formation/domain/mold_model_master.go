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

type MoldModelMasterDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    moldModelMasterCache cache.MoldModelMasterDomainCache
    namespaceName string
    moldName string
}

func NewMoldModelMasterDomain(
    session *core.Gs2RestSession,
    moldModelMasterCache cache.MoldModelMasterDomainCache,
    namespaceName string,
    moldName string,
) MoldModelMasterDomain {
    return MoldModelMasterDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        moldModelMasterCache: moldModelMasterCache,
        namespaceName: namespaceName,
        moldName: moldName,
    }
}

func (p MoldModelMasterDomain) Load(
    request formation.GetMoldModelMasterRequest,
) (*formation.GetMoldModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MoldName = &p.moldName
    r, err := p.client.GetMoldModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MoldModelMasterDomain) Update(
    request formation.UpdateMoldModelMasterRequest,
) (*formation.UpdateMoldModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MoldName = &p.moldName
    r, err := p.client.UpdateMoldModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p MoldModelMasterDomain) Delete(
    request formation.DeleteMoldModelMasterRequest,
) (*formation.DeleteMoldModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.MoldName = &p.moldName
    r, err := p.client.DeleteMoldModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldModelMasterCache.Delete(*r.Item)
    return r, nil
}
