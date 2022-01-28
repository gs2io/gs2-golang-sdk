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
    "github.com/gs2io/gs2-golang-sdk/friend"
    "github.com/gs2io/gs2-golang-sdk/friend/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeFollowsByUserIdIterator struct {
    followUserCache cache.FollowUserDomainCache
    client friend.Gs2FriendRestClient
    namespaceName string
    userId string
    withProfile *bool

    index int64
    pageToken *string
    last bool
    result []friend.FollowUser

    FetchSize *int32
}

func NewDescribeFollowsByUserIdIterator(
    followUserCache cache.FollowUserDomainCache,
    client friend.Gs2FriendRestClient,
    namespaceName string,
    userId string,
    withProfile *bool,
) DescribeFollowsByUserIdIterator {
    it := DescribeFollowsByUserIdIterator{
        followUserCache: followUserCache,
        client: client,
        namespaceName: namespaceName,
        userId: userId,
        withProfile: withProfile,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]friend.FollowUser, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeFollowsByUserIdIterator) load() error {
    r, err := p.client.DescribeFollowsByUserId(
        &friend.DescribeFollowsByUserIdRequest{
            NamespaceName: &p.namespaceName,
            UserId: &p.userId,
            WithProfile: p.withProfile,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.followUserCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeFollowsByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeFollowsByUserIdIterator) Next(

) (friend.FollowUser, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FollowUser{}, err
        }
    }
    if len(p.result) == 0 {
        return friend.FollowUser{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FollowUser{}, err
        }
    }
    return ret, nil
}
