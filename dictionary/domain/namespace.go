package domaindictionary
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

    "github.com/gs2io/gs2-golang-sdk/dictionary"
    "github.com/gs2io/gs2-golang-sdk/dictionary/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/dictionary/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    entryModelCache cache.EntryModelDomainCache
    entryModelMasterCache cache.EntryModelMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        entryModelCache: cache.NewEntryModelDomainCache(),
        entryModelMasterCache: cache.NewEntryModelMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request dictionary.GetNamespaceStatusRequest,
) (*dictionary.GetNamespaceStatusResult, error) {
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
    request dictionary.GetNamespaceRequest,
) (*dictionary.GetNamespaceResult, error) {
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
    request dictionary.UpdateNamespaceRequest,
) (*dictionary.UpdateNamespaceResult, error) {
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
    request dictionary.DeleteNamespaceRequest,
) (*dictionary.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateEntryModelMaster(
    request dictionary.CreateEntryModelMasterRequest,
) (*dictionary.CreateEntryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateEntryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) EntryModels(
) iterator.DescribeEntryModelsIterator {
    return iterator.NewDescribeEntryModelsIterator(
        p.entryModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) EntryModelMasters(
) iterator.DescribeEntryModelMastersIterator {
    return iterator.NewDescribeEntryModelMastersIterator(
        p.entryModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentEntryMaster(
) CurrentEntryMasterDomain {
    return NewCurrentEntryMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) EntryModel(
    entryName string,
) EntryModelDomain {
    return NewEntryModelDomain(
        p.session,
        p.entryModelCache,
        p.namespaceName,
        entryName,
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

func (p NamespaceDomain) EntryModelMaster(
    entryName string,
) EntryModelMasterDomain {
    return NewEntryModelMasterDomain(
        p.session,
        p.entryModelMasterCache,
        p.namespaceName,
        entryName,
    )
}
