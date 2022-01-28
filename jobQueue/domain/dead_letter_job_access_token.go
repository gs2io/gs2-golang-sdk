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

type DeadLetterJobAccessTokenDomain struct {
    session *core.Gs2RestSession
    client job_queue.Gs2JobQueueRestClient
    deadLetterJobCache cache.DeadLetterJobDomainCache
    namespaceName string
    accessToken auth.AccessToken
    deadLetterJobName string
}

func NewDeadLetterJobAccessTokenDomain(
    session *core.Gs2RestSession,
    deadLetterJobCache cache.DeadLetterJobDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    deadLetterJobName string,
) DeadLetterJobAccessTokenDomain {
    return DeadLetterJobAccessTokenDomain {
        session: session,
        client: job_queue.Gs2JobQueueRestClient{
            Session: session,
        },
        deadLetterJobCache: deadLetterJobCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        deadLetterJobName: deadLetterJobName,
    }
}