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


type DescribeBlackListIterator struct {
    client friend.Gs2FriendRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    last bool
    result []string

    FetchSize *int32
}

func NewDescribeBlackListIterator(
    client friend.Gs2FriendRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeBlackListIterator {
    it := DescribeBlackListIterator{
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        last: false,
        result: make([]string, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeBlackListIterator) load() error {
    r, err := p.client.DescribeBlackList(
        &friend.DescribeBlackListRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
        },
    )
    if err != nil {
        return err
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeBlackListIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeBlackListIterator) Next(

) (string, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return "", err
        }
    }
    if len(p.result) == 0 {
        return "", errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return "", err
        }
    }
    return ret, nil
}
