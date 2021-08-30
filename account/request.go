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

package account

import "core"

type DescribeNamespacesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	RequestId                *string        `json:"requestId"`
	ContextStack             *string        `json:"contextStack"`
	DuplicationAvoider       *string        `json:"duplicationAvoider"`
	Name                     *string        `json:"name"`
	Description              *string        `json:"description"`
	ChangePasswordIfTakeOver *bool          `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting `json:"doTakeOverScript"`
	LogSetting               *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		ChangePasswordIfTakeOver: core.CastBool(data["changePasswordIfTakeOver"]),
		CreateAccountScript:      NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer(),
		AuthenticationScript:     NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer(),
		CreateTakeOverScript:     NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer(),
		DoTakeOverScript:         NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                     p.Name,
		"description":              p.Description,
		"changePasswordIfTakeOver": p.ChangePasswordIfTakeOver,
		"createAccountScript":      p.CreateAccountScript.ToDict(),
		"authenticationScript":     p.AuthenticationScript.ToDict(),
		"createTakeOverScript":     p.CreateTakeOverScript.ToDict(),
		"doTakeOverScript":         p.DoTakeOverScript.ToDict(),
		"logSetting":               p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	RequestId                *string        `json:"requestId"`
	ContextStack             *string        `json:"contextStack"`
	DuplicationAvoider       *string        `json:"duplicationAvoider"`
	NamespaceName            *string        `json:"namespaceName"`
	Description              *string        `json:"description"`
	ChangePasswordIfTakeOver *bool          `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting `json:"doTakeOverScript"`
	LogSetting               *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		Description:              core.CastString(data["description"]),
		ChangePasswordIfTakeOver: core.CastBool(data["changePasswordIfTakeOver"]),
		CreateAccountScript:      NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer(),
		AuthenticationScript:     NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer(),
		CreateTakeOverScript:     NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer(),
		DoTakeOverScript:         NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"description":              p.Description,
		"changePasswordIfTakeOver": p.ChangePasswordIfTakeOver,
		"createAccountScript":      p.CreateAccountScript.ToDict(),
		"authenticationScript":     p.AuthenticationScript.ToDict(),
		"createTakeOverScript":     p.CreateTakeOverScript.ToDict(),
		"doTakeOverScript":         p.DoTakeOverScript.ToDict(),
		"logSetting":               p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type DescribeAccountsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeAccountsRequestFromDict(data map[string]interface{}) DescribeAccountsRequest {
	return DescribeAccountsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAccountsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAccountsRequest) Pointer() *DescribeAccountsRequest {
	return &p
}

type CreateAccountRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewCreateAccountRequestFromDict(data map[string]interface{}) CreateAccountRequest {
	return CreateAccountRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p CreateAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p CreateAccountRequest) Pointer() *CreateAccountRequest {
	return &p
}

type UpdateTimeOffsetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffset         *int32  `json:"timeOffset"`
}

func NewUpdateTimeOffsetRequestFromDict(data map[string]interface{}) UpdateTimeOffsetRequest {
	return UpdateTimeOffsetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		TimeOffset:    core.CastInt32(data["timeOffset"]),
	}
}

func (p UpdateTimeOffsetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"timeOffset":    p.TimeOffset,
	}
}

func (p UpdateTimeOffsetRequest) Pointer() *UpdateTimeOffsetRequest {
	return &p
}

type GetAccountRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewGetAccountRequestFromDict(data map[string]interface{}) GetAccountRequest {
	return GetAccountRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p GetAccountRequest) Pointer() *GetAccountRequest {
	return &p
}

type DeleteAccountRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewDeleteAccountRequestFromDict(data map[string]interface{}) DeleteAccountRequest {
	return DeleteAccountRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p DeleteAccountRequest) Pointer() *DeleteAccountRequest {
	return &p
}

type AuthenticationRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	KeyId              *string `json:"keyId"`
	Password           *string `json:"password"`
}

func NewAuthenticationRequestFromDict(data map[string]interface{}) AuthenticationRequest {
	return AuthenticationRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		KeyId:         core.CastString(data["keyId"]),
		Password:      core.CastString(data["password"]),
	}
}

func (p AuthenticationRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"keyId":         p.KeyId,
		"password":      p.Password,
	}
}

func (p AuthenticationRequest) Pointer() *AuthenticationRequest {
	return &p
}

type DescribeTakeOversRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeTakeOversRequestFromDict(data map[string]interface{}) DescribeTakeOversRequest {
	return DescribeTakeOversRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeTakeOversRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeTakeOversRequest) Pointer() *DescribeTakeOversRequest {
	return &p
}

type DescribeTakeOversByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeTakeOversByUserIdRequestFromDict(data map[string]interface{}) DescribeTakeOversByUserIdRequest {
	return DescribeTakeOversByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeTakeOversByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeTakeOversByUserIdRequest) Pointer() *DescribeTakeOversByUserIdRequest {
	return &p
}

type CreateTakeOverRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
	Password           *string `json:"password"`
}

func NewCreateTakeOverRequestFromDict(data map[string]interface{}) CreateTakeOverRequest {
	return CreateTakeOverRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		Password:       core.CastString(data["password"]),
	}
}

func (p CreateTakeOverRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
		"password":       p.Password,
	}
}

func (p CreateTakeOverRequest) Pointer() *CreateTakeOverRequest {
	return &p
}

type CreateTakeOverByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
	Password           *string `json:"password"`
}

func NewCreateTakeOverByUserIdRequestFromDict(data map[string]interface{}) CreateTakeOverByUserIdRequest {
	return CreateTakeOverByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		Password:       core.CastString(data["password"]),
	}
}

func (p CreateTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
		"password":       p.Password,
	}
}

func (p CreateTakeOverByUserIdRequest) Pointer() *CreateTakeOverByUserIdRequest {
	return &p
}

type GetTakeOverRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
}

func NewGetTakeOverRequestFromDict(data map[string]interface{}) GetTakeOverRequest {
	return GetTakeOverRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Type:          core.CastInt32(data["type"]),
	}
}

func (p GetTakeOverRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"type":          p.Type,
	}
}

func (p GetTakeOverRequest) Pointer() *GetTakeOverRequest {
	return &p
}

type GetTakeOverByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
}

func NewGetTakeOverByUserIdRequestFromDict(data map[string]interface{}) GetTakeOverByUserIdRequest {
	return GetTakeOverByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Type:          core.CastInt32(data["type"]),
	}
}

func (p GetTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"type":          p.Type,
	}
}

func (p GetTakeOverByUserIdRequest) Pointer() *GetTakeOverByUserIdRequest {
	return &p
}

type UpdateTakeOverRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	OldPassword        *string `json:"oldPassword"`
	Password           *string `json:"password"`
}

func NewUpdateTakeOverRequestFromDict(data map[string]interface{}) UpdateTakeOverRequest {
	return UpdateTakeOverRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Type:          core.CastInt32(data["type"]),
		OldPassword:   core.CastString(data["oldPassword"]),
		Password:      core.CastString(data["password"]),
	}
}

func (p UpdateTakeOverRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"type":          p.Type,
		"oldPassword":   p.OldPassword,
		"password":      p.Password,
	}
}

func (p UpdateTakeOverRequest) Pointer() *UpdateTakeOverRequest {
	return &p
}

type UpdateTakeOverByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
	OldPassword        *string `json:"oldPassword"`
	Password           *string `json:"password"`
}

func NewUpdateTakeOverByUserIdRequestFromDict(data map[string]interface{}) UpdateTakeOverByUserIdRequest {
	return UpdateTakeOverByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Type:          core.CastInt32(data["type"]),
		OldPassword:   core.CastString(data["oldPassword"]),
		Password:      core.CastString(data["password"]),
	}
}

func (p UpdateTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"type":          p.Type,
		"oldPassword":   p.OldPassword,
		"password":      p.Password,
	}
}

func (p UpdateTakeOverByUserIdRequest) Pointer() *UpdateTakeOverByUserIdRequest {
	return &p
}

type DeleteTakeOverRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
}

func NewDeleteTakeOverRequestFromDict(data map[string]interface{}) DeleteTakeOverRequest {
	return DeleteTakeOverRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
	}
}

func (p DeleteTakeOverRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
	}
}

func (p DeleteTakeOverRequest) Pointer() *DeleteTakeOverRequest {
	return &p
}

type DeleteTakeOverByUserIdentifierRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
}

func NewDeleteTakeOverByUserIdentifierRequestFromDict(data map[string]interface{}) DeleteTakeOverByUserIdentifierRequest {
	return DeleteTakeOverByUserIdentifierRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
	}
}

func (p DeleteTakeOverByUserIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
	}
}

func (p DeleteTakeOverByUserIdentifierRequest) Pointer() *DeleteTakeOverByUserIdentifierRequest {
	return &p
}

type DoTakeOverRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
	Password           *string `json:"password"`
}

func NewDoTakeOverRequestFromDict(data map[string]interface{}) DoTakeOverRequest {
	return DoTakeOverRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		Password:       core.CastString(data["password"]),
	}
}

func (p DoTakeOverRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
		"password":       p.Password,
	}
}

func (p DoTakeOverRequest) Pointer() *DoTakeOverRequest {
	return &p
}
