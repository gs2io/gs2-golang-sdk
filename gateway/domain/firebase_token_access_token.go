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

type FirebaseTokenAccessTokenDomain struct {
    session *core.Gs2RestSession
    client gateway.Gs2GatewayRestClient
    namespaceName string
    accessToken auth.AccessToken
}

func NewFirebaseTokenAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) FirebaseTokenAccessTokenDomain {
    return FirebaseTokenAccessTokenDomain {
        session: session,
        client: gateway.Gs2GatewayRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p FirebaseTokenAccessTokenDomain) Set(
    request gateway.SetFirebaseTokenRequest,
) (*gateway.SetFirebaseTokenResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.SetFirebaseToken(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FirebaseTokenAccessTokenDomain) Load(
    request gateway.GetFirebaseTokenRequest,
) (*gateway.GetFirebaseTokenResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.GetFirebaseToken(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FirebaseTokenAccessTokenDomain) Delete(
    request gateway.DeleteFirebaseTokenRequest,
) (*gateway.DeleteFirebaseTokenResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.DeleteFirebaseToken(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
