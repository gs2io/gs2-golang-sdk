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

package money

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
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	Name               *string            `json:"name"`
	Description        *string            `json:"description"`
	Priority           *string            `json:"priority"`
	ShareFree          *bool              `json:"shareFree"`
	Currency           *string            `json:"currency"`
	AppleKey           *string            `json:"appleKey"`
	GoogleKey          *string            `json:"googleKey"`
	EnableFakeReceipt  *bool              `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting     `json:"createWalletScript"`
	DepositScript      *ScriptSetting     `json:"depositScript"`
	WithdrawScript     *ScriptSetting     `json:"withdrawScript"`
	LogSetting         *LogSetting        `json:"logSetting"`
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
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	Description        *string            `json:"description"`
	Priority           *string            `json:"priority"`
	AppleKey           *string            `json:"appleKey"`
	GoogleKey          *string            `json:"googleKey"`
	EnableFakeReceipt  *bool              `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting     `json:"createWalletScript"`
	DepositScript      *ScriptSetting     `json:"depositScript"`
	WithdrawScript     *ScriptSetting     `json:"withdrawScript"`
	LogSetting         *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeWalletsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeWalletsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type QueryWalletsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetWalletRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Slot                   *int32             `json:"slot"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetWalletByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Slot                   *int32             `json:"slot"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DepositByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Slot                   *int32             `json:"slot"`
	Price                  *float32           `json:"price"`
	Count                  *int32             `json:"count"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type WithdrawRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Slot                   *int32             `json:"slot"`
	Count                  *int32             `json:"count"`
	PaidOnly               *bool              `json:"paidOnly"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type WithdrawByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Slot                   *int32             `json:"slot"`
	Count                  *int32             `json:"count"`
	PaidOnly               *bool              `json:"paidOnly"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DepositByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type WithdrawByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeReceiptsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Slot                   *int32             `json:"slot"`
	Begin                  *int64             `json:"begin"`
	End                    *int64             `json:"end"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetByUserIdAndTransactionIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TransactionId          *string            `json:"transactionId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type RecordReceiptRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ContentsId             *string            `json:"contentsId"`
	Receipt                *string            `json:"receipt"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type RecordReceiptByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
