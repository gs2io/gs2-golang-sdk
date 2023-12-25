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

package grade

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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ChangeGradeScript  *ScriptSetting      `json:"changeGradeScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
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
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ChangeGradeScript:  NewScriptSettingFromDict(core.CastMap(data["changeGradeScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"changeGradeScript":  p.ChangeGradeScript.ToDict(),
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ChangeGradeScript  *ScriptSetting      `json:"changeGradeScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
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
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ChangeGradeScript:  NewScriptSettingFromDict(core.CastMap(data["changeGradeScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"changeGradeScript":  p.ChangeGradeScript.ToDict(),
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
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

type DescribeGradeModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeGradeModelMastersRequestFromJson(data string) DescribeGradeModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelMastersRequestFromDict(dict)
}

func NewDescribeGradeModelMastersRequestFromDict(data map[string]interface{}) DescribeGradeModelMastersRequest {
	return DescribeGradeModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeGradeModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGradeModelMastersRequest) Pointer() *DescribeGradeModelMastersRequest {
	return &p
}

type CreateGradeModelMasterRequest struct {
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
}

func NewCreateGradeModelMasterRequestFromJson(data string) CreateGradeModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGradeModelMasterRequestFromDict(dict)
}

func NewCreateGradeModelMasterRequestFromDict(data map[string]interface{}) CreateGradeModelMasterRequest {
	return CreateGradeModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultGrades:      CastDefaultGradeModels(core.CastArray(data["defaultGrades"])),
		ExperienceModelId:  core.CastString(data["experienceModelId"]),
		GradeEntries:       CastGradeEntryModels(core.CastArray(data["gradeEntries"])),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
	}
}

func (p CreateGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"defaultGrades": CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		),
		"experienceModelId": p.ExperienceModelId,
		"gradeEntries": CastGradeEntryModelsFromDict(
			p.GradeEntries,
		),
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p CreateGradeModelMasterRequest) Pointer() *CreateGradeModelMasterRequest {
	return &p
}

type GetGradeModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
}

func NewGetGradeModelMasterRequestFromJson(data string) GetGradeModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelMasterRequestFromDict(dict)
}

func NewGetGradeModelMasterRequestFromDict(data map[string]interface{}) GetGradeModelMasterRequest {
	return GetGradeModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GradeName:     core.CastString(data["gradeName"]),
	}
}

func (p GetGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p GetGradeModelMasterRequest) Pointer() *GetGradeModelMasterRequest {
	return &p
}

type UpdateGradeModelMasterRequest struct {
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	GradeName          *string             `json:"gradeName"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
}

func NewUpdateGradeModelMasterRequestFromJson(data string) UpdateGradeModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGradeModelMasterRequestFromDict(dict)
}

func NewUpdateGradeModelMasterRequestFromDict(data map[string]interface{}) UpdateGradeModelMasterRequest {
	return UpdateGradeModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		GradeName:          core.CastString(data["gradeName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultGrades:      CastDefaultGradeModels(core.CastArray(data["defaultGrades"])),
		ExperienceModelId:  core.CastString(data["experienceModelId"]),
		GradeEntries:       CastGradeEntryModels(core.CastArray(data["gradeEntries"])),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
	}
}

func (p UpdateGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"defaultGrades": CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		),
		"experienceModelId": p.ExperienceModelId,
		"gradeEntries": CastGradeEntryModelsFromDict(
			p.GradeEntries,
		),
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p UpdateGradeModelMasterRequest) Pointer() *UpdateGradeModelMasterRequest {
	return &p
}

type DeleteGradeModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
}

func NewDeleteGradeModelMasterRequestFromJson(data string) DeleteGradeModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGradeModelMasterRequestFromDict(dict)
}

func NewDeleteGradeModelMasterRequestFromDict(data map[string]interface{}) DeleteGradeModelMasterRequest {
	return DeleteGradeModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GradeName:     core.CastString(data["gradeName"]),
	}
}

func (p DeleteGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p DeleteGradeModelMasterRequest) Pointer() *DeleteGradeModelMasterRequest {
	return &p
}

type DescribeGradeModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeGradeModelsRequestFromJson(data string) DescribeGradeModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelsRequestFromDict(dict)
}

func NewDescribeGradeModelsRequestFromDict(data map[string]interface{}) DescribeGradeModelsRequest {
	return DescribeGradeModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeGradeModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGradeModelsRequest) Pointer() *DescribeGradeModelsRequest {
	return &p
}

type GetGradeModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
}

func NewGetGradeModelRequestFromJson(data string) GetGradeModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelRequestFromDict(dict)
}

func NewGetGradeModelRequestFromDict(data map[string]interface{}) GetGradeModelRequest {
	return GetGradeModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GradeName:     core.CastString(data["gradeName"]),
	}
}

func (p GetGradeModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p GetGradeModelRequest) Pointer() *GetGradeModelRequest {
	return &p
}

type DescribeStatusesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStatusesRequestFromJson(data string) DescribeStatusesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesRequestFromDict(dict)
}

func NewDescribeStatusesRequestFromDict(data map[string]interface{}) DescribeStatusesRequest {
	return DescribeStatusesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GradeName:     core.CastString(data["gradeName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
	return &p
}

type DescribeStatusesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStatusesByUserIdRequestFromJson(data string) DescribeStatusesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeStatusesByUserIdRequest {
	return DescribeStatusesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GradeName:     core.CastString(data["gradeName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusesByUserIdRequest) Pointer() *DescribeStatusesByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	GradeName     *string `json:"gradeName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetStatusRequestFromJson(data string) GetStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusRequestFromDict(dict)
}

func NewGetStatusRequestFromDict(data map[string]interface{}) GetStatusRequest {
	return GetStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
	return &p
}

type GetStatusByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	GradeName     *string `json:"gradeName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetStatusByUserIdRequestFromJson(data string) GetStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusByUserIdRequestFromDict(dict)
}

func NewGetStatusByUserIdRequestFromDict(data map[string]interface{}) GetStatusByUserIdRequest {
	return GetStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type AddGradeByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
}

func NewAddGradeByUserIdRequestFromJson(data string) AddGradeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByUserIdRequestFromDict(dict)
}

func NewAddGradeByUserIdRequestFromDict(data map[string]interface{}) AddGradeByUserIdRequest {
	return AddGradeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		GradeValue:    core.CastInt64(data["gradeValue"]),
	}
}

func (p AddGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p AddGradeByUserIdRequest) Pointer() *AddGradeByUserIdRequest {
	return &p
}

type SubGradeByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
}

func NewSubGradeByUserIdRequestFromJson(data string) SubGradeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByUserIdRequestFromDict(dict)
}

func NewSubGradeByUserIdRequestFromDict(data map[string]interface{}) SubGradeByUserIdRequest {
	return SubGradeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		GradeValue:    core.CastInt64(data["gradeValue"]),
	}
}

func (p SubGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p SubGradeByUserIdRequest) Pointer() *SubGradeByUserIdRequest {
	return &p
}

type SetGradeByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
}

func NewSetGradeByUserIdRequestFromJson(data string) SetGradeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetGradeByUserIdRequestFromDict(dict)
}

func NewSetGradeByUserIdRequestFromDict(data map[string]interface{}) SetGradeByUserIdRequest {
	return SetGradeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		GradeValue:    core.CastInt64(data["gradeValue"]),
	}
}

func (p SetGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p SetGradeByUserIdRequest) Pointer() *SetGradeByUserIdRequest {
	return &p
}

type ApplyRankCapRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
}

func NewApplyRankCapRequestFromJson(data string) ApplyRankCapRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapRequestFromDict(dict)
}

func NewApplyRankCapRequestFromDict(data map[string]interface{}) ApplyRankCapRequest {
	return ApplyRankCapRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p ApplyRankCapRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p ApplyRankCapRequest) Pointer() *ApplyRankCapRequest {
	return &p
}

type ApplyRankCapByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
}

func NewApplyRankCapByUserIdRequestFromJson(data string) ApplyRankCapByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByUserIdRequestFromDict(dict)
}

func NewApplyRankCapByUserIdRequestFromDict(data map[string]interface{}) ApplyRankCapByUserIdRequest {
	return ApplyRankCapByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p ApplyRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p ApplyRankCapByUserIdRequest) Pointer() *ApplyRankCapByUserIdRequest {
	return &p
}

type DeleteStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
}

func NewDeleteStatusByUserIdRequestFromJson(data string) DeleteStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStatusByUserIdRequestFromDict(dict)
}

func NewDeleteStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteStatusByUserIdRequest {
	return DeleteStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p DeleteStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
	return &p
}

type VerifyGradeRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
}

func NewVerifyGradeRequestFromJson(data string) VerifyGradeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeRequestFromDict(dict)
}

func NewVerifyGradeRequestFromDict(data map[string]interface{}) VerifyGradeRequest {
	return VerifyGradeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		GradeName:     core.CastString(data["gradeName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		PropertyId:    core.CastString(data["propertyId"]),
		GradeValue:    core.CastInt64(data["gradeValue"]),
	}
}

func (p VerifyGradeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"verifyType":    p.VerifyType,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p VerifyGradeRequest) Pointer() *VerifyGradeRequest {
	return &p
}

type VerifyGradeByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
}

func NewVerifyGradeByUserIdRequestFromJson(data string) VerifyGradeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByUserIdRequestFromDict(dict)
}

func NewVerifyGradeByUserIdRequestFromDict(data map[string]interface{}) VerifyGradeByUserIdRequest {
	return VerifyGradeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		GradeName:     core.CastString(data["gradeName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		PropertyId:    core.CastString(data["propertyId"]),
		GradeValue:    core.CastInt64(data["gradeValue"]),
	}
}

func (p VerifyGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"verifyType":    p.VerifyType,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p VerifyGradeByUserIdRequest) Pointer() *VerifyGradeByUserIdRequest {
	return &p
}

type VerifyGradeUpMaterialRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	MaterialPropertyId *string `json:"materialPropertyId"`
}

func NewVerifyGradeUpMaterialRequestFromJson(data string) VerifyGradeUpMaterialRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialRequestFromDict(dict)
}

func NewVerifyGradeUpMaterialRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialRequest {
	return VerifyGradeUpMaterialRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		AccessToken:        core.CastString(data["accessToken"]),
		GradeName:          core.CastString(data["gradeName"]),
		VerifyType:         core.CastString(data["verifyType"]),
		PropertyId:         core.CastString(data["propertyId"]),
		MaterialPropertyId: core.CastString(data["materialPropertyId"]),
	}
}

func (p VerifyGradeUpMaterialRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"accessToken":        p.AccessToken,
		"gradeName":          p.GradeName,
		"verifyType":         p.VerifyType,
		"propertyId":         p.PropertyId,
		"materialPropertyId": p.MaterialPropertyId,
	}
}

func (p VerifyGradeUpMaterialRequest) Pointer() *VerifyGradeUpMaterialRequest {
	return &p
}

type VerifyGradeUpMaterialByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	MaterialPropertyId *string `json:"materialPropertyId"`
}

func NewVerifyGradeUpMaterialByUserIdRequestFromJson(data string) VerifyGradeUpMaterialByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByUserIdRequestFromDict(dict)
}

func NewVerifyGradeUpMaterialByUserIdRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialByUserIdRequest {
	return VerifyGradeUpMaterialByUserIdRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		UserId:             core.CastString(data["userId"]),
		GradeName:          core.CastString(data["gradeName"]),
		VerifyType:         core.CastString(data["verifyType"]),
		PropertyId:         core.CastString(data["propertyId"]),
		MaterialPropertyId: core.CastString(data["materialPropertyId"]),
	}
}

func (p VerifyGradeUpMaterialByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"userId":             p.UserId,
		"gradeName":          p.GradeName,
		"verifyType":         p.VerifyType,
		"propertyId":         p.PropertyId,
		"materialPropertyId": p.MaterialPropertyId,
	}
}

func (p VerifyGradeUpMaterialByUserIdRequest) Pointer() *VerifyGradeUpMaterialByUserIdRequest {
	return &p
}

type AddGradeByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddGradeByStampSheetRequestFromJson(data string) AddGradeByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByStampSheetRequestFromDict(dict)
}

func NewAddGradeByStampSheetRequestFromDict(data map[string]interface{}) AddGradeByStampSheetRequest {
	return AddGradeByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddGradeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddGradeByStampSheetRequest) Pointer() *AddGradeByStampSheetRequest {
	return &p
}

type ApplyRankCapByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewApplyRankCapByStampSheetRequestFromJson(data string) ApplyRankCapByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByStampSheetRequestFromDict(dict)
}

func NewApplyRankCapByStampSheetRequestFromDict(data map[string]interface{}) ApplyRankCapByStampSheetRequest {
	return ApplyRankCapByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p ApplyRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ApplyRankCapByStampSheetRequest) Pointer() *ApplyRankCapByStampSheetRequest {
	return &p
}

type SubGradeByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewSubGradeByStampTaskRequestFromJson(data string) SubGradeByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByStampTaskRequestFromDict(dict)
}

func NewSubGradeByStampTaskRequestFromDict(data map[string]interface{}) SubGradeByStampTaskRequest {
	return SubGradeByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p SubGradeByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubGradeByStampTaskRequest) Pointer() *SubGradeByStampTaskRequest {
	return &p
}

type MultiplyAcquireActionsByUserIdRequest struct {
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	GradeName          *string         `json:"gradeName"`
	PropertyId         *string         `json:"propertyId"`
	RateName           *string         `json:"rateName"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
}

func NewMultiplyAcquireActionsByUserIdRequestFromJson(data string) MultiplyAcquireActionsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByUserIdRequestFromDict(dict)
}

func NewMultiplyAcquireActionsByUserIdRequestFromDict(data map[string]interface{}) MultiplyAcquireActionsByUserIdRequest {
	return MultiplyAcquireActionsByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		GradeName:      core.CastString(data["gradeName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RateName:       core.CastString(data["rateName"]),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p MultiplyAcquireActionsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"rateName":      p.RateName,
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p MultiplyAcquireActionsByUserIdRequest) Pointer() *MultiplyAcquireActionsByUserIdRequest {
	return &p
}

type MultiplyAcquireActionsByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewMultiplyAcquireActionsByStampSheetRequestFromJson(data string) MultiplyAcquireActionsByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByStampSheetRequestFromDict(dict)
}

func NewMultiplyAcquireActionsByStampSheetRequestFromDict(data map[string]interface{}) MultiplyAcquireActionsByStampSheetRequest {
	return MultiplyAcquireActionsByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p MultiplyAcquireActionsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p MultiplyAcquireActionsByStampSheetRequest) Pointer() *MultiplyAcquireActionsByStampSheetRequest {
	return &p
}

type VerifyGradeByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyGradeByStampTaskRequestFromJson(data string) VerifyGradeByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByStampTaskRequestFromDict(dict)
}

func NewVerifyGradeByStampTaskRequestFromDict(data map[string]interface{}) VerifyGradeByStampTaskRequest {
	return VerifyGradeByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyGradeByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyGradeByStampTaskRequest) Pointer() *VerifyGradeByStampTaskRequest {
	return &p
}

type VerifyGradeUpMaterialByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyGradeUpMaterialByStampTaskRequestFromJson(data string) VerifyGradeUpMaterialByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByStampTaskRequestFromDict(dict)
}

func NewVerifyGradeUpMaterialByStampTaskRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialByStampTaskRequest {
	return VerifyGradeUpMaterialByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyGradeUpMaterialByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyGradeUpMaterialByStampTaskRequest) Pointer() *VerifyGradeUpMaterialByStampTaskRequest {
	return &p
}

type ExportMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
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

type GetCurrentGradeMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentGradeMasterRequestFromJson(data string) GetCurrentGradeMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentGradeMasterRequestFromDict(dict)
}

func NewGetCurrentGradeMasterRequestFromDict(data map[string]interface{}) GetCurrentGradeMasterRequest {
	return GetCurrentGradeMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentGradeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentGradeMasterRequest) Pointer() *GetCurrentGradeMasterRequest {
	return &p
}

type UpdateCurrentGradeMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentGradeMasterRequestFromJson(data string) UpdateCurrentGradeMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterRequestFromDict(dict)
}

func NewUpdateCurrentGradeMasterRequestFromDict(data map[string]interface{}) UpdateCurrentGradeMasterRequest {
	return UpdateCurrentGradeMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentGradeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentGradeMasterRequest) Pointer() *UpdateCurrentGradeMasterRequest {
	return &p
}

type UpdateCurrentGradeMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentGradeMasterFromGitHubRequestFromJson(data string) UpdateCurrentGradeMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentGradeMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentGradeMasterFromGitHubRequest {
	return UpdateCurrentGradeMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentGradeMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentGradeMasterFromGitHubRequest) Pointer() *UpdateCurrentGradeMasterFromGitHubRequest {
	return &p
}
