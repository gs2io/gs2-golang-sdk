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
    "github.com/gs2io/gs2-golang-sdk/experience"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/cache"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}


type DescribeExperienceModelsIterator struct {
    experienceModelCache cache.ExperienceModelDomainCache
    client experience.Gs2ExperienceRestClient
    namespaceName string

    index int64
    last bool
    result []experience.ExperienceModel

    FetchSize *int32
}

func NewDescribeExperienceModelsIterator(
    experienceModelCache cache.ExperienceModelDomainCache,
    client experience.Gs2ExperienceRestClient,
    namespaceName string,
) DescribeExperienceModelsIterator {
    it := DescribeExperienceModelsIterator{
        experienceModelCache: experienceModelCache,
        client: client,
        namespaceName: namespaceName,

        index: 0,
        last: false,
        result: make([]experience.ExperienceModel, 0),

        FetchSize: nil,
    }
    _ = it.load()
    return it
}

func (p *DescribeExperienceModelsIterator) load() error {
    r, err := p.client.DescribeExperienceModels(
        &experience.DescribeExperienceModelsRequest{
            NamespaceName: &p.namespaceName,
        },
    )
    if err != nil {
        return err
    }
    for _, item := range r.Items {
        p.experienceModelCache.Update(item)
    }
    p.result = r.Items
    p.last = true
    return nil
}

func (p *DescribeExperienceModelsIterator) HasNext(

) bool {
    return len(p.result) != 0 || !p.last
}

func (p *DescribeExperienceModelsIterator) Next(

) (experience.ExperienceModel, error) {
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.ExperienceModel{}, err
        }
    }
    if len(p.result) == 0 {
        return experience.ExperienceModel{}, errors.New("out of range")
    }
    ret := p.result[0]
    p.result = p.result[1:]
    if len(p.result) == 0 && !p.last {
        err := p.load()
        if err != nil {
            return experience.ExperienceModel{}, err
        }
    }
    return ret, nil
}
