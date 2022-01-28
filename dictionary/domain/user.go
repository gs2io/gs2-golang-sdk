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

type UserDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    namespaceName string
    userId string
    entryCache cache.EntryDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        entryCache: cache.NewEntryDomainCache(),
    }
}

func (p UserDomain) AddEntries(
    request dictionary.AddEntriesByUserIdRequest,
) (*dictionary.AddEntriesByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.AddEntriesByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Reset(
    request dictionary.ResetByUserIdRequest,
) (*dictionary.ResetByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.ResetByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Entries(
) iterator.DescribeEntriesByUserIdIterator {
    return iterator.NewDescribeEntriesByUserIdIterator(
        p.entryCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Entry(
    entryModelName string,
) EntryDomain {
    return NewEntryDomain(
        p.session,
        p.entryCache,
        p.namespaceName,
        p.userId,
        entryModelName,
    )
}
