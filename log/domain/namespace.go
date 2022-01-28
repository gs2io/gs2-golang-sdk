package domainlog
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

    "github.com/gs2io/gs2-golang-sdk/log"
    "github.com/gs2io/gs2-golang-sdk/log/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/log/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client log.Gs2LogRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: log.Gs2LogRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
    }
}

func (p NamespaceDomain) GetStatus(
    request log.GetNamespaceStatusRequest,
) (*log.GetNamespaceStatusResult, error) {
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
    request log.GetNamespaceRequest,
) (*log.GetNamespaceResult, error) {
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
    request log.UpdateNamespaceRequest,
) (*log.UpdateNamespaceResult, error) {
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
    request log.DeleteNamespaceRequest,
) (*log.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) QueryAccessLog(
    request log.QueryAccessLogRequest,
) (*log.QueryAccessLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.QueryAccessLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CountAccessLog(
    request log.CountAccessLogRequest,
) (*log.CountAccessLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CountAccessLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) QueryIssueStampSheetLog(
    request log.QueryIssueStampSheetLogRequest,
) (*log.QueryIssueStampSheetLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.QueryIssueStampSheetLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CountIssueStampSheetLog(
    request log.CountIssueStampSheetLogRequest,
) (*log.CountIssueStampSheetLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CountIssueStampSheetLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) QueryExecuteStampSheetLog(
    request log.QueryExecuteStampSheetLogRequest,
) (*log.QueryExecuteStampSheetLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.QueryExecuteStampSheetLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CountExecuteStampSheetLog(
    request log.CountExecuteStampSheetLogRequest,
) (*log.CountExecuteStampSheetLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CountExecuteStampSheetLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) QueryExecuteStampTaskLog(
    request log.QueryExecuteStampTaskLogRequest,
) (*log.QueryExecuteStampTaskLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.QueryExecuteStampTaskLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CountExecuteStampTaskLog(
    request log.CountExecuteStampTaskLogRequest,
) (*log.CountExecuteStampTaskLogResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CountExecuteStampTaskLog(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
