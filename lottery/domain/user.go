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
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    namespaceName string
    userId string
    boxCache cache.BoxDomainCache
    probabilityCache cache.ProbabilityDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        boxCache: cache.NewBoxDomainCache(),
        probabilityCache: cache.NewProbabilityDomainCache(),
    }
}

func (p UserDomain) Draw(
    request lottery.DrawByUserIdRequest,
) (*lottery.DrawByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DrawByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Boxes(
) iterator.DescribeBoxesByUserIdIterator {
    return iterator.NewDescribeBoxesByUserIdIterator(
        p.boxCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Probabilities(
    lotteryName string,
) iterator.DescribeProbabilitiesByUserIdIterator {
    return iterator.NewDescribeProbabilitiesByUserIdIterator(
        p.probabilityCache,
        p.client,
        p.namespaceName,
        lotteryName,
        p.userId,
    )
}

func (p UserDomain) Box(
    prizeTableName string,
) BoxDomain {
    return NewBoxDomain(
        p.session,
        p.boxCache,
        p.namespaceName,
        p.userId,
        prizeTableName,
    )
}
