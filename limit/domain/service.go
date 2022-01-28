package domainlimit
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

    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Limit struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Limit(
    session *core.Gs2RestSession,
) Gs2Limit {
    return Gs2Limit {
        session: session,
        client: limit.Gs2LimitRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Limit) CreateNamespace(
    request limit.CreateNamespaceRequest,
) (*limit.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Limit) CountUpByStampTask(
    request limit.CountUpByStampTaskRequest,
) (*limit.CountUpByStampTaskResult, error) {
    r, err := p.client.CountUpByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Limit) DeleteByStampSheet(
    request limit.DeleteByStampSheetRequest,
) (*limit.DeleteByStampSheetResult, error) {
    r, err := p.client.DeleteByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Limit) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Limit) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
