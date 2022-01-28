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

type DataObjectDomain struct {
    session *core.Gs2RestSession
    client datastore.Gs2DatastoreRestClient
    dataObjectCache cache.DataObjectDomainCache
    namespaceName string
    userId string
    dataObjectName string
    dataObjectHistoryCache cache.DataObjectHistoryDomainCache
}

func NewDataObjectDomain(
    session *core.Gs2RestSession,
    dataObjectCache cache.DataObjectDomainCache,
    namespaceName string,
    userId string,
    dataObjectName string,
) DataObjectDomain {
    return DataObjectDomain {
        session: session,
        client: datastore.Gs2DatastoreRestClient{
            Session: session,
        },
        dataObjectCache: dataObjectCache,
        namespaceName: namespaceName,
        userId: userId,
        dataObjectName: dataObjectName,
        dataObjectHistoryCache: cache.NewDataObjectHistoryDomainCache(),
    }
}

func (p DataObjectDomain) Update(
    request datastore.UpdateDataObjectByUserIdRequest,
) (*datastore.UpdateDataObjectByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.UpdateDataObjectByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Update(*r.Item)
    return r, nil
}

func (p DataObjectDomain) DoneUpload(
    request datastore.DoneUploadByUserIdRequest,
) (*datastore.DoneUploadByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.DoneUploadByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Update(*r.Item)
    return r, nil
}

func (p DataObjectDomain) Delete(
    request datastore.DeleteDataObjectByUserIdRequest,
) (*datastore.DeleteDataObjectByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.DataObjectName = &p.dataObjectName
    r, err := p.client.DeleteDataObjectByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.dataObjectCache.Delete(*r.Item)
    return r, nil
}

func (p DataObjectDomain) DataObjectHistories(
) iterator.DescribeDataObjectHistoriesByUserIdIterator {
    return iterator.NewDescribeDataObjectHistoriesByUserIdIterator(
        p.dataObjectHistoryCache,
        p.client,
        p.namespaceName,
        p.userId,
        p.dataObjectName,
    )
}

func (p DataObjectDomain) DataObjectHistory(
    generation string,
) DataObjectHistoryDomain {
    return NewDataObjectHistoryDomain(
        p.session,
        p.dataObjectHistoryCache,
        p.namespaceName,
        p.userId,
        p.dataObjectName,
        generation,
    )
}
