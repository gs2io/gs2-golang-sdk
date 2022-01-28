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

type EntryModelMasterDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    entryModelMasterCache cache.EntryModelMasterDomainCache
    namespaceName string
    entryName string
}

func NewEntryModelMasterDomain(
    session *core.Gs2RestSession,
    entryModelMasterCache cache.EntryModelMasterDomainCache,
    namespaceName string,
    entryName string,
) EntryModelMasterDomain {
    return EntryModelMasterDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        entryModelMasterCache: entryModelMasterCache,
        namespaceName: namespaceName,
        entryName: entryName,
    }
}

func (p EntryModelMasterDomain) Load(
    request dictionary.GetEntryModelMasterRequest,
) (*dictionary.GetEntryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EntryName = &p.entryName
    r, err := p.client.GetEntryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p EntryModelMasterDomain) Update(
    request dictionary.UpdateEntryModelMasterRequest,
) (*dictionary.UpdateEntryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EntryName = &p.entryName
    r, err := p.client.UpdateEntryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p EntryModelMasterDomain) Delete(
    request dictionary.DeleteEntryModelMasterRequest,
) (*dictionary.DeleteEntryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EntryName = &p.entryName
    r, err := p.client.DeleteEntryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryModelMasterCache.Delete(*r.Item)
    return r, nil
}
