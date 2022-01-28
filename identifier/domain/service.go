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

deny overwrite
*/

import (
    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/core"

    "github.com/gs2io/gs2-golang-sdk/identifier"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeUsersIterator{}

type Gs2Identifier struct {
    session *core.Gs2RestSession
    client identifier.Gs2IdentifierRestClient
    userCache cache.UserDomainCache
    securityPolicyCache cache.SecurityPolicyDomainCache
}

func NewGs2Identifier(
    session *core.Gs2RestSession,
) Gs2Identifier {
    return Gs2Identifier {
        session: session,
        client: identifier.Gs2IdentifierRestClient {
            Session: session,
        },
        userCache: cache.NewUserDomainCache(),
        securityPolicyCache: cache.NewSecurityPolicyDomainCache(),
    }
}

func (p Gs2Identifier) CreateUser(
    request identifier.CreateUserRequest,
) (*identifier.CreateUserResult, error) {
    r, err := p.client.CreateUser(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.userCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Identifier) UpdateUser(
    request identifier.UpdateUserRequest,
) (*identifier.UpdateUserResult, error) {
    r, err := p.client.UpdateUser(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.userCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Identifier) DeleteUser(
    request identifier.DeleteUserRequest,
) (*identifier.DeleteUserResult, error) {
    r, err := p.client.DeleteUser(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) CreateSecurityPolicy(
    request identifier.CreateSecurityPolicyRequest,
) (*identifier.CreateSecurityPolicyResult, error) {
    r, err := p.client.CreateSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.securityPolicyCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Identifier) UpdateSecurityPolicy(
    request identifier.UpdateSecurityPolicyRequest,
) (*identifier.UpdateSecurityPolicyResult, error) {
    r, err := p.client.UpdateSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.securityPolicyCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Identifier) DeleteSecurityPolicy(
    request identifier.DeleteSecurityPolicyRequest,
) (*identifier.DeleteSecurityPolicyResult, error) {
    r, err := p.client.DeleteSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) CreateIdentifier(
    request identifier.CreateIdentifierRequest,
) (*identifier.CreateIdentifierResult, error) {
    r, err := p.client.CreateIdentifier(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) DeleteIdentifier(
    request identifier.DeleteIdentifierRequest,
) (*identifier.DeleteIdentifierResult, error) {
    r, err := p.client.DeleteIdentifier(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) Login(
    request identifier.LoginRequest,
) (*identifier.LoginResult, error) {
    r, err := p.client.Login(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) LoginByUser(
    request identifier.LoginByUserRequest,
) (*identifier.LoginByUserResult, error) {
    r, err := p.client.LoginByUser(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Identifier) Users(
) iterator.DescribeUsersIterator {
    return iterator.NewDescribeUsersIterator(
        p.userCache,
        p.client,
    )
}

func (p Gs2Identifier) SecurityPolicies(
) iterator.DescribeSecurityPoliciesIterator {
    return iterator.NewDescribeSecurityPoliciesIterator(
        p.securityPolicyCache,
        p.client,
    )
}

func (p Gs2Identifier) CommonSecurityPolicies(
) iterator.DescribeCommonSecurityPoliciesIterator {
    return iterator.NewDescribeCommonSecurityPoliciesIterator(
        p.securityPolicyCache,
        p.client,
    )
}

func (p Gs2Identifier) User(
    userName string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.userCache,
        userName,
    )
}

func (p Gs2Identifier) SecurityPolicy(
    securityPolicyName string,
) SecurityPolicyDomain {
    return NewSecurityPolicyDomain(
        p.session,
        p.securityPolicyCache,
        securityPolicyName,
    )
}
