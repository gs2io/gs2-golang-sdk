package domainkey
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

    "github.com/gs2io/gs2-golang-sdk/key"
    "github.com/gs2io/gs2-golang-sdk/key/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/key/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client key.Gs2KeyRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    keyCache cache.KeyDomainCache
    gitHubApiKeyCache cache.GitHubApiKeyDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: key.Gs2KeyRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        keyCache: cache.NewKeyDomainCache(),
        gitHubApiKeyCache: cache.NewGitHubApiKeyDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request key.GetNamespaceStatusRequest,
) (*key.GetNamespaceStatusResult, error) {
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
    request key.GetNamespaceRequest,
) (*key.GetNamespaceResult, error) {
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
    request key.UpdateNamespaceRequest,
) (*key.UpdateNamespaceResult, error) {
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
    request key.DeleteNamespaceRequest,
) (*key.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateKey(
    request key.CreateKeyRequest,
) (*key.CreateKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateGitHubApiKey(
    request key.CreateGitHubApiKeyRequest,
) (*key.CreateGitHubApiKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateGitHubApiKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Keys(
) iterator.DescribeKeysIterator {
    return iterator.NewDescribeKeysIterator(
        p.keyCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) GitHubApiKeys(
) iterator.DescribeGitHubApiKeysIterator {
    return iterator.NewDescribeGitHubApiKeysIterator(
        p.gitHubApiKeyCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) Key(
    keyName string,
) KeyDomain {
    return NewKeyDomain(
        p.session,
        p.keyCache,
        p.namespaceName,
        keyName,
    )
}

func (p NamespaceDomain) GitHubApiKey(
    apiKeyName string,
) GitHubApiKeyDomain {
    return NewGitHubApiKeyDomain(
        p.session,
        p.gitHubApiKeyCache,
        p.namespaceName,
        apiKeyName,
    )
}
