package domainlimit
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

    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type LimitModelDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    limitModelCache cache.LimitModelDomainCache
    namespaceName string
    limitName string
}

func NewLimitModelDomain(
    session *core.Gs2RestSession,
    limitModelCache cache.LimitModelDomainCache,
    namespaceName string,
    limitName string,
) LimitModelDomain {
    return LimitModelDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        limitModelCache: limitModelCache,
        namespaceName: namespaceName,
        limitName: limitName,
    }
}

func (p LimitModelDomain) Load(
    request limit.GetLimitModelRequest,
) (*limit.GetLimitModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.LimitName = &p.limitName
    r, err := p.client.GetLimitModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.limitModelCache.Update(*r.Item)
    return r, nil
}
