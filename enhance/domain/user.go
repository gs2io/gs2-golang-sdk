package domainenhance
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

    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    namespaceName string
    userId string
    progressCache cache.ProgressDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        progressCache: cache.NewProgressDomainCache(),
    }
}

func (p UserDomain) DirectEnhance(
    request enhance.DirectEnhanceByUserIdRequest,
) (*enhance.DirectEnhanceByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DirectEnhanceByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) CreateProgress(
    request enhance.CreateProgressByUserIdRequest,
) (*enhance.CreateProgressByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CreateProgressByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Start(
    request enhance.StartByUserIdRequest,
) (*enhance.StartByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.StartByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) End(
    request enhance.EndByUserIdRequest,
) (*enhance.EndByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.EndByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) DeleteProgress(
    request enhance.DeleteProgressByUserIdRequest,
) (*enhance.DeleteProgressByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteProgressByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Progress(
) ProgressDomain {
    return NewProgressDomain(
        p.session,
        p.progressCache,
        p.namespaceName,
        p.userId,
    )
}
