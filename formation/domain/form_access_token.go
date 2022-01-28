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

type FormAccessTokenDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    formCache cache.FormDomainCache
    namespaceName string
    accessToken auth.AccessToken
    moldName string
    index int32
}

func NewFormAccessTokenDomain(
    session *core.Gs2RestSession,
    formCache cache.FormDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    moldName string,
    index int32,
) FormAccessTokenDomain {
    return FormAccessTokenDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        formCache: formCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        moldName: moldName,
        index: index,
    }
}

func (p FormAccessTokenDomain) Load(
    request formation.GetFormRequest,
) (*formation.GetFormResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.GetForm(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormAccessTokenDomain) GetWithSignature(
    request formation.GetFormWithSignatureRequest,
) (*formation.GetFormWithSignatureResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.GetFormWithSignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormAccessTokenDomain) SetWithSignature(
    request formation.SetFormWithSignatureRequest,
) (*formation.SetFormWithSignatureResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.SetFormWithSignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Update(*r.Item)
    return r, nil
}

func (p FormAccessTokenDomain) Delete(
    request formation.DeleteFormRequest,
) (*formation.DeleteFormResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    request.Index = &p.index
    r, err := p.client.DeleteForm(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.formCache.Delete(*r.Item)
    return r, nil
}
