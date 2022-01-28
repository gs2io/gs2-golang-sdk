package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/datastore"
    "github.com/gs2io/gs2-golang-sdk/datastore/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeDataObjectsIterator struct {
    dataObjectCache cache.DataObjectDomainCache
    client datastore.Gs2DatastoreRestClient
    namespaceName string
    accessToken auth.AccessToken
    status *string

    index int64
    pageToken *string
    last bool
    result []datastore.DataObject

    FetchSize *int32
}

func NewDescribeDataObjectsIterator(
    dataObjectCache cache.DataObjectDomainCache,
    client datastore.Gs2DatastoreRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
    status *string,
) DescribeDataObjectsIterator {
    it := DescribeDataObjectsIterator{
        dataObjectCache: dataObjectCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,
        status: status,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]datastore.DataObject, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeDataObjectsIterator) load() error {
    r, err := p.client.DescribeDataObjects(
        &datastore.DescribeDataObjectsRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
            Status: p.status,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.dataObjectCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeDataObjectsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeDataObjectsIterator) Next(

) (datastore.DataObject, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return datastore.DataObject{}, err
        }
    }
    if len(p.result) == 0 {
        return datastore.DataObject{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return datastore.DataObject{}, err
        }
    }
    return ret, nil
}
