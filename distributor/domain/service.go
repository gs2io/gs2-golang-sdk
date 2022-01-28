package domaindistributor
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

    "github.com/gs2io/gs2-golang-sdk/distributor"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/distributor/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Distributor struct {
    session *core.Gs2RestSession
    client distributor.Gs2DistributorRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Distributor(
    session *core.Gs2RestSession,
) Gs2Distributor {
    return Gs2Distributor {
        session: session,
        client: distributor.Gs2DistributorRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Distributor) CreateNamespace(
    request distributor.CreateNamespaceRequest,
) (*distributor.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Distributor) DistributeWithoutOverflowProcess(
    request distributor.DistributeWithoutOverflowProcessRequest,
) (*distributor.DistributeWithoutOverflowProcessResult, error) {
    r, err := p.client.DistributeWithoutOverflowProcess(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Distributor) RunStampTaskWithoutNamespace(
    request distributor.RunStampTaskWithoutNamespaceRequest,
) (*distributor.RunStampTaskWithoutNamespaceResult, error) {
    r, err := p.client.RunStampTaskWithoutNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Distributor) RunStampSheetWithoutNamespace(
    request distributor.RunStampSheetWithoutNamespaceRequest,
) (*distributor.RunStampSheetWithoutNamespaceResult, error) {
    r, err := p.client.RunStampSheetWithoutNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Distributor) RunStampSheetExpressWithoutNamespace(
    request distributor.RunStampSheetExpressWithoutNamespaceRequest,
) (*distributor.RunStampSheetExpressWithoutNamespaceResult, error) {
    r, err := p.client.RunStampSheetExpressWithoutNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Distributor) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Distributor) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
