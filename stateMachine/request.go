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

package stateMachine

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
	RequestId                   *string             `json:"requestId"`
	ContextStack                *string             `json:"contextStack"`
	Name                        *string             `json:"name"`
	Description                 *string             `json:"description"`
	SupportSpeculativeExecution *string             `json:"supportSpeculativeExecution"`
	TransactionSetting          *TransactionSetting `json:"transactionSetting"`
	StartScript                 *ScriptSetting      `json:"startScript"`
	PassScript                  *ScriptSetting      `json:"passScript"`
	ErrorScript                 *ScriptSetting      `json:"errorScript"`
	LowestStateMachineVersion   *int64              `json:"lowestStateMachineVersion"`
	LogSetting                  *LogSetting         `json:"logSetting"`
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
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		SupportSpeculativeExecution: core.CastString(data["supportSpeculativeExecution"]),
		TransactionSetting:          NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		StartScript:                 NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
		PassScript:                  NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
		ErrorScript:                 NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
		LowestStateMachineVersion:   core.CastInt64(data["lowestStateMachineVersion"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                        p.Name,
		"description":                 p.Description,
		"supportSpeculativeExecution": p.SupportSpeculativeExecution,
		"transactionSetting":          p.TransactionSetting.ToDict(),
		"startScript":                 p.StartScript.ToDict(),
		"passScript":                  p.PassScript.ToDict(),
		"errorScript":                 p.ErrorScript.ToDict(),
		"lowestStateMachineVersion":   p.LowestStateMachineVersion,
		"logSetting":                  p.LogSetting.ToDict(),
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
	RequestId                   *string             `json:"requestId"`
	ContextStack                *string             `json:"contextStack"`
	NamespaceName               *string             `json:"namespaceName"`
	Description                 *string             `json:"description"`
	SupportSpeculativeExecution *string             `json:"supportSpeculativeExecution"`
	TransactionSetting          *TransactionSetting `json:"transactionSetting"`
	StartScript                 *ScriptSetting      `json:"startScript"`
	PassScript                  *ScriptSetting      `json:"passScript"`
	ErrorScript                 *ScriptSetting      `json:"errorScript"`
	LowestStateMachineVersion   *int64              `json:"lowestStateMachineVersion"`
	LogSetting                  *LogSetting         `json:"logSetting"`
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
		NamespaceName:               core.CastString(data["namespaceName"]),
		Description:                 core.CastString(data["description"]),
		SupportSpeculativeExecution: core.CastString(data["supportSpeculativeExecution"]),
		TransactionSetting:          NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		StartScript:                 NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
		PassScript:                  NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
		ErrorScript:                 NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
		LowestStateMachineVersion:   core.CastInt64(data["lowestStateMachineVersion"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":               p.NamespaceName,
		"description":                 p.Description,
		"supportSpeculativeExecution": p.SupportSpeculativeExecution,
		"transactionSetting":          p.TransactionSetting.ToDict(),
		"startScript":                 p.StartScript.ToDict(),
		"passScript":                  p.PassScript.ToDict(),
		"errorScript":                 p.ErrorScript.ToDict(),
		"lowestStateMachineVersion":   p.LowestStateMachineVersion,
		"logSetting":                  p.LogSetting.ToDict(),
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

type DescribeStateMachineMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStateMachineMastersRequestFromJson(data string) DescribeStateMachineMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeStateMachineMastersRequestFromDict(dict2)
}

func NewDescribeStateMachineMastersRequestFromDict(data map[string]interface{}) DescribeStateMachineMastersRequest {
	return DescribeStateMachineMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStateMachineMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStateMachineMastersRequest) Pointer() *DescribeStateMachineMastersRequest {
	return &p
}

type UpdateStateMachineMasterRequest struct {
	RequestId            *string `json:"requestId"`
	ContextStack         *string `json:"contextStack"`
	NamespaceName        *string `json:"namespaceName"`
	MainStateMachineName *string `json:"mainStateMachineName"`
	Payload              *string `json:"payload"`
}

func NewUpdateStateMachineMasterRequestFromJson(data string) UpdateStateMachineMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateStateMachineMasterRequestFromDict(dict2)
}

func NewUpdateStateMachineMasterRequestFromDict(data map[string]interface{}) UpdateStateMachineMasterRequest {
	return UpdateStateMachineMasterRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		MainStateMachineName: core.CastString(data["mainStateMachineName"]),
		Payload:              core.CastString(data["payload"]),
	}
}

func (p UpdateStateMachineMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"mainStateMachineName": p.MainStateMachineName,
		"payload":              p.Payload,
	}
}

func (p UpdateStateMachineMasterRequest) Pointer() *UpdateStateMachineMasterRequest {
	return &p
}

type GetStateMachineMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Version       *int64  `json:"version"`
}

func NewGetStateMachineMasterRequestFromJson(data string) GetStateMachineMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetStateMachineMasterRequestFromDict(dict2)
}

func NewGetStateMachineMasterRequestFromDict(data map[string]interface{}) GetStateMachineMasterRequest {
	return GetStateMachineMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Version:       core.CastInt64(data["version"]),
	}
}

func (p GetStateMachineMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"version":       p.Version,
	}
}

func (p GetStateMachineMasterRequest) Pointer() *GetStateMachineMasterRequest {
	return &p
}

type DeleteStateMachineMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Version       *int64  `json:"version"`
}

func NewDeleteStateMachineMasterRequestFromJson(data string) DeleteStateMachineMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteStateMachineMasterRequestFromDict(dict2)
}

func NewDeleteStateMachineMasterRequestFromDict(data map[string]interface{}) DeleteStateMachineMasterRequest {
	return DeleteStateMachineMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Version:       core.CastInt64(data["version"]),
	}
}

func (p DeleteStateMachineMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"version":       p.Version,
	}
}

func (p DeleteStateMachineMasterRequest) Pointer() *DeleteStateMachineMasterRequest {
	return &p
}

type DescribeStatusesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	Status        *string `json:"status"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
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
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Status:        core.CastString(data["status"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"status":        p.Status,
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
	UserId        *string `json:"userId"`
	Status        *string `json:"status"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
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
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Status:        core.CastString(data["status"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"status":        p.Status,
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
	StatusName    *string `json:"statusName"`
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
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		StatusName:    core.CastString(data["statusName"]),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"statusName":    p.StatusName,
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
	StatusName    *string `json:"statusName"`
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
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		StatusName:    core.CastString(data["statusName"]),
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"statusName":    p.StatusName,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type StartStateMachineByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Args               *string `json:"args"`
	Ttl                *int32  `json:"ttl"`
}

func NewStartStateMachineByUserIdRequestFromJson(data string) StartStateMachineByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewStartStateMachineByUserIdRequestFromDict(dict2)
}

func NewStartStateMachineByUserIdRequestFromDict(data map[string]interface{}) StartStateMachineByUserIdRequest {
	return StartStateMachineByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Args:          core.CastString(data["args"]),
		Ttl:           core.CastInt32(data["ttl"]),
	}
}

func (p StartStateMachineByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"args":          p.Args,
		"ttl":           p.Ttl,
	}
}

func (p StartStateMachineByUserIdRequest) Pointer() *StartStateMachineByUserIdRequest {
	return &p
}

type StartStateMachineByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewStartStateMachineByStampSheetRequestFromJson(data string) StartStateMachineByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewStartStateMachineByStampSheetRequestFromDict(dict2)
}

func NewStartStateMachineByStampSheetRequestFromDict(data map[string]interface{}) StartStateMachineByStampSheetRequest {
	return StartStateMachineByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p StartStateMachineByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p StartStateMachineByStampSheetRequest) Pointer() *StartStateMachineByStampSheetRequest {
	return &p
}

type EmitRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	StatusName         *string `json:"statusName"`
	EventName          *string `json:"eventName"`
	Args               *string `json:"args"`
}

func NewEmitRequestFromJson(data string) EmitRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewEmitRequestFromDict(dict2)
}

func NewEmitRequestFromDict(data map[string]interface{}) EmitRequest {
	return EmitRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		StatusName:    core.CastString(data["statusName"]),
		EventName:     core.CastString(data["eventName"]),
		Args:          core.CastString(data["args"]),
	}
}

func (p EmitRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"statusName":    p.StatusName,
		"eventName":     p.EventName,
		"args":          p.Args,
	}
}

func (p EmitRequest) Pointer() *EmitRequest {
	return &p
}

type EmitByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	StatusName         *string `json:"statusName"`
	EventName          *string `json:"eventName"`
	Args               *string `json:"args"`
}

func NewEmitByUserIdRequestFromJson(data string) EmitByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewEmitByUserIdRequestFromDict(dict2)
}

func NewEmitByUserIdRequestFromDict(data map[string]interface{}) EmitByUserIdRequest {
	return EmitByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		StatusName:    core.CastString(data["statusName"]),
		EventName:     core.CastString(data["eventName"]),
		Args:          core.CastString(data["args"]),
	}
}

func (p EmitByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"statusName":    p.StatusName,
		"eventName":     p.EventName,
		"args":          p.Args,
	}
}

func (p EmitByUserIdRequest) Pointer() *EmitByUserIdRequest {
	return &p
}

type ReportRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	StatusName         *string `json:"statusName"`
	Events             []Event `json:"events"`
}

func NewReportRequestFromJson(data string) ReportRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewReportRequestFromDict(dict2)
}

func NewReportRequestFromDict(data map[string]interface{}) ReportRequest {
	return ReportRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		StatusName:    core.CastString(data["statusName"]),
		Events:        CastEvents(core.CastArray(data["events"])),
	}
}

func (p ReportRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"statusName":    p.StatusName,
		"events": CastEventsFromDict(
			p.Events,
		),
	}
}

func (p ReportRequest) Pointer() *ReportRequest {
	return &p
}

type ReportByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	StatusName         *string `json:"statusName"`
	Events             []Event `json:"events"`
}

func NewReportByUserIdRequestFromJson(data string) ReportByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewReportByUserIdRequestFromDict(dict2)
}

func NewReportByUserIdRequestFromDict(data map[string]interface{}) ReportByUserIdRequest {
	return ReportByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		StatusName:    core.CastString(data["statusName"]),
		Events:        CastEvents(core.CastArray(data["events"])),
	}
}

func (p ReportByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"statusName":    p.StatusName,
		"events": CastEventsFromDict(
			p.Events,
		),
	}
}

func (p ReportByUserIdRequest) Pointer() *ReportByUserIdRequest {
	return &p
}

type DeleteStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	StatusName         *string `json:"statusName"`
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
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		StatusName:    core.CastString(data["statusName"]),
	}
}

func (p DeleteStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"statusName":    p.StatusName,
	}
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
	return &p
}

type ExitStateMachineRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	StatusName         *string `json:"statusName"`
}

func NewExitStateMachineRequestFromJson(data string) ExitStateMachineRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewExitStateMachineRequestFromDict(dict2)
}

func NewExitStateMachineRequestFromDict(data map[string]interface{}) ExitStateMachineRequest {
	return ExitStateMachineRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		StatusName:    core.CastString(data["statusName"]),
	}
}

func (p ExitStateMachineRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"statusName":    p.StatusName,
	}
}

func (p ExitStateMachineRequest) Pointer() *ExitStateMachineRequest {
	return &p
}

type ExitStateMachineByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	StatusName         *string `json:"statusName"`
}

func NewExitStateMachineByUserIdRequestFromJson(data string) ExitStateMachineByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewExitStateMachineByUserIdRequestFromDict(dict2)
}

func NewExitStateMachineByUserIdRequestFromDict(data map[string]interface{}) ExitStateMachineByUserIdRequest {
	return ExitStateMachineByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		StatusName:    core.CastString(data["statusName"]),
	}
}

func (p ExitStateMachineByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"statusName":    p.StatusName,
	}
}

func (p ExitStateMachineByUserIdRequest) Pointer() *ExitStateMachineByUserIdRequest {
	return &p
}
