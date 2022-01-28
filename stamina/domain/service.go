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
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Stamina struct {
    session *core.Gs2RestSession
    client stamina.Gs2StaminaRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Stamina(
    session *core.Gs2RestSession,
) Gs2Stamina {
    return Gs2Stamina {
        session: session,
        client: stamina.Gs2StaminaRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Stamina) CreateNamespace(
    request stamina.CreateNamespaceRequest,
) (*stamina.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Stamina) RecoverStaminaByStampSheet(
    request stamina.RecoverStaminaByStampSheetRequest,
) (*stamina.RecoverStaminaByStampSheetResult, error) {
    r, err := p.client.RecoverStaminaByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) RaiseMaxValueByStampSheet(
    request stamina.RaiseMaxValueByStampSheetRequest,
) (*stamina.RaiseMaxValueByStampSheetResult, error) {
    r, err := p.client.RaiseMaxValueByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) SetMaxValueByStampSheet(
    request stamina.SetMaxValueByStampSheetRequest,
) (*stamina.SetMaxValueByStampSheetResult, error) {
    r, err := p.client.SetMaxValueByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) SetRecoverIntervalByStampSheet(
    request stamina.SetRecoverIntervalByStampSheetRequest,
) (*stamina.SetRecoverIntervalByStampSheetResult, error) {
    r, err := p.client.SetRecoverIntervalByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) SetRecoverValueByStampSheet(
    request stamina.SetRecoverValueByStampSheetRequest,
) (*stamina.SetRecoverValueByStampSheetResult, error) {
    r, err := p.client.SetRecoverValueByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) ConsumeStaminaByStampTask(
    request stamina.ConsumeStaminaByStampTaskRequest,
) (*stamina.ConsumeStaminaByStampTaskResult, error) {
    r, err := p.client.ConsumeStaminaByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Stamina) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Stamina) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
