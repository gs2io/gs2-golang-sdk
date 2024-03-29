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

package formation

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
	SourceRequestId          *string             `json:"sourceRequestId"`
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	Name                     *string             `json:"name"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	UpdateMoldScript         *ScriptSetting      `json:"updateMoldScript"`
	UpdateFormScript         *ScriptSetting      `json:"updateFormScript"`
	UpdatePropertyFormScript *ScriptSetting      `json:"updatePropertyFormScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		UpdateMoldScript:         NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
		UpdateFormScript:         NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
		UpdatePropertyFormScript: NewScriptSettingFromDict(core.CastMap(data["updatePropertyFormScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                     p.Name,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"updateMoldScript":         p.UpdateMoldScript.ToDict(),
		"updateFormScript":         p.UpdateFormScript.ToDict(),
		"updatePropertyFormScript": p.UpdatePropertyFormScript.ToDict(),
		"logSetting":               p.LogSetting.ToDict(),
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
	SourceRequestId          *string             `json:"sourceRequestId"`
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	NamespaceName            *string             `json:"namespaceName"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	UpdateMoldScript         *ScriptSetting      `json:"updateMoldScript"`
	UpdateFormScript         *ScriptSetting      `json:"updateFormScript"`
	UpdatePropertyFormScript *ScriptSetting      `json:"updatePropertyFormScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
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
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		UpdateMoldScript:         NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
		UpdateFormScript:         NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
		UpdatePropertyFormScript: NewScriptSettingFromDict(core.CastMap(data["updatePropertyFormScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"updateMoldScript":         p.UpdateMoldScript.ToDict(),
		"updateFormScript":         p.UpdateFormScript.ToDict(),
		"updatePropertyFormScript": p.UpdatePropertyFormScript.ToDict(),
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

type GetFormModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
}

func NewGetFormModelRequestFromJson(data string) GetFormModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormModelRequestFromDict(dict)
}

func NewGetFormModelRequestFromDict(data map[string]interface{}) GetFormModelRequest {
	return GetFormModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p GetFormModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
	}
}

func (p GetFormModelRequest) Pointer() *GetFormModelRequest {
	return &p
}

type DescribeFormModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeFormModelMastersRequestFromJson(data string) DescribeFormModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormModelMastersRequestFromDict(dict)
}

func NewDescribeFormModelMastersRequestFromDict(data map[string]interface{}) DescribeFormModelMastersRequest {
	return DescribeFormModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeFormModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeFormModelMastersRequest) Pointer() *DescribeFormModelMastersRequest {
	return &p
}

type CreateFormModelMasterRequest struct {
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	Name            *string     `json:"name"`
	Description     *string     `json:"description"`
	Metadata        *string     `json:"metadata"`
	Slots           []SlotModel `json:"slots"`
}

func NewCreateFormModelMasterRequestFromJson(data string) CreateFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateFormModelMasterRequestFromDict(dict)
}

func NewCreateFormModelMasterRequestFromDict(data map[string]interface{}) CreateFormModelMasterRequest {
	return CreateFormModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Slots:         CastSlotModels(core.CastArray(data["slots"])),
	}
}

func (p CreateFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"slots": CastSlotModelsFromDict(
			p.Slots,
		),
	}
}

func (p CreateFormModelMasterRequest) Pointer() *CreateFormModelMasterRequest {
	return &p
}

type GetFormModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	FormModelName   *string `json:"formModelName"`
}

func NewGetFormModelMasterRequestFromJson(data string) GetFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormModelMasterRequestFromDict(dict)
}

func NewGetFormModelMasterRequestFromDict(data map[string]interface{}) GetFormModelMasterRequest {
	return GetFormModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		FormModelName: core.CastString(data["formModelName"]),
	}
}

func (p GetFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"formModelName": p.FormModelName,
	}
}

func (p GetFormModelMasterRequest) Pointer() *GetFormModelMasterRequest {
	return &p
}

type UpdateFormModelMasterRequest struct {
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	FormModelName   *string     `json:"formModelName"`
	Description     *string     `json:"description"`
	Metadata        *string     `json:"metadata"`
	Slots           []SlotModel `json:"slots"`
}

func NewUpdateFormModelMasterRequestFromJson(data string) UpdateFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateFormModelMasterRequestFromDict(dict)
}

func NewUpdateFormModelMasterRequestFromDict(data map[string]interface{}) UpdateFormModelMasterRequest {
	return UpdateFormModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		FormModelName: core.CastString(data["formModelName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Slots:         CastSlotModels(core.CastArray(data["slots"])),
	}
}

func (p UpdateFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"formModelName": p.FormModelName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"slots": CastSlotModelsFromDict(
			p.Slots,
		),
	}
}

func (p UpdateFormModelMasterRequest) Pointer() *UpdateFormModelMasterRequest {
	return &p
}

type DeleteFormModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	FormModelName   *string `json:"formModelName"`
}

func NewDeleteFormModelMasterRequestFromJson(data string) DeleteFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormModelMasterRequestFromDict(dict)
}

func NewDeleteFormModelMasterRequestFromDict(data map[string]interface{}) DeleteFormModelMasterRequest {
	return DeleteFormModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		FormModelName: core.CastString(data["formModelName"]),
	}
}

func (p DeleteFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"formModelName": p.FormModelName,
	}
}

func (p DeleteFormModelMasterRequest) Pointer() *DeleteFormModelMasterRequest {
	return &p
}

type DescribeMoldModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeMoldModelsRequestFromJson(data string) DescribeMoldModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldModelsRequestFromDict(dict)
}

func NewDescribeMoldModelsRequestFromDict(data map[string]interface{}) DescribeMoldModelsRequest {
	return DescribeMoldModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMoldModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMoldModelsRequest) Pointer() *DescribeMoldModelsRequest {
	return &p
}

type GetMoldModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
}

func NewGetMoldModelRequestFromJson(data string) GetMoldModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldModelRequestFromDict(dict)
}

func NewGetMoldModelRequestFromDict(data map[string]interface{}) GetMoldModelRequest {
	return GetMoldModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p GetMoldModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
	}
}

func (p GetMoldModelRequest) Pointer() *GetMoldModelRequest {
	return &p
}

type DescribeMoldModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMoldModelMastersRequestFromJson(data string) DescribeMoldModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldModelMastersRequestFromDict(dict)
}

func NewDescribeMoldModelMastersRequestFromDict(data map[string]interface{}) DescribeMoldModelMastersRequest {
	return DescribeMoldModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMoldModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMoldModelMastersRequest) Pointer() *DescribeMoldModelMastersRequest {
	return &p
}

type CreateMoldModelMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	FormModelName      *string `json:"formModelName"`
	InitialMaxCapacity *int32  `json:"initialMaxCapacity"`
	MaxCapacity        *int32  `json:"maxCapacity"`
}

func NewCreateMoldModelMasterRequestFromJson(data string) CreateMoldModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMoldModelMasterRequestFromDict(dict)
}

func NewCreateMoldModelMasterRequestFromDict(data map[string]interface{}) CreateMoldModelMasterRequest {
	return CreateMoldModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		FormModelName:      core.CastString(data["formModelName"]),
		InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
		MaxCapacity:        core.CastInt32(data["maxCapacity"]),
	}
}

func (p CreateMoldModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"name":               p.Name,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"formModelName":      p.FormModelName,
		"initialMaxCapacity": p.InitialMaxCapacity,
		"maxCapacity":        p.MaxCapacity,
	}
}

func (p CreateMoldModelMasterRequest) Pointer() *CreateMoldModelMasterRequest {
	return &p
}

type GetMoldModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
}

func NewGetMoldModelMasterRequestFromJson(data string) GetMoldModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldModelMasterRequestFromDict(dict)
}

func NewGetMoldModelMasterRequestFromDict(data map[string]interface{}) GetMoldModelMasterRequest {
	return GetMoldModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p GetMoldModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
	}
}

func (p GetMoldModelMasterRequest) Pointer() *GetMoldModelMasterRequest {
	return &p
}

type UpdateMoldModelMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	MoldModelName      *string `json:"moldModelName"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	FormModelName      *string `json:"formModelName"`
	InitialMaxCapacity *int32  `json:"initialMaxCapacity"`
	MaxCapacity        *int32  `json:"maxCapacity"`
}

func NewUpdateMoldModelMasterRequestFromJson(data string) UpdateMoldModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMoldModelMasterRequestFromDict(dict)
}

func NewUpdateMoldModelMasterRequestFromDict(data map[string]interface{}) UpdateMoldModelMasterRequest {
	return UpdateMoldModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		MoldModelName:      core.CastString(data["moldModelName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		FormModelName:      core.CastString(data["formModelName"]),
		InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
		MaxCapacity:        core.CastInt32(data["maxCapacity"]),
	}
}

func (p UpdateMoldModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"moldModelName":      p.MoldModelName,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"formModelName":      p.FormModelName,
		"initialMaxCapacity": p.InitialMaxCapacity,
		"maxCapacity":        p.MaxCapacity,
	}
}

func (p UpdateMoldModelMasterRequest) Pointer() *UpdateMoldModelMasterRequest {
	return &p
}

type DeleteMoldModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
}

func NewDeleteMoldModelMasterRequestFromJson(data string) DeleteMoldModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldModelMasterRequestFromDict(dict)
}

func NewDeleteMoldModelMasterRequestFromDict(data map[string]interface{}) DeleteMoldModelMasterRequest {
	return DeleteMoldModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p DeleteMoldModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
	}
}

func (p DeleteMoldModelMasterRequest) Pointer() *DeleteMoldModelMasterRequest {
	return &p
}

type DescribePropertyFormModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribePropertyFormModelsRequestFromJson(data string) DescribePropertyFormModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormModelsRequestFromDict(dict)
}

func NewDescribePropertyFormModelsRequestFromDict(data map[string]interface{}) DescribePropertyFormModelsRequest {
	return DescribePropertyFormModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribePropertyFormModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribePropertyFormModelsRequest) Pointer() *DescribePropertyFormModelsRequest {
	return &p
}

type GetPropertyFormModelRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewGetPropertyFormModelRequestFromJson(data string) GetPropertyFormModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormModelRequestFromDict(dict)
}

func NewGetPropertyFormModelRequestFromDict(data map[string]interface{}) GetPropertyFormModelRequest {
	return GetPropertyFormModelRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
	}
}

func (p GetPropertyFormModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"propertyFormModelName": p.PropertyFormModelName,
	}
}

func (p GetPropertyFormModelRequest) Pointer() *GetPropertyFormModelRequest {
	return &p
}

type DescribePropertyFormModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribePropertyFormModelMastersRequestFromJson(data string) DescribePropertyFormModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormModelMastersRequestFromDict(dict)
}

func NewDescribePropertyFormModelMastersRequestFromDict(data map[string]interface{}) DescribePropertyFormModelMastersRequest {
	return DescribePropertyFormModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribePropertyFormModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribePropertyFormModelMastersRequest) Pointer() *DescribePropertyFormModelMastersRequest {
	return &p
}

type CreatePropertyFormModelMasterRequest struct {
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	Name            *string     `json:"name"`
	Description     *string     `json:"description"`
	Metadata        *string     `json:"metadata"`
	Slots           []SlotModel `json:"slots"`
}

func NewCreatePropertyFormModelMasterRequestFromJson(data string) CreatePropertyFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePropertyFormModelMasterRequestFromDict(dict)
}

func NewCreatePropertyFormModelMasterRequestFromDict(data map[string]interface{}) CreatePropertyFormModelMasterRequest {
	return CreatePropertyFormModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Slots:         CastSlotModels(core.CastArray(data["slots"])),
	}
}

func (p CreatePropertyFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"slots": CastSlotModelsFromDict(
			p.Slots,
		),
	}
}

func (p CreatePropertyFormModelMasterRequest) Pointer() *CreatePropertyFormModelMasterRequest {
	return &p
}

type GetPropertyFormModelMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewGetPropertyFormModelMasterRequestFromJson(data string) GetPropertyFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormModelMasterRequestFromDict(dict)
}

func NewGetPropertyFormModelMasterRequestFromDict(data map[string]interface{}) GetPropertyFormModelMasterRequest {
	return GetPropertyFormModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
	}
}

func (p GetPropertyFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"propertyFormModelName": p.PropertyFormModelName,
	}
}

func (p GetPropertyFormModelMasterRequest) Pointer() *GetPropertyFormModelMasterRequest {
	return &p
}

type UpdatePropertyFormModelMasterRequest struct {
	SourceRequestId       *string     `json:"sourceRequestId"`
	RequestId             *string     `json:"requestId"`
	ContextStack          *string     `json:"contextStack"`
	NamespaceName         *string     `json:"namespaceName"`
	PropertyFormModelName *string     `json:"propertyFormModelName"`
	Description           *string     `json:"description"`
	Metadata              *string     `json:"metadata"`
	Slots                 []SlotModel `json:"slots"`
}

func NewUpdatePropertyFormModelMasterRequestFromJson(data string) UpdatePropertyFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdatePropertyFormModelMasterRequestFromDict(dict)
}

func NewUpdatePropertyFormModelMasterRequestFromDict(data map[string]interface{}) UpdatePropertyFormModelMasterRequest {
	return UpdatePropertyFormModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		Slots:                 CastSlotModels(core.CastArray(data["slots"])),
	}
}

func (p UpdatePropertyFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"propertyFormModelName": p.PropertyFormModelName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"slots": CastSlotModelsFromDict(
			p.Slots,
		),
	}
}

func (p UpdatePropertyFormModelMasterRequest) Pointer() *UpdatePropertyFormModelMasterRequest {
	return &p
}

type DeletePropertyFormModelMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewDeletePropertyFormModelMasterRequestFromJson(data string) DeletePropertyFormModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormModelMasterRequestFromDict(dict)
}

func NewDeletePropertyFormModelMasterRequestFromDict(data map[string]interface{}) DeletePropertyFormModelMasterRequest {
	return DeletePropertyFormModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
	}
}

func (p DeletePropertyFormModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"propertyFormModelName": p.PropertyFormModelName,
	}
}

func (p DeletePropertyFormModelMasterRequest) Pointer() *DeletePropertyFormModelMasterRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExportMasterRequestFromDict(dict)
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
	return ExportMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
	return &p
}

type GetCurrentFormMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentFormMasterRequestFromJson(data string) GetCurrentFormMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentFormMasterRequestFromDict(dict)
}

func NewGetCurrentFormMasterRequestFromDict(data map[string]interface{}) GetCurrentFormMasterRequest {
	return GetCurrentFormMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentFormMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentFormMasterRequest) Pointer() *GetCurrentFormMasterRequest {
	return &p
}

type UpdateCurrentFormMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentFormMasterRequestFromJson(data string) UpdateCurrentFormMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFormMasterRequestFromDict(dict)
}

func NewUpdateCurrentFormMasterRequestFromDict(data map[string]interface{}) UpdateCurrentFormMasterRequest {
	return UpdateCurrentFormMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentFormMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentFormMasterRequest) Pointer() *UpdateCurrentFormMasterRequest {
	return &p
}

type UpdateCurrentFormMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentFormMasterFromGitHubRequestFromJson(data string) UpdateCurrentFormMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFormMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentFormMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentFormMasterFromGitHubRequest {
	return UpdateCurrentFormMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentFormMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentFormMasterFromGitHubRequest) Pointer() *UpdateCurrentFormMasterFromGitHubRequest {
	return &p
}

type DescribeMoldsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMoldsRequestFromJson(data string) DescribeMoldsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldsRequestFromDict(dict)
}

func NewDescribeMoldsRequestFromDict(data map[string]interface{}) DescribeMoldsRequest {
	return DescribeMoldsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMoldsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMoldsRequest) Pointer() *DescribeMoldsRequest {
	return &p
}

type DescribeMoldsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeMoldsByUserIdRequestFromJson(data string) DescribeMoldsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldsByUserIdRequestFromDict(dict)
}

func NewDescribeMoldsByUserIdRequestFromDict(data map[string]interface{}) DescribeMoldsByUserIdRequest {
	return DescribeMoldsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeMoldsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeMoldsByUserIdRequest) Pointer() *DescribeMoldsByUserIdRequest {
	return &p
}

type GetMoldRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	MoldModelName   *string `json:"moldModelName"`
}

func NewGetMoldRequestFromJson(data string) GetMoldRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldRequestFromDict(dict)
}

func NewGetMoldRequestFromDict(data map[string]interface{}) GetMoldRequest {
	return GetMoldRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p GetMoldRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
	}
}

func (p GetMoldRequest) Pointer() *GetMoldRequest {
	return &p
}

type GetMoldByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	MoldModelName   *string `json:"moldModelName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetMoldByUserIdRequestFromJson(data string) GetMoldByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldByUserIdRequestFromDict(dict)
}

func NewGetMoldByUserIdRequestFromDict(data map[string]interface{}) GetMoldByUserIdRequest {
	return GetMoldByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetMoldByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetMoldByUserIdRequest) Pointer() *GetMoldByUserIdRequest {
	return &p
}

type SetMoldCapacityByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewSetMoldCapacityByUserIdRequestFromJson(data string) SetMoldCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMoldCapacityByUserIdRequestFromDict(dict)
}

func NewSetMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) SetMoldCapacityByUserIdRequest {
	return SetMoldCapacityByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Capacity:        core.CastInt32(data["capacity"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"capacity":        p.Capacity,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetMoldCapacityByUserIdRequest) Pointer() *SetMoldCapacityByUserIdRequest {
	return &p
}

type AddMoldCapacityByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewAddMoldCapacityByUserIdRequestFromJson(data string) AddMoldCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddMoldCapacityByUserIdRequestFromDict(dict)
}

func NewAddMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) AddMoldCapacityByUserIdRequest {
	return AddMoldCapacityByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Capacity:        core.CastInt32(data["capacity"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AddMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"capacity":        p.Capacity,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddMoldCapacityByUserIdRequest) Pointer() *AddMoldCapacityByUserIdRequest {
	return &p
}

type SubMoldCapacityByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewSubMoldCapacityByUserIdRequestFromJson(data string) SubMoldCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubMoldCapacityByUserIdRequestFromDict(dict)
}

func NewSubMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) SubMoldCapacityByUserIdRequest {
	return SubMoldCapacityByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Capacity:        core.CastInt32(data["capacity"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SubMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"capacity":        p.Capacity,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SubMoldCapacityByUserIdRequest) Pointer() *SubMoldCapacityByUserIdRequest {
	return &p
}

type DeleteMoldRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MoldModelName      *string `json:"moldModelName"`
}

func NewDeleteMoldRequestFromJson(data string) DeleteMoldRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldRequestFromDict(dict)
}

func NewDeleteMoldRequestFromDict(data map[string]interface{}) DeleteMoldRequest {
	return DeleteMoldRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p DeleteMoldRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
	}
}

func (p DeleteMoldRequest) Pointer() *DeleteMoldRequest {
	return &p
}

type DeleteMoldByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteMoldByUserIdRequestFromJson(data string) DeleteMoldByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldByUserIdRequestFromDict(dict)
}

func NewDeleteMoldByUserIdRequestFromDict(data map[string]interface{}) DeleteMoldByUserIdRequest {
	return DeleteMoldByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteMoldByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteMoldByUserIdRequest) Pointer() *DeleteMoldByUserIdRequest {
	return &p
}

type AddCapacityByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewAddCapacityByStampSheetRequestFromJson(data string) AddCapacityByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddCapacityByStampSheetRequestFromDict(dict)
}

func NewAddCapacityByStampSheetRequestFromDict(data map[string]interface{}) AddCapacityByStampSheetRequest {
	return AddCapacityByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddCapacityByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddCapacityByStampSheetRequest) Pointer() *AddCapacityByStampSheetRequest {
	return &p
}

type SubCapacityByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewSubCapacityByStampTaskRequestFromJson(data string) SubCapacityByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubCapacityByStampTaskRequestFromDict(dict)
}

func NewSubCapacityByStampTaskRequestFromDict(data map[string]interface{}) SubCapacityByStampTaskRequest {
	return SubCapacityByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p SubCapacityByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubCapacityByStampTaskRequest) Pointer() *SubCapacityByStampTaskRequest {
	return &p
}

type SetCapacityByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewSetCapacityByStampSheetRequestFromJson(data string) SetCapacityByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetCapacityByStampSheetRequestFromDict(dict)
}

func NewSetCapacityByStampSheetRequestFromDict(data map[string]interface{}) SetCapacityByStampSheetRequest {
	return SetCapacityByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetCapacityByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetCapacityByStampSheetRequest) Pointer() *SetCapacityByStampSheetRequest {
	return &p
}

type DescribeFormsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeFormsRequestFromJson(data string) DescribeFormsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormsRequestFromDict(dict)
}

func NewDescribeFormsRequestFromDict(data map[string]interface{}) DescribeFormsRequest {
	return DescribeFormsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeFormsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeFormsRequest) Pointer() *DescribeFormsRequest {
	return &p
}

type DescribeFormsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeFormsByUserIdRequestFromJson(data string) DescribeFormsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormsByUserIdRequestFromDict(dict)
}

func NewDescribeFormsByUserIdRequestFromDict(data map[string]interface{}) DescribeFormsByUserIdRequest {
	return DescribeFormsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeFormsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"moldModelName":   p.MoldModelName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeFormsByUserIdRequest) Pointer() *DescribeFormsByUserIdRequest {
	return &p
}

type GetFormRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	MoldModelName   *string `json:"moldModelName"`
	Index           *int32  `json:"index"`
}

func NewGetFormRequestFromJson(data string) GetFormRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormRequestFromDict(dict)
}

func NewGetFormRequestFromDict(data map[string]interface{}) GetFormRequest {
	return GetFormRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
	}
}

func (p GetFormRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
	}
}

func (p GetFormRequest) Pointer() *GetFormRequest {
	return &p
}

type GetFormByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	MoldModelName   *string `json:"moldModelName"`
	Index           *int32  `json:"index"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetFormByUserIdRequestFromJson(data string) GetFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormByUserIdRequestFromDict(dict)
}

func NewGetFormByUserIdRequestFromDict(data map[string]interface{}) GetFormByUserIdRequest {
	return GetFormByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Index:           core.CastInt32(data["index"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"index":           p.Index,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetFormByUserIdRequest) Pointer() *GetFormByUserIdRequest {
	return &p
}

type GetFormWithSignatureRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	MoldModelName   *string `json:"moldModelName"`
	Index           *int32  `json:"index"`
	KeyId           *string `json:"keyId"`
}

func NewGetFormWithSignatureRequestFromJson(data string) GetFormWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormWithSignatureRequestFromDict(dict)
}

func NewGetFormWithSignatureRequestFromDict(data map[string]interface{}) GetFormWithSignatureRequest {
	return GetFormWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetFormWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"keyId":         p.KeyId,
	}
}

func (p GetFormWithSignatureRequest) Pointer() *GetFormWithSignatureRequest {
	return &p
}

type GetFormWithSignatureByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	MoldModelName   *string `json:"moldModelName"`
	Index           *int32  `json:"index"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetFormWithSignatureByUserIdRequestFromJson(data string) GetFormWithSignatureByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetFormWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetFormWithSignatureByUserIdRequest {
	return GetFormWithSignatureByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Index:           core.CastInt32(data["index"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetFormWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"index":           p.Index,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetFormWithSignatureByUserIdRequest) Pointer() *GetFormWithSignatureByUserIdRequest {
	return &p
}

type SetFormByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
	Slots              []Slot  `json:"slots"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewSetFormByUserIdRequestFromJson(data string) SetFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFormByUserIdRequestFromDict(dict)
}

func NewSetFormByUserIdRequestFromDict(data map[string]interface{}) SetFormByUserIdRequest {
	return SetFormByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Index:           core.CastInt32(data["index"]),
		Slots:           CastSlots(core.CastArray(data["slots"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"slots": CastSlotsFromDict(
			p.Slots,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetFormByUserIdRequest) Pointer() *SetFormByUserIdRequest {
	return &p
}

type SetFormWithSignatureRequest struct {
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	DuplicationAvoider *string             `json:"duplicationAvoider"`
	NamespaceName      *string             `json:"namespaceName"`
	AccessToken        *string             `json:"accessToken"`
	MoldModelName      *string             `json:"moldModelName"`
	Index              *int32              `json:"index"`
	Slots              []SlotWithSignature `json:"slots"`
	KeyId              *string             `json:"keyId"`
}

func NewSetFormWithSignatureRequestFromJson(data string) SetFormWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFormWithSignatureRequestFromDict(dict)
}

func NewSetFormWithSignatureRequestFromDict(data map[string]interface{}) SetFormWithSignatureRequest {
	return SetFormWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
		Slots:         CastSlotWithSignatures(core.CastArray(data["slots"])),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p SetFormWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"slots": CastSlotWithSignaturesFromDict(
			p.Slots,
		),
		"keyId": p.KeyId,
	}
}

func (p SetFormWithSignatureRequest) Pointer() *SetFormWithSignatureRequest {
	return &p
}

type AcquireActionsToFormPropertiesRequest struct {
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	UserId             *string        `json:"userId"`
	MoldModelName      *string        `json:"moldModelName"`
	Index              *int32         `json:"index"`
	AcquireAction      *AcquireAction `json:"acquireAction"`
	Config             []Config       `json:"config"`
	TimeOffsetToken    *string        `json:"timeOffsetToken"`
}

func NewAcquireActionsToFormPropertiesRequestFromJson(data string) AcquireActionsToFormPropertiesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionsToFormPropertiesRequestFromDict(dict)
}

func NewAcquireActionsToFormPropertiesRequestFromDict(data map[string]interface{}) AcquireActionsToFormPropertiesRequest {
	return AcquireActionsToFormPropertiesRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Index:           core.CastInt32(data["index"]),
		AcquireAction:   NewAcquireActionFromDict(core.CastMap(data["acquireAction"])).Pointer(),
		Config:          CastConfigs(core.CastArray(data["config"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireActionsToFormPropertiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"acquireAction": p.AcquireAction.ToDict(),
		"config": CastConfigsFromDict(
			p.Config,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcquireActionsToFormPropertiesRequest) Pointer() *AcquireActionsToFormPropertiesRequest {
	return &p
}

type DeleteFormRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
}

func NewDeleteFormRequestFromJson(data string) DeleteFormRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormRequestFromDict(dict)
}

func NewDeleteFormRequestFromDict(data map[string]interface{}) DeleteFormRequest {
	return DeleteFormRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
	}
}

func (p DeleteFormRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
	}
}

func (p DeleteFormRequest) Pointer() *DeleteFormRequest {
	return &p
}

type DeleteFormByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteFormByUserIdRequestFromJson(data string) DeleteFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormByUserIdRequestFromDict(dict)
}

func NewDeleteFormByUserIdRequestFromDict(data map[string]interface{}) DeleteFormByUserIdRequest {
	return DeleteFormByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MoldModelName:   core.CastString(data["moldModelName"]),
		Index:           core.CastInt32(data["index"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"moldModelName":   p.MoldModelName,
		"index":           p.Index,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteFormByUserIdRequest) Pointer() *DeleteFormByUserIdRequest {
	return &p
}

type AcquireActionToFormPropertiesByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewAcquireActionToFormPropertiesByStampSheetRequestFromJson(data string) AcquireActionToFormPropertiesByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionToFormPropertiesByStampSheetRequestFromDict(dict)
}

func NewAcquireActionToFormPropertiesByStampSheetRequestFromDict(data map[string]interface{}) AcquireActionToFormPropertiesByStampSheetRequest {
	return AcquireActionToFormPropertiesByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireActionToFormPropertiesByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireActionToFormPropertiesByStampSheetRequest) Pointer() *AcquireActionToFormPropertiesByStampSheetRequest {
	return &p
}

type SetFormByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewSetFormByStampSheetRequestFromJson(data string) SetFormByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFormByStampSheetRequestFromDict(dict)
}

func NewSetFormByStampSheetRequestFromDict(data map[string]interface{}) SetFormByStampSheetRequest {
	return SetFormByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetFormByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetFormByStampSheetRequest) Pointer() *SetFormByStampSheetRequest {
	return &p
}

type DescribePropertyFormsRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PageToken             *string `json:"pageToken"`
	Limit                 *int32  `json:"limit"`
}

func NewDescribePropertyFormsRequestFromJson(data string) DescribePropertyFormsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormsRequestFromDict(dict)
}

func NewDescribePropertyFormsRequestFromDict(data map[string]interface{}) DescribePropertyFormsRequest {
	return DescribePropertyFormsRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PageToken:             core.CastString(data["pageToken"]),
		Limit:                 core.CastInt32(data["limit"]),
	}
}

func (p DescribePropertyFormsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"accessToken":           p.AccessToken,
		"propertyFormModelName": p.PropertyFormModelName,
		"pageToken":             p.PageToken,
		"limit":                 p.Limit,
	}
}

func (p DescribePropertyFormsRequest) Pointer() *DescribePropertyFormsRequest {
	return &p
}

type DescribePropertyFormsByUserIdRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PageToken             *string `json:"pageToken"`
	Limit                 *int32  `json:"limit"`
	TimeOffsetToken       *string `json:"timeOffsetToken"`
}

func NewDescribePropertyFormsByUserIdRequestFromJson(data string) DescribePropertyFormsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormsByUserIdRequestFromDict(dict)
}

func NewDescribePropertyFormsByUserIdRequestFromDict(data map[string]interface{}) DescribePropertyFormsByUserIdRequest {
	return DescribePropertyFormsByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PageToken:             core.CastString(data["pageToken"]),
		Limit:                 core.CastInt32(data["limit"]),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribePropertyFormsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"pageToken":             p.PageToken,
		"limit":                 p.Limit,
		"timeOffsetToken":       p.TimeOffsetToken,
	}
}

func (p DescribePropertyFormsByUserIdRequest) Pointer() *DescribePropertyFormsByUserIdRequest {
	return &p
}

type GetPropertyFormRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewGetPropertyFormRequestFromJson(data string) GetPropertyFormRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormRequestFromDict(dict)
}

func NewGetPropertyFormRequestFromDict(data map[string]interface{}) GetPropertyFormRequest {
	return GetPropertyFormRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
	}
}

func (p GetPropertyFormRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"accessToken":           p.AccessToken,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
	}
}

func (p GetPropertyFormRequest) Pointer() *GetPropertyFormRequest {
	return &p
}

type GetPropertyFormByUserIdRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	TimeOffsetToken       *string `json:"timeOffsetToken"`
}

func NewGetPropertyFormByUserIdRequestFromJson(data string) GetPropertyFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormByUserIdRequestFromDict(dict)
}

func NewGetPropertyFormByUserIdRequestFromDict(data map[string]interface{}) GetPropertyFormByUserIdRequest {
	return GetPropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetPropertyFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"timeOffsetToken":       p.TimeOffsetToken,
	}
}

func (p GetPropertyFormByUserIdRequest) Pointer() *GetPropertyFormByUserIdRequest {
	return &p
}

type GetPropertyFormWithSignatureRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	KeyId                 *string `json:"keyId"`
}

func NewGetPropertyFormWithSignatureRequestFromJson(data string) GetPropertyFormWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormWithSignatureRequestFromDict(dict)
}

func NewGetPropertyFormWithSignatureRequestFromDict(data map[string]interface{}) GetPropertyFormWithSignatureRequest {
	return GetPropertyFormWithSignatureRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		KeyId:                 core.CastString(data["keyId"]),
	}
}

func (p GetPropertyFormWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"accessToken":           p.AccessToken,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"keyId":                 p.KeyId,
	}
}

func (p GetPropertyFormWithSignatureRequest) Pointer() *GetPropertyFormWithSignatureRequest {
	return &p
}

type GetPropertyFormWithSignatureByUserIdRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	KeyId                 *string `json:"keyId"`
	TimeOffsetToken       *string `json:"timeOffsetToken"`
}

func NewGetPropertyFormWithSignatureByUserIdRequestFromJson(data string) GetPropertyFormWithSignatureByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetPropertyFormWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetPropertyFormWithSignatureByUserIdRequest {
	return GetPropertyFormWithSignatureByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		KeyId:                 core.CastString(data["keyId"]),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetPropertyFormWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"keyId":                 p.KeyId,
		"timeOffsetToken":       p.TimeOffsetToken,
	}
}

func (p GetPropertyFormWithSignatureByUserIdRequest) Pointer() *GetPropertyFormWithSignatureByUserIdRequest {
	return &p
}

type SetPropertyFormByUserIdRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	Slots                 []Slot  `json:"slots"`
	TimeOffsetToken       *string `json:"timeOffsetToken"`
}

func NewSetPropertyFormByUserIdRequestFromJson(data string) SetPropertyFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetPropertyFormByUserIdRequestFromDict(dict)
}

func NewSetPropertyFormByUserIdRequestFromDict(data map[string]interface{}) SetPropertyFormByUserIdRequest {
	return SetPropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		Slots:                 CastSlots(core.CastArray(data["slots"])),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetPropertyFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"slots": CastSlotsFromDict(
			p.Slots,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetPropertyFormByUserIdRequest) Pointer() *SetPropertyFormByUserIdRequest {
	return &p
}

type SetPropertyFormWithSignatureRequest struct {
	SourceRequestId       *string             `json:"sourceRequestId"`
	RequestId             *string             `json:"requestId"`
	ContextStack          *string             `json:"contextStack"`
	DuplicationAvoider    *string             `json:"duplicationAvoider"`
	NamespaceName         *string             `json:"namespaceName"`
	AccessToken           *string             `json:"accessToken"`
	PropertyFormModelName *string             `json:"propertyFormModelName"`
	PropertyId            *string             `json:"propertyId"`
	Slots                 []SlotWithSignature `json:"slots"`
	KeyId                 *string             `json:"keyId"`
}

func NewSetPropertyFormWithSignatureRequestFromJson(data string) SetPropertyFormWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetPropertyFormWithSignatureRequestFromDict(dict)
}

func NewSetPropertyFormWithSignatureRequestFromDict(data map[string]interface{}) SetPropertyFormWithSignatureRequest {
	return SetPropertyFormWithSignatureRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		Slots:                 CastSlotWithSignatures(core.CastArray(data["slots"])),
		KeyId:                 core.CastString(data["keyId"]),
	}
}

func (p SetPropertyFormWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"accessToken":           p.AccessToken,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"slots": CastSlotWithSignaturesFromDict(
			p.Slots,
		),
		"keyId": p.KeyId,
	}
}

func (p SetPropertyFormWithSignatureRequest) Pointer() *SetPropertyFormWithSignatureRequest {
	return &p
}

type AcquireActionsToPropertyFormPropertiesRequest struct {
	SourceRequestId       *string        `json:"sourceRequestId"`
	RequestId             *string        `json:"requestId"`
	ContextStack          *string        `json:"contextStack"`
	DuplicationAvoider    *string        `json:"duplicationAvoider"`
	NamespaceName         *string        `json:"namespaceName"`
	UserId                *string        `json:"userId"`
	PropertyFormModelName *string        `json:"propertyFormModelName"`
	PropertyId            *string        `json:"propertyId"`
	AcquireAction         *AcquireAction `json:"acquireAction"`
	Config                []Config       `json:"config"`
	TimeOffsetToken       *string        `json:"timeOffsetToken"`
}

func NewAcquireActionsToPropertyFormPropertiesRequestFromJson(data string) AcquireActionsToPropertyFormPropertiesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionsToPropertyFormPropertiesRequestFromDict(dict)
}

func NewAcquireActionsToPropertyFormPropertiesRequestFromDict(data map[string]interface{}) AcquireActionsToPropertyFormPropertiesRequest {
	return AcquireActionsToPropertyFormPropertiesRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		AcquireAction:         NewAcquireActionFromDict(core.CastMap(data["acquireAction"])).Pointer(),
		Config:                CastConfigs(core.CastArray(data["config"])),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireActionsToPropertyFormPropertiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"acquireAction":         p.AcquireAction.ToDict(),
		"config": CastConfigsFromDict(
			p.Config,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcquireActionsToPropertyFormPropertiesRequest) Pointer() *AcquireActionsToPropertyFormPropertiesRequest {
	return &p
}

type DeletePropertyFormRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewDeletePropertyFormRequestFromJson(data string) DeletePropertyFormRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormRequestFromDict(dict)
}

func NewDeletePropertyFormRequestFromDict(data map[string]interface{}) DeletePropertyFormRequest {
	return DeletePropertyFormRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
	}
}

func (p DeletePropertyFormRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"accessToken":           p.AccessToken,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
	}
}

func (p DeletePropertyFormRequest) Pointer() *DeletePropertyFormRequest {
	return &p
}

type DeletePropertyFormByUserIdRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	TimeOffsetToken       *string `json:"timeOffsetToken"`
}

func NewDeletePropertyFormByUserIdRequestFromJson(data string) DeletePropertyFormByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormByUserIdRequestFromDict(dict)
}

func NewDeletePropertyFormByUserIdRequestFromDict(data map[string]interface{}) DeletePropertyFormByUserIdRequest {
	return DeletePropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		TimeOffsetToken:       core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeletePropertyFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"timeOffsetToken":       p.TimeOffsetToken,
	}
}

func (p DeletePropertyFormByUserIdRequest) Pointer() *DeletePropertyFormByUserIdRequest {
	return &p
}

type AcquireActionToPropertyFormPropertiesByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewAcquireActionToPropertyFormPropertiesByStampSheetRequestFromJson(data string) AcquireActionToPropertyFormPropertiesByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionToPropertyFormPropertiesByStampSheetRequestFromDict(dict)
}

func NewAcquireActionToPropertyFormPropertiesByStampSheetRequestFromDict(data map[string]interface{}) AcquireActionToPropertyFormPropertiesByStampSheetRequest {
	return AcquireActionToPropertyFormPropertiesByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireActionToPropertyFormPropertiesByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireActionToPropertyFormPropertiesByStampSheetRequest) Pointer() *AcquireActionToPropertyFormPropertiesByStampSheetRequest {
	return &p
}
