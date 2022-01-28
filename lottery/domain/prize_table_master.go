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

type PrizeTableMasterDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    prizeTableMasterCache cache.PrizeTableMasterDomainCache
    namespaceName string
    prizeTableName string
}

func NewPrizeTableMasterDomain(
    session *core.Gs2RestSession,
    prizeTableMasterCache cache.PrizeTableMasterDomainCache,
    namespaceName string,
    prizeTableName string,
) PrizeTableMasterDomain {
    return PrizeTableMasterDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        prizeTableMasterCache: prizeTableMasterCache,
        namespaceName: namespaceName,
        prizeTableName: prizeTableName,
    }
}

func (p PrizeTableMasterDomain) Load(
    request lottery.GetPrizeTableMasterRequest,
) (*lottery.GetPrizeTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.GetPrizeTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.prizeTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p PrizeTableMasterDomain) Update(
    request lottery.UpdatePrizeTableMasterRequest,
) (*lottery.UpdatePrizeTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.UpdatePrizeTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.prizeTableMasterCache.Update(*r.Item)
    return r, nil
}

func (p PrizeTableMasterDomain) Delete(
    request lottery.DeletePrizeTableMasterRequest,
) (*lottery.DeletePrizeTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.DeletePrizeTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.prizeTableMasterCache.Delete(*r.Item)
    return r, nil
}
