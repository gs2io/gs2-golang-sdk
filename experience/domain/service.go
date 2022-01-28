package domainexperience
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

    "github.com/gs2io/gs2-golang-sdk/experience"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Experience struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Experience(
    session *core.Gs2RestSession,
) Gs2Experience {
    return Gs2Experience {
        session: session,
        client: experience.Gs2ExperienceRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Experience) CreateNamespace(
    request experience.CreateNamespaceRequest,
) (*experience.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Experience) AddExperienceByStampSheet(
    request experience.AddExperienceByStampSheetRequest,
) (*experience.AddExperienceByStampSheetResult, error) {
    r, err := p.client.AddExperienceByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Experience) AddRankCapByStampSheet(
    request experience.AddRankCapByStampSheetRequest,
) (*experience.AddRankCapByStampSheetResult, error) {
    r, err := p.client.AddRankCapByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Experience) SetRankCapByStampSheet(
    request experience.SetRankCapByStampSheetRequest,
) (*experience.SetRankCapByStampSheetResult, error) {
    r, err := p.client.SetRankCapByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Experience) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Experience) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
