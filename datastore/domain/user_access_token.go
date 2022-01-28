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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client datastore.Gs2DatastoreRestClient
    namespaceName string
    accessToken auth.AccessToken
    dataObjectCache cache.DataObjectDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: datastore.Gs2DatastoreRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        dataObjectCache: cache.NewDataObjectDomainCache(),
    }
}

func (p UserAccessTokenDomain) PrepareUpload(
    request datastore.PrepareUploadRequest,
) (*datastore.PrepareUploadResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareUpload(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) PrepareReUpload(
    request datastore.PrepareReUploadRequest,
) (*datastore.PrepareReUploadResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareReUpload(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) PrepareDownload(
    request datastore.PrepareDownloadRequest,
) (*datastore.PrepareDownloadResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareDownload(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) PrepareDownloadByGeneration(
    request datastore.PrepareDownloadByGenerationRequest,
) (*datastore.PrepareDownloadByGenerationResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareDownloadByGeneration(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) PrepareDownloadOwnData(
    request datastore.PrepareDownloadOwnDataRequest,
) (*datastore.PrepareDownloadOwnDataResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareDownloadOwnData(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) PrepareDownloadOwnDataByGeneration(
    request datastore.PrepareDownloadOwnDataByGenerationRequest,
) (*datastore.PrepareDownloadOwnDataByGenerationResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.PrepareDownloadOwnDataByGeneration(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) DataObjects(
    status *string,
) iterator.DescribeDataObjectsIterator {
    return iterator.NewDescribeDataObjectsIterator(
        p.dataObjectCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        status,
    )
}

func (p UserAccessTokenDomain) DataObject(
    dataObjectName string,
) DataObjectAccessTokenDomain {
    return NewDataObjectAccessTokenDomain(
        p.session,
        p.dataObjectCache,
        p.namespaceName,
        p.accessToken,
        dataObjectName,
    )
}
