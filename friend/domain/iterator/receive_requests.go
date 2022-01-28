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


type DescribeReceiveRequestsIterator struct {
    friendRequestCache cache.FriendRequestDomainCache
    client friend.Gs2FriendRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []friend.FriendRequest

    FetchSize *int32
}

func NewDescribeReceiveRequestsIterator(
    friendRequestCache cache.FriendRequestDomainCache,
    client friend.Gs2FriendRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeReceiveRequestsIterator {
    it := DescribeReceiveRequestsIterator{
        friendRequestCache: friendRequestCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]friend.FriendRequest, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeReceiveRequestsIterator) load() error {
    r, err := p.client.DescribeReceiveRequests(
        &friend.DescribeReceiveRequestsRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.friendRequestCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeReceiveRequestsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeReceiveRequestsIterator) Next(

) (friend.FriendRequest, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FriendRequest{}, err
        }
    }
    if len(p.result) == 0 {
        return friend.FriendRequest{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FriendRequest{}, err
        }
    }
    return ret, nil
}
