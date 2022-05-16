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

package identifier

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeUsersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeUsersRequestFromJson(data string) DescribeUsersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeUsersRequestFromDict(dict)
}

func NewDescribeUsersRequestFromDict(data map[string]interface{}) DescribeUsersRequest {
    return DescribeUsersRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeUsersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeUsersRequest) Pointer() *DescribeUsersRequest {
    return &p
}

type CreateUserRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
}

func NewCreateUserRequestFromJson(data string) CreateUserRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateUserRequestFromDict(dict)
}

func NewCreateUserRequestFromDict(data map[string]interface{}) CreateUserRequest {
    return CreateUserRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
    }
}

func (p CreateUserRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
    }
}

func (p CreateUserRequest) Pointer() *CreateUserRequest {
    return &p
}

type UpdateUserRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    Description *string `json:"description"`
}

func NewUpdateUserRequestFromJson(data string) UpdateUserRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateUserRequestFromDict(dict)
}

func NewUpdateUserRequestFromDict(data map[string]interface{}) UpdateUserRequest {
    return UpdateUserRequest {
        UserName: core.CastString(data["userName"]),
        Description: core.CastString(data["description"]),
    }
}

func (p UpdateUserRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "description": p.Description,
    }
}

func (p UpdateUserRequest) Pointer() *UpdateUserRequest {
    return &p
}

type GetUserRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewGetUserRequestFromJson(data string) GetUserRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetUserRequestFromDict(dict)
}

func NewGetUserRequestFromDict(data map[string]interface{}) GetUserRequest {
    return GetUserRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p GetUserRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p GetUserRequest) Pointer() *GetUserRequest {
    return &p
}

type DeleteUserRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewDeleteUserRequestFromJson(data string) DeleteUserRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteUserRequestFromDict(dict)
}

func NewDeleteUserRequestFromDict(data map[string]interface{}) DeleteUserRequest {
    return DeleteUserRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p DeleteUserRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p DeleteUserRequest) Pointer() *DeleteUserRequest {
    return &p
}

type DescribeSecurityPoliciesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeSecurityPoliciesRequestFromJson(data string) DescribeSecurityPoliciesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSecurityPoliciesRequestFromDict(dict)
}

func NewDescribeSecurityPoliciesRequestFromDict(data map[string]interface{}) DescribeSecurityPoliciesRequest {
    return DescribeSecurityPoliciesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeSecurityPoliciesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeSecurityPoliciesRequest) Pointer() *DescribeSecurityPoliciesRequest {
    return &p
}

type DescribeCommonSecurityPoliciesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCommonSecurityPoliciesRequestFromJson(data string) DescribeCommonSecurityPoliciesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCommonSecurityPoliciesRequestFromDict(dict)
}

func NewDescribeCommonSecurityPoliciesRequestFromDict(data map[string]interface{}) DescribeCommonSecurityPoliciesRequest {
    return DescribeCommonSecurityPoliciesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCommonSecurityPoliciesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCommonSecurityPoliciesRequest) Pointer() *DescribeCommonSecurityPoliciesRequest {
    return &p
}

type CreateSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Policy *string `json:"policy"`
}

func NewCreateSecurityPolicyRequestFromJson(data string) CreateSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSecurityPolicyRequestFromDict(dict)
}

func NewCreateSecurityPolicyRequestFromDict(data map[string]interface{}) CreateSecurityPolicyRequest {
    return CreateSecurityPolicyRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Policy: core.CastString(data["policy"]),
    }
}

func (p CreateSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "policy": p.Policy,
    }
}

func (p CreateSecurityPolicyRequest) Pointer() *CreateSecurityPolicyRequest {
    return &p
}

type UpdateSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    SecurityPolicyName *string `json:"securityPolicyName"`
    Description *string `json:"description"`
    Policy *string `json:"policy"`
}

func NewUpdateSecurityPolicyRequestFromJson(data string) UpdateSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSecurityPolicyRequestFromDict(dict)
}

func NewUpdateSecurityPolicyRequestFromDict(data map[string]interface{}) UpdateSecurityPolicyRequest {
    return UpdateSecurityPolicyRequest {
        SecurityPolicyName: core.CastString(data["securityPolicyName"]),
        Description: core.CastString(data["description"]),
        Policy: core.CastString(data["policy"]),
    }
}

func (p UpdateSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "securityPolicyName": p.SecurityPolicyName,
        "description": p.Description,
        "policy": p.Policy,
    }
}

func (p UpdateSecurityPolicyRequest) Pointer() *UpdateSecurityPolicyRequest {
    return &p
}

type GetSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    SecurityPolicyName *string `json:"securityPolicyName"`
}

func NewGetSecurityPolicyRequestFromJson(data string) GetSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSecurityPolicyRequestFromDict(dict)
}

func NewGetSecurityPolicyRequestFromDict(data map[string]interface{}) GetSecurityPolicyRequest {
    return GetSecurityPolicyRequest {
        SecurityPolicyName: core.CastString(data["securityPolicyName"]),
    }
}

func (p GetSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "securityPolicyName": p.SecurityPolicyName,
    }
}

func (p GetSecurityPolicyRequest) Pointer() *GetSecurityPolicyRequest {
    return &p
}

type DeleteSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    SecurityPolicyName *string `json:"securityPolicyName"`
}

func NewDeleteSecurityPolicyRequestFromJson(data string) DeleteSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSecurityPolicyRequestFromDict(dict)
}

func NewDeleteSecurityPolicyRequestFromDict(data map[string]interface{}) DeleteSecurityPolicyRequest {
    return DeleteSecurityPolicyRequest {
        SecurityPolicyName: core.CastString(data["securityPolicyName"]),
    }
}

func (p DeleteSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "securityPolicyName": p.SecurityPolicyName,
    }
}

func (p DeleteSecurityPolicyRequest) Pointer() *DeleteSecurityPolicyRequest {
    return &p
}

type DescribeIdentifiersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeIdentifiersRequestFromJson(data string) DescribeIdentifiersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeIdentifiersRequestFromDict(dict)
}

func NewDescribeIdentifiersRequestFromDict(data map[string]interface{}) DescribeIdentifiersRequest {
    return DescribeIdentifiersRequest {
        UserName: core.CastString(data["userName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeIdentifiersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeIdentifiersRequest) Pointer() *DescribeIdentifiersRequest {
    return &p
}

type CreateIdentifierRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewCreateIdentifierRequestFromJson(data string) CreateIdentifierRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateIdentifierRequestFromDict(dict)
}

func NewCreateIdentifierRequestFromDict(data map[string]interface{}) CreateIdentifierRequest {
    return CreateIdentifierRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p CreateIdentifierRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p CreateIdentifierRequest) Pointer() *CreateIdentifierRequest {
    return &p
}

type GetIdentifierRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    ClientId *string `json:"clientId"`
}

func NewGetIdentifierRequestFromJson(data string) GetIdentifierRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetIdentifierRequestFromDict(dict)
}

func NewGetIdentifierRequestFromDict(data map[string]interface{}) GetIdentifierRequest {
    return GetIdentifierRequest {
        UserName: core.CastString(data["userName"]),
        ClientId: core.CastString(data["clientId"]),
    }
}

func (p GetIdentifierRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "clientId": p.ClientId,
    }
}

func (p GetIdentifierRequest) Pointer() *GetIdentifierRequest {
    return &p
}

type DeleteIdentifierRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    ClientId *string `json:"clientId"`
}

func NewDeleteIdentifierRequestFromJson(data string) DeleteIdentifierRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteIdentifierRequestFromDict(dict)
}

func NewDeleteIdentifierRequestFromDict(data map[string]interface{}) DeleteIdentifierRequest {
    return DeleteIdentifierRequest {
        UserName: core.CastString(data["userName"]),
        ClientId: core.CastString(data["clientId"]),
    }
}

func (p DeleteIdentifierRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "clientId": p.ClientId,
    }
}

func (p DeleteIdentifierRequest) Pointer() *DeleteIdentifierRequest {
    return &p
}

type DescribePasswordsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribePasswordsRequestFromJson(data string) DescribePasswordsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribePasswordsRequestFromDict(dict)
}

func NewDescribePasswordsRequestFromDict(data map[string]interface{}) DescribePasswordsRequest {
    return DescribePasswordsRequest {
        UserName: core.CastString(data["userName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribePasswordsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribePasswordsRequest) Pointer() *DescribePasswordsRequest {
    return &p
}

type CreatePasswordRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    Password *string `json:"password"`
}

func NewCreatePasswordRequestFromJson(data string) CreatePasswordRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreatePasswordRequestFromDict(dict)
}

func NewCreatePasswordRequestFromDict(data map[string]interface{}) CreatePasswordRequest {
    return CreatePasswordRequest {
        UserName: core.CastString(data["userName"]),
        Password: core.CastString(data["password"]),
    }
}

func (p CreatePasswordRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "password": p.Password,
    }
}

func (p CreatePasswordRequest) Pointer() *CreatePasswordRequest {
    return &p
}

type GetPasswordRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewGetPasswordRequestFromJson(data string) GetPasswordRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetPasswordRequestFromDict(dict)
}

func NewGetPasswordRequestFromDict(data map[string]interface{}) GetPasswordRequest {
    return GetPasswordRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p GetPasswordRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p GetPasswordRequest) Pointer() *GetPasswordRequest {
    return &p
}

type DeletePasswordRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewDeletePasswordRequestFromJson(data string) DeletePasswordRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeletePasswordRequestFromDict(dict)
}

func NewDeletePasswordRequestFromDict(data map[string]interface{}) DeletePasswordRequest {
    return DeletePasswordRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p DeletePasswordRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p DeletePasswordRequest) Pointer() *DeletePasswordRequest {
    return &p
}

type GetHasSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
}

func NewGetHasSecurityPolicyRequestFromJson(data string) GetHasSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetHasSecurityPolicyRequestFromDict(dict)
}

func NewGetHasSecurityPolicyRequestFromDict(data map[string]interface{}) GetHasSecurityPolicyRequest {
    return GetHasSecurityPolicyRequest {
        UserName: core.CastString(data["userName"]),
    }
}

func (p GetHasSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
    }
}

func (p GetHasSecurityPolicyRequest) Pointer() *GetHasSecurityPolicyRequest {
    return &p
}

type AttachSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    SecurityPolicyId *string `json:"securityPolicyId"`
}

func NewAttachSecurityPolicyRequestFromJson(data string) AttachSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAttachSecurityPolicyRequestFromDict(dict)
}

func NewAttachSecurityPolicyRequestFromDict(data map[string]interface{}) AttachSecurityPolicyRequest {
    return AttachSecurityPolicyRequest {
        UserName: core.CastString(data["userName"]),
        SecurityPolicyId: core.CastString(data["securityPolicyId"]),
    }
}

func (p AttachSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "securityPolicyId": p.SecurityPolicyId,
    }
}

func (p AttachSecurityPolicyRequest) Pointer() *AttachSecurityPolicyRequest {
    return &p
}

type DetachSecurityPolicyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    SecurityPolicyId *string `json:"securityPolicyId"`
}

func NewDetachSecurityPolicyRequestFromJson(data string) DetachSecurityPolicyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDetachSecurityPolicyRequestFromDict(dict)
}

func NewDetachSecurityPolicyRequestFromDict(data map[string]interface{}) DetachSecurityPolicyRequest {
    return DetachSecurityPolicyRequest {
        UserName: core.CastString(data["userName"]),
        SecurityPolicyId: core.CastString(data["securityPolicyId"]),
    }
}

func (p DetachSecurityPolicyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "securityPolicyId": p.SecurityPolicyId,
    }
}

func (p DetachSecurityPolicyRequest) Pointer() *DetachSecurityPolicyRequest {
    return &p
}

type LoginRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    ClientId *string `json:"clientId"`
    ClientSecret *string `json:"clientSecret"`
}

func NewLoginRequestFromJson(data string) LoginRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLoginRequestFromDict(dict)
}

func NewLoginRequestFromDict(data map[string]interface{}) LoginRequest {
    return LoginRequest {
        ClientId: core.CastString(data["clientId"]),
        ClientSecret: core.CastString(data["clientSecret"]),
    }
}

func (p LoginRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "clientId": p.ClientId,
        "clientSecret": p.ClientSecret,
    }
}

func (p LoginRequest) Pointer() *LoginRequest {
    return &p
}

type LoginByUserRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    UserName *string `json:"userName"`
    Password *string `json:"password"`
}

func NewLoginByUserRequestFromJson(data string) LoginByUserRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLoginByUserRequestFromDict(dict)
}

func NewLoginByUserRequestFromDict(data map[string]interface{}) LoginByUserRequest {
    return LoginByUserRequest {
        UserName: core.CastString(data["userName"]),
        Password: core.CastString(data["password"]),
    }
}

func (p LoginByUserRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userName": p.UserName,
        "password": p.Password,
    }
}

func (p LoginByUserRequest) Pointer() *LoginByUserRequest {
    return &p
}