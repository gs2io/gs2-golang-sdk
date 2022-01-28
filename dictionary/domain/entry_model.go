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

type EntryModelDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    entryModelCache cache.EntryModelDomainCache
    namespaceName string
    entryName string
}

func NewEntryModelDomain(
    session *core.Gs2RestSession,
    entryModelCache cache.EntryModelDomainCache,
    namespaceName string,
    entryName string,
) EntryModelDomain {
    return EntryModelDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        entryModelCache: entryModelCache,
        namespaceName: namespaceName,
        entryName: entryName,
    }
}

func (p EntryModelDomain) Load(
    request dictionary.GetEntryModelRequest,
) (*dictionary.GetEntryModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EntryName = &p.entryName
    r, err := p.client.GetEntryModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryModelCache.Update(*r.Item)
    return r, nil
}
