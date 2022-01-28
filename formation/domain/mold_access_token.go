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

type MoldAccessTokenDomain struct {
    session *core.Gs2RestSession
    client formation.Gs2FormationRestClient
    moldCache cache.MoldDomainCache
    namespaceName string
    accessToken auth.AccessToken
    moldName string
    formCache cache.FormDomainCache
}

func NewMoldAccessTokenDomain(
    session *core.Gs2RestSession,
    moldCache cache.MoldDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    moldName string,
) MoldAccessTokenDomain {
    return MoldAccessTokenDomain {
        session: session,
        client: formation.Gs2FormationRestClient{
            Session: session,
        },
        moldCache: moldCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        moldName: moldName,
        formCache: cache.NewFormDomainCache(),
    }
}

func (p MoldAccessTokenDomain) Load(
    request formation.GetMoldRequest,
) (*formation.GetMoldResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    r, err := p.client.GetMold(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Update(*r.Item)
    return r, nil
}

func (p MoldAccessTokenDomain) Delete(
    request formation.DeleteMoldRequest,
) (*formation.DeleteMoldResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MoldName = &p.moldName
    r, err := p.client.DeleteMold(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.moldCache.Delete(*r.Item)
    return r, nil
}

func (p MoldAccessTokenDomain) Forms(
) iterator.DescribeFormsIterator {
    return iterator.NewDescribeFormsIterator(
        p.formCache,
        p.client,
        p.namespaceName,
        p.moldName,
        p.accessToken,
    )
}

func (p MoldAccessTokenDomain) Form(
    index int32,
) FormAccessTokenDomain {
    return NewFormAccessTokenDomain(
        p.session,
        p.formCache,
        p.namespaceName,
        p.accessToken,
        p.moldName,
        index,
    )
}
