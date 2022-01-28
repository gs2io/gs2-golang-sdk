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
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type StatusAccessTokenDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    statusCache cache.StatusDomainCache
    namespaceName string
    accessToken auth.AccessToken
    experienceName string
    propertyId string
}

func NewStatusAccessTokenDomain(
    session *core.Gs2RestSession,
    statusCache cache.StatusDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    experienceName string,
    propertyId string,
) StatusAccessTokenDomain {
    return StatusAccessTokenDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        statusCache: statusCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        experienceName: experienceName,
        propertyId: propertyId,
    }
}

func (p StatusAccessTokenDomain) Load(
    request experience.GetStatusRequest,
) (*experience.GetStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.GetStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}

func (p StatusAccessTokenDomain) GetWithSignature(
    request experience.GetStatusWithSignatureRequest,
) (*experience.GetStatusWithSignatureResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.ExperienceName = &p.experienceName
    request.PropertyId = &p.propertyId
    r, err := p.client.GetStatusWithSignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.statusCache.Update(*r.Item)
    return r, nil
}
