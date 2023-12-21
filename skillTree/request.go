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

package skillTree

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
	ReleaseScript      *ScriptSetting      `json:"releaseScript"`
	RestrainScript     *ScriptSetting      `json:"restrainScript"`
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
		ReleaseScript:      NewScriptSettingFromDict(core.CastMap(data["releaseScript"])).Pointer(),
		RestrainScript:     NewScriptSettingFromDict(core.CastMap(data["restrainScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"releaseScript":      p.ReleaseScript.ToDict(),
		"restrainScript":     p.RestrainScript.ToDict(),
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
	ReleaseScript      *ScriptSetting      `json:"releaseScript"`
	RestrainScript     *ScriptSetting      `json:"restrainScript"`
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
		ReleaseScript:      NewScriptSettingFromDict(core.CastMap(data["releaseScript"])).Pointer(),
		RestrainScript:     NewScriptSettingFromDict(core.CastMap(data["restrainScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"releaseScript":      p.ReleaseScript.ToDict(),
		"restrainScript":     p.RestrainScript.ToDict(),
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

type DescribeNodeModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeNodeModelsRequestFromJson(data string) DescribeNodeModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNodeModelsRequestFromDict(dict2)
}

func NewDescribeNodeModelsRequestFromDict(data map[string]interface{}) DescribeNodeModelsRequest {
	return DescribeNodeModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeNodeModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeNodeModelsRequest) Pointer() *DescribeNodeModelsRequest {
	return &p
}

type GetNodeModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewGetNodeModelRequestFromJson(data string) GetNodeModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNodeModelRequestFromDict(dict2)
}

func NewGetNodeModelRequestFromDict(data map[string]interface{}) GetNodeModelRequest {
	return GetNodeModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p GetNodeModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p GetNodeModelRequest) Pointer() *GetNodeModelRequest {
	return &p
}

type DescribeNodeModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeNodeModelMastersRequestFromJson(data string) DescribeNodeModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNodeModelMastersRequestFromDict(dict2)
}

func NewDescribeNodeModelMastersRequestFromDict(data map[string]interface{}) DescribeNodeModelMastersRequest {
	return DescribeNodeModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeNodeModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeNodeModelMastersRequest) Pointer() *DescribeNodeModelMastersRequest {
	return &p
}

type CreateNodeModelMasterRequest struct {
	RequestId             *string         `json:"requestId"`
	ContextStack          *string         `json:"contextStack"`
	NamespaceName         *string         `json:"namespaceName"`
	Name                  *string         `json:"name"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
}

func NewCreateNodeModelMasterRequestFromJson(data string) CreateNodeModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateNodeModelMasterRequestFromDict(dict2)
}

func NewCreateNodeModelMasterRequestFromDict(data map[string]interface{}) CreateNodeModelMasterRequest {
	return CreateNodeModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
	}
}

func (p CreateNodeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"releaseConsumeActions": CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		),
		"restrainReturnRate": p.RestrainReturnRate,
		"premiseNodeNames": core.CastStringsFromDict(
			p.PremiseNodeNames,
		),
	}
}

func (p CreateNodeModelMasterRequest) Pointer() *CreateNodeModelMasterRequest {
	return &p
}

type GetNodeModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewGetNodeModelMasterRequestFromJson(data string) GetNodeModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNodeModelMasterRequestFromDict(dict2)
}

func NewGetNodeModelMasterRequestFromDict(data map[string]interface{}) GetNodeModelMasterRequest {
	return GetNodeModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p GetNodeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p GetNodeModelMasterRequest) Pointer() *GetNodeModelMasterRequest {
	return &p
}

type UpdateNodeModelMasterRequest struct {
	RequestId             *string         `json:"requestId"`
	ContextStack          *string         `json:"contextStack"`
	NamespaceName         *string         `json:"namespaceName"`
	NodeModelName         *string         `json:"nodeModelName"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
}

func NewUpdateNodeModelMasterRequestFromJson(data string) UpdateNodeModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateNodeModelMasterRequestFromDict(dict2)
}

func NewUpdateNodeModelMasterRequestFromDict(data map[string]interface{}) UpdateNodeModelMasterRequest {
	return UpdateNodeModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		NodeModelName:         core.CastString(data["nodeModelName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
	}
}

func (p UpdateNodeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"releaseConsumeActions": CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		),
		"restrainReturnRate": p.RestrainReturnRate,
		"premiseNodeNames": core.CastStringsFromDict(
			p.PremiseNodeNames,
		),
	}
}

func (p UpdateNodeModelMasterRequest) Pointer() *UpdateNodeModelMasterRequest {
	return &p
}

type DeleteNodeModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewDeleteNodeModelMasterRequestFromJson(data string) DeleteNodeModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteNodeModelMasterRequestFromDict(dict2)
}

func NewDeleteNodeModelMasterRequestFromDict(data map[string]interface{}) DeleteNodeModelMasterRequest {
	return DeleteNodeModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p DeleteNodeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p DeleteNodeModelMasterRequest) Pointer() *DeleteNodeModelMasterRequest {
	return &p
}

type MarkReleaseByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
}

func NewMarkReleaseByUserIdRequestFromJson(data string) MarkReleaseByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMarkReleaseByUserIdRequestFromDict(dict2)
}

func NewMarkReleaseByUserIdRequestFromDict(data map[string]interface{}) MarkReleaseByUserIdRequest {
	return MarkReleaseByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
	}
}

func (p MarkReleaseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
	}
}

func (p MarkReleaseByUserIdRequest) Pointer() *MarkReleaseByUserIdRequest {
	return &p
}

type ReleaseRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewReleaseRequestFromJson(data string) ReleaseRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewReleaseRequestFromDict(dict2)
}

func NewReleaseRequestFromDict(data map[string]interface{}) ReleaseRequest {
	return ReleaseRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReleaseRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReleaseRequest) Pointer() *ReleaseRequest {
	return &p
}

type ReleaseByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewReleaseByUserIdRequestFromJson(data string) ReleaseByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewReleaseByUserIdRequestFromDict(dict2)
}

func NewReleaseByUserIdRequestFromDict(data map[string]interface{}) ReleaseByUserIdRequest {
	return ReleaseByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReleaseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReleaseByUserIdRequest) Pointer() *ReleaseByUserIdRequest {
	return &p
}

type MarkRestrainByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
}

func NewMarkRestrainByUserIdRequestFromJson(data string) MarkRestrainByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMarkRestrainByUserIdRequestFromDict(dict2)
}

func NewMarkRestrainByUserIdRequestFromDict(data map[string]interface{}) MarkRestrainByUserIdRequest {
	return MarkRestrainByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
	}
}

func (p MarkRestrainByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
	}
}

func (p MarkRestrainByUserIdRequest) Pointer() *MarkRestrainByUserIdRequest {
	return &p
}

type RestrainRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewRestrainRequestFromJson(data string) RestrainRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewRestrainRequestFromDict(dict2)
}

func NewRestrainRequestFromDict(data map[string]interface{}) RestrainRequest {
	return RestrainRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RestrainRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RestrainRequest) Pointer() *RestrainRequest {
	return &p
}

type RestrainByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewRestrainByUserIdRequestFromJson(data string) RestrainByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewRestrainByUserIdRequestFromDict(dict2)
}

func NewRestrainByUserIdRequestFromDict(data map[string]interface{}) RestrainByUserIdRequest {
	return RestrainByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RestrainByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RestrainByUserIdRequest) Pointer() *RestrainByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
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
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
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
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type ResetRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	Config             []Config `json:"config"`
}

func NewResetRequestFromJson(data string) ResetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewResetRequestFromDict(dict2)
}

func NewResetRequestFromDict(data map[string]interface{}) ResetRequest {
	return ResetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ResetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ResetRequest) Pointer() *ResetRequest {
	return &p
}

type ResetByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	Config             []Config `json:"config"`
}

func NewResetByUserIdRequestFromJson(data string) ResetByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewResetByUserIdRequestFromDict(dict2)
}

func NewResetByUserIdRequestFromDict(data map[string]interface{}) ResetByUserIdRequest {
	return ResetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ResetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ResetByUserIdRequest) Pointer() *ResetByUserIdRequest {
	return &p
}

type MarkReleaseByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewMarkReleaseByStampSheetRequestFromJson(data string) MarkReleaseByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMarkReleaseByStampSheetRequestFromDict(dict2)
}

func NewMarkReleaseByStampSheetRequestFromDict(data map[string]interface{}) MarkReleaseByStampSheetRequest {
	return MarkReleaseByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p MarkReleaseByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p MarkReleaseByStampSheetRequest) Pointer() *MarkReleaseByStampSheetRequest {
	return &p
}

type MarkRestrainByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewMarkRestrainByStampTaskRequestFromJson(data string) MarkRestrainByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewMarkRestrainByStampTaskRequestFromDict(dict2)
}

func NewMarkRestrainByStampTaskRequestFromDict(data map[string]interface{}) MarkRestrainByStampTaskRequest {
	return MarkRestrainByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p MarkRestrainByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p MarkRestrainByStampTaskRequest) Pointer() *MarkRestrainByStampTaskRequest {
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

type GetCurrentTreeMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentTreeMasterRequestFromJson(data string) GetCurrentTreeMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentTreeMasterRequestFromDict(dict2)
}

func NewGetCurrentTreeMasterRequestFromDict(data map[string]interface{}) GetCurrentTreeMasterRequest {
	return GetCurrentTreeMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentTreeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentTreeMasterRequest) Pointer() *GetCurrentTreeMasterRequest {
	return &p
}

type UpdateCurrentTreeMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentTreeMasterRequestFromJson(data string) UpdateCurrentTreeMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentTreeMasterRequestFromDict(dict2)
}

func NewUpdateCurrentTreeMasterRequestFromDict(data map[string]interface{}) UpdateCurrentTreeMasterRequest {
	return UpdateCurrentTreeMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentTreeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentTreeMasterRequest) Pointer() *UpdateCurrentTreeMasterRequest {
	return &p
}

type UpdateCurrentTreeMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentTreeMasterFromGitHubRequestFromJson(data string) UpdateCurrentTreeMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentTreeMasterFromGitHubRequestFromDict(dict2)
}

func NewUpdateCurrentTreeMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentTreeMasterFromGitHubRequest {
	return UpdateCurrentTreeMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentTreeMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentTreeMasterFromGitHubRequest) Pointer() *UpdateCurrentTreeMasterFromGitHubRequest {
	return &p
}
