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

type GitHubApiKeyDomain struct {
    session *core.Gs2RestSession
    client key.Gs2KeyRestClient
    gitHubApiKeyCache cache.GitHubApiKeyDomainCache
    namespaceName string
    apiKeyName string
}

func NewGitHubApiKeyDomain(
    session *core.Gs2RestSession,
    gitHubApiKeyCache cache.GitHubApiKeyDomainCache,
    namespaceName string,
    apiKeyName string,
) GitHubApiKeyDomain {
    return GitHubApiKeyDomain {
        session: session,
        client: key.Gs2KeyRestClient{
            Session: session,
        },
        gitHubApiKeyCache: gitHubApiKeyCache,
        namespaceName: namespaceName,
        apiKeyName: apiKeyName,
    }
}

func (p GitHubApiKeyDomain) Update(
    request key.UpdateGitHubApiKeyRequest,
) (*key.UpdateGitHubApiKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ApiKeyName = &p.apiKeyName
    r, err := p.client.UpdateGitHubApiKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gitHubApiKeyCache.Update(*r.Item)
    return r, nil
}

func (p GitHubApiKeyDomain) Load(
    request key.GetGitHubApiKeyRequest,
) (*key.GetGitHubApiKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ApiKeyName = &p.apiKeyName
    r, err := p.client.GetGitHubApiKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gitHubApiKeyCache.Update(*r.Item)
    return r, nil
}

func (p GitHubApiKeyDomain) Delete(
    request key.DeleteGitHubApiKeyRequest,
) (*key.DeleteGitHubApiKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.ApiKeyName = &p.apiKeyName
    r, err := p.client.DeleteGitHubApiKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.gitHubApiKeyCache.Delete(*r.Item)
    return r, nil
}
