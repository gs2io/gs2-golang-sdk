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


type DescribeFriendsIterator struct {
    friendUserCache cache.FriendUserDomainCache
    client friend.Gs2FriendRestClient
    namespaceName string
    accessToken auth.AccessToken
    withProfile *bool

    index int64
    pageToken *string
    last bool
    result []friend.FriendUser

    FetchSize *int32
}

func NewDescribeFriendsIterator(
    friendUserCache cache.FriendUserDomainCache,
    client friend.Gs2FriendRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
    withProfile *bool,
) DescribeFriendsIterator {
    it := DescribeFriendsIterator{
        friendUserCache: friendUserCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,
        withProfile: withProfile,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]friend.FriendUser, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeFriendsIterator) load() error {
    r, err := p.client.DescribeFriends(
        &friend.DescribeFriendsRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
            WithProfile: p.withProfile,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.friendUserCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeFriendsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeFriendsIterator) Next(

) (friend.FriendUser, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FriendUser{}, err
        }
    }
    if len(p.result) == 0 {
        return friend.FriendUser{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return friend.FriendUser{}, err
        }
    }
    return ret, nil
}
