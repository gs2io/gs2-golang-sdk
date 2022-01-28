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
    "github.com/gs2io/gs2-golang-sdk/identifier"
    "github.com/gs2io/gs2-golang-sdk/identifier/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.UserDomainCache{}


type DescribeUsersIterator struct {
    userCache cache.UserDomainCache
    client identifier.Gs2IdentifierRestClient

    index int64
    pageToken *string
    last bool
    result []identifier.User

    FetchSize *int32
}

func NewDescribeUsersIterator(
    userCache cache.UserDomainCache,
    client identifier.Gs2IdentifierRestClient,
) DescribeUsersIterator {
    it := DescribeUsersIterator{
        userCache: userCache,
        client: client,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]identifier.User, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeUsersIterator) load() error {
    r, err := p.client.DescribeUsers(
        &identifier.DescribeUsersRequest{
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.userCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeUsersIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeUsersIterator) Next(

) (identifier.User, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return identifier.User{}, err
        }
    }
    if len(p.result) == 0 {
        return identifier.User{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return identifier.User{}, err
        }
    }
    return ret, nil
}
