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

type RecoverIntervalTableMasterDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    recoverIntervalTableMasterCache cache.RecoverIntervalTableMasterDomainCache
    namespaceName string
    recoverIntervalTableName string
}

func NewRecoverIntervalTableMasterDomain(
    session *core.Gs2RestSession,
    recoverIntervalTableMasterCache cache.RecoverIntervalTableMasterDomainCache,
    namespaceName string,
    recoverIntervalTableName string,
) RecoverIntervalTableMasterDomain {
    return RecoverIntervalTableMasterDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        recoverIntervalTableMasterCache: recoverIntervalTableMasterCache,
        namespaceName: namespaceName,
        recoverIntervalTableName: recoverIntervalTableName,
    }
}

func (p RecoverIntervalTableMasterDomain) Load(
    request stamina.GetRecoverIntervalTableMasterRequest,
) (*stamina.GetRecoverIntervalTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverIntervalTableName = &p.recoverIntervalTableName
    r, err := p.client.GetRecoverIntervalTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverIntervalTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p RecoverIntervalTableMasterDomain) Update(
    request stamina.UpdateRecoverIntervalTableMasterRequest,
) (*stamina.UpdateRecoverIntervalTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverIntervalTableName = &p.recoverIntervalTableName
    r, err := p.client.UpdateRecoverIntervalTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverIntervalTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p RecoverIntervalTableMasterDomain) Delete(
    request stamina.DeleteRecoverIntervalTableMasterRequest,
) (*stamina.DeleteRecoverIntervalTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RecoverIntervalTableName = &p.recoverIntervalTableName
    r, err := p.client.DeleteRecoverIntervalTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.recoverIntervalTableMasterCache.Delete(*r.Item)
    return r, nil
}
