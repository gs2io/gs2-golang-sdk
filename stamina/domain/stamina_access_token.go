package domainstamina
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

    "github.com/gs2io/gs2-golang-sdk/stamina"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/stamina/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type StaminaAccessTokenDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    staminaCache cache.StaminaDomainCache
    namespaceName string
    accessToken auth.AccessToken
    staminaName string
}

func NewStaminaAccessTokenDomain(
    session *core.Gs2RestSession,
    staminaCache cache.StaminaDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    staminaName string,
) StaminaAccessTokenDomain {
    return StaminaAccessTokenDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        staminaCache: staminaCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        staminaName: staminaName,
    }
}

func (p StaminaAccessTokenDomain) Load(
    request stamina.GetStaminaRequest,
) (*stamina.GetStaminaResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.StaminaName = &p.staminaName
    r, err := p.client.GetStamina(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaAccessTokenDomain) Consume(
    request stamina.ConsumeStaminaRequest,
) (*stamina.ConsumeStaminaResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.StaminaName = &p.staminaName
    r, err := p.client.ConsumeStamina(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaAccessTokenDomain) SetMaxValueByStatus(
    request stamina.SetMaxValueByStatusRequest,
) (*stamina.SetMaxValueByStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.StaminaName = &p.staminaName
    r, err := p.client.SetMaxValueByStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaAccessTokenDomain) SetRecoverIntervalByStatus(
    request stamina.SetRecoverIntervalByStatusRequest,
) (*stamina.SetRecoverIntervalByStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.StaminaName = &p.staminaName
    r, err := p.client.SetRecoverIntervalByStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaAccessTokenDomain) SetRecoverValueByStatus(
    request stamina.SetRecoverValueByStatusRequest,
) (*stamina.SetRecoverValueByStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.StaminaName = &p.staminaName
    r, err := p.client.SetRecoverValueByStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}
