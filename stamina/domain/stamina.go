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

type StaminaDomain struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    staminaCache cache.StaminaDomainCache
    namespaceName string
    userId string
    staminaName string
}

func NewStaminaDomain(
    session *core.Gs2RestSession,
    staminaCache cache.StaminaDomainCache,
    namespaceName string,
    userId string,
    staminaName string,
) StaminaDomain {
    return StaminaDomain {
        session: session,
        client: stamina.Gs2StaminaRestClient{
            Session: session,
        },
        staminaCache: staminaCache,
        namespaceName: namespaceName,
        userId: userId,
        staminaName: staminaName,
    }
}

func (p StaminaDomain) Load(
    request stamina.GetStaminaByUserIdRequest,
) (*stamina.GetStaminaByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.GetStaminaByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) Update(
    request stamina.UpdateStaminaByUserIdRequest,
) (*stamina.UpdateStaminaByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.UpdateStaminaByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) Consume(
    request stamina.ConsumeStaminaByUserIdRequest,
) (*stamina.ConsumeStaminaByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.ConsumeStaminaByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) Recover(
    request stamina.RecoverStaminaByUserIdRequest,
) (*stamina.RecoverStaminaByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.RecoverStaminaByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) RaiseMaxValue(
    request stamina.RaiseMaxValueByUserIdRequest,
) (*stamina.RaiseMaxValueByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.RaiseMaxValueByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) SetMaxValue(
    request stamina.SetMaxValueByUserIdRequest,
) (*stamina.SetMaxValueByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.SetMaxValueByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) SetRecoverInterval(
    request stamina.SetRecoverIntervalByUserIdRequest,
) (*stamina.SetRecoverIntervalByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.SetRecoverIntervalByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) SetRecoverValue(
    request stamina.SetRecoverValueByUserIdRequest,
) (*stamina.SetRecoverValueByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.SetRecoverValueByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.staminaCache.Update(*r.Item)
    return r, nil
}

func (p StaminaDomain) Delete(
    request stamina.DeleteStaminaByUserIdRequest,
) (*stamina.DeleteStaminaByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.StaminaName = &p.staminaName
    r, err := p.client.DeleteStaminaByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
