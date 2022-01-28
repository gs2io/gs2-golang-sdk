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

type EntryDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    entryCache cache.EntryDomainCache
    namespaceName string
    userId string
    entryModelName string
}

func NewEntryDomain(
    session *core.Gs2RestSession,
    entryCache cache.EntryDomainCache,
    namespaceName string,
    userId string,
    entryModelName string,
) EntryDomain {
    return EntryDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        entryCache: entryCache,
        namespaceName: namespaceName,
        userId: userId,
        entryModelName: entryModelName,
    }
}

func (p EntryDomain) Load(
    request dictionary.GetEntryByUserIdRequest,
) (*dictionary.GetEntryByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.EntryModelName = &p.entryModelName
    r, err := p.client.GetEntryByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryCache.Update(*r.Item)
    return r, nil
}

func (p EntryDomain) GetWithSignature(
    request dictionary.GetEntryWithSignatureByUserIdRequest,
) (*dictionary.GetEntryWithSignatureByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.EntryModelName = &p.entryModelName
    r, err := p.client.GetEntryWithSignatureByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryCache.Update(*r.Item)
    return r, nil
}
