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

package showcase

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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	BuyScript          *ScriptSetting      `json:"buyScript"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId      *string     `json:"keyId"`
	LogSetting *LogSetting `json:"logSetting"`
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
		BuyScript:          NewScriptSettingFromDict(core.CastMap(data["buyScript"])).Pointer(),
		QueueNamespaceId:   core.CastString(data["queueNamespaceId"]),
		KeyId:              core.CastString(data["keyId"]),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"buyScript":          p.BuyScript.ToDict(),
		"queueNamespaceId":   p.QueueNamespaceId,
		"keyId":              p.KeyId,
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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	BuyScript          *ScriptSetting      `json:"buyScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
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
		BuyScript:          NewScriptSettingFromDict(core.CastMap(data["buyScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:   core.CastString(data["queueNamespaceId"]),
		KeyId:              core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"buyScript":          p.BuyScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
		"queueNamespaceId":   p.QueueNamespaceId,
		"keyId":              p.KeyId,
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
	SourceRequestId    *string `json:"sourceRequestId"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
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

type DescribeSalesItemMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSalesItemMastersRequestFromJson(data string) DescribeSalesItemMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSalesItemMastersRequestFromDict(dict)
}

func NewDescribeSalesItemMastersRequestFromDict(data map[string]interface{}) DescribeSalesItemMastersRequest {
	return DescribeSalesItemMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSalesItemMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSalesItemMastersRequest) Pointer() *DescribeSalesItemMastersRequest {
	return &p
}

type CreateSalesItemMasterRequest struct {
	SourceRequestId *string         `json:"sourceRequestId"`
	RequestId       *string         `json:"requestId"`
	ContextStack    *string         `json:"contextStack"`
	NamespaceName   *string         `json:"namespaceName"`
	Name            *string         `json:"name"`
	Description     *string         `json:"description"`
	Metadata        *string         `json:"metadata"`
	ConsumeActions  []ConsumeAction `json:"consumeActions"`
	AcquireActions  []AcquireAction `json:"acquireActions"`
}

func NewCreateSalesItemMasterRequestFromJson(data string) CreateSalesItemMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSalesItemMasterRequestFromDict(dict)
}

func NewCreateSalesItemMasterRequestFromDict(data map[string]interface{}) CreateSalesItemMasterRequest {
	return CreateSalesItemMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p CreateSalesItemMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p CreateSalesItemMasterRequest) Pointer() *CreateSalesItemMasterRequest {
	return &p
}

type GetSalesItemMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SalesItemName   *string `json:"salesItemName"`
}

func NewGetSalesItemMasterRequestFromJson(data string) GetSalesItemMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSalesItemMasterRequestFromDict(dict)
}

func NewGetSalesItemMasterRequestFromDict(data map[string]interface{}) GetSalesItemMasterRequest {
	return GetSalesItemMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SalesItemName: core.CastString(data["salesItemName"]),
	}
}

func (p GetSalesItemMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"salesItemName": p.SalesItemName,
	}
}

func (p GetSalesItemMasterRequest) Pointer() *GetSalesItemMasterRequest {
	return &p
}

type UpdateSalesItemMasterRequest struct {
	SourceRequestId *string         `json:"sourceRequestId"`
	RequestId       *string         `json:"requestId"`
	ContextStack    *string         `json:"contextStack"`
	NamespaceName   *string         `json:"namespaceName"`
	SalesItemName   *string         `json:"salesItemName"`
	Description     *string         `json:"description"`
	Metadata        *string         `json:"metadata"`
	ConsumeActions  []ConsumeAction `json:"consumeActions"`
	AcquireActions  []AcquireAction `json:"acquireActions"`
}

func NewUpdateSalesItemMasterRequestFromJson(data string) UpdateSalesItemMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSalesItemMasterRequestFromDict(dict)
}

func NewUpdateSalesItemMasterRequestFromDict(data map[string]interface{}) UpdateSalesItemMasterRequest {
	return UpdateSalesItemMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		SalesItemName:  core.CastString(data["salesItemName"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p UpdateSalesItemMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"salesItemName": p.SalesItemName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p UpdateSalesItemMasterRequest) Pointer() *UpdateSalesItemMasterRequest {
	return &p
}

type DeleteSalesItemMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SalesItemName   *string `json:"salesItemName"`
}

func NewDeleteSalesItemMasterRequestFromJson(data string) DeleteSalesItemMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSalesItemMasterRequestFromDict(dict)
}

func NewDeleteSalesItemMasterRequestFromDict(data map[string]interface{}) DeleteSalesItemMasterRequest {
	return DeleteSalesItemMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SalesItemName: core.CastString(data["salesItemName"]),
	}
}

func (p DeleteSalesItemMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"salesItemName": p.SalesItemName,
	}
}

func (p DeleteSalesItemMasterRequest) Pointer() *DeleteSalesItemMasterRequest {
	return &p
}

type DescribeSalesItemGroupMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSalesItemGroupMastersRequestFromJson(data string) DescribeSalesItemGroupMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSalesItemGroupMastersRequestFromDict(dict)
}

func NewDescribeSalesItemGroupMastersRequestFromDict(data map[string]interface{}) DescribeSalesItemGroupMastersRequest {
	return DescribeSalesItemGroupMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSalesItemGroupMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSalesItemGroupMastersRequest) Pointer() *DescribeSalesItemGroupMastersRequest {
	return &p
}

type CreateSalesItemGroupMasterRequest struct {
	SourceRequestId *string   `json:"sourceRequestId"`
	RequestId       *string   `json:"requestId"`
	ContextStack    *string   `json:"contextStack"`
	NamespaceName   *string   `json:"namespaceName"`
	Name            *string   `json:"name"`
	Description     *string   `json:"description"`
	Metadata        *string   `json:"metadata"`
	SalesItemNames  []*string `json:"salesItemNames"`
}

func NewCreateSalesItemGroupMasterRequestFromJson(data string) CreateSalesItemGroupMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSalesItemGroupMasterRequestFromDict(dict)
}

func NewCreateSalesItemGroupMasterRequestFromDict(data map[string]interface{}) CreateSalesItemGroupMasterRequest {
	return CreateSalesItemGroupMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		SalesItemNames: core.CastStrings(core.CastArray(data["salesItemNames"])),
	}
}

func (p CreateSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"salesItemNames": core.CastStringsFromDict(
			p.SalesItemNames,
		),
	}
}

func (p CreateSalesItemGroupMasterRequest) Pointer() *CreateSalesItemGroupMasterRequest {
	return &p
}

type GetSalesItemGroupMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	SalesItemGroupName *string `json:"salesItemGroupName"`
}

func NewGetSalesItemGroupMasterRequestFromJson(data string) GetSalesItemGroupMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSalesItemGroupMasterRequestFromDict(dict)
}

func NewGetSalesItemGroupMasterRequestFromDict(data map[string]interface{}) GetSalesItemGroupMasterRequest {
	return GetSalesItemGroupMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
	}
}

func (p GetSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"salesItemGroupName": p.SalesItemGroupName,
	}
}

func (p GetSalesItemGroupMasterRequest) Pointer() *GetSalesItemGroupMasterRequest {
	return &p
}

type UpdateSalesItemGroupMasterRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	NamespaceName      *string   `json:"namespaceName"`
	SalesItemGroupName *string   `json:"salesItemGroupName"`
	Description        *string   `json:"description"`
	Metadata           *string   `json:"metadata"`
	SalesItemNames     []*string `json:"salesItemNames"`
}

func NewUpdateSalesItemGroupMasterRequestFromJson(data string) UpdateSalesItemGroupMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSalesItemGroupMasterRequestFromDict(dict)
}

func NewUpdateSalesItemGroupMasterRequestFromDict(data map[string]interface{}) UpdateSalesItemGroupMasterRequest {
	return UpdateSalesItemGroupMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		SalesItemNames:     core.CastStrings(core.CastArray(data["salesItemNames"])),
	}
}

func (p UpdateSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"salesItemGroupName": p.SalesItemGroupName,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"salesItemNames": core.CastStringsFromDict(
			p.SalesItemNames,
		),
	}
}

func (p UpdateSalesItemGroupMasterRequest) Pointer() *UpdateSalesItemGroupMasterRequest {
	return &p
}

type DeleteSalesItemGroupMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	SalesItemGroupName *string `json:"salesItemGroupName"`
}

func NewDeleteSalesItemGroupMasterRequestFromJson(data string) DeleteSalesItemGroupMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSalesItemGroupMasterRequestFromDict(dict)
}

func NewDeleteSalesItemGroupMasterRequestFromDict(data map[string]interface{}) DeleteSalesItemGroupMasterRequest {
	return DeleteSalesItemGroupMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
	}
}

func (p DeleteSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"salesItemGroupName": p.SalesItemGroupName,
	}
}

func (p DeleteSalesItemGroupMasterRequest) Pointer() *DeleteSalesItemGroupMasterRequest {
	return &p
}

type DescribeShowcaseMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeShowcaseMastersRequestFromJson(data string) DescribeShowcaseMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcaseMastersRequestFromDict(dict)
}

func NewDescribeShowcaseMastersRequestFromDict(data map[string]interface{}) DescribeShowcaseMastersRequest {
	return DescribeShowcaseMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeShowcaseMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeShowcaseMastersRequest) Pointer() *DescribeShowcaseMastersRequest {
	return &p
}

type CreateShowcaseMasterRequest struct {
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DisplayItems       []DisplayItemMaster `json:"displayItems"`
	SalesPeriodEventId *string             `json:"salesPeriodEventId"`
}

func NewCreateShowcaseMasterRequestFromJson(data string) CreateShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateShowcaseMasterRequestFromDict(dict)
}

func NewCreateShowcaseMasterRequestFromDict(data map[string]interface{}) CreateShowcaseMasterRequest {
	return CreateShowcaseMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DisplayItems:       CastDisplayItemMasters(core.CastArray(data["displayItems"])),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
	}
}

func (p CreateShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"displayItems": CastDisplayItemMastersFromDict(
			p.DisplayItems,
		),
		"salesPeriodEventId": p.SalesPeriodEventId,
	}
}

func (p CreateShowcaseMasterRequest) Pointer() *CreateShowcaseMasterRequest {
	return &p
}

type GetShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func NewGetShowcaseMasterRequestFromJson(data string) GetShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseMasterRequestFromDict(dict)
}

func NewGetShowcaseMasterRequestFromDict(data map[string]interface{}) GetShowcaseMasterRequest {
	return GetShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p GetShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p GetShowcaseMasterRequest) Pointer() *GetShowcaseMasterRequest {
	return &p
}

type UpdateShowcaseMasterRequest struct {
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	ShowcaseName       *string             `json:"showcaseName"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DisplayItems       []DisplayItemMaster `json:"displayItems"`
	SalesPeriodEventId *string             `json:"salesPeriodEventId"`
}

func NewUpdateShowcaseMasterRequestFromJson(data string) UpdateShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateShowcaseMasterRequestFromDict(dict)
}

func NewUpdateShowcaseMasterRequestFromDict(data map[string]interface{}) UpdateShowcaseMasterRequest {
	return UpdateShowcaseMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		ShowcaseName:       core.CastString(data["showcaseName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DisplayItems:       CastDisplayItemMasters(core.CastArray(data["displayItems"])),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
	}
}

func (p UpdateShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"displayItems": CastDisplayItemMastersFromDict(
			p.DisplayItems,
		),
		"salesPeriodEventId": p.SalesPeriodEventId,
	}
}

func (p UpdateShowcaseMasterRequest) Pointer() *UpdateShowcaseMasterRequest {
	return &p
}

type DeleteShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func NewDeleteShowcaseMasterRequestFromJson(data string) DeleteShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteShowcaseMasterRequestFromDict(dict)
}

func NewDeleteShowcaseMasterRequestFromDict(data map[string]interface{}) DeleteShowcaseMasterRequest {
	return DeleteShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p DeleteShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p DeleteShowcaseMasterRequest) Pointer() *DeleteShowcaseMasterRequest {
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

type GetCurrentShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentShowcaseMasterRequestFromJson(data string) GetCurrentShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentShowcaseMasterRequestFromDict(dict)
}

func NewGetCurrentShowcaseMasterRequestFromDict(data map[string]interface{}) GetCurrentShowcaseMasterRequest {
	return GetCurrentShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentShowcaseMasterRequest) Pointer() *GetCurrentShowcaseMasterRequest {
	return &p
}

type UpdateCurrentShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentShowcaseMasterRequestFromJson(data string) UpdateCurrentShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentShowcaseMasterRequestFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterRequestFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterRequest {
	return UpdateCurrentShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentShowcaseMasterRequest) Pointer() *UpdateCurrentShowcaseMasterRequest {
	return &p
}

type UpdateCurrentShowcaseMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentShowcaseMasterFromGitHubRequestFromJson(data string) UpdateCurrentShowcaseMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentShowcaseMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterFromGitHubRequest {
	return UpdateCurrentShowcaseMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentShowcaseMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentShowcaseMasterFromGitHubRequest) Pointer() *UpdateCurrentShowcaseMasterFromGitHubRequest {
	return &p
}

type DescribeShowcasesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
}

func NewDescribeShowcasesRequestFromJson(data string) DescribeShowcasesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcasesRequestFromDict(dict)
}

func NewDescribeShowcasesRequestFromDict(data map[string]interface{}) DescribeShowcasesRequest {
	return DescribeShowcasesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DescribeShowcasesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeShowcasesRequest) Pointer() *DescribeShowcasesRequest {
	return &p
}

type DescribeShowcasesByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
}

func NewDescribeShowcasesByUserIdRequestFromJson(data string) DescribeShowcasesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcasesByUserIdRequestFromDict(dict)
}

func NewDescribeShowcasesByUserIdRequestFromDict(data map[string]interface{}) DescribeShowcasesByUserIdRequest {
	return DescribeShowcasesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DescribeShowcasesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p DescribeShowcasesByUserIdRequest) Pointer() *DescribeShowcasesByUserIdRequest {
	return &p
}

type GetShowcaseRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	AccessToken     *string `json:"accessToken"`
}

func NewGetShowcaseRequestFromJson(data string) GetShowcaseRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseRequestFromDict(dict)
}

func NewGetShowcaseRequestFromDict(data map[string]interface{}) GetShowcaseRequest {
	return GetShowcaseRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetShowcaseRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetShowcaseRequest) Pointer() *GetShowcaseRequest {
	return &p
}

type GetShowcaseByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	UserId          *string `json:"userId"`
}

func NewGetShowcaseByUserIdRequestFromJson(data string) GetShowcaseByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseByUserIdRequestFromDict(dict)
}

func NewGetShowcaseByUserIdRequestFromDict(data map[string]interface{}) GetShowcaseByUserIdRequest {
	return GetShowcaseByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetShowcaseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"userId":        p.UserId,
	}
}

func (p GetShowcaseByUserIdRequest) Pointer() *GetShowcaseByUserIdRequest {
	return &p
}

type BuyRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	ShowcaseName       *string  `json:"showcaseName"`
	DisplayItemId      *string  `json:"displayItemId"`
	AccessToken        *string  `json:"accessToken"`
	Quantity           *int32   `json:"quantity"`
	Config             []Config `json:"config"`
}

func NewBuyRequestFromJson(data string) BuyRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuyRequestFromDict(dict)
}

func NewBuyRequestFromDict(data map[string]interface{}) BuyRequest {
	return BuyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		DisplayItemId: core.CastString(data["displayItemId"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Quantity:      core.CastInt32(data["quantity"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p BuyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"displayItemId": p.DisplayItemId,
		"accessToken":   p.AccessToken,
		"quantity":      p.Quantity,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p BuyRequest) Pointer() *BuyRequest {
	return &p
}

type BuyByUserIdRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	ShowcaseName       *string  `json:"showcaseName"`
	DisplayItemId      *string  `json:"displayItemId"`
	UserId             *string  `json:"userId"`
	Quantity           *int32   `json:"quantity"`
	Config             []Config `json:"config"`
}

func NewBuyByUserIdRequestFromJson(data string) BuyByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuyByUserIdRequestFromDict(dict)
}

func NewBuyByUserIdRequestFromDict(data map[string]interface{}) BuyByUserIdRequest {
	return BuyByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		DisplayItemId: core.CastString(data["displayItemId"]),
		UserId:        core.CastString(data["userId"]),
		Quantity:      core.CastInt32(data["quantity"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p BuyByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"displayItemId": p.DisplayItemId,
		"userId":        p.UserId,
		"quantity":      p.Quantity,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p BuyByUserIdRequest) Pointer() *BuyByUserIdRequest {
	return &p
}

type DescribeRandomShowcaseMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeRandomShowcaseMastersRequestFromJson(data string) DescribeRandomShowcaseMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomShowcaseMastersRequestFromDict(dict)
}

func NewDescribeRandomShowcaseMastersRequestFromDict(data map[string]interface{}) DescribeRandomShowcaseMastersRequest {
	return DescribeRandomShowcaseMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRandomShowcaseMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRandomShowcaseMastersRequest) Pointer() *DescribeRandomShowcaseMastersRequest {
	return &p
}

type CreateRandomShowcaseMasterRequest struct {
	SourceRequestId       *string                  `json:"sourceRequestId"`
	RequestId             *string                  `json:"requestId"`
	ContextStack          *string                  `json:"contextStack"`
	NamespaceName         *string                  `json:"namespaceName"`
	Name                  *string                  `json:"name"`
	Description           *string                  `json:"description"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
}

func NewCreateRandomShowcaseMasterRequestFromJson(data string) CreateRandomShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRandomShowcaseMasterRequestFromDict(dict)
}

func NewCreateRandomShowcaseMasterRequestFromDict(data map[string]interface{}) CreateRandomShowcaseMasterRequest {
	return CreateRandomShowcaseMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
	}
}

func (p CreateRandomShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"name":                  p.Name,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"maximumNumberOfChoice": p.MaximumNumberOfChoice,
		"displayItems": CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		),
		"baseTimestamp":      p.BaseTimestamp,
		"resetIntervalHours": p.ResetIntervalHours,
		"salesPeriodEventId": p.SalesPeriodEventId,
	}
}

func (p CreateRandomShowcaseMasterRequest) Pointer() *CreateRandomShowcaseMasterRequest {
	return &p
}

type GetRandomShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func NewGetRandomShowcaseMasterRequestFromJson(data string) GetRandomShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomShowcaseMasterRequestFromDict(dict)
}

func NewGetRandomShowcaseMasterRequestFromDict(data map[string]interface{}) GetRandomShowcaseMasterRequest {
	return GetRandomShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p GetRandomShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p GetRandomShowcaseMasterRequest) Pointer() *GetRandomShowcaseMasterRequest {
	return &p
}

type UpdateRandomShowcaseMasterRequest struct {
	SourceRequestId       *string                  `json:"sourceRequestId"`
	RequestId             *string                  `json:"requestId"`
	ContextStack          *string                  `json:"contextStack"`
	NamespaceName         *string                  `json:"namespaceName"`
	ShowcaseName          *string                  `json:"showcaseName"`
	Description           *string                  `json:"description"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
}

func NewUpdateRandomShowcaseMasterRequestFromJson(data string) UpdateRandomShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRandomShowcaseMasterRequestFromDict(dict)
}

func NewUpdateRandomShowcaseMasterRequestFromDict(data map[string]interface{}) UpdateRandomShowcaseMasterRequest {
	return UpdateRandomShowcaseMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		ShowcaseName:          core.CastString(data["showcaseName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
	}
}

func (p UpdateRandomShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"showcaseName":          p.ShowcaseName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"maximumNumberOfChoice": p.MaximumNumberOfChoice,
		"displayItems": CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		),
		"baseTimestamp":      p.BaseTimestamp,
		"resetIntervalHours": p.ResetIntervalHours,
		"salesPeriodEventId": p.SalesPeriodEventId,
	}
}

func (p UpdateRandomShowcaseMasterRequest) Pointer() *UpdateRandomShowcaseMasterRequest {
	return &p
}

type DeleteRandomShowcaseMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func NewDeleteRandomShowcaseMasterRequestFromJson(data string) DeleteRandomShowcaseMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRandomShowcaseMasterRequestFromDict(dict)
}

func NewDeleteRandomShowcaseMasterRequestFromDict(data map[string]interface{}) DeleteRandomShowcaseMasterRequest {
	return DeleteRandomShowcaseMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p DeleteRandomShowcaseMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p DeleteRandomShowcaseMasterRequest) Pointer() *DeleteRandomShowcaseMasterRequest {
	return &p
}

type IncrementPurchaseCountByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	ShowcaseName       *string `json:"showcaseName"`
	DisplayItemName    *string `json:"displayItemName"`
	UserId             *string `json:"userId"`
	Count              *int32  `json:"count"`
}

func NewIncrementPurchaseCountByUserIdRequestFromJson(data string) IncrementPurchaseCountByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementPurchaseCountByUserIdRequestFromDict(dict)
}

func NewIncrementPurchaseCountByUserIdRequestFromDict(data map[string]interface{}) IncrementPurchaseCountByUserIdRequest {
	return IncrementPurchaseCountByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		UserId:          core.CastString(data["userId"]),
		Count:           core.CastInt32(data["count"]),
	}
}

func (p IncrementPurchaseCountByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"userId":          p.UserId,
		"count":           p.Count,
	}
}

func (p IncrementPurchaseCountByUserIdRequest) Pointer() *IncrementPurchaseCountByUserIdRequest {
	return &p
}

type DecrementPurchaseCountByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	ShowcaseName       *string `json:"showcaseName"`
	DisplayItemName    *string `json:"displayItemName"`
	UserId             *string `json:"userId"`
	Count              *int32  `json:"count"`
}

func NewDecrementPurchaseCountByUserIdRequestFromJson(data string) DecrementPurchaseCountByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecrementPurchaseCountByUserIdRequestFromDict(dict)
}

func NewDecrementPurchaseCountByUserIdRequestFromDict(data map[string]interface{}) DecrementPurchaseCountByUserIdRequest {
	return DecrementPurchaseCountByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		UserId:          core.CastString(data["userId"]),
		Count:           core.CastInt32(data["count"]),
	}
}

func (p DecrementPurchaseCountByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"userId":          p.UserId,
		"count":           p.Count,
	}
}

func (p DecrementPurchaseCountByUserIdRequest) Pointer() *DecrementPurchaseCountByUserIdRequest {
	return &p
}

type IncrementPurchaseCountByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewIncrementPurchaseCountByStampTaskRequestFromJson(data string) IncrementPurchaseCountByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementPurchaseCountByStampTaskRequestFromDict(dict)
}

func NewIncrementPurchaseCountByStampTaskRequestFromDict(data map[string]interface{}) IncrementPurchaseCountByStampTaskRequest {
	return IncrementPurchaseCountByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p IncrementPurchaseCountByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p IncrementPurchaseCountByStampTaskRequest) Pointer() *IncrementPurchaseCountByStampTaskRequest {
	return &p
}

type DecrementPurchaseCountByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewDecrementPurchaseCountByStampSheetRequestFromJson(data string) DecrementPurchaseCountByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecrementPurchaseCountByStampSheetRequestFromDict(dict)
}

func NewDecrementPurchaseCountByStampSheetRequestFromDict(data map[string]interface{}) DecrementPurchaseCountByStampSheetRequest {
	return DecrementPurchaseCountByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DecrementPurchaseCountByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DecrementPurchaseCountByStampSheetRequest) Pointer() *DecrementPurchaseCountByStampSheetRequest {
	return &p
}

type ForceReDrawByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	ShowcaseName       *string `json:"showcaseName"`
	UserId             *string `json:"userId"`
}

func NewForceReDrawByUserIdRequestFromJson(data string) ForceReDrawByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForceReDrawByUserIdRequestFromDict(dict)
}

func NewForceReDrawByUserIdRequestFromDict(data map[string]interface{}) ForceReDrawByUserIdRequest {
	return ForceReDrawByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p ForceReDrawByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"userId":        p.UserId,
	}
}

func (p ForceReDrawByUserIdRequest) Pointer() *ForceReDrawByUserIdRequest {
	return &p
}

type ForceReDrawByUserIdByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewForceReDrawByUserIdByStampSheetRequestFromJson(data string) ForceReDrawByUserIdByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForceReDrawByUserIdByStampSheetRequestFromDict(dict)
}

func NewForceReDrawByUserIdByStampSheetRequestFromDict(data map[string]interface{}) ForceReDrawByUserIdByStampSheetRequest {
	return ForceReDrawByUserIdByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p ForceReDrawByUserIdByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ForceReDrawByUserIdByStampSheetRequest) Pointer() *ForceReDrawByUserIdByStampSheetRequest {
	return &p
}

type DescribeRandomDisplayItemsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	AccessToken     *string `json:"accessToken"`
}

func NewDescribeRandomDisplayItemsRequestFromJson(data string) DescribeRandomDisplayItemsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomDisplayItemsRequestFromDict(dict)
}

func NewDescribeRandomDisplayItemsRequestFromDict(data map[string]interface{}) DescribeRandomDisplayItemsRequest {
	return DescribeRandomDisplayItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DescribeRandomDisplayItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeRandomDisplayItemsRequest) Pointer() *DescribeRandomDisplayItemsRequest {
	return &p
}

type DescribeRandomDisplayItemsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	UserId          *string `json:"userId"`
}

func NewDescribeRandomDisplayItemsByUserIdRequestFromJson(data string) DescribeRandomDisplayItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomDisplayItemsByUserIdRequestFromDict(dict)
}

func NewDescribeRandomDisplayItemsByUserIdRequestFromDict(data map[string]interface{}) DescribeRandomDisplayItemsByUserIdRequest {
	return DescribeRandomDisplayItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DescribeRandomDisplayItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"userId":        p.UserId,
	}
}

func (p DescribeRandomDisplayItemsByUserIdRequest) Pointer() *DescribeRandomDisplayItemsByUserIdRequest {
	return &p
}

type GetRandomDisplayItemRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	DisplayItemName *string `json:"displayItemName"`
	AccessToken     *string `json:"accessToken"`
}

func NewGetRandomDisplayItemRequestFromJson(data string) GetRandomDisplayItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomDisplayItemRequestFromDict(dict)
}

func NewGetRandomDisplayItemRequestFromDict(data map[string]interface{}) GetRandomDisplayItemRequest {
	return GetRandomDisplayItemRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		AccessToken:     core.CastString(data["accessToken"]),
	}
}

func (p GetRandomDisplayItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"accessToken":     p.AccessToken,
	}
}

func (p GetRandomDisplayItemRequest) Pointer() *GetRandomDisplayItemRequest {
	return &p
}

type GetRandomDisplayItemByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	DisplayItemName *string `json:"displayItemName"`
	UserId          *string `json:"userId"`
}

func NewGetRandomDisplayItemByUserIdRequestFromJson(data string) GetRandomDisplayItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomDisplayItemByUserIdRequestFromDict(dict)
}

func NewGetRandomDisplayItemByUserIdRequestFromDict(data map[string]interface{}) GetRandomDisplayItemByUserIdRequest {
	return GetRandomDisplayItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		UserId:          core.CastString(data["userId"]),
	}
}

func (p GetRandomDisplayItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"userId":          p.UserId,
	}
}

func (p GetRandomDisplayItemByUserIdRequest) Pointer() *GetRandomDisplayItemByUserIdRequest {
	return &p
}

type RandomShowcaseBuyRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	ShowcaseName       *string  `json:"showcaseName"`
	DisplayItemName    *string  `json:"displayItemName"`
	AccessToken        *string  `json:"accessToken"`
	Quantity           *int32   `json:"quantity"`
	Config             []Config `json:"config"`
}

func NewRandomShowcaseBuyRequestFromJson(data string) RandomShowcaseBuyRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseBuyRequestFromDict(dict)
}

func NewRandomShowcaseBuyRequestFromDict(data map[string]interface{}) RandomShowcaseBuyRequest {
	return RandomShowcaseBuyRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		Quantity:        core.CastInt32(data["quantity"]),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RandomShowcaseBuyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"accessToken":     p.AccessToken,
		"quantity":        p.Quantity,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RandomShowcaseBuyRequest) Pointer() *RandomShowcaseBuyRequest {
	return &p
}

type RandomShowcaseBuyByUserIdRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	ShowcaseName       *string  `json:"showcaseName"`
	DisplayItemName    *string  `json:"displayItemName"`
	UserId             *string  `json:"userId"`
	Quantity           *int32   `json:"quantity"`
	Config             []Config `json:"config"`
}

func NewRandomShowcaseBuyByUserIdRequestFromJson(data string) RandomShowcaseBuyByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseBuyByUserIdRequestFromDict(dict)
}

func NewRandomShowcaseBuyByUserIdRequestFromDict(data map[string]interface{}) RandomShowcaseBuyByUserIdRequest {
	return RandomShowcaseBuyByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ShowcaseName:    core.CastString(data["showcaseName"]),
		DisplayItemName: core.CastString(data["displayItemName"]),
		UserId:          core.CastString(data["userId"]),
		Quantity:        core.CastInt32(data["quantity"]),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RandomShowcaseBuyByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"showcaseName":    p.ShowcaseName,
		"displayItemName": p.DisplayItemName,
		"userId":          p.UserId,
		"quantity":        p.Quantity,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RandomShowcaseBuyByUserIdRequest) Pointer() *RandomShowcaseBuyByUserIdRequest {
	return &p
}
