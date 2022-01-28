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
    "github.com/gs2io/gs2-golang-sdk/chat"
    "github.com/gs2io/gs2-golang-sdk/chat/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeMessagesByUserIdIterator struct {
    messageCache cache.MessageDomainCache
    client chat.Gs2ChatRestClient
    namespaceName string
    roomName string
    password *string
    userId string

    index int64
    startAt *int64
    last bool
    result []chat.Message

    FetchSize *int32
}

func NewDescribeMessagesByUserIdIterator(
    messageCache cache.MessageDomainCache,
    client chat.Gs2ChatRestClient,
    namespaceName string,
    roomName string,
    password *string,
    userId string,
) DescribeMessagesByUserIdIterator {
    it := DescribeMessagesByUserIdIterator{
        messageCache: messageCache,
        client: client,
        namespaceName: namespaceName,
        roomName: roomName,
        password: password,
        userId: userId,

        index: 0,
        startAt: nil,
        last: false,
        result: make([]chat.Message, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeMessagesByUserIdIterator) load() error {
    r, err := p.client.DescribeMessagesByUserId(
        &chat.DescribeMessagesByUserIdRequest{
            NamespaceName: &p.namespaceName,
            RoomName: &p.roomName,
            Password: p.password,
            UserId: &p.userId,
            StartAt: p.startAt,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.messageCache.Update(item)
    }
    p.result = r.Items
    if len(p.result) > 0 {
        nextCreatedAt := *p.result[len(p.result)-1].CreatedAt + 1
        p.startAt = &nextCreatedAt
    }
    return nil
}

func (p *DescribeMessagesByUserIdIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeMessagesByUserIdIterator) Next(

) (chat.Message, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return chat.Message{}, err
        }
    }
    if len(p.result) == 0 {
        return chat.Message{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return chat.Message{}, err
        }
    }
    return ret, nil
}
