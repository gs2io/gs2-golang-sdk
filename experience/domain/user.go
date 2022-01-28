package domainexperience
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
    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/core"

    "github.com/gs2io/gs2-golang-sdk/experience"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/experience/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client experience.Gs2ExperienceRestClient
    namespaceName string
    userId string
    statusCache cache.StatusDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: experience.Gs2ExperienceRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        statusCache: cache.NewStatusDomainCache(),
    }
}

func (p UserDomain) Statuses(
    experienceName *string,
) iterator.DescribeStatusesByUserIdIterator {
    return iterator.NewDescribeStatusesByUserIdIterator(
        p.statusCache,
        p.client,
        p.namespaceName,
        experienceName,
        p.userId,
    )
}

func (p UserDomain) Status(
    experienceName string,
    propertyId string,
) StatusDomain {
    return NewStatusDomain(
        p.session,
        p.statusCache,
        p.namespaceName,
        p.userId,
        experienceName,
        propertyId,
    )
}
