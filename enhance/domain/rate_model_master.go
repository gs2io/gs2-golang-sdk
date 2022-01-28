package domainenhance
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

    "github.com/gs2io/gs2-golang-sdk/enhance"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/enhance/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type RateModelMasterDomain struct {
    session *core.Gs2RestSession
    client enhance.Gs2EnhanceRestClient
    rateModelMasterCache cache.RateModelMasterDomainCache
    namespaceName string
    rateName string
}

func NewRateModelMasterDomain(
    session *core.Gs2RestSession,
    rateModelMasterCache cache.RateModelMasterDomainCache,
    namespaceName string,
    rateName string,
) RateModelMasterDomain {
    return RateModelMasterDomain {
        session: session,
        client: enhance.Gs2EnhanceRestClient{
            Session: session,
        },
        rateModelMasterCache: rateModelMasterCache,
        namespaceName: namespaceName,
        rateName: rateName,
    }
}

func (p RateModelMasterDomain) Load(
    request enhance.GetRateModelMasterRequest,
) (*enhance.GetRateModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RateName = &p.rateName
    r, err := p.client.GetRateModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.rateModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p RateModelMasterDomain) Update(
    request enhance.UpdateRateModelMasterRequest,
) (*enhance.UpdateRateModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RateName = &p.rateName
    r, err := p.client.UpdateRateModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.rateModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p RateModelMasterDomain) Delete(
    request enhance.DeleteRateModelMasterRequest,
) (*enhance.DeleteRateModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RateName = &p.rateName
    r, err := p.client.DeleteRateModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.rateModelMasterCache.Delete(*r.Item)
    return r, nil
}
