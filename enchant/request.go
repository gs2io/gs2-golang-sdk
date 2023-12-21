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

package enchant

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
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
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
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
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

type DescribeBalanceParameterModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeBalanceParameterModelsRequestFromJson(data string) DescribeBalanceParameterModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterModelsRequestFromDict(dict)
}

func NewDescribeBalanceParameterModelsRequestFromDict(data map[string]interface{}) DescribeBalanceParameterModelsRequest {
	return DescribeBalanceParameterModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeBalanceParameterModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeBalanceParameterModelsRequest) Pointer() *DescribeBalanceParameterModelsRequest {
	return &p
}

type GetBalanceParameterModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewGetBalanceParameterModelRequestFromJson(data string) GetBalanceParameterModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterModelRequestFromDict(dict)
}

func NewGetBalanceParameterModelRequestFromDict(data map[string]interface{}) GetBalanceParameterModelRequest {
	return GetBalanceParameterModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p GetBalanceParameterModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p GetBalanceParameterModelRequest) Pointer() *GetBalanceParameterModelRequest {
	return &p
}

type DescribeBalanceParameterModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBalanceParameterModelMastersRequestFromJson(data string) DescribeBalanceParameterModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterModelMastersRequestFromDict(dict)
}

func NewDescribeBalanceParameterModelMastersRequestFromDict(data map[string]interface{}) DescribeBalanceParameterModelMastersRequest {
	return DescribeBalanceParameterModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBalanceParameterModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBalanceParameterModelMastersRequest) Pointer() *DescribeBalanceParameterModelMastersRequest {
	return &p
}

type CreateBalanceParameterModelMasterRequest struct {
	RequestId            *string                      `json:"requestId"`
	ContextStack         *string                      `json:"contextStack"`
	NamespaceName        *string                      `json:"namespaceName"`
	Name                 *string                      `json:"name"`
	Description          *string                      `json:"description"`
	Metadata             *string                      `json:"metadata"`
	TotalValue           *int64                       `json:"totalValue"`
	InitialValueStrategy *string                      `json:"initialValueStrategy"`
	Parameters           []BalanceParameterValueModel `json:"parameters"`
}

func NewCreateBalanceParameterModelMasterRequestFromJson(data string) CreateBalanceParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBalanceParameterModelMasterRequestFromDict(dict)
}

func NewCreateBalanceParameterModelMasterRequestFromDict(data map[string]interface{}) CreateBalanceParameterModelMasterRequest {
	return CreateBalanceParameterModelMasterRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		Metadata:             core.CastString(data["metadata"]),
		TotalValue:           core.CastInt64(data["totalValue"]),
		InitialValueStrategy: core.CastString(data["initialValueStrategy"]),
		Parameters:           CastBalanceParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p CreateBalanceParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"name":                 p.Name,
		"description":          p.Description,
		"metadata":             p.Metadata,
		"totalValue":           p.TotalValue,
		"initialValueStrategy": p.InitialValueStrategy,
		"parameters": CastBalanceParameterValueModelsFromDict(
			p.Parameters,
		),
	}
}

func (p CreateBalanceParameterModelMasterRequest) Pointer() *CreateBalanceParameterModelMasterRequest {
	return &p
}

type GetBalanceParameterModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewGetBalanceParameterModelMasterRequestFromJson(data string) GetBalanceParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterModelMasterRequestFromDict(dict)
}

func NewGetBalanceParameterModelMasterRequestFromDict(data map[string]interface{}) GetBalanceParameterModelMasterRequest {
	return GetBalanceParameterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p GetBalanceParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p GetBalanceParameterModelMasterRequest) Pointer() *GetBalanceParameterModelMasterRequest {
	return &p
}

type UpdateBalanceParameterModelMasterRequest struct {
	RequestId            *string                      `json:"requestId"`
	ContextStack         *string                      `json:"contextStack"`
	NamespaceName        *string                      `json:"namespaceName"`
	ParameterName        *string                      `json:"parameterName"`
	Description          *string                      `json:"description"`
	Metadata             *string                      `json:"metadata"`
	TotalValue           *int64                       `json:"totalValue"`
	InitialValueStrategy *string                      `json:"initialValueStrategy"`
	Parameters           []BalanceParameterValueModel `json:"parameters"`
}

func NewUpdateBalanceParameterModelMasterRequestFromJson(data string) UpdateBalanceParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBalanceParameterModelMasterRequestFromDict(dict)
}

func NewUpdateBalanceParameterModelMasterRequestFromDict(data map[string]interface{}) UpdateBalanceParameterModelMasterRequest {
	return UpdateBalanceParameterModelMasterRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		ParameterName:        core.CastString(data["parameterName"]),
		Description:          core.CastString(data["description"]),
		Metadata:             core.CastString(data["metadata"]),
		TotalValue:           core.CastInt64(data["totalValue"]),
		InitialValueStrategy: core.CastString(data["initialValueStrategy"]),
		Parameters:           CastBalanceParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p UpdateBalanceParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"parameterName":        p.ParameterName,
		"description":          p.Description,
		"metadata":             p.Metadata,
		"totalValue":           p.TotalValue,
		"initialValueStrategy": p.InitialValueStrategy,
		"parameters": CastBalanceParameterValueModelsFromDict(
			p.Parameters,
		),
	}
}

func (p UpdateBalanceParameterModelMasterRequest) Pointer() *UpdateBalanceParameterModelMasterRequest {
	return &p
}

type DeleteBalanceParameterModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewDeleteBalanceParameterModelMasterRequestFromJson(data string) DeleteBalanceParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBalanceParameterModelMasterRequestFromDict(dict)
}

func NewDeleteBalanceParameterModelMasterRequestFromDict(data map[string]interface{}) DeleteBalanceParameterModelMasterRequest {
	return DeleteBalanceParameterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p DeleteBalanceParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p DeleteBalanceParameterModelMasterRequest) Pointer() *DeleteBalanceParameterModelMasterRequest {
	return &p
}

type DescribeRarityParameterModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeRarityParameterModelsRequestFromJson(data string) DescribeRarityParameterModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterModelsRequestFromDict(dict)
}

func NewDescribeRarityParameterModelsRequestFromDict(data map[string]interface{}) DescribeRarityParameterModelsRequest {
	return DescribeRarityParameterModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeRarityParameterModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRarityParameterModelsRequest) Pointer() *DescribeRarityParameterModelsRequest {
	return &p
}

type GetRarityParameterModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewGetRarityParameterModelRequestFromJson(data string) GetRarityParameterModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterModelRequestFromDict(dict)
}

func NewGetRarityParameterModelRequestFromDict(data map[string]interface{}) GetRarityParameterModelRequest {
	return GetRarityParameterModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p GetRarityParameterModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p GetRarityParameterModelRequest) Pointer() *GetRarityParameterModelRequest {
	return &p
}

type DescribeRarityParameterModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeRarityParameterModelMastersRequestFromJson(data string) DescribeRarityParameterModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterModelMastersRequestFromDict(dict)
}

func NewDescribeRarityParameterModelMastersRequestFromDict(data map[string]interface{}) DescribeRarityParameterModelMastersRequest {
	return DescribeRarityParameterModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRarityParameterModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRarityParameterModelMastersRequest) Pointer() *DescribeRarityParameterModelMastersRequest {
	return &p
}

type CreateRarityParameterModelMasterRequest struct {
	RequestId             *string                     `json:"requestId"`
	ContextStack          *string                     `json:"contextStack"`
	NamespaceName         *string                     `json:"namespaceName"`
	Name                  *string                     `json:"name"`
	Description           *string                     `json:"description"`
	Metadata              *string                     `json:"metadata"`
	MaximumParameterCount *int32                      `json:"maximumParameterCount"`
	ParameterCounts       []RarityParameterCountModel `json:"parameterCounts"`
	Parameters            []RarityParameterValueModel `json:"parameters"`
}

func NewCreateRarityParameterModelMasterRequestFromJson(data string) CreateRarityParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRarityParameterModelMasterRequestFromDict(dict)
}

func NewCreateRarityParameterModelMasterRequestFromDict(data map[string]interface{}) CreateRarityParameterModelMasterRequest {
	return CreateRarityParameterModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumParameterCount: core.CastInt32(data["maximumParameterCount"]),
		ParameterCounts:       CastRarityParameterCountModels(core.CastArray(data["parameterCounts"])),
		Parameters:            CastRarityParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p CreateRarityParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"name":                  p.Name,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"maximumParameterCount": p.MaximumParameterCount,
		"parameterCounts": CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		),
		"parameters": CastRarityParameterValueModelsFromDict(
			p.Parameters,
		),
	}
}

func (p CreateRarityParameterModelMasterRequest) Pointer() *CreateRarityParameterModelMasterRequest {
	return &p
}

type GetRarityParameterModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewGetRarityParameterModelMasterRequestFromJson(data string) GetRarityParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterModelMasterRequestFromDict(dict)
}

func NewGetRarityParameterModelMasterRequestFromDict(data map[string]interface{}) GetRarityParameterModelMasterRequest {
	return GetRarityParameterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p GetRarityParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p GetRarityParameterModelMasterRequest) Pointer() *GetRarityParameterModelMasterRequest {
	return &p
}

type UpdateRarityParameterModelMasterRequest struct {
	RequestId             *string                     `json:"requestId"`
	ContextStack          *string                     `json:"contextStack"`
	NamespaceName         *string                     `json:"namespaceName"`
	ParameterName         *string                     `json:"parameterName"`
	Description           *string                     `json:"description"`
	Metadata              *string                     `json:"metadata"`
	MaximumParameterCount *int32                      `json:"maximumParameterCount"`
	ParameterCounts       []RarityParameterCountModel `json:"parameterCounts"`
	Parameters            []RarityParameterValueModel `json:"parameters"`
}

func NewUpdateRarityParameterModelMasterRequestFromJson(data string) UpdateRarityParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRarityParameterModelMasterRequestFromDict(dict)
}

func NewUpdateRarityParameterModelMasterRequestFromDict(data map[string]interface{}) UpdateRarityParameterModelMasterRequest {
	return UpdateRarityParameterModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		ParameterName:         core.CastString(data["parameterName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumParameterCount: core.CastInt32(data["maximumParameterCount"]),
		ParameterCounts:       CastRarityParameterCountModels(core.CastArray(data["parameterCounts"])),
		Parameters:            CastRarityParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p UpdateRarityParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"parameterName":         p.ParameterName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"maximumParameterCount": p.MaximumParameterCount,
		"parameterCounts": CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		),
		"parameters": CastRarityParameterValueModelsFromDict(
			p.Parameters,
		),
	}
}

func (p UpdateRarityParameterModelMasterRequest) Pointer() *UpdateRarityParameterModelMasterRequest {
	return &p
}

type DeleteRarityParameterModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ParameterName *string `json:"parameterName"`
}

func NewDeleteRarityParameterModelMasterRequestFromJson(data string) DeleteRarityParameterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRarityParameterModelMasterRequestFromDict(dict)
}

func NewDeleteRarityParameterModelMasterRequestFromDict(data map[string]interface{}) DeleteRarityParameterModelMasterRequest {
	return DeleteRarityParameterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ParameterName: core.CastString(data["parameterName"]),
	}
}

func (p DeleteRarityParameterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"parameterName": p.ParameterName,
	}
}

func (p DeleteRarityParameterModelMasterRequest) Pointer() *DeleteRarityParameterModelMasterRequest {
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

type GetCurrentParameterMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentParameterMasterRequestFromJson(data string) GetCurrentParameterMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentParameterMasterRequestFromDict(dict)
}

func NewGetCurrentParameterMasterRequestFromDict(data map[string]interface{}) GetCurrentParameterMasterRequest {
	return GetCurrentParameterMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentParameterMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentParameterMasterRequest) Pointer() *GetCurrentParameterMasterRequest {
	return &p
}

type UpdateCurrentParameterMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentParameterMasterRequestFromJson(data string) UpdateCurrentParameterMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentParameterMasterRequestFromDict(dict)
}

func NewUpdateCurrentParameterMasterRequestFromDict(data map[string]interface{}) UpdateCurrentParameterMasterRequest {
	return UpdateCurrentParameterMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentParameterMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentParameterMasterRequest) Pointer() *UpdateCurrentParameterMasterRequest {
	return &p
}

type UpdateCurrentParameterMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentParameterMasterFromGitHubRequestFromJson(data string) UpdateCurrentParameterMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentParameterMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentParameterMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentParameterMasterFromGitHubRequest {
	return UpdateCurrentParameterMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentParameterMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentParameterMasterFromGitHubRequest) Pointer() *UpdateCurrentParameterMasterFromGitHubRequest {
	return &p
}

type DescribeBalanceParameterStatusesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	ParameterName *string `json:"parameterName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBalanceParameterStatusesRequestFromJson(data string) DescribeBalanceParameterStatusesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterStatusesRequestFromDict(dict)
}

func NewDescribeBalanceParameterStatusesRequestFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesRequest {
	return DescribeBalanceParameterStatusesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ParameterName: core.CastString(data["parameterName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBalanceParameterStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"parameterName": p.ParameterName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBalanceParameterStatusesRequest) Pointer() *DescribeBalanceParameterStatusesRequest {
	return &p
}

type DescribeBalanceParameterStatusesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	ParameterName *string `json:"parameterName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBalanceParameterStatusesByUserIdRequestFromJson(data string) DescribeBalanceParameterStatusesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeBalanceParameterStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesByUserIdRequest {
	return DescribeBalanceParameterStatusesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBalanceParameterStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBalanceParameterStatusesByUserIdRequest) Pointer() *DescribeBalanceParameterStatusesByUserIdRequest {
	return &p
}

type GetBalanceParameterStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	ParameterName *string `json:"parameterName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetBalanceParameterStatusRequestFromJson(data string) GetBalanceParameterStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterStatusRequestFromDict(dict)
}

func NewGetBalanceParameterStatusRequestFromDict(data map[string]interface{}) GetBalanceParameterStatusRequest {
	return GetBalanceParameterStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetBalanceParameterStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetBalanceParameterStatusRequest) Pointer() *GetBalanceParameterStatusRequest {
	return &p
}

type GetBalanceParameterStatusByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	ParameterName *string `json:"parameterName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetBalanceParameterStatusByUserIdRequestFromJson(data string) GetBalanceParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterStatusByUserIdRequestFromDict(dict)
}

func NewGetBalanceParameterStatusByUserIdRequestFromDict(data map[string]interface{}) GetBalanceParameterStatusByUserIdRequest {
	return GetBalanceParameterStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetBalanceParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetBalanceParameterStatusByUserIdRequest) Pointer() *GetBalanceParameterStatusByUserIdRequest {
	return &p
}

type DeleteBalanceParameterStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ParameterName      *string `json:"parameterName"`
	PropertyId         *string `json:"propertyId"`
}

func NewDeleteBalanceParameterStatusByUserIdRequestFromJson(data string) DeleteBalanceParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBalanceParameterStatusByUserIdRequestFromDict(dict)
}

func NewDeleteBalanceParameterStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteBalanceParameterStatusByUserIdRequest {
	return DeleteBalanceParameterStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p DeleteBalanceParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p DeleteBalanceParameterStatusByUserIdRequest) Pointer() *DeleteBalanceParameterStatusByUserIdRequest {
	return &p
}

type ReDrawBalanceParameterStatusByUserIdRequest struct {
	RequestId           *string   `json:"requestId"`
	ContextStack        *string   `json:"contextStack"`
	DuplicationAvoider  *string   `json:"duplicationAvoider"`
	NamespaceName       *string   `json:"namespaceName"`
	UserId              *string   `json:"userId"`
	ParameterName       *string   `json:"parameterName"`
	PropertyId          *string   `json:"propertyId"`
	FixedParameterNames []*string `json:"fixedParameterNames"`
}

func NewReDrawBalanceParameterStatusByUserIdRequestFromJson(data string) ReDrawBalanceParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawBalanceParameterStatusByUserIdRequestFromDict(dict)
}

func NewReDrawBalanceParameterStatusByUserIdRequestFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByUserIdRequest {
	return ReDrawBalanceParameterStatusByUserIdRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		UserId:              core.CastString(data["userId"]),
		ParameterName:       core.CastString(data["parameterName"]),
		PropertyId:          core.CastString(data["propertyId"]),
		FixedParameterNames: core.CastStrings(core.CastArray(data["fixedParameterNames"])),
	}
}

func (p ReDrawBalanceParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
		"fixedParameterNames": core.CastStringsFromDict(
			p.FixedParameterNames,
		),
	}
}

func (p ReDrawBalanceParameterStatusByUserIdRequest) Pointer() *ReDrawBalanceParameterStatusByUserIdRequest {
	return &p
}

type ReDrawBalanceParameterStatusByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewReDrawBalanceParameterStatusByStampSheetRequestFromJson(data string) ReDrawBalanceParameterStatusByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawBalanceParameterStatusByStampSheetRequestFromDict(dict)
}

func NewReDrawBalanceParameterStatusByStampSheetRequestFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByStampSheetRequest {
	return ReDrawBalanceParameterStatusByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p ReDrawBalanceParameterStatusByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ReDrawBalanceParameterStatusByStampSheetRequest) Pointer() *ReDrawBalanceParameterStatusByStampSheetRequest {
	return &p
}

type SetBalanceParameterStatusByUserIdRequest struct {
	RequestId          *string                 `json:"requestId"`
	ContextStack       *string                 `json:"contextStack"`
	DuplicationAvoider *string                 `json:"duplicationAvoider"`
	NamespaceName      *string                 `json:"namespaceName"`
	UserId             *string                 `json:"userId"`
	ParameterName      *string                 `json:"parameterName"`
	PropertyId         *string                 `json:"propertyId"`
	ParameterValues    []BalanceParameterValue `json:"parameterValues"`
}

func NewSetBalanceParameterStatusByUserIdRequestFromJson(data string) SetBalanceParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBalanceParameterStatusByUserIdRequestFromDict(dict)
}

func NewSetBalanceParameterStatusByUserIdRequestFromDict(data map[string]interface{}) SetBalanceParameterStatusByUserIdRequest {
	return SetBalanceParameterStatusByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ParameterName:   core.CastString(data["parameterName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ParameterValues: CastBalanceParameterValues(core.CastArray(data["parameterValues"])),
	}
}

func (p SetBalanceParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
		"parameterValues": CastBalanceParameterValuesFromDict(
			p.ParameterValues,
		),
	}
}

func (p SetBalanceParameterStatusByUserIdRequest) Pointer() *SetBalanceParameterStatusByUserIdRequest {
	return &p
}

type SetBalanceParameterStatusByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetBalanceParameterStatusByStampSheetRequestFromJson(data string) SetBalanceParameterStatusByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBalanceParameterStatusByStampSheetRequestFromDict(dict)
}

func NewSetBalanceParameterStatusByStampSheetRequestFromDict(data map[string]interface{}) SetBalanceParameterStatusByStampSheetRequest {
	return SetBalanceParameterStatusByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetBalanceParameterStatusByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetBalanceParameterStatusByStampSheetRequest) Pointer() *SetBalanceParameterStatusByStampSheetRequest {
	return &p
}

type DescribeRarityParameterStatusesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	ParameterName *string `json:"parameterName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeRarityParameterStatusesRequestFromJson(data string) DescribeRarityParameterStatusesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterStatusesRequestFromDict(dict)
}

func NewDescribeRarityParameterStatusesRequestFromDict(data map[string]interface{}) DescribeRarityParameterStatusesRequest {
	return DescribeRarityParameterStatusesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ParameterName: core.CastString(data["parameterName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRarityParameterStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"parameterName": p.ParameterName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRarityParameterStatusesRequest) Pointer() *DescribeRarityParameterStatusesRequest {
	return &p
}

type DescribeRarityParameterStatusesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	ParameterName *string `json:"parameterName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeRarityParameterStatusesByUserIdRequestFromJson(data string) DescribeRarityParameterStatusesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeRarityParameterStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeRarityParameterStatusesByUserIdRequest {
	return DescribeRarityParameterStatusesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRarityParameterStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRarityParameterStatusesByUserIdRequest) Pointer() *DescribeRarityParameterStatusesByUserIdRequest {
	return &p
}

type GetRarityParameterStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	ParameterName *string `json:"parameterName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetRarityParameterStatusRequestFromJson(data string) GetRarityParameterStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterStatusRequestFromDict(dict)
}

func NewGetRarityParameterStatusRequestFromDict(data map[string]interface{}) GetRarityParameterStatusRequest {
	return GetRarityParameterStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetRarityParameterStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetRarityParameterStatusRequest) Pointer() *GetRarityParameterStatusRequest {
	return &p
}

type GetRarityParameterStatusByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	ParameterName *string `json:"parameterName"`
	PropertyId    *string `json:"propertyId"`
}

func NewGetRarityParameterStatusByUserIdRequestFromJson(data string) GetRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewGetRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) GetRarityParameterStatusByUserIdRequest {
	return GetRarityParameterStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetRarityParameterStatusByUserIdRequest) Pointer() *GetRarityParameterStatusByUserIdRequest {
	return &p
}

type DeleteRarityParameterStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ParameterName      *string `json:"parameterName"`
	PropertyId         *string `json:"propertyId"`
}

func NewDeleteRarityParameterStatusByUserIdRequestFromJson(data string) DeleteRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewDeleteRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteRarityParameterStatusByUserIdRequest {
	return DeleteRarityParameterStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p DeleteRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
	}
}

func (p DeleteRarityParameterStatusByUserIdRequest) Pointer() *DeleteRarityParameterStatusByUserIdRequest {
	return &p
}

type ReDrawRarityParameterStatusByUserIdRequest struct {
	RequestId           *string   `json:"requestId"`
	ContextStack        *string   `json:"contextStack"`
	DuplicationAvoider  *string   `json:"duplicationAvoider"`
	NamespaceName       *string   `json:"namespaceName"`
	UserId              *string   `json:"userId"`
	ParameterName       *string   `json:"parameterName"`
	PropertyId          *string   `json:"propertyId"`
	FixedParameterNames []*string `json:"fixedParameterNames"`
}

func NewReDrawRarityParameterStatusByUserIdRequestFromJson(data string) ReDrawRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewReDrawRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByUserIdRequest {
	return ReDrawRarityParameterStatusByUserIdRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		UserId:              core.CastString(data["userId"]),
		ParameterName:       core.CastString(data["parameterName"]),
		PropertyId:          core.CastString(data["propertyId"]),
		FixedParameterNames: core.CastStrings(core.CastArray(data["fixedParameterNames"])),
	}
}

func (p ReDrawRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
		"fixedParameterNames": core.CastStringsFromDict(
			p.FixedParameterNames,
		),
	}
}

func (p ReDrawRarityParameterStatusByUserIdRequest) Pointer() *ReDrawRarityParameterStatusByUserIdRequest {
	return &p
}

type ReDrawRarityParameterStatusByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewReDrawRarityParameterStatusByStampSheetRequestFromJson(data string) ReDrawRarityParameterStatusByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawRarityParameterStatusByStampSheetRequestFromDict(dict)
}

func NewReDrawRarityParameterStatusByStampSheetRequestFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByStampSheetRequest {
	return ReDrawRarityParameterStatusByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p ReDrawRarityParameterStatusByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ReDrawRarityParameterStatusByStampSheetRequest) Pointer() *ReDrawRarityParameterStatusByStampSheetRequest {
	return &p
}

type AddRarityParameterStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ParameterName      *string `json:"parameterName"`
	PropertyId         *string `json:"propertyId"`
	Count              *int32  `json:"count"`
}

func NewAddRarityParameterStatusByUserIdRequestFromJson(data string) AddRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewAddRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) AddRarityParameterStatusByUserIdRequest {
	return AddRarityParameterStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		ParameterName: core.CastString(data["parameterName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		Count:         core.CastInt32(data["count"]),
	}
}

func (p AddRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
		"count":         p.Count,
	}
}

func (p AddRarityParameterStatusByUserIdRequest) Pointer() *AddRarityParameterStatusByUserIdRequest {
	return &p
}

type AddRarityParameterStatusByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddRarityParameterStatusByStampSheetRequestFromJson(data string) AddRarityParameterStatusByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddRarityParameterStatusByStampSheetRequestFromDict(dict)
}

func NewAddRarityParameterStatusByStampSheetRequestFromDict(data map[string]interface{}) AddRarityParameterStatusByStampSheetRequest {
	return AddRarityParameterStatusByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddRarityParameterStatusByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddRarityParameterStatusByStampSheetRequest) Pointer() *AddRarityParameterStatusByStampSheetRequest {
	return &p
}

type VerifyRarityParameterStatusRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	ParameterName      *string `json:"parameterName"`
	AccessToken        *string `json:"accessToken"`
	PropertyId         *string `json:"propertyId"`
	VerifyType         *string `json:"verifyType"`
	ParameterValueName *string `json:"parameterValueName"`
	ParameterCount     *int32  `json:"parameterCount"`
}

func NewVerifyRarityParameterStatusRequestFromJson(data string) VerifyRarityParameterStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusRequestFromDict(dict)
}

func NewVerifyRarityParameterStatusRequestFromDict(data map[string]interface{}) VerifyRarityParameterStatusRequest {
	return VerifyRarityParameterStatusRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		ParameterName:      core.CastString(data["parameterName"]),
		AccessToken:        core.CastString(data["accessToken"]),
		PropertyId:         core.CastString(data["propertyId"]),
		VerifyType:         core.CastString(data["verifyType"]),
		ParameterValueName: core.CastString(data["parameterValueName"]),
		ParameterCount:     core.CastInt32(data["parameterCount"]),
	}
}

func (p VerifyRarityParameterStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"parameterName":      p.ParameterName,
		"accessToken":        p.AccessToken,
		"propertyId":         p.PropertyId,
		"verifyType":         p.VerifyType,
		"parameterValueName": p.ParameterValueName,
		"parameterCount":     p.ParameterCount,
	}
}

func (p VerifyRarityParameterStatusRequest) Pointer() *VerifyRarityParameterStatusRequest {
	return &p
}

type VerifyRarityParameterStatusByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	ParameterName      *string `json:"parameterName"`
	UserId             *string `json:"userId"`
	PropertyId         *string `json:"propertyId"`
	VerifyType         *string `json:"verifyType"`
	ParameterValueName *string `json:"parameterValueName"`
	ParameterCount     *int32  `json:"parameterCount"`
}

func NewVerifyRarityParameterStatusByUserIdRequestFromJson(data string) VerifyRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewVerifyRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) VerifyRarityParameterStatusByUserIdRequest {
	return VerifyRarityParameterStatusByUserIdRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		ParameterName:      core.CastString(data["parameterName"]),
		UserId:             core.CastString(data["userId"]),
		PropertyId:         core.CastString(data["propertyId"]),
		VerifyType:         core.CastString(data["verifyType"]),
		ParameterValueName: core.CastString(data["parameterValueName"]),
		ParameterCount:     core.CastInt32(data["parameterCount"]),
	}
}

func (p VerifyRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"parameterName":      p.ParameterName,
		"userId":             p.UserId,
		"propertyId":         p.PropertyId,
		"verifyType":         p.VerifyType,
		"parameterValueName": p.ParameterValueName,
		"parameterCount":     p.ParameterCount,
	}
}

func (p VerifyRarityParameterStatusByUserIdRequest) Pointer() *VerifyRarityParameterStatusByUserIdRequest {
	return &p
}

type VerifyRarityParameterStatusByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyRarityParameterStatusByStampTaskRequestFromJson(data string) VerifyRarityParameterStatusByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusByStampTaskRequestFromDict(dict)
}

func NewVerifyRarityParameterStatusByStampTaskRequestFromDict(data map[string]interface{}) VerifyRarityParameterStatusByStampTaskRequest {
	return VerifyRarityParameterStatusByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyRarityParameterStatusByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyRarityParameterStatusByStampTaskRequest) Pointer() *VerifyRarityParameterStatusByStampTaskRequest {
	return &p
}

type SetRarityParameterStatusByUserIdRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	UserId             *string                `json:"userId"`
	ParameterName      *string                `json:"parameterName"`
	PropertyId         *string                `json:"propertyId"`
	ParameterValues    []RarityParameterValue `json:"parameterValues"`
}

func NewSetRarityParameterStatusByUserIdRequestFromJson(data string) SetRarityParameterStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRarityParameterStatusByUserIdRequestFromDict(dict)
}

func NewSetRarityParameterStatusByUserIdRequestFromDict(data map[string]interface{}) SetRarityParameterStatusByUserIdRequest {
	return SetRarityParameterStatusByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		ParameterName:   core.CastString(data["parameterName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ParameterValues: CastRarityParameterValues(core.CastArray(data["parameterValues"])),
	}
}

func (p SetRarityParameterStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"parameterName": p.ParameterName,
		"propertyId":    p.PropertyId,
		"parameterValues": CastRarityParameterValuesFromDict(
			p.ParameterValues,
		),
	}
}

func (p SetRarityParameterStatusByUserIdRequest) Pointer() *SetRarityParameterStatusByUserIdRequest {
	return &p
}

type SetRarityParameterStatusByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetRarityParameterStatusByStampSheetRequestFromJson(data string) SetRarityParameterStatusByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRarityParameterStatusByStampSheetRequestFromDict(dict)
}

func NewSetRarityParameterStatusByStampSheetRequestFromDict(data map[string]interface{}) SetRarityParameterStatusByStampSheetRequest {
	return SetRarityParameterStatusByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRarityParameterStatusByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRarityParameterStatusByStampSheetRequest) Pointer() *SetRarityParameterStatusByStampSheetRequest {
	return &p
}
