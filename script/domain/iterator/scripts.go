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
    "github.com/gs2io/gs2-golang-sdk/script"
    "github.com/gs2io/gs2-golang-sdk/script/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeScriptsIterator struct {
    scriptCache cache.ScriptDomainCache
    client script.Gs2ScriptRestClient
    namespaceName string

    index int64
    pageToken *string
    last bool
    result []script.Script

    FetchSize *int32
}

func NewDescribeScriptsIterator(
    scriptCache cache.ScriptDomainCache,
    client script.Gs2ScriptRestClient,
    namespaceName string,
) DescribeScriptsIterator {
    it := DescribeScriptsIterator{
        scriptCache: scriptCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]script.Script, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeScriptsIterator) load() error {
    r, err := p.client.DescribeScripts(
        &script.DescribeScriptsRequest{
            NamespaceName: &p.namespaceName,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.scriptCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeScriptsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeScriptsIterator) Next(

) (script.Script, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return script.Script{}, err
        }
    }
    if len(p.result) == 0 {
        return script.Script{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return script.Script{}, err
        }
    }
    return ret, nil
}
