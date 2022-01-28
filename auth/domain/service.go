package domainauth
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

deny overwrite
*/

import (
    "github.com/gs2io/gs2-golang-sdk/account/domain/iterator"
    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/core"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Auth struct {
    session *core.Gs2RestSession
    client auth.Gs2AuthRestClient
}

func NewGs2Auth(
    session *core.Gs2RestSession,
) Gs2Auth {
    return Gs2Auth {
        session: session,
        client: auth.Gs2AuthRestClient {
            Session: session,
        },
    }
}

func (p Gs2Auth) Login(
    request auth.LoginRequest,
) (*auth.LoginResult, error) {
    r, err := p.client.Login(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Auth) LoginBySignature(
    request auth.LoginBySignatureRequest,
) (*auth.LoginBySignatureResult, error) {
    r, err := p.client.LoginBySignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
