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
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type DeadLetterJobDomain struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    deadLetterJobCache cache.DeadLetterJobDomainCache
    namespaceName string
    userId string
    deadLetterJobName string
}

func NewDeadLetterJobDomain(
    session *core.Gs2RestSession,
    deadLetterJobCache cache.DeadLetterJobDomainCache,
    namespaceName string,
    userId string,
    deadLetterJobName string,
) DeadLetterJobDomain {
    return DeadLetterJobDomain {
        session: session,
        client: job_queue.Gs2JobQueueRestClient{
            Session: session,
        },
        deadLetterJobCache: deadLetterJobCache,
        namespaceName: namespaceName,
        userId: userId,
        deadLetterJobName: deadLetterJobName,
    }
}

func (p DeadLetterJobDomain) Load(
    request job_queue.GetDeadLetterJobByUserIdRequest,
) (*job_queue.GetDeadLetterJobByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.DeadLetterJobName = &p.deadLetterJobName
    r, err := p.client.GetDeadLetterJobByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.deadLetterJobCache.Update(*r.Item)
    return r, nil
}

func (p DeadLetterJobDomain) Delete(
    request job_queue.DeleteDeadLetterJobByUserIdRequest,
) (*job_queue.DeleteDeadLetterJobByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.DeadLetterJobName = &p.deadLetterJobName
    r, err := p.client.DeleteDeadLetterJobByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.deadLetterJobCache.Delete(*r.Item)
    return r, nil
}
