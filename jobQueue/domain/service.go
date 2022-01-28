package domainjob_queue
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

    "github.com/gs2io/gs2-golang-sdk/jobQueue"
    "github.com/gs2io/gs2-golang-sdk/jobQueue/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/jobQueue/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2JobQueue struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2JobQueue(
    session *core.Gs2RestSession,
) Gs2JobQueue {
    return Gs2JobQueue {
        session: session,
        client: job_queue.Gs2JobQueueRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2JobQueue) CreateNamespace(
    request job_queue.CreateNamespaceRequest,
) (*job_queue.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2JobQueue) PushByStampSheet(
    request job_queue.PushByStampSheetRequest,
) (*job_queue.PushByStampSheetResult, error) {
    r, err := p.client.PushByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2JobQueue) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2JobQueue) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
