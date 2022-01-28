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
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Script struct {
    session *core.Gs2RestSession
    client script.Gs2ScriptRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Script(
    session *core.Gs2RestSession,
) Gs2Script {
    return Gs2Script {
        session: session,
        client: script.Gs2ScriptRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Script) CreateNamespace(
    request script.CreateNamespaceRequest,
) (*script.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Script) InvokeScript(
    request script.InvokeScriptRequest,
) (*script.InvokeScriptResult, error) {
    r, err := p.client.InvokeScript(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Script) DebugInvoke(
    request script.DebugInvokeRequest,
) (*script.DebugInvokeResult, error) {
    r, err := p.client.DebugInvoke(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Script) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Script) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
