package domaindatastore
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

    "github.com/gs2io/gs2-golang-sdk/datastore"
    "github.com/gs2io/gs2-golang-sdk/datastore/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/datastore/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type DataObjectHistoryAccessTokenDomain struct {
    session *core.Gs2RestSession
    client datastore.Gs2DatastoreRestClient
    dataObjectHistoryCache cache.DataObjectHistoryDomainCache
    namespaceName string
    accessToken auth.AccessToken
    dataObjectName string
    generation string
}

func NewDataObjectHistoryAccessTokenDomain(
    session *core.Gs2RestSession,
    dataObjectHistoryCache cache.DataObjectHistoryDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    dataObjectName string,
    generation string,
) DataObjectHistoryAccessTokenDomain {
    return DataObjectHistoryAccessTokenDomain {
        session: session,
        client: datastore.Gs2DatastoreRestClient{
            Session: session,
        },
        dataObjectHistoryCache: dataObjectHistoryCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        dataObjectName: dataObjectName,
        generation: generation,
    }
}

func (p DataObjectHistoryAccessTokenDomain) Load(
    request datastore.GetDataObjectHistoryRequest,
) (*datastore.GetDataObjectHistoryResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.DataObjectName = &p.dataObjectName
    request.Generation = &p.generation
    r, err := p.client.GetDataObjectHistory(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectHistoryCache.Update(*r.Item)
    return r, nil
}
