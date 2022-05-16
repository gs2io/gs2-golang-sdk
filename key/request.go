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

package key

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesRequestFromDict(dict)
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
    return DescribeNamespacesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
    return &p
}

type CreateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusRequestFromDict(dict)
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
    return GetNamespaceStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
    return &p
}

type GetNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceRequestFromDict(dict)
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
    return GetNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
    return &p
}

type UpdateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceRequestFromDict(dict)
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
    return DeleteNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
    return &p
}

type DescribeKeysRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeKeysRequestFromJson(data string) DescribeKeysRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeKeysRequestFromDict(dict)
}

func NewDescribeKeysRequestFromDict(data map[string]interface{}) DescribeKeysRequest {
    return DescribeKeysRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeKeysRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeKeysRequest) Pointer() *DescribeKeysRequest {
    return &p
}

type CreateKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
}

func NewCreateKeyRequestFromJson(data string) CreateKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateKeyRequestFromDict(dict)
}

func NewCreateKeyRequestFromDict(data map[string]interface{}) CreateKeyRequest {
    return CreateKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
    }
}

func (p CreateKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
    }
}

func (p CreateKeyRequest) Pointer() *CreateKeyRequest {
    return &p
}

type UpdateKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    KeyName *string `json:"keyName"`
    Description *string `json:"description"`
}

func NewUpdateKeyRequestFromJson(data string) UpdateKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateKeyRequestFromDict(dict)
}

func NewUpdateKeyRequestFromDict(data map[string]interface{}) UpdateKeyRequest {
    return UpdateKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        KeyName: core.CastString(data["keyName"]),
        Description: core.CastString(data["description"]),
    }
}

func (p UpdateKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "keyName": p.KeyName,
        "description": p.Description,
    }
}

func (p UpdateKeyRequest) Pointer() *UpdateKeyRequest {
    return &p
}

type GetKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    KeyName *string `json:"keyName"`
}

func NewGetKeyRequestFromJson(data string) GetKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetKeyRequestFromDict(dict)
}

func NewGetKeyRequestFromDict(data map[string]interface{}) GetKeyRequest {
    return GetKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        KeyName: core.CastString(data["keyName"]),
    }
}

func (p GetKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "keyName": p.KeyName,
    }
}

func (p GetKeyRequest) Pointer() *GetKeyRequest {
    return &p
}

type DeleteKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    KeyName *string `json:"keyName"`
}

func NewDeleteKeyRequestFromJson(data string) DeleteKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteKeyRequestFromDict(dict)
}

func NewDeleteKeyRequestFromDict(data map[string]interface{}) DeleteKeyRequest {
    return DeleteKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        KeyName: core.CastString(data["keyName"]),
    }
}

func (p DeleteKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "keyName": p.KeyName,
    }
}

func (p DeleteKeyRequest) Pointer() *DeleteKeyRequest {
    return &p
}

type EncryptRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    KeyName *string `json:"keyName"`
    Data *string `json:"data"`
}

func NewEncryptRequestFromJson(data string) EncryptRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEncryptRequestFromDict(dict)
}

func NewEncryptRequestFromDict(data map[string]interface{}) EncryptRequest {
    return EncryptRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        KeyName: core.CastString(data["keyName"]),
        Data: core.CastString(data["data"]),
    }
}

func (p EncryptRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "keyName": p.KeyName,
        "data": p.Data,
    }
}

func (p EncryptRequest) Pointer() *EncryptRequest {
    return &p
}

type DecryptRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    KeyName *string `json:"keyName"`
    Data *string `json:"data"`
}

func NewDecryptRequestFromJson(data string) DecryptRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDecryptRequestFromDict(dict)
}

func NewDecryptRequestFromDict(data map[string]interface{}) DecryptRequest {
    return DecryptRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        KeyName: core.CastString(data["keyName"]),
        Data: core.CastString(data["data"]),
    }
}

func (p DecryptRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "keyName": p.KeyName,
        "data": p.Data,
    }
}

func (p DecryptRequest) Pointer() *DecryptRequest {
    return &p
}

type DescribeGitHubApiKeysRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeGitHubApiKeysRequestFromJson(data string) DescribeGitHubApiKeysRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeGitHubApiKeysRequestFromDict(dict)
}

func NewDescribeGitHubApiKeysRequestFromDict(data map[string]interface{}) DescribeGitHubApiKeysRequest {
    return DescribeGitHubApiKeysRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeGitHubApiKeysRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeGitHubApiKeysRequest) Pointer() *DescribeGitHubApiKeysRequest {
    return &p
}

type CreateGitHubApiKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    ApiKey *string `json:"apiKey"`
    EncryptionKeyName *string `json:"encryptionKeyName"`
}

func NewCreateGitHubApiKeyRequestFromJson(data string) CreateGitHubApiKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateGitHubApiKeyRequestFromDict(dict)
}

func NewCreateGitHubApiKeyRequestFromDict(data map[string]interface{}) CreateGitHubApiKeyRequest {
    return CreateGitHubApiKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        ApiKey: core.CastString(data["apiKey"]),
        EncryptionKeyName: core.CastString(data["encryptionKeyName"]),
    }
}

func (p CreateGitHubApiKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "apiKey": p.ApiKey,
        "encryptionKeyName": p.EncryptionKeyName,
    }
}

func (p CreateGitHubApiKeyRequest) Pointer() *CreateGitHubApiKeyRequest {
    return &p
}

type UpdateGitHubApiKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ApiKeyName *string `json:"apiKeyName"`
    Description *string `json:"description"`
    ApiKey *string `json:"apiKey"`
    EncryptionKeyName *string `json:"encryptionKeyName"`
}

func NewUpdateGitHubApiKeyRequestFromJson(data string) UpdateGitHubApiKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateGitHubApiKeyRequestFromDict(dict)
}

func NewUpdateGitHubApiKeyRequestFromDict(data map[string]interface{}) UpdateGitHubApiKeyRequest {
    return UpdateGitHubApiKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ApiKeyName: core.CastString(data["apiKeyName"]),
        Description: core.CastString(data["description"]),
        ApiKey: core.CastString(data["apiKey"]),
        EncryptionKeyName: core.CastString(data["encryptionKeyName"]),
    }
}

func (p UpdateGitHubApiKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "apiKeyName": p.ApiKeyName,
        "description": p.Description,
        "apiKey": p.ApiKey,
        "encryptionKeyName": p.EncryptionKeyName,
    }
}

func (p UpdateGitHubApiKeyRequest) Pointer() *UpdateGitHubApiKeyRequest {
    return &p
}

type GetGitHubApiKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ApiKeyName *string `json:"apiKeyName"`
}

func NewGetGitHubApiKeyRequestFromJson(data string) GetGitHubApiKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetGitHubApiKeyRequestFromDict(dict)
}

func NewGetGitHubApiKeyRequestFromDict(data map[string]interface{}) GetGitHubApiKeyRequest {
    return GetGitHubApiKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ApiKeyName: core.CastString(data["apiKeyName"]),
    }
}

func (p GetGitHubApiKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "apiKeyName": p.ApiKeyName,
    }
}

func (p GetGitHubApiKeyRequest) Pointer() *GetGitHubApiKeyRequest {
    return &p
}

type DeleteGitHubApiKeyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ApiKeyName *string `json:"apiKeyName"`
}

func NewDeleteGitHubApiKeyRequestFromJson(data string) DeleteGitHubApiKeyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteGitHubApiKeyRequestFromDict(dict)
}

func NewDeleteGitHubApiKeyRequestFromDict(data map[string]interface{}) DeleteGitHubApiKeyRequest {
    return DeleteGitHubApiKeyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ApiKeyName: core.CastString(data["apiKeyName"]),
    }
}

func (p DeleteGitHubApiKeyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "apiKeyName": p.ApiKeyName,
    }
}

func (p DeleteGitHubApiKeyRequest) Pointer() *DeleteGitHubApiKeyRequest {
    return &p
}