package domaindictionary
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

    "github.com/gs2io/gs2-golang-sdk/dictionary"
    "github.com/gs2io/gs2-golang-sdk/dictionary/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/dictionary/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type CurrentEntryMasterDomain struct {
    session *core.Gs2RestSession
    client dictionary.Gs2DictionaryRestClient
    namespaceName string
}

func NewCurrentEntryMasterDomain(
    session *core.Gs2RestSession,
    namespaceName string,
) CurrentEntryMasterDomain {
    return CurrentEntryMasterDomain {
        session: session,
        client: dictionary.Gs2DictionaryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
    }
}

func (p CurrentEntryMasterDomain) ExportMaster(
    request dictionary.ExportMasterRequest,
) (*dictionary.ExportMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.ExportMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEntryMasterDomain) Load(
    request dictionary.GetCurrentEntryMasterRequest,
) (*dictionary.GetCurrentEntryMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetCurrentEntryMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEntryMasterDomain) Update(
    request dictionary.UpdateCurrentEntryMasterRequest,
) (*dictionary.UpdateCurrentEntryMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentEntryMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p CurrentEntryMasterDomain) UpdateFromGitHub(
    request dictionary.UpdateCurrentEntryMasterFromGitHubRequest,
) (*dictionary.UpdateCurrentEntryMasterFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateCurrentEntryMasterFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
