package iterator
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
    "github.com/pkg/errors"

    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/jobQueue"
    "github.com/gs2io/gs2-golang-sdk/jobQueue/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeJobsByUserIdIterator struct {
    jobCache cache.JobDomainCache
    client job_queue.Gs2JobQueueRestClient
    namespaceName string
    userId string

    index int64
    pageToken *string
    last bool
    result []job_queue.Job

    FetchSize *int32
}

func NewDescribeJobsByUserIdIterator(
    jobCache cache.JobDomainCache,
    client job_queue.Gs2JobQueueRestClient,
    namespaceName string,
    userId string,
) DescribeJobsByUserIdIterator {
    it := DescribeJobsByUserIdIterator{
        jobCache: jobCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]job_queue.Job, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeJobsByUserIdIterator) load() error {
    r, err := p.client.DescribeJobsByUserId(
        &job_queue.DescribeJobsByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.jobCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeJobsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeJobsByUserIdIterator) Next(

) (job_queue.Job, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return job_queue.Job{}, err
        }
    }
    if len(p.result) == 0 {
        return job_queue.Job{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return job_queue.Job{}, err
        }
    }
    return ret, nil
}
