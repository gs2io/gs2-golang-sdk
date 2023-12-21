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
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNamespacesRequestFromDict(dict2)
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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	UpdateMoldScript   *ScriptSetting      `json:"updateMoldScript"`
	UpdateFormScript   *ScriptSetting      `json:"updateFormScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateNamespaceRequestFromDict(dict2)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		UpdateMoldScript:   NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
		UpdateFormScript:   NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"updateMoldScript":   p.UpdateMoldScript.ToDict(),
		"updateFormScript":   p.UpdateFormScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceStatusRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceRequestFromDict(dict2)
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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	UpdateMoldScript   *ScriptSetting      `json:"updateMoldScript"`
	UpdateFormScript   *ScriptSetting      `json:"updateFormScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateNamespaceRequestFromDict(dict2)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		UpdateMoldScript:   NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
		UpdateFormScript:   NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"updateMoldScript":   p.UpdateMoldScript.ToDict(),
		"updateFormScript":   p.UpdateFormScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteNamespaceRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDumpUserDataByUserIdRequestFromDict(dict2)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCleanUserDataByUserIdRequestFromDict(dict2)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict2)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewImportUserDataByUserIdRequestFromDict(dict2)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type GetFormModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
}

func NewGetFormModelRequestFromJson(data string) GetFormModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormModelRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeFormModelMastersRequestFromJson(data string) DescribeFormModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeFormModelMastersRequestFromDict(dict2)
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Name          *string     `json:"name"`
	Description   *string     `json:"description"`
	Metadata      *string     `json:"metadata"`
	Slots         []SlotModel `json:"slots"`
}

func NewCreateFormModelMasterRequestFromJson(data string) CreateFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateFormModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	FormModelName *string `json:"formModelName"`
}

func NewGetFormModelMasterRequestFromJson(data string) GetFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormModelMasterRequestFromDict(dict2)
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	FormModelName *string     `json:"formModelName"`
	Description   *string     `json:"description"`
	Metadata      *string     `json:"metadata"`
	Slots         []SlotModel `json:"slots"`
}

func NewUpdateFormModelMasterRequestFromJson(data string) UpdateFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateFormModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	FormModelName *string `json:"formModelName"`
}

func NewDeleteFormModelMasterRequestFromJson(data string) DeleteFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteFormModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMoldModelsRequestFromJson(data string) DescribeMoldModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeMoldModelsRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
}

func NewGetMoldModelRequestFromJson(data string) GetMoldModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetMoldModelRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeMoldModelMastersRequestFromJson(data string) DescribeMoldModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeMoldModelMastersRequestFromDict(dict2)
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
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateMoldModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
}

func NewGetMoldModelMasterRequestFromJson(data string) GetMoldModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetMoldModelMasterRequestFromDict(dict2)
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
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateMoldModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
}

func NewDeleteMoldModelMasterRequestFromJson(data string) DeleteMoldModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteMoldModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribePropertyFormModelsRequestFromJson(data string) DescribePropertyFormModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribePropertyFormModelsRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewGetPropertyFormModelRequestFromJson(data string) GetPropertyFormModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormModelRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribePropertyFormModelMastersRequestFromJson(data string) DescribePropertyFormModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribePropertyFormModelMastersRequestFromDict(dict2)
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Name          *string     `json:"name"`
	Description   *string     `json:"description"`
	Metadata      *string     `json:"metadata"`
	Slots         []SlotModel `json:"slots"`
}

func NewCreatePropertyFormModelMasterRequestFromJson(data string) CreatePropertyFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreatePropertyFormModelMasterRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewGetPropertyFormModelMasterRequestFromJson(data string) GetPropertyFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormModelMasterRequestFromDict(dict2)
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
	RequestId             *string     `json:"requestId"`
	ContextStack          *string     `json:"contextStack"`
	NamespaceName         *string     `json:"namespaceName"`
	PropertyFormModelName *string     `json:"propertyFormModelName"`
	Description           *string     `json:"description"`
	Metadata              *string     `json:"metadata"`
	Slots                 []SlotModel `json:"slots"`
}

func NewUpdatePropertyFormModelMasterRequestFromJson(data string) UpdatePropertyFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdatePropertyFormModelMasterRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
}

func NewDeletePropertyFormModelMasterRequestFromJson(data string) DeletePropertyFormModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeletePropertyFormModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewExportMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentFormMasterRequestFromJson(data string) GetCurrentFormMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentFormMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentFormMasterRequestFromJson(data string) UpdateCurrentFormMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentFormMasterRequestFromDict(dict2)
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
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentFormMasterFromGitHubRequestFromJson(data string) UpdateCurrentFormMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentFormMasterFromGitHubRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeMoldsRequestFromJson(data string) DescribeMoldsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeMoldsRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeMoldsByUserIdRequestFromJson(data string) DescribeMoldsByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeMoldsByUserIdRequestFromDict(dict2)
}

func NewDescribeMoldsByUserIdRequestFromDict(data map[string]interface{}) DescribeMoldsByUserIdRequest {
	return DescribeMoldsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMoldsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMoldsByUserIdRequest) Pointer() *DescribeMoldsByUserIdRequest {
	return &p
}

type GetMoldRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	MoldModelName *string `json:"moldModelName"`
}

func NewGetMoldRequestFromJson(data string) GetMoldRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetMoldRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	MoldModelName *string `json:"moldModelName"`
}

func NewGetMoldByUserIdRequestFromJson(data string) GetMoldByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetMoldByUserIdRequestFromDict(dict2)
}

func NewGetMoldByUserIdRequestFromDict(data map[string]interface{}) GetMoldByUserIdRequest {
	return GetMoldByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p GetMoldByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
	}
}

func (p GetMoldByUserIdRequest) Pointer() *GetMoldByUserIdRequest {
	return &p
}

type SetMoldCapacityByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
}

func NewSetMoldCapacityByUserIdRequestFromJson(data string) SetMoldCapacityByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetMoldCapacityByUserIdRequestFromDict(dict2)
}

func NewSetMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) SetMoldCapacityByUserIdRequest {
	return SetMoldCapacityByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Capacity:      core.CastInt32(data["capacity"]),
	}
}

func (p SetMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"capacity":      p.Capacity,
	}
}

func (p SetMoldCapacityByUserIdRequest) Pointer() *SetMoldCapacityByUserIdRequest {
	return &p
}

type AddMoldCapacityByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
}

func NewAddMoldCapacityByUserIdRequestFromJson(data string) AddMoldCapacityByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddMoldCapacityByUserIdRequestFromDict(dict2)
}

func NewAddMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) AddMoldCapacityByUserIdRequest {
	return AddMoldCapacityByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Capacity:      core.CastInt32(data["capacity"]),
	}
}

func (p AddMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"capacity":      p.Capacity,
	}
}

func (p AddMoldCapacityByUserIdRequest) Pointer() *AddMoldCapacityByUserIdRequest {
	return &p
}

type SubMoldCapacityByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Capacity           *int32  `json:"capacity"`
}

func NewSubMoldCapacityByUserIdRequestFromJson(data string) SubMoldCapacityByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubMoldCapacityByUserIdRequestFromDict(dict2)
}

func NewSubMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) SubMoldCapacityByUserIdRequest {
	return SubMoldCapacityByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Capacity:      core.CastInt32(data["capacity"]),
	}
}

func (p SubMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"capacity":      p.Capacity,
	}
}

func (p SubMoldCapacityByUserIdRequest) Pointer() *SubMoldCapacityByUserIdRequest {
	return &p
}

type DeleteMoldRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MoldModelName      *string `json:"moldModelName"`
}

func NewDeleteMoldRequestFromJson(data string) DeleteMoldRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteMoldRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
}

func NewDeleteMoldByUserIdRequestFromJson(data string) DeleteMoldByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteMoldByUserIdRequestFromDict(dict2)
}

func NewDeleteMoldByUserIdRequestFromDict(data map[string]interface{}) DeleteMoldByUserIdRequest {
	return DeleteMoldByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p DeleteMoldByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
	}
}

func (p DeleteMoldByUserIdRequest) Pointer() *DeleteMoldByUserIdRequest {
	return &p
}

type AddCapacityByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddCapacityByStampSheetRequestFromJson(data string) AddCapacityByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddCapacityByStampSheetRequestFromDict(dict2)
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
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewSubCapacityByStampTaskRequestFromJson(data string) SubCapacityByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubCapacityByStampTaskRequestFromDict(dict2)
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
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetCapacityByStampSheetRequestFromJson(data string) SetCapacityByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetCapacityByStampSheetRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeFormsRequestFromJson(data string) DescribeFormsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeFormsRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	MoldModelName *string `json:"moldModelName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeFormsByUserIdRequestFromJson(data string) DescribeFormsByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeFormsByUserIdRequestFromDict(dict2)
}

func NewDescribeFormsByUserIdRequestFromDict(data map[string]interface{}) DescribeFormsByUserIdRequest {
	return DescribeFormsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeFormsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeFormsByUserIdRequest) Pointer() *DescribeFormsByUserIdRequest {
	return &p
}

type GetFormRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	MoldModelName *string `json:"moldModelName"`
	Index         *int32  `json:"index"`
}

func NewGetFormRequestFromJson(data string) GetFormRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	MoldModelName *string `json:"moldModelName"`
	Index         *int32  `json:"index"`
}

func NewGetFormByUserIdRequestFromJson(data string) GetFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormByUserIdRequestFromDict(dict2)
}

func NewGetFormByUserIdRequestFromDict(data map[string]interface{}) GetFormByUserIdRequest {
	return GetFormByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
	}
}

func (p GetFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
	}
}

func (p GetFormByUserIdRequest) Pointer() *GetFormByUserIdRequest {
	return &p
}

type GetFormWithSignatureRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	MoldModelName *string `json:"moldModelName"`
	Index         *int32  `json:"index"`
	KeyId         *string `json:"keyId"`
}

func NewGetFormWithSignatureRequestFromJson(data string) GetFormWithSignatureRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormWithSignatureRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	MoldModelName *string `json:"moldModelName"`
	Index         *int32  `json:"index"`
	KeyId         *string `json:"keyId"`
}

func NewGetFormWithSignatureByUserIdRequestFromJson(data string) GetFormWithSignatureByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetFormWithSignatureByUserIdRequestFromDict(dict2)
}

func NewGetFormWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetFormWithSignatureByUserIdRequest {
	return GetFormWithSignatureByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetFormWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"keyId":         p.KeyId,
	}
}

func (p GetFormWithSignatureByUserIdRequest) Pointer() *GetFormWithSignatureByUserIdRequest {
	return &p
}

type SetFormByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
	Slots              []Slot  `json:"slots"`
}

func NewSetFormByUserIdRequestFromJson(data string) SetFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetFormByUserIdRequestFromDict(dict2)
}

func NewSetFormByUserIdRequestFromDict(data map[string]interface{}) SetFormByUserIdRequest {
	return SetFormByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
		Slots:         CastSlots(core.CastArray(data["slots"])),
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
	}
}

func (p SetFormByUserIdRequest) Pointer() *SetFormByUserIdRequest {
	return &p
}

type SetFormWithSignatureRequest struct {
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
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetFormWithSignatureRequestFromDict(dict2)
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
	RequestId          *string               `json:"requestId"`
	ContextStack       *string               `json:"contextStack"`
	DuplicationAvoider *string               `json:"duplicationAvoider"`
	NamespaceName      *string               `json:"namespaceName"`
	UserId             *string               `json:"userId"`
	MoldModelName      *string               `json:"moldModelName"`
	Index              *int32                `json:"index"`
	AcquireAction      *AcquireAction        `json:"acquireAction"`
	Config             []AcquireActionConfig `json:"config"`
}

func NewAcquireActionsToFormPropertiesRequestFromJson(data string) AcquireActionsToFormPropertiesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAcquireActionsToFormPropertiesRequestFromDict(dict2)
}

func NewAcquireActionsToFormPropertiesRequestFromDict(data map[string]interface{}) AcquireActionsToFormPropertiesRequest {
	return AcquireActionsToFormPropertiesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
		AcquireAction: NewAcquireActionFromDict(core.CastMap(data["acquireAction"])).Pointer(),
		Config:        CastAcquireActionConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireActionsToFormPropertiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
		"acquireAction": p.AcquireAction.ToDict(),
		"config": CastAcquireActionConfigsFromDict(
			p.Config,
		),
	}
}

func (p AcquireActionsToFormPropertiesRequest) Pointer() *AcquireActionsToFormPropertiesRequest {
	return &p
}

type DeleteFormRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
}

func NewDeleteFormRequestFromJson(data string) DeleteFormRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteFormRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MoldModelName      *string `json:"moldModelName"`
	Index              *int32  `json:"index"`
}

func NewDeleteFormByUserIdRequestFromJson(data string) DeleteFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteFormByUserIdRequestFromDict(dict2)
}

func NewDeleteFormByUserIdRequestFromDict(data map[string]interface{}) DeleteFormByUserIdRequest {
	return DeleteFormByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Index:         core.CastInt32(data["index"]),
	}
}

func (p DeleteFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"moldModelName": p.MoldModelName,
		"index":         p.Index,
	}
}

func (p DeleteFormByUserIdRequest) Pointer() *DeleteFormByUserIdRequest {
	return &p
}

type AcquireActionToFormPropertiesByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquireActionToFormPropertiesByStampSheetRequestFromJson(data string) AcquireActionToFormPropertiesByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAcquireActionToFormPropertiesByStampSheetRequestFromDict(dict2)
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

type DescribePropertyFormsRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PageToken             *string `json:"pageToken"`
	Limit                 *int32  `json:"limit"`
}

func NewDescribePropertyFormsRequestFromJson(data string) DescribePropertyFormsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribePropertyFormsRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PageToken             *string `json:"pageToken"`
	Limit                 *int32  `json:"limit"`
}

func NewDescribePropertyFormsByUserIdRequestFromJson(data string) DescribePropertyFormsByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribePropertyFormsByUserIdRequestFromDict(dict2)
}

func NewDescribePropertyFormsByUserIdRequestFromDict(data map[string]interface{}) DescribePropertyFormsByUserIdRequest {
	return DescribePropertyFormsByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PageToken:             core.CastString(data["pageToken"]),
		Limit:                 core.CastInt32(data["limit"]),
	}
}

func (p DescribePropertyFormsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"pageToken":             p.PageToken,
		"limit":                 p.Limit,
	}
}

func (p DescribePropertyFormsByUserIdRequest) Pointer() *DescribePropertyFormsByUserIdRequest {
	return &p
}

type GetPropertyFormRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewGetPropertyFormRequestFromJson(data string) GetPropertyFormRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewGetPropertyFormByUserIdRequestFromJson(data string) GetPropertyFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormByUserIdRequestFromDict(dict2)
}

func NewGetPropertyFormByUserIdRequestFromDict(data map[string]interface{}) GetPropertyFormByUserIdRequest {
	return GetPropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
	}
}

func (p GetPropertyFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
	}
}

func (p GetPropertyFormByUserIdRequest) Pointer() *GetPropertyFormByUserIdRequest {
	return &p
}

type GetPropertyFormWithSignatureRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	KeyId                 *string `json:"keyId"`
}

func NewGetPropertyFormWithSignatureRequestFromJson(data string) GetPropertyFormWithSignatureRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormWithSignatureRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	KeyId                 *string `json:"keyId"`
}

func NewGetPropertyFormWithSignatureByUserIdRequestFromJson(data string) GetPropertyFormWithSignatureByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetPropertyFormWithSignatureByUserIdRequestFromDict(dict2)
}

func NewGetPropertyFormWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetPropertyFormWithSignatureByUserIdRequest {
	return GetPropertyFormWithSignatureByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		KeyId:                 core.CastString(data["keyId"]),
	}
}

func (p GetPropertyFormWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"keyId":                 p.KeyId,
	}
}

func (p GetPropertyFormWithSignatureByUserIdRequest) Pointer() *GetPropertyFormWithSignatureByUserIdRequest {
	return &p
}

type SetPropertyFormByUserIdRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
	Slots                 []Slot  `json:"slots"`
}

func NewSetPropertyFormByUserIdRequestFromJson(data string) SetPropertyFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetPropertyFormByUserIdRequestFromDict(dict2)
}

func NewSetPropertyFormByUserIdRequestFromDict(data map[string]interface{}) SetPropertyFormByUserIdRequest {
	return SetPropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		Slots:                 CastSlots(core.CastArray(data["slots"])),
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
	}
}

func (p SetPropertyFormByUserIdRequest) Pointer() *SetPropertyFormByUserIdRequest {
	return &p
}

type SetPropertyFormWithSignatureRequest struct {
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
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetPropertyFormWithSignatureRequestFromDict(dict2)
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
	RequestId             *string               `json:"requestId"`
	ContextStack          *string               `json:"contextStack"`
	DuplicationAvoider    *string               `json:"duplicationAvoider"`
	NamespaceName         *string               `json:"namespaceName"`
	UserId                *string               `json:"userId"`
	PropertyFormModelName *string               `json:"propertyFormModelName"`
	PropertyId            *string               `json:"propertyId"`
	AcquireAction         *AcquireAction        `json:"acquireAction"`
	Config                []AcquireActionConfig `json:"config"`
}

func NewAcquireActionsToPropertyFormPropertiesRequestFromJson(data string) AcquireActionsToPropertyFormPropertiesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAcquireActionsToPropertyFormPropertiesRequestFromDict(dict2)
}

func NewAcquireActionsToPropertyFormPropertiesRequestFromDict(data map[string]interface{}) AcquireActionsToPropertyFormPropertiesRequest {
	return AcquireActionsToPropertyFormPropertiesRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
		AcquireAction:         NewAcquireActionFromDict(core.CastMap(data["acquireAction"])).Pointer(),
		Config:                CastAcquireActionConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireActionsToPropertyFormPropertiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
		"acquireAction":         p.AcquireAction.ToDict(),
		"config": CastAcquireActionConfigsFromDict(
			p.Config,
		),
	}
}

func (p AcquireActionsToPropertyFormPropertiesRequest) Pointer() *AcquireActionsToPropertyFormPropertiesRequest {
	return &p
}

type DeletePropertyFormRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	AccessToken           *string `json:"accessToken"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewDeletePropertyFormRequestFromJson(data string) DeletePropertyFormRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeletePropertyFormRequestFromDict(dict2)
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
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	UserId                *string `json:"userId"`
	PropertyFormModelName *string `json:"propertyFormModelName"`
	PropertyId            *string `json:"propertyId"`
}

func NewDeletePropertyFormByUserIdRequestFromJson(data string) DeletePropertyFormByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeletePropertyFormByUserIdRequestFromDict(dict2)
}

func NewDeletePropertyFormByUserIdRequestFromDict(data map[string]interface{}) DeletePropertyFormByUserIdRequest {
	return DeletePropertyFormByUserIdRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		UserId:                core.CastString(data["userId"]),
		PropertyFormModelName: core.CastString(data["propertyFormModelName"]),
		PropertyId:            core.CastString(data["propertyId"]),
	}
}

func (p DeletePropertyFormByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"userId":                p.UserId,
		"propertyFormModelName": p.PropertyFormModelName,
		"propertyId":            p.PropertyId,
	}
}

func (p DeletePropertyFormByUserIdRequest) Pointer() *DeletePropertyFormByUserIdRequest {
	return &p
}

type AcquireActionToPropertyFormPropertiesByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquireActionToPropertyFormPropertiesByStampSheetRequestFromJson(data string) AcquireActionToPropertyFormPropertiesByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAcquireActionToPropertyFormPropertiesByStampSheetRequestFromDict(dict2)
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
