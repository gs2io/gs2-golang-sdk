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

type FormDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    formCache cache.FormDomainCache
    namespaceName string
    userId string
    moldName string
    index int32
}

func NewFormDomain(
    session *core.Gs2RestSession,
    formCache cache.FormDomainCache,
    namespaceName string,
    userId string,
    moldName string,
    index int32,
) FormDomain {
    return FormDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        formCache: formCache,
        namespaceName: namespaceName,
        userId: userId,
        moldName: moldName,
        index: index,
    }
}

func (p FormDomain) Load(
    request formation.GetFormByUserIdRequest,
) (*formation.GetFormByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.GetFormByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormDomain) GetWithSignature(
    request formation.GetFormWithSignatureByUserIdRequest,
) (*formation.GetFormWithSignatureByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.GetFormWithSignatureByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormDomain) Set(
    request formation.SetFormByUserIdRequest,
) (*formation.SetFormByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.SetFormByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormDomain) AcquireActionsToProperties(
    request formation.AcquireActionsToFormPropertiesRequest,
) (*formation.AcquireActionsToFormPropertiesResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.AcquireActionsToFormProperties(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormDomain) Delete(
    request formation.DeleteFormByUserIdRequest,
) (*formation.DeleteFormByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.DeleteFormByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Delete(*r.Item)
    return r, nil
}
