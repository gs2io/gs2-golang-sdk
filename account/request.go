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
	SourceRequestId                         *string        `json:"sourceRequestId"`
	RequestId                               *string        `json:"requestId"`
	ContextStack                            *string        `json:"contextStack"`
	Name                                    *string        `json:"name"`
	Description                             *string        `json:"description"`
	ChangePasswordIfTakeOver                *bool          `json:"changePasswordIfTakeOver"`
	DifferentUserIdForLoginAndDataRetention *bool          `json:"differentUserIdForLoginAndDataRetention"`
	CreateAccountScript                     *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript                    *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript                    *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript                        *ScriptSetting `json:"doTakeOverScript"`
	LogSetting                              *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                                    core.CastString(data["name"]),
		Description:                             core.CastString(data["description"]),
		ChangePasswordIfTakeOver:                core.CastBool(data["changePasswordIfTakeOver"]),
		DifferentUserIdForLoginAndDataRetention: core.CastBool(data["differentUserIdForLoginAndDataRetention"]),
		CreateAccountScript:                     NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer(),
		AuthenticationScript:                    NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer(),
		CreateTakeOverScript:                    NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer(),
		DoTakeOverScript:                        NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer(),
		LogSetting:                              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                     p.Name,
		"description":              p.Description,
		"changePasswordIfTakeOver": p.ChangePasswordIfTakeOver,
		"differentUserIdForLoginAndDataRetention": p.DifferentUserIdForLoginAndDataRetention,
		"createAccountScript":                     p.CreateAccountScript.ToDict(),
		"authenticationScript":                    p.AuthenticationScript.ToDict(),
		"createTakeOverScript":                    p.CreateTakeOverScript.ToDict(),
		"doTakeOverScript":                        p.DoTakeOverScript.ToDict(),
		"logSetting":                              p.LogSetting.ToDict(),
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
	SourceRequestId          *string        `json:"sourceRequestId"`
	RequestId                *string        `json:"requestId"`
	ContextStack             *string        `json:"contextStack"`
	NamespaceName            *string        `json:"namespaceName"`
	Description              *string        `json:"description"`
	ChangePasswordIfTakeOver *bool          `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting `json:"doTakeOverScript"`
	LogSetting               *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
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

type DescribeAccountsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeAccountsRequestFromJson(data string) DescribeAccountsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAccountsRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewCreateAccountRequestFromJson(data string) CreateAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAccountRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffset         *int32  `json:"timeOffset"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUpdateTimeOffsetRequestFromJson(data string) UpdateTimeOffsetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTimeOffsetRequestFromDict(dict)
}

func NewUpdateTimeOffsetRequestFromDict(data map[string]interface{}) UpdateTimeOffsetRequest {
	return UpdateTimeOffsetRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffset:      core.CastInt32(data["timeOffset"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateTimeOffsetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffset":      p.TimeOffset,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateTimeOffsetRequest) Pointer() *UpdateTimeOffsetRequest {
	return &p
}

type UpdateBannedRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Banned             *bool   `json:"banned"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUpdateBannedRequestFromJson(data string) UpdateBannedRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBannedRequestFromDict(dict)
}

func NewUpdateBannedRequestFromDict(data map[string]interface{}) UpdateBannedRequest {
	return UpdateBannedRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Banned:          core.CastBool(data["banned"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateBannedRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"banned":          p.Banned,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateBannedRequest) Pointer() *UpdateBannedRequest {
	return &p
}

type AddBanRequest struct {
	SourceRequestId    *string    `json:"sourceRequestId"`
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	UserId             *string    `json:"userId"`
	BanStatus          *BanStatus `json:"banStatus"`
	TimeOffsetToken    *string    `json:"timeOffsetToken"`
}

func NewAddBanRequestFromJson(data string) AddBanRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddBanRequestFromDict(dict)
}

func NewAddBanRequestFromDict(data map[string]interface{}) AddBanRequest {
	return AddBanRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		BanStatus:       NewBanStatusFromDict(core.CastMap(data["banStatus"])).Pointer(),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AddBanRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"banStatus":       p.BanStatus.ToDict(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddBanRequest) Pointer() *AddBanRequest {
	return &p
}

type RemoveBanRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	BanStatusName      *string `json:"banStatusName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRemoveBanRequestFromJson(data string) RemoveBanRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRemoveBanRequestFromDict(dict)
}

func NewRemoveBanRequestFromDict(data map[string]interface{}) RemoveBanRequest {
	return RemoveBanRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		BanStatusName:   core.CastString(data["banStatusName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RemoveBanRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"banStatusName":   p.BanStatusName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RemoveBanRequest) Pointer() *RemoveBanRequest {
	return &p
}

type GetAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetAccountRequestFromJson(data string) GetAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAccountRequestFromDict(dict)
}

func NewGetAccountRequestFromDict(data map[string]interface{}) GetAccountRequest {
	return GetAccountRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetAccountRequest) Pointer() *GetAccountRequest {
	return &p
}

type DeleteAccountRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteAccountRequestFromJson(data string) DeleteAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAccountRequestFromDict(dict)
}

func NewDeleteAccountRequestFromDict(data map[string]interface{}) DeleteAccountRequest {
	return DeleteAccountRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteAccountRequest) Pointer() *DeleteAccountRequest {
	return &p
}

type AuthenticationRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	KeyId              *string `json:"keyId"`
	Password           *string `json:"password"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewAuthenticationRequestFromJson(data string) AuthenticationRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAuthenticationRequestFromDict(dict)
}

func NewAuthenticationRequestFromDict(data map[string]interface{}) AuthenticationRequest {
	return AuthenticationRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		KeyId:           core.CastString(data["keyId"]),
		Password:        core.CastString(data["password"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AuthenticationRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"keyId":           p.KeyId,
		"password":        p.Password,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AuthenticationRequest) Pointer() *AuthenticationRequest {
	return &p
}

type DescribeTakeOversRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeTakeOversRequestFromJson(data string) DescribeTakeOversRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeTakeOversByUserIdRequestFromJson(data string) DescribeTakeOversByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversByUserIdRequestFromDict(dict)
}

func NewDescribeTakeOversByUserIdRequestFromDict(data map[string]interface{}) DescribeTakeOversByUserIdRequest {
	return DescribeTakeOversByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeTakeOversByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeTakeOversByUserIdRequest) Pointer() *DescribeTakeOversByUserIdRequest {
	return &p
}

type CreateTakeOverRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
	Password           *string `json:"password"`
}

func NewCreateTakeOverRequestFromJson(data string) CreateTakeOverRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
	Password           *string `json:"password"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCreateTakeOverByUserIdRequestFromJson(data string) CreateTakeOverByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverByUserIdRequestFromDict(dict)
}

func NewCreateTakeOverByUserIdRequestFromDict(data map[string]interface{}) CreateTakeOverByUserIdRequest {
	return CreateTakeOverByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Type:            core.CastInt32(data["type"]),
		UserIdentifier:  core.CastString(data["userIdentifier"]),
		Password:        core.CastString(data["password"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CreateTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"type":            p.Type,
		"userIdentifier":  p.UserIdentifier,
		"password":        p.Password,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CreateTakeOverByUserIdRequest) Pointer() *CreateTakeOverByUserIdRequest {
	return &p
}

type GetTakeOverRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	Type            *int32  `json:"type"`
}

func NewGetTakeOverRequestFromJson(data string) GetTakeOverRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	Type            *int32  `json:"type"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetTakeOverByUserIdRequestFromJson(data string) GetTakeOverByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverByUserIdRequestFromDict(dict)
}

func NewGetTakeOverByUserIdRequestFromDict(data map[string]interface{}) GetTakeOverByUserIdRequest {
	return GetTakeOverByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Type:            core.CastInt32(data["type"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"type":            p.Type,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetTakeOverByUserIdRequest) Pointer() *GetTakeOverByUserIdRequest {
	return &p
}

type UpdateTakeOverRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	OldPassword        *string `json:"oldPassword"`
	Password           *string `json:"password"`
}

func NewUpdateTakeOverRequestFromJson(data string) UpdateTakeOverRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
	OldPassword        *string `json:"oldPassword"`
	Password           *string `json:"password"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUpdateTakeOverByUserIdRequestFromJson(data string) UpdateTakeOverByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverByUserIdRequestFromDict(dict)
}

func NewUpdateTakeOverByUserIdRequestFromDict(data map[string]interface{}) UpdateTakeOverByUserIdRequest {
	return UpdateTakeOverByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Type:            core.CastInt32(data["type"]),
		OldPassword:     core.CastString(data["oldPassword"]),
		Password:        core.CastString(data["password"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"type":            p.Type,
		"oldPassword":     p.OldPassword,
		"password":        p.Password,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateTakeOverByUserIdRequest) Pointer() *UpdateTakeOverByUserIdRequest {
	return &p
}

type DeleteTakeOverRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
}

func NewDeleteTakeOverRequestFromJson(data string) DeleteTakeOverRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Type               *int32  `json:"type"`
	UserIdentifier     *string `json:"userIdentifier"`
}

func NewDeleteTakeOverByUserIdentifierRequestFromJson(data string) DeleteTakeOverByUserIdentifierRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdentifierRequestFromDict(dict)
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

type DeleteTakeOverByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Type               *int32  `json:"type"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteTakeOverByUserIdRequestFromJson(data string) DeleteTakeOverByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdRequestFromDict(dict)
}

func NewDeleteTakeOverByUserIdRequestFromDict(data map[string]interface{}) DeleteTakeOverByUserIdRequest {
	return DeleteTakeOverByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Type:            core.CastInt32(data["type"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteTakeOverByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"type":            p.Type,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteTakeOverByUserIdRequest) Pointer() *DeleteTakeOverByUserIdRequest {
	return &p
}

type DoTakeOverRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Type            *int32  `json:"type"`
	UserIdentifier  *string `json:"userIdentifier"`
	Password        *string `json:"password"`
}

func NewDoTakeOverRequestFromJson(data string) DoTakeOverRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoTakeOverRequestFromDict(dict)
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

type GetDataOwnerByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetDataOwnerByUserIdRequestFromJson(data string) GetDataOwnerByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDataOwnerByUserIdRequestFromDict(dict)
}

func NewGetDataOwnerByUserIdRequestFromDict(data map[string]interface{}) GetDataOwnerByUserIdRequest {
	return GetDataOwnerByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetDataOwnerByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetDataOwnerByUserIdRequest) Pointer() *GetDataOwnerByUserIdRequest {
	return &p
}

type DeleteDataOwnerByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteDataOwnerByUserIdRequestFromJson(data string) DeleteDataOwnerByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDataOwnerByUserIdRequestFromDict(dict)
}

func NewDeleteDataOwnerByUserIdRequestFromDict(data map[string]interface{}) DeleteDataOwnerByUserIdRequest {
	return DeleteDataOwnerByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteDataOwnerByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteDataOwnerByUserIdRequest) Pointer() *DeleteDataOwnerByUserIdRequest {
	return &p
}
