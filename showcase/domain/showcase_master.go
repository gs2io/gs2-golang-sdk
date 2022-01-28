package domainshowcase
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

    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type ShowcaseMasterDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    showcaseMasterCache cache.ShowcaseMasterDomainCache
    namespaceName string
    showcaseName string
}

func NewShowcaseMasterDomain(
    session *core.Gs2RestSession,
    showcaseMasterCache cache.ShowcaseMasterDomainCache,
    namespaceName string,
    showcaseName string,
) ShowcaseMasterDomain {
    return ShowcaseMasterDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        showcaseMasterCache: showcaseMasterCache,
        namespaceName: namespaceName,
        showcaseName: showcaseName,
    }
}

func (p ShowcaseMasterDomain) Load(
    request showcase.GetShowcaseMasterRequest,
) (*showcase.GetShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ShowcaseName = &p.showcaseName
    r, err := p.client.GetShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.showcaseMasterCache.Update(*r.Item)
    return r, nil
}

func (p ShowcaseMasterDomain) Update(
    request showcase.UpdateShowcaseMasterRequest,
) (*showcase.UpdateShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ShowcaseName = &p.showcaseName
    r, err := p.client.UpdateShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.showcaseMasterCache.Update(*r.Item)
    return r, nil
}

func (p ShowcaseMasterDomain) Delete(
    request showcase.DeleteShowcaseMasterRequest,
) (*showcase.DeleteShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ShowcaseName = &p.showcaseName
    r, err := p.client.DeleteShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.showcaseMasterCache.Delete(*r.Item)
    return r, nil
}
