package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/identifier"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.UserDomainCache{}


type DescribeCommonSecurityPoliciesIterator struct {
    securityPolicyCache cache.SecurityPolicyDomainCache
    client identifier.Gs2IdentifierRestClient

    index int64
    pageToken *string
    last bool
    result []identifier.SecurityPolicy

    FetchSize *int32
}

func NewDescribeCommonSecurityPoliciesIterator(
    securityPolicyCache cache.SecurityPolicyDomainCache,
    client identifier.Gs2IdentifierRestClient,
) DescribeCommonSecurityPoliciesIterator {
    it := DescribeCommonSecurityPoliciesIterator{
        securityPolicyCache: securityPolicyCache,
        client: client,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]identifier.SecurityPolicy, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeCommonSecurityPoliciesIterator) load() error {
    r, err := p.client.DescribeCommonSecurityPolicies(
        &identifier.DescribeCommonSecurityPoliciesRequest{
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.securityPolicyCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeCommonSecurityPoliciesIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeCommonSecurityPoliciesIterator) Next(

) (identifier.SecurityPolicy, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return identifier.SecurityPolicy{}, err
        }
    }
    if len(p.result) == 0 {
        return identifier.SecurityPolicy{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return identifier.SecurityPolicy{}, err
        }
    }
    return ret, nil
}
