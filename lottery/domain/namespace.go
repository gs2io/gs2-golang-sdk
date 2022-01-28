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

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    lotteryModelMasterCache cache.LotteryModelMasterDomainCache
    prizeTableMasterCache cache.PrizeTableMasterDomainCache
    lotteryModelCache cache.LotteryModelDomainCache
    prizeTableCache cache.PrizeTableDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        lotteryModelMasterCache: cache.NewLotteryModelMasterDomainCache(),
        prizeTableMasterCache: cache.NewPrizeTableMasterDomainCache(),
        lotteryModelCache: cache.NewLotteryModelDomainCache(),
        prizeTableCache: cache.NewPrizeTableDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request lottery.GetNamespaceStatusRequest,
) (*lottery.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request lottery.GetNamespaceRequest,
) (*lottery.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request lottery.UpdateNamespaceRequest,
) (*lottery.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request lottery.DeleteNamespaceRequest,
) (*lottery.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateLotteryModelMaster(
    request lottery.CreateLotteryModelMasterRequest,
) (*lottery.CreateLotteryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateLotteryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreatePrizeTableMaster(
    request lottery.CreatePrizeTableMasterRequest,
) (*lottery.CreatePrizeTableMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreatePrizeTableMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) LotteryModelMasters(
) iterator.DescribeLotteryModelMastersIterator {
    return iterator.NewDescribeLotteryModelMastersIterator(
        p.lotteryModelMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) PrizeTableMasters(
) iterator.DescribePrizeTableMastersIterator {
    return iterator.NewDescribePrizeTableMastersIterator(
        p.prizeTableMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) LotteryModels(
) iterator.DescribeLotteryModelsIterator {
    return iterator.NewDescribeLotteryModelsIterator(
        p.lotteryModelCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) PrizeTables(
) iterator.DescribePrizeTablesIterator {
    return iterator.NewDescribePrizeTablesIterator(
        p.prizeTableCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) CurrentLotteryMaster(
) CurrentLotteryMasterDomain {
    return NewCurrentLotteryMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) LotteryModel(
    lotteryName string,
) LotteryModelDomain {
    return NewLotteryModelDomain(
        p.session,
        p.lotteryModelCache,
        p.namespaceName,
        lotteryName,
    )
}

func (p NamespaceDomain) PrizeTableMaster(
    prizeTableName string,
) PrizeTableMasterDomain {
    return NewPrizeTableMasterDomain(
        p.session,
        p.prizeTableMasterCache,
        p.namespaceName,
        prizeTableName,
    )
}

func (p NamespaceDomain) LotteryModelMaster(
    lotteryName string,
) LotteryModelMasterDomain {
    return NewLotteryModelMasterDomain(
        p.session,
        p.lotteryModelMasterCache,
        p.namespaceName,
        lotteryName,
    )
}

func (p NamespaceDomain) PrizeTable(
    prizeTableName string,
) PrizeTableDomain {
    return NewPrizeTableDomain(
        p.session,
        p.prizeTableCache,
        p.namespaceName,
        prizeTableName,
    )
}
