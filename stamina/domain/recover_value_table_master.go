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

type RecoverValueTableMasterDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    recoverValueTableMasterCache cache.RecoverValueTableMasterDomainCache
    namespaceName string
    recoverValueTableName string
}

func NewRecoverValueTableMasterDomain(
    session *core.Gs2RestSession,
    recoverValueTableMasterCache cache.RecoverValueTableMasterDomainCache,
    namespaceName string,
    recoverValueTableName string,
) RecoverValueTableMasterDomain {
    return RecoverValueTableMasterDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        recoverValueTableMasterCache: recoverValueTableMasterCache,
        namespaceName: namespaceName,
        recoverValueTableName: recoverValueTableName,
    }
}

func (p RecoverValueTableMasterDomain) Load(
    request stamina.GetRecoverValueTableMasterRequest,
) (*stamina.GetRecoverValueTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverValueTableName = &p.recoverValueTableName
    r, err := p.client.GetRecoverValueTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverValueTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p RecoverValueTableMasterDomain) Update(
    request stamina.UpdateRecoverValueTableMasterRequest,
) (*stamina.UpdateRecoverValueTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverValueTableName = &p.recoverValueTableName
    r, err := p.client.UpdateRecoverValueTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverValueTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p RecoverValueTableMasterDomain) Delete(
    request stamina.DeleteRecoverValueTableMasterRequest,
) (*stamina.DeleteRecoverValueTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverValueTableName = &p.recoverValueTableName
    r, err := p.client.DeleteRecoverValueTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverValueTableMasterCache.Delete(*r.Item)
    return r, nil
}
