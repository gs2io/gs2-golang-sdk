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

type UserDomain struct {
    session *core.Gs2RestSession
    client datastore.Gs2DatastoreRestClient
    namespaceName string
    userId string
    dataObjectCache cache.DataObjectDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: datastore.Gs2DatastoreRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        dataObjectCache: cache.NewDataObjectDomainCache(),
    }
}

func (p UserDomain) PrepareUpload(
    request datastore.PrepareUploadByUserIdRequest,
) (*datastore.PrepareUploadByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareUploadByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) PrepareReUpload(
    request datastore.PrepareReUploadByUserIdRequest,
) (*datastore.PrepareReUploadByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareReUploadByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) PrepareDownload(
    request datastore.PrepareDownloadByUserIdRequest,
) (*datastore.PrepareDownloadByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareDownloadByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) PrepareDownloadByGenerationAndId(
    request datastore.PrepareDownloadByGenerationAndUserIdRequest,
) (*datastore.PrepareDownloadByGenerationAndUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareDownloadByGenerationAndUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) PrepareDownloadByIdAndDataObjectName(
    request datastore.PrepareDownloadByUserIdAndDataObjectNameRequest,
) (*datastore.PrepareDownloadByUserIdAndDataObjectNameResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareDownloadByUserIdAndDataObjectName(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) PrepareDownloadByIdAndDataObjectNameAndGeneration(
    request datastore.PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest,
) (*datastore.PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PrepareDownloadByUserIdAndDataObjectNameAndGeneration(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) DataObjects(
    status *string,
) iterator.DescribeDataObjectsByUserIdIterator {
    return iterator.NewDescribeDataObjectsByUserIdIterator(
        p.dataObjectCache,
        p.client,
        p.namespaceName,
        p.userId,
        status,
    )
}

func (p UserDomain) DataObject(
    dataObjectName string,
) DataObjectDomain {
    return NewDataObjectDomain(
        p.session,
        p.dataObjectCache,
        p.namespaceName,
        p.userId,
        dataObjectName,
    )
}
