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
    "github.com/gs2io/gs2-golang-sdk/news"
    "github.com/gs2io/gs2-golang-sdk/news/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeNewsIterator struct {
    newsCache cache.NewsDomainCache
    client news.Gs2NewsRestClient
    namespaceName string
    accessToken auth.AccessToken

    index int64
    last bool
    result []news.News

    FetchSize *int32
}

func NewDescribeNewsIterator(
    newsCache cache.NewsDomainCache,
    client news.Gs2NewsRestClient,
    namespaceName string,
    accessToken auth.AccessToken,
) DescribeNewsIterator {
    it := DescribeNewsIterator{
        newsCache: newsCache,
        client: client,
        namespaceName: namespaceName,
        accessToken: accessToken,

        index: 0,
        last: false,
        result: make([]news.News, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeNewsIterator) load() error {
    r, err := p.client.DescribeNews(
        &news.DescribeNewsRequest{
            NamespaceName: &p.namespaceName,
            AccessToken: p.accessToken.Token,
        },
    )
    if err != nil {
        return err
    }
    p.newsCache.Update(r.Items)
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeNewsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeNewsIterator) Next(

) (news.News, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return news.News{}, err
        }
    }
    if len(p.result) == 0 {
        return news.News{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return news.News{}, err
        }
    }
    return ret, nil
}
