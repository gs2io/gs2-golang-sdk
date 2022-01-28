package domainschedule
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

    "github.com/gs2io/gs2-golang-sdk/schedule"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type TriggerAccessTokenDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    triggerCache cache.TriggerDomainCache
    namespaceName string
    accessToken auth.AccessToken
    triggerName string
}

func NewTriggerAccessTokenDomain(
    session *core.Gs2RestSession,
    triggerCache cache.TriggerDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    triggerName string,
) TriggerAccessTokenDomain {
    return TriggerAccessTokenDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        triggerCache: triggerCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        triggerName: triggerName,
    }
}

func (p TriggerAccessTokenDomain) Load(
    request schedule.GetTriggerRequest,
) (*schedule.GetTriggerResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.TriggerName = &p.triggerName
    r, err := p.client.GetTrigger(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.triggerCache.Update(*r.Item)
    return r, nil
}

func (p TriggerAccessTokenDomain) Delete(
    request schedule.DeleteTriggerRequest,
) (*schedule.DeleteTriggerResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.TriggerName = &p.triggerName
    r, err := p.client.DeleteTrigger(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.triggerCache.Delete(*r.Item)
    return r, nil
}
