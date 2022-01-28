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
    "github.com/gs2io/gs2-golang-sdk/formation"
    "github.com/gs2io/gs2-golang-sdk/formation/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeFormsIterator struct {
    formCache cache.FormDomainCache
    client formation.Gs2FormationRestClient
    namespaceName string
    moldName string
    accessToken auth.AccessToken

    index int64
    pageToken *string
    last bool
    result []formation.Form

    FetchSize *int32
}

func NewDescribeFormsIterator(
    formCache cache.FormDomainCache,
    client formation.Gs2FormationRestClient,
    namespaceName string,
    moldName string,
    accessToken auth.AccessToken,
) DescribeFormsIterator {
    it := DescribeFormsIterator{
        formCache: formCache,
        client: client,
        namespaceName: namespaceName,
        moldName: moldName,
        accessToken: accessToken,

        index: 0,
        pageToken: nil,
        last: false,
        result: make([]formation.Form, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeFormsIterator) load() error {
    r, err := p.client.DescribeForms(
        &formation.DescribeFormsRequest{
            NamespaceName: &p.namespaceName,
            MoldName: &p.moldName,
            AccessToken: p.accessToken.Token,
            PageToken: p.pageToken,
            Limit: p.FetchSize,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.formCache.Update(item)
    }
    p.result = r.Items
    p.pageToken = r.NextPageToken
    p.last = p.pageToken == nil
    return nil
}

func (p *DescribeFormsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeFormsIterator) Next(

) (formation.Form, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Form{}, err
        }
    }
    if len(p.result) == 0 {
        return formation.Form{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return formation.Form{}, err
        }
    }
    return ret, nil
}
