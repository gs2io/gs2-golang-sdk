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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    namespaceName string
    accessToken auth.AccessToken
    entryCache cache.EntryDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        entryCache: cache.NewEntryDomainCache(),
    }
}

func (p UserAccessTokenDomain) Entries(
) iterator.DescribeEntriesIterator {
    return iterator.NewDescribeEntriesIterator(
        p.entryCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Entry(
    entryModelName string,
) EntryAccessTokenDomain {
    return NewEntryAccessTokenDomain(
        p.session,
        p.entryCache,
        p.namespaceName,
        p.accessToken,
        entryModelName,
    )
}
