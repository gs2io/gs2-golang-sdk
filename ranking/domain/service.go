package domainranking
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

    "github.com/gs2io/gs2-golang-sdk/ranking"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/ranking/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Ranking struct {
    session *core.Gs2RestSession
    client ranking.Gs2RankingRestClient
    namespaceCache cache.NamespaceDomainCache
    subscribeUserCache cache.SubscribeUserDomainCache
}

func NewGs2Ranking(
    session *core.Gs2RestSession,
) Gs2Ranking {
    return Gs2Ranking {
        session: session,
        client: ranking.Gs2RankingRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
        subscribeUserCache: cache.NewSubscribeUserDomainCache(),
    }
}

func (p Gs2Ranking) CreateNamespace(
    request ranking.CreateNamespaceRequest,
) (*ranking.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Ranking) CalcRanking(
    request ranking.CalcRankingRequest,
) (*ranking.CalcRankingResult, error) {
    r, err := p.client.CalcRanking(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Ranking) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Ranking) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
