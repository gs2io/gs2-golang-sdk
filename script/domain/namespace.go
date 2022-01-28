package domainscript
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

    "github.com/gs2io/gs2-golang-sdk/script"
    "github.com/gs2io/gs2-golang-sdk/script/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/script/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client script.Gs2ScriptRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    scriptCache cache.ScriptDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: script.Gs2ScriptRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        scriptCache: cache.NewScriptDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request script.GetNamespaceStatusRequest,
) (*script.GetNamespaceStatusResult, error) {
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
    request script.GetNamespaceRequest,
) (*script.GetNamespaceResult, error) {
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
    request script.UpdateNamespaceRequest,
) (*script.UpdateNamespaceResult, error) {
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
    request script.DeleteNamespaceRequest,
) (*script.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateScript(
    request script.CreateScriptRequest,
) (*script.CreateScriptResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateScript(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateScriptFromGitHub(
    request script.CreateScriptFromGitHubRequest,
) (*script.CreateScriptFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateScriptFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Scripts(
) iterator.DescribeScriptsIterator {
    return iterator.NewDescribeScriptsIterator(
        p.scriptCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) Script(
    scriptName string,
) ScriptDomain {
    return NewScriptDomain(
        p.session,
        p.scriptCache,
        p.namespaceName,
        scriptName,
    )
}
