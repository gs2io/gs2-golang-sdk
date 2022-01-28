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

type CurrentLotteryMasterDomain struct {
    session *core.Gs2RestSession
    client lottery.Gs2LotteryRestClient
    namespaceName string
}

func NewCurrentLotteryMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentLotteryMasterDomain {
    return CurrentLotteryMasterDomain {
        session: session,
        client: lottery.Gs2LotteryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentLotteryMasterDomain) ExportMaster(
    request lottery.ExportMasterRequest,
) (*lottery.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLotteryMasterDomain) Load(
    request lottery.GetCurrentLotteryMasterRequest,
) (*lottery.GetCurrentLotteryMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentLotteryMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLotteryMasterDomain) Update(
    request lottery.UpdateCurrentLotteryMasterRequest,
) (*lottery.UpdateCurrentLotteryMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentLotteryMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentLotteryMasterDomain) UpdateFromGitHub(
    request lottery.UpdateCurrentLotteryMasterFromGitHubRequest,
) (*lottery.UpdateCurrentLotteryMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentLotteryMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
