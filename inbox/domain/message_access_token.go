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

type MessageAccessTokenDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    messageCache cache.MessageDomainCache
    namespaceName string
    accessToken auth.AccessToken
    messageName string
}

func NewMessageAccessTokenDomain(
    session *core.Gs2RestSession,
    messageCache cache.MessageDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    messageName string,
) MessageAccessTokenDomain {
    return MessageAccessTokenDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        messageCache: messageCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        messageName: messageName,
    }
}

func (p MessageAccessTokenDomain) Load(
    request inbox.GetMessageRequest,
) (*inbox.GetMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MessageName = &p.messageName
    r, err := p.client.GetMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageAccessTokenDomain) Open(
    request inbox.OpenMessageRequest,
) (*inbox.OpenMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MessageName = &p.messageName
    r, err := p.client.OpenMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageAccessTokenDomain) Read(
    request inbox.ReadMessageRequest,
) (*inbox.ReadMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MessageName = &p.messageName
    r, err := p.client.ReadMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageAccessTokenDomain) Delete(
    request inbox.DeleteMessageRequest,
) (*inbox.DeleteMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.MessageName = &p.messageName
    r, err := p.client.DeleteMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
