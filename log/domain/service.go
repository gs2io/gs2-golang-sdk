package domainlog
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

    "github.com/gs2io/gs2-golang-sdk/log"
    "github.com/gs2io/gs2-golang-sdk/log/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/log/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Log struct {
    session *core.Gs2RestSession
    client log.Gs2LogRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Log(
    session *core.Gs2RestSession,
) Gs2Log {
    return Gs2Log {
        session: session,
        client: log.Gs2LogRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Log) CreateNamespace(
    request log.CreateNamespaceRequest,
) (*log.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Log) PutLog(
    request log.PutLogRequest,
) (*log.PutLogResult, error) {
    r, err := p.client.PutLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Log) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Log) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
