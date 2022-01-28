package domainmatchmaking
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

    "github.com/gs2io/gs2-golang-sdk/matchmaking"
    "github.com/gs2io/gs2-golang-sdk/matchmaking/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/matchmaking/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Matchmaking struct {
    session *core.Gs2RestSession
    client matchmaking.Gs2MatchmakingRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Matchmaking(
    session *core.Gs2RestSession,
) Gs2Matchmaking {
    return Gs2Matchmaking {
        session: session,
        client: matchmaking.Gs2MatchmakingRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Matchmaking) CreateNamespace(
    request matchmaking.CreateNamespaceRequest,
) (*matchmaking.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Matchmaking) DoMatchmakingByPlayer(
    request matchmaking.DoMatchmakingByPlayerRequest,
) (*matchmaking.DoMatchmakingByPlayerResult, error) {
    r, err := p.client.DoMatchmakingByPlayer(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) DeleteGathering(
    request matchmaking.DeleteGatheringRequest,
) (*matchmaking.DeleteGatheringResult, error) {
    r, err := p.client.DeleteGathering(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) PutResult(
    request matchmaking.PutResultRequest,
) (*matchmaking.PutResultResult, error) {
    r, err := p.client.PutResult(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) Vote(
    request matchmaking.VoteRequest,
) (*matchmaking.VoteResult, error) {
    r, err := p.client.Vote(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) VoteMultiple(
    request matchmaking.VoteMultipleRequest,
) (*matchmaking.VoteMultipleResult, error) {
    r, err := p.client.VoteMultiple(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) CommitVote(
    request matchmaking.CommitVoteRequest,
) (*matchmaking.CommitVoteResult, error) {
    r, err := p.client.CommitVote(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Matchmaking) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Matchmaking) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
