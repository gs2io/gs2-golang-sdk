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

type EntryAccessTokenDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    entryCache cache.EntryDomainCache
    namespaceName string
    accessToken auth.AccessToken
    entryModelName string
}

func NewEntryAccessTokenDomain(
    session *core.Gs2RestSession,
    entryCache cache.EntryDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    entryModelName string,
) EntryAccessTokenDomain {
    return EntryAccessTokenDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        entryCache: entryCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        entryModelName: entryModelName,
    }
}

func (p EntryAccessTokenDomain) Load(
    request dictionary.GetEntryRequest,
) (*dictionary.GetEntryResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.EntryModelName = &p.entryModelName
    r, err := p.client.GetEntry(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryCache.Update(*r.Item)
    return r, nil
}

func (p EntryAccessTokenDomain) GetWithSignature(
    request dictionary.GetEntryWithSignatureRequest,
) (*dictionary.GetEntryWithSignatureResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.EntryModelName = &p.entryModelName
    r, err := p.client.GetEntryWithSignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.entryCache.Update(*r.Item)
    return r, nil
}
