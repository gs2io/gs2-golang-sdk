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

type CurrentEventMasterDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    namespaceName string
}

func NewCurrentEventMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentEventMasterDomain {
    return CurrentEventMasterDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentEventMasterDomain) ExportMaster(
    request schedule.ExportMasterRequest,
) (*schedule.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEventMasterDomain) Load(
    request schedule.GetCurrentEventMasterRequest,
) (*schedule.GetCurrentEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEventMasterDomain) Update(
    request schedule.UpdateCurrentEventMasterRequest,
) (*schedule.UpdateCurrentEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEventMasterDomain) UpdateFromGitHub(
    request schedule.UpdateCurrentEventMasterFromGitHubRequest,
) (*schedule.UpdateCurrentEventMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentEventMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
