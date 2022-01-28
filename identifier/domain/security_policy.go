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

type SecurityPolicyDomain struct {
    session *core.Gs2RestSession
    client identifier.Gs2IdentifierRestClient
    securityPolicyCache cache.SecurityPolicyDomainCache
    securityPolicyName string
}

func NewSecurityPolicyDomain(
    session *core.Gs2RestSession,
    securityPolicyCache cache.SecurityPolicyDomainCache,
    securityPolicyName string,
) SecurityPolicyDomain {
    return SecurityPolicyDomain {
        session: session,
        client: identifier.Gs2IdentifierRestClient{
            Session: session,
        },
        securityPolicyCache: securityPolicyCache,
        securityPolicyName: securityPolicyName,
    }
}

func (p SecurityPolicyDomain) Load(
    request identifier.GetSecurityPolicyRequest,
) (*identifier.GetSecurityPolicyResult, error) {
    request.SecurityPolicyName = &p.securityPolicyName
    r, err := p.client.GetSecurityPolicy(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.securityPolicyCache.Update(*r.Item)
    return r, nil
}
