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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	Name                     *string            `json:"name"`
	Description              *string            `json:"description"`
	ChangePasswordIfTakeOver *bool              `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting     `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting     `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting     `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting     `json:"doTakeOverScript"`
	LogSetting               *LogSetting        `json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	Description              *string            `json:"description"`
	ChangePasswordIfTakeOver *bool              `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting     `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting     `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting     `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting     `json:"doTakeOverScript"`
	LogSetting               *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeAccountsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateAccountRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateTimeOffsetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TimeOffset             *int32             `json:"timeOffset"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetAccountRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteAccountRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AuthenticationRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	KeyId                  *string            `json:"keyId"`
	Password               *string            `json:"password"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeTakeOversRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeTakeOversByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateTakeOverRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Type                   *int32             `json:"type"`
	UserIdentifier         *string            `json:"userIdentifier"`
	Password               *string            `json:"password"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type CreateTakeOverByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Type                   *int32             `json:"type"`
	UserIdentifier         *string            `json:"userIdentifier"`
	Password               *string            `json:"password"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetTakeOverRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Type                   *int32             `json:"type"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetTakeOverByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Type                   *int32             `json:"type"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UpdateTakeOverRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Type                   *int32             `json:"type"`
	OldPassword            *string            `json:"oldPassword"`
	Password               *string            `json:"password"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UpdateTakeOverByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Type                   *int32             `json:"type"`
	OldPassword            *string            `json:"oldPassword"`
	Password               *string            `json:"password"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteTakeOverRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Type                   *int32             `json:"type"`
	UserIdentifier         *string            `json:"userIdentifier"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteTakeOverByUserIdentifierRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Type                   *int32             `json:"type"`
	UserIdentifier         *string            `json:"userIdentifier"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DoTakeOverRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	Type           *int32             `json:"type"`
	UserIdentifier *string            `json:"userIdentifier"`
	Password       *string            `json:"password"`
}
