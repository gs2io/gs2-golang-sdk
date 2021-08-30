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

import "github.com/gs2io/gs2-golang-sdk/core"

type DescribeUsersResult struct {
	Items         []User  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeUsersAsyncResult struct {
	result *DescribeUsersResult
	err    error
}

func NewDescribeUsersResultFromDict(data map[string]interface{}) DescribeUsersResult {
	return DescribeUsersResult{
		Items:         CastUsers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeUsersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastUsersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeUsersResult) Pointer() *DescribeUsersResult {
	return &p
}

type CreateUserResult struct {
	Item *User `json:"item"`
}

type CreateUserAsyncResult struct {
	result *CreateUserResult
	err    error
}

func NewCreateUserResultFromDict(data map[string]interface{}) CreateUserResult {
	return CreateUserResult{
		Item: NewUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateUserResult) Pointer() *CreateUserResult {
	return &p
}

type UpdateUserResult struct {
	Item *User `json:"item"`
}

type UpdateUserAsyncResult struct {
	result *UpdateUserResult
	err    error
}

func NewUpdateUserResultFromDict(data map[string]interface{}) UpdateUserResult {
	return UpdateUserResult{
		Item: NewUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateUserResult) Pointer() *UpdateUserResult {
	return &p
}

type GetUserResult struct {
	Item *User `json:"item"`
}

type GetUserAsyncResult struct {
	result *GetUserResult
	err    error
}

func NewGetUserResultFromDict(data map[string]interface{}) GetUserResult {
	return GetUserResult{
		Item: NewUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetUserResult) Pointer() *GetUserResult {
	return &p
}

type DeleteUserResult struct {
}

type DeleteUserAsyncResult struct {
	result *DeleteUserResult
	err    error
}

func NewDeleteUserResultFromDict(data map[string]interface{}) DeleteUserResult {
	return DeleteUserResult{}
}

func (p DeleteUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteUserResult) Pointer() *DeleteUserResult {
	return &p
}

type DescribeSecurityPoliciesResult struct {
	Items         []SecurityPolicy `json:"items"`
	NextPageToken *string          `json:"nextPageToken"`
}

type DescribeSecurityPoliciesAsyncResult struct {
	result *DescribeSecurityPoliciesResult
	err    error
}

func NewDescribeSecurityPoliciesResultFromDict(data map[string]interface{}) DescribeSecurityPoliciesResult {
	return DescribeSecurityPoliciesResult{
		Items:         CastSecurityPolicies(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeSecurityPoliciesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSecurityPoliciesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSecurityPoliciesResult) Pointer() *DescribeSecurityPoliciesResult {
	return &p
}

type DescribeCommonSecurityPoliciesResult struct {
	Items         []SecurityPolicy `json:"items"`
	NextPageToken *string          `json:"nextPageToken"`
}

type DescribeCommonSecurityPoliciesAsyncResult struct {
	result *DescribeCommonSecurityPoliciesResult
	err    error
}

func NewDescribeCommonSecurityPoliciesResultFromDict(data map[string]interface{}) DescribeCommonSecurityPoliciesResult {
	return DescribeCommonSecurityPoliciesResult{
		Items:         CastSecurityPolicies(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCommonSecurityPoliciesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSecurityPoliciesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCommonSecurityPoliciesResult) Pointer() *DescribeCommonSecurityPoliciesResult {
	return &p
}

type CreateSecurityPolicyResult struct {
	Item *SecurityPolicy `json:"item"`
}

type CreateSecurityPolicyAsyncResult struct {
	result *CreateSecurityPolicyResult
	err    error
}

func NewCreateSecurityPolicyResultFromDict(data map[string]interface{}) CreateSecurityPolicyResult {
	return CreateSecurityPolicyResult{
		Item: NewSecurityPolicyFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateSecurityPolicyResult) Pointer() *CreateSecurityPolicyResult {
	return &p
}

type UpdateSecurityPolicyResult struct {
	Item *SecurityPolicy `json:"item"`
}

type UpdateSecurityPolicyAsyncResult struct {
	result *UpdateSecurityPolicyResult
	err    error
}

func NewUpdateSecurityPolicyResultFromDict(data map[string]interface{}) UpdateSecurityPolicyResult {
	return UpdateSecurityPolicyResult{
		Item: NewSecurityPolicyFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateSecurityPolicyResult) Pointer() *UpdateSecurityPolicyResult {
	return &p
}

type GetSecurityPolicyResult struct {
	Item *SecurityPolicy `json:"item"`
}

type GetSecurityPolicyAsyncResult struct {
	result *GetSecurityPolicyResult
	err    error
}

func NewGetSecurityPolicyResultFromDict(data map[string]interface{}) GetSecurityPolicyResult {
	return GetSecurityPolicyResult{
		Item: NewSecurityPolicyFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetSecurityPolicyResult) Pointer() *GetSecurityPolicyResult {
	return &p
}

type DeleteSecurityPolicyResult struct {
}

type DeleteSecurityPolicyAsyncResult struct {
	result *DeleteSecurityPolicyResult
	err    error
}

func NewDeleteSecurityPolicyResultFromDict(data map[string]interface{}) DeleteSecurityPolicyResult {
	return DeleteSecurityPolicyResult{}
}

func (p DeleteSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteSecurityPolicyResult) Pointer() *DeleteSecurityPolicyResult {
	return &p
}

type DescribeIdentifiersResult struct {
	Items         []Identifier `json:"items"`
	NextPageToken *string      `json:"nextPageToken"`
}

type DescribeIdentifiersAsyncResult struct {
	result *DescribeIdentifiersResult
	err    error
}

func NewDescribeIdentifiersResultFromDict(data map[string]interface{}) DescribeIdentifiersResult {
	return DescribeIdentifiersResult{
		Items:         CastIdentifiers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeIdentifiersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIdentifiersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeIdentifiersResult) Pointer() *DescribeIdentifiersResult {
	return &p
}

type CreateIdentifierResult struct {
	Item         *Identifier `json:"item"`
	ClientSecret *string     `json:"clientSecret"`
}

type CreateIdentifierAsyncResult struct {
	result *CreateIdentifierResult
	err    error
}

func NewCreateIdentifierResultFromDict(data map[string]interface{}) CreateIdentifierResult {
	return CreateIdentifierResult{
		Item:         NewIdentifierFromDict(core.CastMap(data["item"])).Pointer(),
		ClientSecret: core.CastString(data["clientSecret"]),
	}
}

func (p CreateIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"clientSecret": p.ClientSecret,
	}
}

func (p CreateIdentifierResult) Pointer() *CreateIdentifierResult {
	return &p
}

type GetIdentifierResult struct {
	Item *Identifier `json:"item"`
}

type GetIdentifierAsyncResult struct {
	result *GetIdentifierResult
	err    error
}

func NewGetIdentifierResultFromDict(data map[string]interface{}) GetIdentifierResult {
	return GetIdentifierResult{
		Item: NewIdentifierFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetIdentifierResult) Pointer() *GetIdentifierResult {
	return &p
}

type DeleteIdentifierResult struct {
}

type DeleteIdentifierAsyncResult struct {
	result *DeleteIdentifierResult
	err    error
}

func NewDeleteIdentifierResultFromDict(data map[string]interface{}) DeleteIdentifierResult {
	return DeleteIdentifierResult{}
}

func (p DeleteIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteIdentifierResult) Pointer() *DeleteIdentifierResult {
	return &p
}

type DescribePasswordsResult struct {
	Items         []Password `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribePasswordsAsyncResult struct {
	result *DescribePasswordsResult
	err    error
}

func NewDescribePasswordsResultFromDict(data map[string]interface{}) DescribePasswordsResult {
	return DescribePasswordsResult{
		Items:         CastPasswords(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribePasswordsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPasswordsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePasswordsResult) Pointer() *DescribePasswordsResult {
	return &p
}

type CreatePasswordResult struct {
	Item *Password `json:"item"`
}

type CreatePasswordAsyncResult struct {
	result *CreatePasswordResult
	err    error
}

func NewCreatePasswordResultFromDict(data map[string]interface{}) CreatePasswordResult {
	return CreatePasswordResult{
		Item: NewPasswordFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreatePasswordResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreatePasswordResult) Pointer() *CreatePasswordResult {
	return &p
}

type GetPasswordResult struct {
	Item *Password `json:"item"`
}

type GetPasswordAsyncResult struct {
	result *GetPasswordResult
	err    error
}

func NewGetPasswordResultFromDict(data map[string]interface{}) GetPasswordResult {
	return GetPasswordResult{
		Item: NewPasswordFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetPasswordResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetPasswordResult) Pointer() *GetPasswordResult {
	return &p
}

type DeletePasswordResult struct {
}

type DeletePasswordAsyncResult struct {
	result *DeletePasswordResult
	err    error
}

func NewDeletePasswordResultFromDict(data map[string]interface{}) DeletePasswordResult {
	return DeletePasswordResult{}
}

func (p DeletePasswordResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeletePasswordResult) Pointer() *DeletePasswordResult {
	return &p
}

type GetHasSecurityPolicyResult struct {
	Items []SecurityPolicy `json:"items"`
}

type GetHasSecurityPolicyAsyncResult struct {
	result *GetHasSecurityPolicyResult
	err    error
}

func NewGetHasSecurityPolicyResultFromDict(data map[string]interface{}) GetHasSecurityPolicyResult {
	return GetHasSecurityPolicyResult{
		Items: CastSecurityPolicies(core.CastArray(data["items"])),
	}
}

func (p GetHasSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSecurityPoliciesFromDict(
			p.Items,
		),
	}
}

func (p GetHasSecurityPolicyResult) Pointer() *GetHasSecurityPolicyResult {
	return &p
}

type AttachSecurityPolicyResult struct {
	Items []SecurityPolicy `json:"items"`
}

type AttachSecurityPolicyAsyncResult struct {
	result *AttachSecurityPolicyResult
	err    error
}

func NewAttachSecurityPolicyResultFromDict(data map[string]interface{}) AttachSecurityPolicyResult {
	return AttachSecurityPolicyResult{
		Items: CastSecurityPolicies(core.CastArray(data["items"])),
	}
}

func (p AttachSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSecurityPoliciesFromDict(
			p.Items,
		),
	}
}

func (p AttachSecurityPolicyResult) Pointer() *AttachSecurityPolicyResult {
	return &p
}

type DetachSecurityPolicyResult struct {
	Items []SecurityPolicy `json:"items"`
}

type DetachSecurityPolicyAsyncResult struct {
	result *DetachSecurityPolicyResult
	err    error
}

func NewDetachSecurityPolicyResultFromDict(data map[string]interface{}) DetachSecurityPolicyResult {
	return DetachSecurityPolicyResult{
		Items: CastSecurityPolicies(core.CastArray(data["items"])),
	}
}

func (p DetachSecurityPolicyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSecurityPoliciesFromDict(
			p.Items,
		),
	}
}

func (p DetachSecurityPolicyResult) Pointer() *DetachSecurityPolicyResult {
	return &p
}

type LoginResult struct {
	AccessToken *string `json:"accessToken"`
	TokenType   *string `json:"tokenType"`
	ExpiresIn   *int32  `json:"expiresIn"`
}

type LoginAsyncResult struct {
	result *LoginResult
	err    error
}

func NewLoginResultFromDict(data map[string]interface{}) LoginResult {
	return LoginResult{
		AccessToken: core.CastString(data["accessToken"]),
		TokenType:   core.CastString(data["tokenType"]),
		ExpiresIn:   core.CastInt32(data["expiresIn"]),
	}
}

func (p LoginResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accessToken": p.AccessToken,
		"tokenType":   p.TokenType,
		"expiresIn":   p.ExpiresIn,
	}
}

func (p LoginResult) Pointer() *LoginResult {
	return &p
}

type LoginByUserResult struct {
	Item *ProjectToken `json:"item"`
}

type LoginByUserAsyncResult struct {
	result *LoginByUserResult
	err    error
}

func NewLoginByUserResultFromDict(data map[string]interface{}) LoginByUserResult {
	return LoginByUserResult{
		Item: NewProjectTokenFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p LoginByUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p LoginByUserResult) Pointer() *LoginByUserResult {
	return &p
}
