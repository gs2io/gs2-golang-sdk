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

type BoxDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    boxCache cache.BoxDomainCache
    namespaceName string
    userId string
    prizeTableName string
}

func NewBoxDomain(
    session *core.Gs2RestSession,
    boxCache cache.BoxDomainCache,
    namespaceName string,
    userId string,
    prizeTableName string,
) BoxDomain {
    return BoxDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        boxCache: boxCache,
        namespaceName: namespaceName,
        userId: userId,
        prizeTableName: prizeTableName,
    }
}

func (p BoxDomain) Load(
    request lottery.GetBoxByUserIdRequest,
) (*lottery.GetBoxByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.GetBoxByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p BoxDomain) GetRaw(
    request lottery.GetRawBoxByUserIdRequest,
) (*lottery.GetRawBoxByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.GetRawBoxByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.boxCache.Update(*r.Item)
    return r, nil
}

func (p BoxDomain) Reset(
    request lottery.ResetBoxByUserIdRequest,
) (*lottery.ResetBoxByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.PrizeTableName = &p.prizeTableName
    r, err := p.client.ResetBoxByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
