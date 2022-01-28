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

type UserDomain struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    namespaceName string
    userId string
    jobCache cache.JobDomainCache
    deadLetterJobCache cache.DeadLetterJobDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: job_queue.Gs2JobQueueRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        jobCache: cache.NewJobDomainCache(),
        deadLetterJobCache: cache.NewDeadLetterJobDomainCache(),
    }
}

func (p UserDomain) Push(
    request job_queue.PushByUserIdRequest,
) (*job_queue.PushByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.PushByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Run(
    request job_queue.RunByUserIdRequest,
) (*job_queue.RunByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.RunByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Jobs(
) iterator.DescribeJobsByUserIdIterator {
    return iterator.NewDescribeJobsByUserIdIterator(
        p.jobCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) DeadLetterJobs(
) iterator.DescribeDeadLetterJobsByUserIdIterator {
    return iterator.NewDescribeDeadLetterJobsByUserIdIterator(
        p.deadLetterJobCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Job(
    jobName string,
) JobDomain {
    return NewJobDomain(
        p.session,
        p.jobCache,
        p.namespaceName,
        p.userId,
        jobName,
    )
}

func (p UserDomain) DeadLetterJob(
    deadLetterJobName string,
) DeadLetterJobDomain {
    return NewDeadLetterJobDomain(
        p.session,
        p.deadLetterJobCache,
        p.namespaceName,
        p.userId,
        deadLetterJobName,
    )
}
