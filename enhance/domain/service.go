package domainenhance
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

    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Enhance struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Enhance(
    session *core.Gs2RestSession,
) Gs2Enhance {
    return Gs2Enhance {
        session: session,
        client: enhance.Gs2EnhanceRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Enhance) CreateNamespace(
    request enhance.CreateNamespaceRequest,
) (*enhance.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Enhance) DirectEnhanceByStampSheet(
    request enhance.DirectEnhanceByStampSheetRequest,
) (*enhance.DirectEnhanceByStampSheetResult, error) {
    r, err := p.client.DirectEnhanceByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Enhance) CreateProgressByStampSheet(
    request enhance.CreateProgressByStampSheetRequest,
) (*enhance.CreateProgressByStampSheetResult, error) {
    r, err := p.client.CreateProgressByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Enhance) DeleteProgressByStampTask(
    request enhance.DeleteProgressByStampTaskRequest,
) (*enhance.DeleteProgressByStampTaskResult, error) {
    r, err := p.client.DeleteProgressByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Enhance) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Enhance) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
