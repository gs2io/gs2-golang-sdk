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

type UserDomain struct {
    session *core.Gs2RestSession
    client identifier.Gs2IdentifierRestClient
    userCache cache.UserDomainCache
    userName string
    identifierCache cache.IdentifierDomainCache
    passwordCache cache.PasswordDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    userCache cache.UserDomainCache,
    userName string,
) UserDomain {
    return UserDomain {
        session: session,
        client: identifier.Gs2IdentifierRestClient{
            Session: session,
        },
        userCache: userCache,
        userName: userName,
        identifierCache: cache.NewIdentifierDomainCache(),
        passwordCache: cache.NewPasswordDomainCache(),
    }
}

func (p UserDomain) Load(
    request identifier.GetUserRequest,
) (*identifier.GetUserResult, error) {
    request.UserName = &p.userName
    r, err := p.client.GetUser(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.userCache.Update(*r.Item)
    return r, nil
}

func (p UserDomain) CreatePassword(
    request identifier.CreatePasswordRequest,
) (*identifier.CreatePasswordResult, error) {
    request.UserName = &p.userName
    r, err := p.client.CreatePassword(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) DeletePassword(
    request identifier.DeletePasswordRequest,
) (*identifier.DeletePasswordResult, error) {
    request.UserName = &p.userName
    r, err := p.client.DeletePassword(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) GetHasSecurityPolicy(
    request identifier.GetHasSecurityPolicyRequest,
) (*identifier.GetHasSecurityPolicyResult, error) {
    request.UserName = &p.userName
    r, err := p.client.GetHasSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) AttachSecurityPolicy(
    request identifier.AttachSecurityPolicyRequest,
) (*identifier.AttachSecurityPolicyResult, error) {
    request.UserName = &p.userName
    r, err := p.client.AttachSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) DetachSecurityPolicy(
    request identifier.DetachSecurityPolicyRequest,
) (*identifier.DetachSecurityPolicyResult, error) {
    request.UserName = &p.userName
    r, err := p.client.DetachSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Passwords(
) iterator.DescribePasswordsIterator {
    return iterator.NewDescribePasswordsIterator(
        p.passwordCache,
        p.client,
        p.userName,
    )
}

func (p UserDomain) Identifier(
) IdentifierDomain {
    return NewIdentifierDomain(
        p.session,
        p.identifierCache,
        p.userName,
    )
}

func (p UserDomain) Password(
) PasswordDomain {
    return NewPasswordDomain(
        p.session,
        p.passwordCache,
        p.userName,
    )
}
