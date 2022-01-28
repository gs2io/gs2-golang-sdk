package domaininbox
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

    "github.com/gs2io/gs2-golang-sdk/inbox"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type GlobalMessageMasterDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    globalMessageMasterCache cache.GlobalMessageMasterDomainCache
    namespaceName string
    globalMessageName string
}

func NewGlobalMessageMasterDomain(
    session *core.Gs2RestSession,
    globalMessageMasterCache cache.GlobalMessageMasterDomainCache,
    namespaceName string,
    globalMessageName string,
) GlobalMessageMasterDomain {
    return GlobalMessageMasterDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        globalMessageMasterCache: globalMessageMasterCache,
        namespaceName: namespaceName,
        globalMessageName: globalMessageName,
    }
}

func (p GlobalMessageMasterDomain) Load(
    request inbox.GetGlobalMessageMasterRequest,
) (*inbox.GetGlobalMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.GlobalMessageName = &p.globalMessageName
    r, err := p.client.GetGlobalMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.globalMessageMasterCache.Update(*r.Item)
    return r, nil
}

func (p GlobalMessageMasterDomain) Update(
    request inbox.UpdateGlobalMessageMasterRequest,
) (*inbox.UpdateGlobalMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.GlobalMessageName = &p.globalMessageName
    r, err := p.client.UpdateGlobalMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.globalMessageMasterCache.Update(*r.Item)
    return r, nil
}

func (p GlobalMessageMasterDomain) Delete(
    request inbox.DeleteGlobalMessageMasterRequest,
) (*inbox.DeleteGlobalMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.GlobalMessageName = &p.globalMessageName
    r, err := p.client.DeleteGlobalMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
