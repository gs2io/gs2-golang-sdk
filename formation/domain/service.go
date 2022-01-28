package domainformation
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

    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Formation struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Formation(
    session *core.Gs2RestSession,
) Gs2Formation {
    return Gs2Formation {
        session: session,
        client: formation.Gs2FormationRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Formation) CreateNamespace(
    request formation.CreateNamespaceRequest,
) (*formation.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Formation) AddCapacityByStampSheet(
    request formation.AddCapacityByStampSheetRequest,
) (*formation.AddCapacityByStampSheetResult, error) {
    r, err := p.client.AddCapacityByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Formation) SetCapacityByStampSheet(
    request formation.SetCapacityByStampSheetRequest,
) (*formation.SetCapacityByStampSheetResult, error) {
    r, err := p.client.SetCapacityByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Formation) AcquireActionToFormPropertiesByStampSheet(
    request formation.AcquireActionToFormPropertiesByStampSheetRequest,
) (*formation.AcquireActionToFormPropertiesByStampSheetResult, error) {
    r, err := p.client.AcquireActionToFormPropertiesByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Formation) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Formation) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
