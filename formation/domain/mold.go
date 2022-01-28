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
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type MoldDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    moldCache cache.MoldDomainCache
    namespaceName string
    userId string
    moldName string
    formCache cache.FormDomainCache
}

func NewMoldDomain(
    session *core.Gs2RestSession,
    moldCache cache.MoldDomainCache,
    namespaceName string,
    userId string,
    moldName string,
) MoldDomain {
    return MoldDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        moldCache: moldCache,
        namespaceName: namespaceName,
        userId: userId,
        moldName: moldName,
        formCache: cache.NewFormDomainCache(),
    }
}

func (p MoldDomain) Load(
    request formation.GetMoldByUserIdRequest,
) (*formation.GetMoldByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    r, err := p.client.GetMoldByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Update(*r.Item)
    return r, nil
}

func (p MoldDomain) SetCapacity(
    request formation.SetMoldCapacityByUserIdRequest,
) (*formation.SetMoldCapacityByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    r, err := p.client.SetMoldCapacityByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Update(*r.Item)
    return r, nil
}

func (p MoldDomain) AddCapacity(
    request formation.AddMoldCapacityByUserIdRequest,
) (*formation.AddMoldCapacityByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    r, err := p.client.AddMoldCapacityByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Update(*r.Item)
    return r, nil
}

func (p MoldDomain) Delete(
    request formation.DeleteMoldByUserIdRequest,
) (*formation.DeleteMoldByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    r, err := p.client.DeleteMoldByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Delete(*r.Item)
    return r, nil
}

func (p MoldDomain) Forms(
) iterator.DescribeFormsByUserIdIterator {
    return iterator.NewDescribeFormsByUserIdIterator(
        p.formCache,
        p.client,
        p.namespaceName,
        p.moldName,
        p.userId,
    )
}

func (p MoldDomain) Form(
    index int32,
) FormDomain {
    return NewFormDomain(
        p.session,
        p.formCache,
        p.namespaceName,
        p.userId,
        p.moldName,
        index,
    )
}
