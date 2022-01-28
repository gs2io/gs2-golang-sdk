package domainkey
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

    "github.com/gs2io/gs2-golang-sdk/key"
    "github.com/gs2io/gs2-golang-sdk/key/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/key/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type KeyDomain struct {
    session *core.Gs2RestSession
    client key.Gs2KeyRestClient
    keyCache cache.KeyDomainCache
    namespaceName string
    keyName string
}

func NewKeyDomain(
    session *core.Gs2RestSession,
    keyCache cache.KeyDomainCache,
    namespaceName string,
    keyName string,
) KeyDomain {
    return KeyDomain {
        session: session,
        client: key.Gs2KeyRestClient{
            Session: session,
        },
        keyCache: keyCache,
        namespaceName: namespaceName,
        keyName: keyName,
    }
}

func (p KeyDomain) Update(
    request key.UpdateKeyRequest,
) (*key.UpdateKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.KeyName = &p.keyName
    r, err := p.client.UpdateKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.keyCache.Update(*r.Item)
    return r, nil
}

func (p KeyDomain) Load(
    request key.GetKeyRequest,
) (*key.GetKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.KeyName = &p.keyName
    r, err := p.client.GetKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.keyCache.Update(*r.Item)
    return r, nil
}

func (p KeyDomain) Delete(
    request key.DeleteKeyRequest,
) (*key.DeleteKeyResult, error) {
    request.NamespaceName = &p.namespaceName
    request.KeyName = &p.keyName
    r, err := p.client.DeleteKey(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.keyCache.Delete(*r.Item)
    return r, nil
}

func (p KeyDomain) Encrypt(
    request key.EncryptRequest,
) (*key.EncryptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.KeyName = &p.keyName
    r, err := p.client.Encrypt(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p KeyDomain) Decrypt(
    request key.DecryptRequest,
) (*key.DecryptResult, error) {
    request.NamespaceName = &p.namespaceName
    request.KeyName = &p.keyName
    r, err := p.client.Decrypt(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}
