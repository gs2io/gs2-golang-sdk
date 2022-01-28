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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    globalMessageMasterCache cache.GlobalMessageMasterDomainCache
    globalMessageCache cache.GlobalMessageDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        globalMessageMasterCache: cache.NewGlobalMessageMasterDomainCache(),
        globalMessageCache: cache.NewGlobalMessageDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request inbox.GetNamespaceStatusRequest,
) (*inbox.GetNamespaceStatusResult, error) {
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
    request inbox.GetNamespaceRequest,
) (*inbox.GetNamespaceResult, error) {
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
    request inbox.UpdateNamespaceRequest,
) (*inbox.UpdateNamespaceResult, error) {
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
    request inbox.DeleteNamespaceRequest,
) (*inbox.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateGlobalMessageMaster(
    request inbox.CreateGlobalMessageMasterRequest,
) (*inbox.CreateGlobalMessageMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateGlobalMessageMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) GlobalMessageMasters(
) iterator.DescribeGlobalMessageMastersIterator {
    return iterator.NewDescribeGlobalMessageMastersIterator(
        p.globalMessageMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) GlobalMessages(
) iterator.DescribeGlobalMessagesIterator {
    return iterator.NewDescribeGlobalMessagesIterator(
        p.globalMessageCache,
        p.client,
        p.namespaceName,
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

func (p NamespaceDomain) CurrentMessageMaster(
) CurrentMessageMasterDomain {
    return NewCurrentMessageMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) GlobalMessage(
    globalMessageName string,
) GlobalMessageDomain {
    return NewGlobalMessageDomain(
        p.session,
        p.globalMessageCache,
        p.namespaceName,
        globalMessageName,
    )
}

func (p NamespaceDomain) GlobalMessageMaster(
    globalMessageName string,
) GlobalMessageMasterDomain {
    return NewGlobalMessageMasterDomain(
        p.session,
        p.globalMessageMasterCache,
        p.namespaceName,
        globalMessageName,
    )
}
