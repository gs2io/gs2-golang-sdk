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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    namespaceName string
    accessToken auth.AccessToken
    boxCache cache.BoxDomainCache
    probabilityCache cache.ProbabilityDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        boxCache: cache.NewBoxDomainCache(),
        probabilityCache: cache.NewProbabilityDomainCache(),
    }
}

func (p UserAccessTokenDomain) Boxes(
) iterator.DescribeBoxesIterator {
    return iterator.NewDescribeBoxesIterator(
        p.boxCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Probabilities(
    lotteryName string,
) iterator.DescribeProbabilitiesIterator {
    return iterator.NewDescribeProbabilitiesIterator(
        p.probabilityCache,
        p.client,
        p.namespaceName,
        lotteryName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Box(
    prizeTableName string,
) BoxAccessTokenDomain {
    return NewBoxAccessTokenDomain(
        p.session,
        p.boxCache,
        p.namespaceName,
        p.accessToken,
        prizeTableName,
    )
}
