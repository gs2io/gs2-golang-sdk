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

type ScriptDomain struct {
    session *core.Gs2RestSession
    client script.Gs2ScriptRestClient
    scriptCache cache.ScriptDomainCache
    namespaceName string
    scriptName string
}

func NewScriptDomain(
    session *core.Gs2RestSession,
    scriptCache cache.ScriptDomainCache,
    namespaceName string,
    scriptName string,
) ScriptDomain {
    return ScriptDomain {
        session: session,
        client: script.Gs2ScriptRestClient{
            Session: session,
        },
        scriptCache: scriptCache,
        namespaceName: namespaceName,
        scriptName: scriptName,
    }
}

func (p ScriptDomain) Load(
    request script.GetScriptRequest,
) (*script.GetScriptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ScriptName = &p.scriptName
    r, err := p.client.GetScript(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.scriptCache.Update(*r.Item)
    return r, nil
}

func (p ScriptDomain) Update(
    request script.UpdateScriptRequest,
) (*script.UpdateScriptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ScriptName = &p.scriptName
    r, err := p.client.UpdateScript(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.scriptCache.Update(*r.Item)
    return r, nil
}

func (p ScriptDomain) UpdateFromGitHub(
    request script.UpdateScriptFromGitHubRequest,
) (*script.UpdateScriptFromGitHubResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ScriptName = &p.scriptName
    r, err := p.client.UpdateScriptFromGitHub(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.scriptCache.Update(*r.Item)
    return r, nil
}

func (p ScriptDomain) Delete(
    request script.DeleteScriptRequest,
) (*script.DeleteScriptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ScriptName = &p.scriptName
    r, err := p.client.DeleteScript(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
