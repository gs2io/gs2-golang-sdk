package domainidentifier
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

    "github.com/gs2io/gs2-golang-sdk/identifier"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.UserDomainCache{}
var _ = iterator.DescribeUsersIterator{}

type IdentifierDomain struct {
    session *core.Gs2RestSession
    client identifier.Gs2IdentifierRestClient
    identifierCache cache.IdentifierDomainCache
    userName string
}

func NewIdentifierDomain(
    session *core.Gs2RestSession,
    identifierCache cache.IdentifierDomainCache,
    userName string,
) IdentifierDomain {
    return IdentifierDomain {
        session: session,
        client: identifier.Gs2IdentifierRestClient{
            Session: session,
        },
        identifierCache: identifierCache,
        userName: userName,
    }
}

func (p IdentifierDomain) Load(
    request identifier.GetIdentifierRequest,
) (*identifier.GetIdentifierResult, error) {
    request.UserName = &p.userName
    r, err := p.client.GetIdentifier(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.identifierCache.Update(*r.Item)
    return r, nil
}

func (p IdentifierDomain) List(
) iterator.DescribeIdentifiersIterator {
    return iterator.NewDescribeIdentifiersIterator(
        p.identifierCache,
        p.client,
        p.userName,
    )
}
