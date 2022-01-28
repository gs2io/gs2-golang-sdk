package domaingateway
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

    "github.com/gs2io/gs2-golang-sdk/gateway"
    "github.com/gs2io/gs2-golang-sdk/gateway/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/gateway/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type FirebaseTokenDomain struct {
    session *core.Gs2RestSession
    client gateway.Gs2GatewayRestClient
    namespaceName string
    userId string
}

func NewFirebaseTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) FirebaseTokenDomain {
    return FirebaseTokenDomain {
        session: session,
        client: gateway.Gs2GatewayRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p FirebaseTokenDomain) Set(
    request gateway.SetFirebaseTokenByUserIdRequest,
) (*gateway.SetFirebaseTokenByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.SetFirebaseTokenByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FirebaseTokenDomain) Load(
    request gateway.GetFirebaseTokenByUserIdRequest,
) (*gateway.GetFirebaseTokenByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetFirebaseTokenByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FirebaseTokenDomain) Delete(
    request gateway.DeleteFirebaseTokenByUserIdRequest,
) (*gateway.DeleteFirebaseTokenByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteFirebaseTokenByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FirebaseTokenDomain) SendMobileNotification(
    request gateway.SendMobileNotificationByUserIdRequest,
) (*gateway.SendMobileNotificationByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.SendMobileNotificationByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
