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

type WebSocketSessionAccessTokenDomain struct {
    session *core.Gs2RestSession
    client gateway.Gs2GatewayRestClient
    webSocketSessionCache cache.WebSocketSessionDomainCache
    namespaceName string
    accessToken auth.AccessToken
}

func NewWebSocketSessionAccessTokenDomain(
    session *core.Gs2RestSession,
    webSocketSessionCache cache.WebSocketSessionDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
) WebSocketSessionAccessTokenDomain {
    return WebSocketSessionAccessTokenDomain {
        session: session,
        client: gateway.Gs2GatewayRestClient{
            Session: session,
        },
        webSocketSessionCache: webSocketSessionCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p WebSocketSessionAccessTokenDomain) SetUserId(
    request gateway.SetUserIdRequest,
) (*gateway.SetUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.SetUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.webSocketSessionCache.Update(*r.Item)
    return r, nil
}
