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

type GlobalMessageDomain struct {
    session *core.Gs2RestSession
    client inbox.Gs2InboxRestClient
    globalMessageCache cache.GlobalMessageDomainCache
    namespaceName string
    globalMessageName string
}

func NewGlobalMessageDomain(
    session *core.Gs2RestSession,
    globalMessageCache cache.GlobalMessageDomainCache,
    namespaceName string,
    globalMessageName string,
) GlobalMessageDomain {
    return GlobalMessageDomain {
        session: session,
        client: inbox.Gs2InboxRestClient{
            Session: session,
        },
        globalMessageCache: globalMessageCache,
        namespaceName: namespaceName,
        globalMessageName: globalMessageName,
    }
}

func (p GlobalMessageDomain) Load(
    request inbox.GetGlobalMessageRequest,
) (*inbox.GetGlobalMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.GlobalMessageName = &p.globalMessageName
    r, err := p.client.GetGlobalMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.globalMessageCache.Update(*r.Item)
    return r, nil
}
