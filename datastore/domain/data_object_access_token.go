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

type DataObjectAccessTokenDomain struct {
    session *core.Gs2RestSession
    client datastore.Gs2DatastoreRestClient
    dataObjectCache cache.DataObjectDomainCache
    namespaceName string
    accessToken auth.AccessToken
    dataObjectName string
    dataObjectHistoryCache cache.DataObjectHistoryDomainCache
}

func NewDataObjectAccessTokenDomain(
    session *core.Gs2RestSession,
    dataObjectCache cache.DataObjectDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    dataObjectName string,
) DataObjectAccessTokenDomain {
    return DataObjectAccessTokenDomain {
        session: session,
        client: datastore.Gs2DatastoreRestClient{
            Session: session,
        },
        dataObjectCache: dataObjectCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        dataObjectName: dataObjectName,
        dataObjectHistoryCache: cache.NewDataObjectHistoryDomainCache(),
    }
}

func (p DataObjectAccessTokenDomain) Update(
    request datastore.UpdateDataObjectRequest,
) (*datastore.UpdateDataObjectResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.UpdateDataObject(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Update(*r.Item)
    return r, nil
}

func (p DataObjectAccessTokenDomain) DoneUpload(
    request datastore.DoneUploadRequest,
) (*datastore.DoneUploadResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.DoneUpload(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Update(*r.Item)
    return r, nil
}

func (p DataObjectAccessTokenDomain) Delete(
    request datastore.DeleteDataObjectRequest,
) (*datastore.DeleteDataObjectResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.DeleteDataObject(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Delete(*r.Item)
    return r, nil
}

func (p DataObjectAccessTokenDomain) DataObjectHistories(
) iterator.DescribeDataObjectHistoriesIterator {
    return iterator.NewDescribeDataObjectHistoriesIterator(
        p.dataObjectHistoryCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        p.dataObjectName,
    )
}

func (p DataObjectAccessTokenDomain) DataObjectHistory(
    generation string,
) DataObjectHistoryAccessTokenDomain {
    return NewDataObjectHistoryAccessTokenDomain(
        p.session,
        p.dataObjectHistoryCache,
        p.namespaceName,
        p.accessToken,
        p.dataObjectName,
        generation,
    )
}
