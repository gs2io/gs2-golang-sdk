package domainquest
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

    "github.com/gs2io/gs2-golang-sdk/quest"
    "github.com/gs2io/gs2-golang-sdk/quest/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/quest/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentQuestMasterDomain struct {
    session *core.Gs2RestSession
    client quest.Gs2QuestRestClient
    namespaceName string
}

func NewCurrentQuestMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentQuestMasterDomain {
    return CurrentQuestMasterDomain {
        session: session,
        client: quest.Gs2QuestRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentQuestMasterDomain) ExportMaster(
    request quest.ExportMasterRequest,
) (*quest.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentQuestMasterDomain) Load(
    request quest.GetCurrentQuestMasterRequest,
) (*quest.GetCurrentQuestMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentQuestMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentQuestMasterDomain) Update(
    request quest.UpdateCurrentQuestMasterRequest,
) (*quest.UpdateCurrentQuestMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentQuestMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentQuestMasterDomain) UpdateFromGitHub(
    request quest.UpdateCurrentQuestMasterFromGitHubRequest,
) (*quest.UpdateCurrentQuestMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentQuestMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
