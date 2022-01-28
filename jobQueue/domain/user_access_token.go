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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    namespaceName string
    accessToken auth.AccessToken
    jobCache cache.JobDomainCache
    deadLetterJobCache cache.DeadLetterJobDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: job_queue.Gs2JobQueueRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        jobCache: cache.NewJobDomainCache(),
        deadLetterJobCache: cache.NewDeadLetterJobDomainCache(),
    }
}

func (p UserAccessTokenDomain) Run(
    request job_queue.RunRequest,
) (*job_queue.RunResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.Run(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Job(
    jobName string,
) JobAccessTokenDomain {
    return NewJobAccessTokenDomain(
        p.session,
        p.jobCache,
        p.namespaceName,
        p.accessToken,
        jobName,
    )
}

func (p UserAccessTokenDomain) DeadLetterJob(
    deadLetterJobName string,
) DeadLetterJobAccessTokenDomain {
    return NewDeadLetterJobAccessTokenDomain(
        p.session,
        p.deadLetterJobCache,
        p.namespaceName,
        p.accessToken,
        deadLetterJobName,
    )
}
