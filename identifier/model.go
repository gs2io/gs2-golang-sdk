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

type User struct {
	UserId *string `json:"userId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewUserFromJson(data string) User {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUserFromDict(dict)
}

func NewUserFromDict(data map[string]interface{}) User {
    return User {
        UserId: core.CastString(data["userId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p User) ToDict() map[string]interface{} {
    
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "userId": userId,
        "name": name,
        "description": description,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p User) Pointer() *User {
    return &p
}

func CastUsers(data []interface{}) []User {
	v := make([]User, 0)
	for _, d := range data {
		v = append(v, NewUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUsersFromDict(data []User) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type SecurityPolicy struct {
	SecurityPolicyId *string `json:"securityPolicyId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Policy *string `json:"policy"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewSecurityPolicyFromJson(data string) SecurityPolicy {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSecurityPolicyFromDict(dict)
}

func NewSecurityPolicyFromDict(data map[string]interface{}) SecurityPolicy {
    return SecurityPolicy {
        SecurityPolicyId: core.CastString(data["securityPolicyId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Policy: core.CastString(data["policy"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p SecurityPolicy) ToDict() map[string]interface{} {
    
    var securityPolicyId *string
    if p.SecurityPolicyId != nil {
        securityPolicyId = p.SecurityPolicyId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var policy *string
    if p.Policy != nil {
        policy = p.Policy
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "securityPolicyId": securityPolicyId,
        "name": name,
        "description": description,
        "policy": policy,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p SecurityPolicy) Pointer() *SecurityPolicy {
    return &p
}

func CastSecurityPolicies(data []interface{}) []SecurityPolicy {
	v := make([]SecurityPolicy, 0)
	for _, d := range data {
		v = append(v, NewSecurityPolicyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSecurityPoliciesFromDict(data []SecurityPolicy) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Identifier struct {
	ClientId *string `json:"clientId"`
	UserName *string `json:"userName"`
	ClientSecret *string `json:"clientSecret"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewIdentifierFromJson(data string) Identifier {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIdentifierFromDict(dict)
}

func NewIdentifierFromDict(data map[string]interface{}) Identifier {
    return Identifier {
        ClientId: core.CastString(data["clientId"]),
        UserName: core.CastString(data["userName"]),
        ClientSecret: core.CastString(data["clientSecret"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Identifier) ToDict() map[string]interface{} {
    
    var clientId *string
    if p.ClientId != nil {
        clientId = p.ClientId
    }
    var userName *string
    if p.UserName != nil {
        userName = p.UserName
    }
    var clientSecret *string
    if p.ClientSecret != nil {
        clientSecret = p.ClientSecret
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "clientId": clientId,
        "userName": userName,
        "clientSecret": clientSecret,
        "createdAt": createdAt,
    }
}

func (p Identifier) Pointer() *Identifier {
    return &p
}

func CastIdentifiers(data []interface{}) []Identifier {
	v := make([]Identifier, 0)
	for _, d := range data {
		v = append(v, NewIdentifierFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIdentifiersFromDict(data []Identifier) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Password struct {
	PasswordId *string `json:"passwordId"`
	UserId *string `json:"userId"`
	UserName *string `json:"userName"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewPasswordFromJson(data string) Password {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPasswordFromDict(dict)
}

func NewPasswordFromDict(data map[string]interface{}) Password {
    return Password {
        PasswordId: core.CastString(data["passwordId"]),
        UserId: core.CastString(data["userId"]),
        UserName: core.CastString(data["userName"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Password) ToDict() map[string]interface{} {
    
    var passwordId *string
    if p.PasswordId != nil {
        passwordId = p.PasswordId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var userName *string
    if p.UserName != nil {
        userName = p.UserName
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "passwordId": passwordId,
        "userId": userId,
        "userName": userName,
        "createdAt": createdAt,
    }
}

func (p Password) Pointer() *Password {
    return &p
}

func CastPasswords(data []interface{}) []Password {
	v := make([]Password, 0)
	for _, d := range data {
		v = append(v, NewPasswordFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPasswordsFromDict(data []Password) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type AttachSecurityPolicy struct {
	UserId *string `json:"userId"`
	SecurityPolicyIds []string `json:"securityPolicyIds"`
	AttachedAt *int64 `json:"attachedAt"`
}

func NewAttachSecurityPolicyFromJson(data string) AttachSecurityPolicy {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAttachSecurityPolicyFromDict(dict)
}

func NewAttachSecurityPolicyFromDict(data map[string]interface{}) AttachSecurityPolicy {
    return AttachSecurityPolicy {
        UserId: core.CastString(data["userId"]),
        SecurityPolicyIds: core.CastStrings(core.CastArray(data["securityPolicyIds"])),
        AttachedAt: core.CastInt64(data["attachedAt"]),
    }
}

func (p AttachSecurityPolicy) ToDict() map[string]interface{} {
    
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var securityPolicyIds []interface{}
    if p.SecurityPolicyIds != nil {
        securityPolicyIds = core.CastStringsFromDict(
            p.SecurityPolicyIds,
        )
    }
    var attachedAt *int64
    if p.AttachedAt != nil {
        attachedAt = p.AttachedAt
    }
    return map[string]interface{} {
        "userId": userId,
        "securityPolicyIds": securityPolicyIds,
        "attachedAt": attachedAt,
    }
}

func (p AttachSecurityPolicy) Pointer() *AttachSecurityPolicy {
    return &p
}

func CastAttachSecurityPolicies(data []interface{}) []AttachSecurityPolicy {
	v := make([]AttachSecurityPolicy, 0)
	for _, d := range data {
		v = append(v, NewAttachSecurityPolicyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttachSecurityPoliciesFromDict(data []AttachSecurityPolicy) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ProjectToken struct {
	Token *string `json:"token"`
}

func NewProjectTokenFromJson(data string) ProjectToken {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewProjectTokenFromDict(dict)
}

func NewProjectTokenFromDict(data map[string]interface{}) ProjectToken {
    return ProjectToken {
        Token: core.CastString(data["token"]),
    }
}

func (p ProjectToken) ToDict() map[string]interface{} {
    
    var token *string
    if p.Token != nil {
        token = p.Token
    }
    return map[string]interface{} {
        "token": token,
    }
}

func (p ProjectToken) Pointer() *ProjectToken {
    return &p
}

func CastProjectTokens(data []interface{}) []ProjectToken {
	v := make([]ProjectToken, 0)
	for _, d := range data {
		v = append(v, NewProjectTokenFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProjectTokensFromDict(data []ProjectToken) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}