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

type JobDomain struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    jobCache cache.JobDomainCache
    namespaceName string
    userId string
    jobName string
}

func NewJobDomain(
    session *core.Gs2RestSession,
    jobCache cache.JobDomainCache,
    namespaceName string,
    userId string,
    jobName string,
) JobDomain {
    return JobDomain {
        session: session,
        client: job_queue.Gs2JobQueueRestClient{
            Session: session,
        },
        jobCache: jobCache,
        namespaceName: namespaceName,
        userId: userId,
        jobName: jobName,
    }
}

func (p JobDomain) Load(
    request job_queue.GetJobByUserIdRequest,
) (*job_queue.GetJobByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.JobName = &p.jobName
    r, err := p.client.GetJobByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.jobCache.Update(*r.Item)
    return r, nil
}

func (p JobDomain) Delete(
    request job_queue.DeleteJobByUserIdRequest,
) (*job_queue.DeleteJobByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.JobName = &p.jobName
    r, err := p.client.DeleteJobByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.jobCache.Delete(*r.Item)
    return r, nil
}
