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

type LotteryModelMasterDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    lotteryModelMasterCache cache.LotteryModelMasterDomainCache
    namespaceName string
    lotteryName string
}

func NewLotteryModelMasterDomain(
    session *core.Gs2RestSession,
    lotteryModelMasterCache cache.LotteryModelMasterDomainCache,
    namespaceName string,
    lotteryName string,
) LotteryModelMasterDomain {
    return LotteryModelMasterDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        lotteryModelMasterCache: lotteryModelMasterCache,
        namespaceName: namespaceName,
        lotteryName: lotteryName,
    }
}

func (p LotteryModelMasterDomain) Load(
    request lottery.GetLotteryModelMasterRequest,
) (*lottery.GetLotteryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.LotteryName = &p.lotteryName
    r, err := p.client.GetLotteryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.lotteryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p LotteryModelMasterDomain) Update(
    request lottery.UpdateLotteryModelMasterRequest,
) (*lottery.UpdateLotteryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.LotteryName = &p.lotteryName
    r, err := p.client.UpdateLotteryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.lotteryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p LotteryModelMasterDomain) Delete(
    request lottery.DeleteLotteryModelMasterRequest,
) (*lottery.DeleteLotteryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.LotteryName = &p.lotteryName
    r, err := p.client.DeleteLotteryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.lotteryModelMasterCache.Delete(*r.Item)
    return r, nil
}
