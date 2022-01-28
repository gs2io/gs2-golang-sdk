package domaininbox
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

    "github.com/gs2io/gs2-golang-sdk/inbox"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/inbox/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    namespaceName string
    accessToken auth.AccessToken
    messageCache cache.MessageDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        messageCache: cache.NewMessageDomainCache(),
    }
}

func (p UserAccessTokenDomain) ReceiveGlobalMessage(
    request inbox.ReceiveGlobalMessageRequest,
) (*inbox.ReceiveGlobalMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.ReceiveGlobalMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Messages(
) iterator.DescribeMessagesIterator {
    return iterator.NewDescribeMessagesIterator(
        p.messageCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Message(
    messageName string,
) MessageAccessTokenDomain {
    return NewMessageAccessTokenDomain(
        p.session,
        p.messageCache,
        p.namespaceName,
        p.accessToken,
        messageName,
    )
}

func (p UserAccessTokenDomain) Received(
) ReceivedAccessTokenDomain {
    return NewReceivedAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}