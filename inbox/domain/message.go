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

type MessageDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    messageCache cache.MessageDomainCache
    namespaceName string
    userId string
    messageName string
}

func NewMessageDomain(
    session *core.Gs2RestSession,
    messageCache cache.MessageDomainCache,
    namespaceName string,
    userId string,
    messageName string,
) MessageDomain {
    return MessageDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        messageCache: messageCache,
        namespaceName: namespaceName,
        userId: userId,
        messageName: messageName,
    }
}

func (p MessageDomain) Load(
    request inbox.GetMessageByUserIdRequest,
) (*inbox.GetMessageByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MessageName = &p.messageName
    r, err := p.client.GetMessageByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageDomain) Open(
    request inbox.OpenMessageByUserIdRequest,
) (*inbox.OpenMessageByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MessageName = &p.messageName
    r, err := p.client.OpenMessageByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageDomain) Read(
    request inbox.ReadMessageByUserIdRequest,
) (*inbox.ReadMessageByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MessageName = &p.messageName
    r, err := p.client.ReadMessageByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageDomain) Delete(
    request inbox.DeleteMessageByUserIdRequest,
) (*inbox.DeleteMessageByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.MessageName = &p.messageName
    r, err := p.client.DeleteMessageByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
