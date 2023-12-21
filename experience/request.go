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

package experience

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
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	Name                     *string             `json:"name"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	RankCapScriptId          *string             `json:"rankCapScriptId"`
	ChangeExperienceScript   *ScriptSetting      `json:"changeExperienceScript"`
	ChangeRankScript         *ScriptSetting      `json:"changeRankScript"`
	ChangeRankCapScript      *ScriptSetting      `json:"changeRankCapScript"`
	OverflowExperienceScript *string             `json:"overflowExperienceScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
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
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		RankCapScriptId:          core.CastString(data["rankCapScriptId"]),
		ChangeExperienceScript:   NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer(),
		ChangeRankScript:         NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer(),
		ChangeRankCapScript:      NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer(),
		OverflowExperienceScript: core.CastString(data["overflowExperienceScript"]),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                     p.Name,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"rankCapScriptId":          p.RankCapScriptId,
		"changeExperienceScript":   p.ChangeExperienceScript.ToDict(),
		"changeRankScript":         p.ChangeRankScript.ToDict(),
		"changeRankCapScript":      p.ChangeRankCapScript.ToDict(),
		"overflowExperienceScript": p.OverflowExperienceScript,
		"logSetting":               p.LogSetting.ToDict(),
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
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	NamespaceName            *string             `json:"namespaceName"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	RankCapScriptId          *string             `json:"rankCapScriptId"`
	ChangeExperienceScript   *ScriptSetting      `json:"changeExperienceScript"`
	ChangeRankScript         *ScriptSetting      `json:"changeRankScript"`
	ChangeRankCapScript      *ScriptSetting      `json:"changeRankCapScript"`
	OverflowExperienceScript *string             `json:"overflowExperienceScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
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
		NamespaceName:            core.CastString(data["namespaceName"]),
		Description:              core.CastString(data["description"]),
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		RankCapScriptId:          core.CastString(data["rankCapScriptId"]),
		ChangeExperienceScript:   NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer(),
		ChangeRankScript:         NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer(),
		ChangeRankCapScript:      NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer(),
		OverflowExperienceScript: core.CastString(data["overflowExperienceScript"]),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"rankCapScriptId":          p.RankCapScriptId,
		"changeExperienceScript":   p.ChangeExperienceScript.ToDict(),
		"changeRankScript":         p.ChangeRankScript.ToDict(),
		"changeRankCapScript":      p.ChangeRankCapScript.ToDict(),
		"overflowExperienceScript": p.OverflowExperienceScript,
		"logSetting":               p.LogSetting.ToDict(),
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

type DescribeExperienceModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeExperienceModelMastersRequestFromJson(data string) DescribeExperienceModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeExperienceModelMastersRequestFromDict(dict2)
}

func NewDescribeExperienceModelMastersRequestFromDict(data map[string]interface{}) DescribeExperienceModelMastersRequest {
	return DescribeExperienceModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeExperienceModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeExperienceModelMastersRequest) Pointer() *DescribeExperienceModelMastersRequest {
	return &p
}

type CreateExperienceModelMasterRequest struct {
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultExperience  *int64              `json:"defaultExperience"`
	DefaultRankCap     *int64              `json:"defaultRankCap"`
	MaxRankCap         *int64              `json:"maxRankCap"`
	RankThresholdName  *string             `json:"rankThresholdName"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
}

func NewCreateExperienceModelMasterRequestFromJson(data string) CreateExperienceModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateExperienceModelMasterRequestFromDict(dict2)
}

func NewCreateExperienceModelMasterRequestFromDict(data map[string]interface{}) CreateExperienceModelMasterRequest {
	return CreateExperienceModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultExperience:  core.CastInt64(data["defaultExperience"]),
		DefaultRankCap:     core.CastInt64(data["defaultRankCap"]),
		MaxRankCap:         core.CastInt64(data["maxRankCap"]),
		RankThresholdName:  core.CastString(data["rankThresholdName"]),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
	}
}

func (p CreateExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThresholdName": p.RankThresholdName,
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p CreateExperienceModelMasterRequest) Pointer() *CreateExperienceModelMasterRequest {
	return &p
}

type GetExperienceModelMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
}

func NewGetExperienceModelMasterRequestFromJson(data string) GetExperienceModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetExperienceModelMasterRequestFromDict(dict2)
}

func NewGetExperienceModelMasterRequestFromDict(data map[string]interface{}) GetExperienceModelMasterRequest {
	return GetExperienceModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
	}
}

func (p GetExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p GetExperienceModelMasterRequest) Pointer() *GetExperienceModelMasterRequest {
	return &p
}

type UpdateExperienceModelMasterRequest struct {
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	ExperienceName     *string             `json:"experienceName"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultExperience  *int64              `json:"defaultExperience"`
	DefaultRankCap     *int64              `json:"defaultRankCap"`
	MaxRankCap         *int64              `json:"maxRankCap"`
	RankThresholdName  *string             `json:"rankThresholdName"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
}

func NewUpdateExperienceModelMasterRequestFromJson(data string) UpdateExperienceModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateExperienceModelMasterRequestFromDict(dict2)
}

func NewUpdateExperienceModelMasterRequestFromDict(data map[string]interface{}) UpdateExperienceModelMasterRequest {
	return UpdateExperienceModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		ExperienceName:     core.CastString(data["experienceName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultExperience:  core.CastInt64(data["defaultExperience"]),
		DefaultRankCap:     core.CastInt64(data["defaultRankCap"]),
		MaxRankCap:         core.CastInt64(data["maxRankCap"]),
		RankThresholdName:  core.CastString(data["rankThresholdName"]),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
	}
}

func (p UpdateExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"experienceName":    p.ExperienceName,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThresholdName": p.RankThresholdName,
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p UpdateExperienceModelMasterRequest) Pointer() *UpdateExperienceModelMasterRequest {
	return &p
}

type DeleteExperienceModelMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
}

func NewDeleteExperienceModelMasterRequestFromJson(data string) DeleteExperienceModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteExperienceModelMasterRequestFromDict(dict2)
}

func NewDeleteExperienceModelMasterRequestFromDict(data map[string]interface{}) DeleteExperienceModelMasterRequest {
	return DeleteExperienceModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
	}
}

func (p DeleteExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p DeleteExperienceModelMasterRequest) Pointer() *DeleteExperienceModelMasterRequest {
	return &p
}

type DescribeExperienceModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeExperienceModelsRequestFromJson(data string) DescribeExperienceModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeExperienceModelsRequestFromDict(dict2)
}

func NewDescribeExperienceModelsRequestFromDict(data map[string]interface{}) DescribeExperienceModelsRequest {
	return DescribeExperienceModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeExperienceModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeExperienceModelsRequest) Pointer() *DescribeExperienceModelsRequest {
	return &p
}

type GetExperienceModelRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
}

func NewGetExperienceModelRequestFromJson(data string) GetExperienceModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetExperienceModelRequestFromDict(dict2)
}

func NewGetExperienceModelRequestFromDict(data map[string]interface{}) GetExperienceModelRequest {
	return GetExperienceModelRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
	}
}

func (p GetExperienceModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p GetExperienceModelRequest) Pointer() *GetExperienceModelRequest {
	return &p
}

type DescribeThresholdMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeThresholdMastersRequestFromJson(data string) DescribeThresholdMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeThresholdMastersRequestFromDict(dict2)
}

func NewDescribeThresholdMastersRequestFromDict(data map[string]interface{}) DescribeThresholdMastersRequest {
	return DescribeThresholdMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeThresholdMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeThresholdMastersRequest) Pointer() *DescribeThresholdMastersRequest {
	return &p
}

type CreateThresholdMasterRequest struct {
	RequestId     *string  `json:"requestId"`
	ContextStack  *string  `json:"contextStack"`
	NamespaceName *string  `json:"namespaceName"`
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Metadata      *string  `json:"metadata"`
	Values        []*int64 `json:"values"`
}

func NewCreateThresholdMasterRequestFromJson(data string) CreateThresholdMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateThresholdMasterRequestFromDict(dict2)
}

func NewCreateThresholdMasterRequestFromDict(data map[string]interface{}) CreateThresholdMasterRequest {
	return CreateThresholdMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Values:        core.CastInt64s(core.CastArray(data["values"])),
	}
}

func (p CreateThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
	}
}

func (p CreateThresholdMasterRequest) Pointer() *CreateThresholdMasterRequest {
	return &p
}

type GetThresholdMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ThresholdName *string `json:"thresholdName"`
}

func NewGetThresholdMasterRequestFromJson(data string) GetThresholdMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetThresholdMasterRequestFromDict(dict2)
}

func NewGetThresholdMasterRequestFromDict(data map[string]interface{}) GetThresholdMasterRequest {
	return GetThresholdMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ThresholdName: core.CastString(data["thresholdName"]),
	}
}

func (p GetThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
	}
}

func (p GetThresholdMasterRequest) Pointer() *GetThresholdMasterRequest {
	return &p
}

type UpdateThresholdMasterRequest struct {
	RequestId     *string  `json:"requestId"`
	ContextStack  *string  `json:"contextStack"`
	NamespaceName *string  `json:"namespaceName"`
	ThresholdName *string  `json:"thresholdName"`
	Description   *string  `json:"description"`
	Metadata      *string  `json:"metadata"`
	Values        []*int64 `json:"values"`
}

func NewUpdateThresholdMasterRequestFromJson(data string) UpdateThresholdMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateThresholdMasterRequestFromDict(dict2)
}

func NewUpdateThresholdMasterRequestFromDict(data map[string]interface{}) UpdateThresholdMasterRequest {
	return UpdateThresholdMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ThresholdName: core.CastString(data["thresholdName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Values:        core.CastInt64s(core.CastArray(data["values"])),
	}
}

func (p UpdateThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
	}
}

func (p UpdateThresholdMasterRequest) Pointer() *UpdateThresholdMasterRequest {
	return &p
}

type DeleteThresholdMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ThresholdName *string `json:"thresholdName"`
}

func NewDeleteThresholdMasterRequestFromJson(data string) DeleteThresholdMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteThresholdMasterRequestFromDict(dict2)
}

func NewDeleteThresholdMasterRequestFromDict(data map[string]interface{}) DeleteThresholdMasterRequest {
	return DeleteThresholdMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ThresholdName: core.CastString(data["thresholdName"]),
	}
}

func (p DeleteThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
	}
}

func (p DeleteThresholdMasterRequest) Pointer() *DeleteThresholdMasterRequest {
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

type GetCurrentExperienceMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentExperienceMasterRequestFromJson(data string) GetCurrentExperienceMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentExperienceMasterRequestFromDict(dict2)
}

func NewGetCurrentExperienceMasterRequestFromDict(data map[string]interface{}) GetCurrentExperienceMasterRequest {
	return GetCurrentExperienceMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentExperienceMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentExperienceMasterRequest) Pointer() *GetCurrentExperienceMasterRequest {
	return &p
}

type UpdateCurrentExperienceMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentExperienceMasterRequestFromJson(data string) UpdateCurrentExperienceMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentExperienceMasterRequestFromDict(dict2)
}

func NewUpdateCurrentExperienceMasterRequestFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterRequest {
	return UpdateCurrentExperienceMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentExperienceMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentExperienceMasterRequest) Pointer() *UpdateCurrentExperienceMasterRequest {
	return &p
}

type UpdateCurrentExperienceMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentExperienceMasterFromGitHubRequestFromJson(data string) UpdateCurrentExperienceMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentExperienceMasterFromGitHubRequestFromDict(dict2)
}

func NewUpdateCurrentExperienceMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterFromGitHubRequest {
	return UpdateCurrentExperienceMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentExperienceMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentExperienceMasterFromGitHubRequest) Pointer() *UpdateCurrentExperienceMasterFromGitHubRequest {
	return &p
}

type DescribeStatusesRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	AccessToken    *string `json:"accessToken"`
	PageToken      *string `json:"pageToken"`
	Limit          *int32  `json:"limit"`
}

func NewDescribeStatusesRequestFromJson(data string) DescribeStatusesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeStatusesRequestFromDict(dict2)
}

func NewDescribeStatusesRequestFromDict(data map[string]interface{}) DescribeStatusesRequest {
	return DescribeStatusesRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
		"accessToken":    p.AccessToken,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
	return &p
}

type DescribeStatusesByUserIdRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	UserId         *string `json:"userId"`
	PageToken      *string `json:"pageToken"`
	Limit          *int32  `json:"limit"`
}

func NewDescribeStatusesByUserIdRequestFromJson(data string) DescribeStatusesByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeStatusesByUserIdRequestFromDict(dict2)
}

func NewDescribeStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeStatusesByUserIdRequest {
	return DescribeStatusesByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
		UserId:         core.CastString(data["userId"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
		"userId":         p.UserId,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeStatusesByUserIdRequest) Pointer() *DescribeStatusesByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AccessToken    *string `json:"accessToken"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
}

func NewGetStatusRequestFromJson(data string) GetStatusRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetStatusRequestFromDict(dict2)
}

func NewGetStatusRequestFromDict(data map[string]interface{}) GetStatusRequest {
	return GetStatusRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
	}
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
	return &p
}

type GetStatusByUserIdRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	UserId         *string `json:"userId"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
}

func NewGetStatusByUserIdRequestFromJson(data string) GetStatusByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetStatusByUserIdRequestFromDict(dict2)
}

func NewGetStatusByUserIdRequestFromDict(data map[string]interface{}) GetStatusByUserIdRequest {
	return GetStatusByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type GetStatusWithSignatureRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AccessToken    *string `json:"accessToken"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
	KeyId          *string `json:"keyId"`
}

func NewGetStatusWithSignatureRequestFromJson(data string) GetStatusWithSignatureRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetStatusWithSignatureRequestFromDict(dict2)
}

func NewGetStatusWithSignatureRequestFromDict(data map[string]interface{}) GetStatusWithSignatureRequest {
	return GetStatusWithSignatureRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetStatusWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"keyId":          p.KeyId,
	}
}

func (p GetStatusWithSignatureRequest) Pointer() *GetStatusWithSignatureRequest {
	return &p
}

type GetStatusWithSignatureByUserIdRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	UserId         *string `json:"userId"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
	KeyId          *string `json:"keyId"`
}

func NewGetStatusWithSignatureByUserIdRequestFromJson(data string) GetStatusWithSignatureByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetStatusWithSignatureByUserIdRequestFromDict(dict2)
}

func NewGetStatusWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetStatusWithSignatureByUserIdRequest {
	return GetStatusWithSignatureByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetStatusWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"keyId":          p.KeyId,
	}
}

func (p GetStatusWithSignatureByUserIdRequest) Pointer() *GetStatusWithSignatureByUserIdRequest {
	return &p
}

type AddExperienceByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
}

func NewAddExperienceByUserIdRequestFromJson(data string) AddExperienceByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddExperienceByUserIdRequestFromDict(dict2)
}

func NewAddExperienceByUserIdRequestFromDict(data map[string]interface{}) AddExperienceByUserIdRequest {
	return AddExperienceByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ExperienceName:  core.CastString(data["experienceName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt64(data["experienceValue"]),
	}
}

func (p AddExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
	}
}

func (p AddExperienceByUserIdRequest) Pointer() *AddExperienceByUserIdRequest {
	return &p
}

type SubExperienceByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
}

func NewSubExperienceByUserIdRequestFromJson(data string) SubExperienceByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubExperienceByUserIdRequestFromDict(dict2)
}

func NewSubExperienceByUserIdRequestFromDict(data map[string]interface{}) SubExperienceByUserIdRequest {
	return SubExperienceByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ExperienceName:  core.CastString(data["experienceName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt64(data["experienceValue"]),
	}
}

func (p SubExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
	}
}

func (p SubExperienceByUserIdRequest) Pointer() *SubExperienceByUserIdRequest {
	return &p
}

type SetExperienceByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
}

func NewSetExperienceByUserIdRequestFromJson(data string) SetExperienceByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetExperienceByUserIdRequestFromDict(dict2)
}

func NewSetExperienceByUserIdRequestFromDict(data map[string]interface{}) SetExperienceByUserIdRequest {
	return SetExperienceByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ExperienceName:  core.CastString(data["experienceName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt64(data["experienceValue"]),
	}
}

func (p SetExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
	}
}

func (p SetExperienceByUserIdRequest) Pointer() *SetExperienceByUserIdRequest {
	return &p
}

type AddRankCapByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
}

func NewAddRankCapByUserIdRequestFromJson(data string) AddRankCapByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddRankCapByUserIdRequestFromDict(dict2)
}

func NewAddRankCapByUserIdRequestFromDict(data map[string]interface{}) AddRankCapByUserIdRequest {
	return AddRankCapByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankCapValue:   core.CastInt64(data["rankCapValue"]),
	}
}

func (p AddRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p AddRankCapByUserIdRequest) Pointer() *AddRankCapByUserIdRequest {
	return &p
}

type SubRankCapByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
}

func NewSubRankCapByUserIdRequestFromJson(data string) SubRankCapByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubRankCapByUserIdRequestFromDict(dict2)
}

func NewSubRankCapByUserIdRequestFromDict(data map[string]interface{}) SubRankCapByUserIdRequest {
	return SubRankCapByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankCapValue:   core.CastInt64(data["rankCapValue"]),
	}
}

func (p SubRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p SubRankCapByUserIdRequest) Pointer() *SubRankCapByUserIdRequest {
	return &p
}

type SetRankCapByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
}

func NewSetRankCapByUserIdRequestFromJson(data string) SetRankCapByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetRankCapByUserIdRequestFromDict(dict2)
}

func NewSetRankCapByUserIdRequestFromDict(data map[string]interface{}) SetRankCapByUserIdRequest {
	return SetRankCapByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankCapValue:   core.CastInt64(data["rankCapValue"]),
	}
}

func (p SetRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p SetRankCapByUserIdRequest) Pointer() *SetRankCapByUserIdRequest {
	return &p
}

type DeleteStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
}

func NewDeleteStatusByUserIdRequestFromJson(data string) DeleteStatusByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteStatusByUserIdRequestFromDict(dict2)
}

func NewDeleteStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteStatusByUserIdRequest {
	return DeleteStatusByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
	}
}

func (p DeleteStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
	}
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
	return &p
}

type VerifyRankRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	ExperienceName     *string `json:"experienceName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	RankValue          *int64  `json:"rankValue"`
}

func NewVerifyRankRequestFromJson(data string) VerifyRankRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankRequestFromDict(dict2)
}

func NewVerifyRankRequestFromDict(data map[string]interface{}) VerifyRankRequest {
	return VerifyRankRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		ExperienceName: core.CastString(data["experienceName"]),
		VerifyType:     core.CastString(data["verifyType"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankValue:      core.CastInt64(data["rankValue"]),
	}
}

func (p VerifyRankRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"verifyType":     p.VerifyType,
		"propertyId":     p.PropertyId,
		"rankValue":      p.RankValue,
	}
}

func (p VerifyRankRequest) Pointer() *VerifyRankRequest {
	return &p
}

type VerifyRankByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	RankValue          *int64  `json:"rankValue"`
}

func NewVerifyRankByUserIdRequestFromJson(data string) VerifyRankByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankByUserIdRequestFromDict(dict2)
}

func NewVerifyRankByUserIdRequestFromDict(data map[string]interface{}) VerifyRankByUserIdRequest {
	return VerifyRankByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		VerifyType:     core.CastString(data["verifyType"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankValue:      core.CastInt64(data["rankValue"]),
	}
}

func (p VerifyRankByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"verifyType":     p.VerifyType,
		"propertyId":     p.PropertyId,
		"rankValue":      p.RankValue,
	}
}

func (p VerifyRankByUserIdRequest) Pointer() *VerifyRankByUserIdRequest {
	return &p
}

type VerifyRankCapRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	ExperienceName     *string `json:"experienceName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
}

func NewVerifyRankCapRequestFromJson(data string) VerifyRankCapRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankCapRequestFromDict(dict2)
}

func NewVerifyRankCapRequestFromDict(data map[string]interface{}) VerifyRankCapRequest {
	return VerifyRankCapRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		ExperienceName: core.CastString(data["experienceName"]),
		VerifyType:     core.CastString(data["verifyType"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankCapValue:   core.CastInt64(data["rankCapValue"]),
	}
}

func (p VerifyRankCapRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"verifyType":     p.VerifyType,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p VerifyRankCapRequest) Pointer() *VerifyRankCapRequest {
	return &p
}

type VerifyRankCapByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
}

func NewVerifyRankCapByUserIdRequestFromJson(data string) VerifyRankCapByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankCapByUserIdRequestFromDict(dict2)
}

func NewVerifyRankCapByUserIdRequestFromDict(data map[string]interface{}) VerifyRankCapByUserIdRequest {
	return VerifyRankCapByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		VerifyType:     core.CastString(data["verifyType"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RankCapValue:   core.CastInt64(data["rankCapValue"]),
	}
}

func (p VerifyRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"verifyType":     p.VerifyType,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p VerifyRankCapByUserIdRequest) Pointer() *VerifyRankCapByUserIdRequest {
	return &p
}

type AddExperienceByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddExperienceByStampSheetRequestFromJson(data string) AddExperienceByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddExperienceByStampSheetRequestFromDict(dict2)
}

func NewAddExperienceByStampSheetRequestFromDict(data map[string]interface{}) AddExperienceByStampSheetRequest {
	return AddExperienceByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddExperienceByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddExperienceByStampSheetRequest) Pointer() *AddExperienceByStampSheetRequest {
	return &p
}

type SubExperienceByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewSubExperienceByStampTaskRequestFromJson(data string) SubExperienceByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubExperienceByStampTaskRequestFromDict(dict2)
}

func NewSubExperienceByStampTaskRequestFromDict(data map[string]interface{}) SubExperienceByStampTaskRequest {
	return SubExperienceByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p SubExperienceByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubExperienceByStampTaskRequest) Pointer() *SubExperienceByStampTaskRequest {
	return &p
}

type AddRankCapByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddRankCapByStampSheetRequestFromJson(data string) AddRankCapByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewAddRankCapByStampSheetRequestFromDict(dict2)
}

func NewAddRankCapByStampSheetRequestFromDict(data map[string]interface{}) AddRankCapByStampSheetRequest {
	return AddRankCapByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddRankCapByStampSheetRequest) Pointer() *AddRankCapByStampSheetRequest {
	return &p
}

type SubRankCapByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewSubRankCapByStampTaskRequestFromJson(data string) SubRankCapByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSubRankCapByStampTaskRequestFromDict(dict2)
}

func NewSubRankCapByStampTaskRequestFromDict(data map[string]interface{}) SubRankCapByStampTaskRequest {
	return SubRankCapByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p SubRankCapByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubRankCapByStampTaskRequest) Pointer() *SubRankCapByStampTaskRequest {
	return &p
}

type SetRankCapByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetRankCapByStampSheetRequestFromJson(data string) SetRankCapByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewSetRankCapByStampSheetRequestFromDict(dict2)
}

func NewSetRankCapByStampSheetRequestFromDict(data map[string]interface{}) SetRankCapByStampSheetRequest {
	return SetRankCapByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRankCapByStampSheetRequest) Pointer() *SetRankCapByStampSheetRequest {
	return &p
}

type MultiplyAcquireActionsByUserIdRequest struct {
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	ExperienceName     *string         `json:"experienceName"`
	PropertyId         *string         `json:"propertyId"`
	RateName           *string         `json:"rateName"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
}

func NewMultiplyAcquireActionsByUserIdRequestFromJson(data string) MultiplyAcquireActionsByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMultiplyAcquireActionsByUserIdRequestFromDict(dict2)
}

func NewMultiplyAcquireActionsByUserIdRequestFromDict(data map[string]interface{}) MultiplyAcquireActionsByUserIdRequest {
	return MultiplyAcquireActionsByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		RateName:       core.CastString(data["rateName"]),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p MultiplyAcquireActionsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rateName":       p.RateName,
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
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMultiplyAcquireActionsByStampSheetRequestFromDict(dict2)
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

type VerifyRankByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyRankByStampTaskRequestFromJson(data string) VerifyRankByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankByStampTaskRequestFromDict(dict2)
}

func NewVerifyRankByStampTaskRequestFromDict(data map[string]interface{}) VerifyRankByStampTaskRequest {
	return VerifyRankByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyRankByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyRankByStampTaskRequest) Pointer() *VerifyRankByStampTaskRequest {
	return &p
}

type VerifyRankCapByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyRankCapByStampTaskRequestFromJson(data string) VerifyRankCapByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyRankCapByStampTaskRequestFromDict(dict2)
}

func NewVerifyRankCapByStampTaskRequestFromDict(data map[string]interface{}) VerifyRankCapByStampTaskRequest {
	return VerifyRankCapByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyRankCapByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyRankCapByStampTaskRequest) Pointer() *VerifyRankCapByStampTaskRequest {
	return &p
}
