package domainlottery
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

    "github.com/gs2io/gs2-golang-sdk/lottery"
    "github.com/gs2io/gs2-golang-sdk/lottery/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/lottery/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Lottery struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    namespaceCache cache.NamespaceDomainCache
    probabilityCache cache.ProbabilityDomainCache
}

func NewGs2Lottery(
    session *core.Gs2RestSession,
) Gs2Lottery {
    return Gs2Lottery {
        session: session,
        client: lottery.Gs2LotteryRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
        probabilityCache: cache.NewProbabilityDomainCache(),
    }
}

func (p Gs2Lottery) CreateNamespace(
    request lottery.CreateNamespaceRequest,
) (*lottery.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Lottery) DrawByStampSheet(
    request lottery.DrawByStampSheetRequest,
) (*lottery.DrawByStampSheetResult, error) {
    r, err := p.client.DrawByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Lottery) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Lottery) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
