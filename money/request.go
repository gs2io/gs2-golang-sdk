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
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesRequestFromDict(dict)
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
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	Name               *string        `json:"name"`
	Description        *string        `json:"description"`
	Priority           *string        `json:"priority"`
	ShareFree          *bool          `json:"shareFree"`
	Currency           *string        `json:"currency"`
	AppleKey           *string        `json:"appleKey"`
	GoogleKey          *string        `json:"googleKey"`
	EnableFakeReceipt  *bool          `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting `json:"createWalletScript"`
	DepositScript      *ScriptSetting `json:"depositScript"`
	WithdrawScript     *ScriptSetting `json:"withdrawScript"`
	LogSetting         *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Priority:           core.CastString(data["priority"]),
		ShareFree:          core.CastBool(data["shareFree"]),
		Currency:           core.CastString(data["currency"]),
		AppleKey:           core.CastString(data["appleKey"]),
		GoogleKey:          core.CastString(data["googleKey"]),
		EnableFakeReceipt:  core.CastBool(data["enableFakeReceipt"]),
		CreateWalletScript: NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer(),
		DepositScript:      NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer(),
		WithdrawScript:     NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"priority":           p.Priority,
		"shareFree":          p.ShareFree,
		"currency":           p.Currency,
		"appleKey":           p.AppleKey,
		"googleKey":          p.GoogleKey,
		"enableFakeReceipt":  p.EnableFakeReceipt,
		"createWalletScript": p.CreateWalletScript.ToDict(),
		"depositScript":      p.DepositScript.ToDict(),
		"withdrawScript":     p.WithdrawScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceRequestFromDict(dict)
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
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	NamespaceName      *string        `json:"namespaceName"`
	Description        *string        `json:"description"`
	Priority           *string        `json:"priority"`
	AppleKey           *string        `json:"appleKey"`
	GoogleKey          *string        `json:"googleKey"`
	EnableFakeReceipt  *bool          `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting `json:"createWalletScript"`
	DepositScript      *ScriptSetting `json:"depositScript"`
	WithdrawScript     *ScriptSetting `json:"withdrawScript"`
	LogSetting         *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Description:        core.CastString(data["description"]),
		Priority:           core.CastString(data["priority"]),
		AppleKey:           core.CastString(data["appleKey"]),
		GoogleKey:          core.CastString(data["googleKey"]),
		EnableFakeReceipt:  core.CastBool(data["enableFakeReceipt"]),
		CreateWalletScript: NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer(),
		DepositScript:      NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer(),
		WithdrawScript:     NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"priority":           p.Priority,
		"appleKey":           p.AppleKey,
		"googleKey":          p.GoogleKey,
		"enableFakeReceipt":  p.EnableFakeReceipt,
		"createWalletScript": p.CreateWalletScript.ToDict(),
		"depositScript":      p.DepositScript.ToDict(),
		"withdrawScript":     p.WithdrawScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceRequestFromDict(dict)
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

type DumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeWalletsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeWalletsRequestFromJson(data string) DescribeWalletsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsRequestFromDict(dict)
}

func NewDescribeWalletsRequestFromDict(data map[string]interface{}) DescribeWalletsRequest {
	return DescribeWalletsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeWalletsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeWalletsRequest) Pointer() *DescribeWalletsRequest {
	return &p
}

type DescribeWalletsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeWalletsByUserIdRequestFromJson(data string) DescribeWalletsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsByUserIdRequestFromDict(dict)
}

func NewDescribeWalletsByUserIdRequestFromDict(data map[string]interface{}) DescribeWalletsByUserIdRequest {
	return DescribeWalletsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeWalletsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeWalletsByUserIdRequest) Pointer() *DescribeWalletsByUserIdRequest {
	return &p
}

type GetWalletRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	Slot            *int32  `json:"slot"`
}

func NewGetWalletRequestFromJson(data string) GetWalletRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletRequestFromDict(dict)
}

func NewGetWalletRequestFromDict(data map[string]interface{}) GetWalletRequest {
	return GetWalletRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Slot:          core.CastInt32(data["slot"]),
	}
}

func (p GetWalletRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"slot":          p.Slot,
	}
}

func (p GetWalletRequest) Pointer() *GetWalletRequest {
	return &p
}

type GetWalletByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	Slot            *int32  `json:"slot"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetWalletByUserIdRequestFromJson(data string) GetWalletByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletByUserIdRequestFromDict(dict)
}

func NewGetWalletByUserIdRequestFromDict(data map[string]interface{}) GetWalletByUserIdRequest {
	return GetWalletByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Slot:            core.CastInt32(data["slot"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetWalletByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetWalletByUserIdRequest) Pointer() *GetWalletByUserIdRequest {
	return &p
}

type DepositByUserIdRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	Slot               *int32   `json:"slot"`
	Price              *float32 `json:"price"`
	Count              *int32   `json:"count"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
}

func NewDepositByUserIdRequestFromJson(data string) DepositByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByUserIdRequestFromDict(dict)
}

func NewDepositByUserIdRequestFromDict(data map[string]interface{}) DepositByUserIdRequest {
	return DepositByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Slot:            core.CastInt32(data["slot"]),
		Price:           core.CastFloat32(data["price"]),
		Count:           core.CastInt32(data["count"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DepositByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"price":           p.Price,
		"count":           p.Count,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DepositByUserIdRequest) Pointer() *DepositByUserIdRequest {
	return &p
}

type WithdrawRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Slot               *int32  `json:"slot"`
	Count              *int32  `json:"count"`
	PaidOnly           *bool   `json:"paidOnly"`
}

func NewWithdrawRequestFromJson(data string) WithdrawRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawRequestFromDict(dict)
}

func NewWithdrawRequestFromDict(data map[string]interface{}) WithdrawRequest {
	return WithdrawRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Slot:          core.CastInt32(data["slot"]),
		Count:         core.CastInt32(data["count"]),
		PaidOnly:      core.CastBool(data["paidOnly"]),
	}
}

func (p WithdrawRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"slot":          p.Slot,
		"count":         p.Count,
		"paidOnly":      p.PaidOnly,
	}
}

func (p WithdrawRequest) Pointer() *WithdrawRequest {
	return &p
}

type WithdrawByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Slot               *int32  `json:"slot"`
	Count              *int32  `json:"count"`
	PaidOnly           *bool   `json:"paidOnly"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewWithdrawByUserIdRequestFromJson(data string) WithdrawByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByUserIdRequestFromDict(dict)
}

func NewWithdrawByUserIdRequestFromDict(data map[string]interface{}) WithdrawByUserIdRequest {
	return WithdrawByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Slot:            core.CastInt32(data["slot"]),
		Count:           core.CastInt32(data["count"]),
		PaidOnly:        core.CastBool(data["paidOnly"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p WithdrawByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"count":           p.Count,
		"paidOnly":        p.PaidOnly,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p WithdrawByUserIdRequest) Pointer() *WithdrawByUserIdRequest {
	return &p
}

type DepositByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewDepositByStampSheetRequestFromJson(data string) DepositByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByStampSheetRequestFromDict(dict)
}

func NewDepositByStampSheetRequestFromDict(data map[string]interface{}) DepositByStampSheetRequest {
	return DepositByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DepositByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DepositByStampSheetRequest) Pointer() *DepositByStampSheetRequest {
	return &p
}

type WithdrawByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewWithdrawByStampTaskRequestFromJson(data string) WithdrawByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByStampTaskRequestFromDict(dict)
}

func NewWithdrawByStampTaskRequestFromDict(data map[string]interface{}) WithdrawByStampTaskRequest {
	return WithdrawByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p WithdrawByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p WithdrawByStampTaskRequest) Pointer() *WithdrawByStampTaskRequest {
	return &p
}

type DescribeReceiptsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	Slot            *int32  `json:"slot"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeReceiptsRequestFromJson(data string) DescribeReceiptsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiptsRequestFromDict(dict)
}

func NewDescribeReceiptsRequestFromDict(data map[string]interface{}) DescribeReceiptsRequest {
	return DescribeReceiptsRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Slot:            core.CastInt32(data["slot"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeReceiptsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"begin":           p.Begin,
		"end":             p.End,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeReceiptsRequest) Pointer() *DescribeReceiptsRequest {
	return &p
}

type GetByUserIdAndTransactionIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TransactionId   *string `json:"transactionId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetByUserIdAndTransactionIdRequestFromJson(data string) GetByUserIdAndTransactionIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetByUserIdAndTransactionIdRequestFromDict(dict)
}

func NewGetByUserIdAndTransactionIdRequestFromDict(data map[string]interface{}) GetByUserIdAndTransactionIdRequest {
	return GetByUserIdAndTransactionIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TransactionId:   core.CastString(data["transactionId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetByUserIdAndTransactionIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"transactionId":   p.TransactionId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetByUserIdAndTransactionIdRequest) Pointer() *GetByUserIdAndTransactionIdRequest {
	return &p
}

type RecordReceiptRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ContentsId         *string `json:"contentsId"`
	Receipt            *string `json:"receipt"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRecordReceiptRequestFromJson(data string) RecordReceiptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecordReceiptRequestFromDict(dict)
}

func NewRecordReceiptRequestFromDict(data map[string]interface{}) RecordReceiptRequest {
	return RecordReceiptRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ContentsId:      core.CastString(data["contentsId"]),
		Receipt:         core.CastString(data["receipt"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RecordReceiptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"contentsId":      p.ContentsId,
		"receipt":         p.Receipt,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RecordReceiptRequest) Pointer() *RecordReceiptRequest {
	return &p
}

type RevertRecordReceiptRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Receipt            *string `json:"receipt"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRevertRecordReceiptRequestFromJson(data string) RevertRecordReceiptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertRecordReceiptRequestFromDict(dict)
}

func NewRevertRecordReceiptRequestFromDict(data map[string]interface{}) RevertRecordReceiptRequest {
	return RevertRecordReceiptRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Receipt:         core.CastString(data["receipt"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RevertRecordReceiptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"receipt":         p.Receipt,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RevertRecordReceiptRequest) Pointer() *RevertRecordReceiptRequest {
	return &p
}

type RecordReceiptByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewRecordReceiptByStampTaskRequestFromJson(data string) RecordReceiptByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecordReceiptByStampTaskRequestFromDict(dict)
}

func NewRecordReceiptByStampTaskRequestFromDict(data map[string]interface{}) RecordReceiptByStampTaskRequest {
	return RecordReceiptByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p RecordReceiptByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p RecordReceiptByStampTaskRequest) Pointer() *RecordReceiptByStampTaskRequest {
	return &p
}

type RevertRecordReceiptByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewRevertRecordReceiptByStampSheetRequestFromJson(data string) RevertRecordReceiptByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertRecordReceiptByStampSheetRequestFromDict(dict)
}

func NewRevertRecordReceiptByStampSheetRequestFromDict(data map[string]interface{}) RevertRecordReceiptByStampSheetRequest {
	return RevertRecordReceiptByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RevertRecordReceiptByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RevertRecordReceiptByStampSheetRequest) Pointer() *RevertRecordReceiptByStampSheetRequest {
	return &p
}
