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

package inventory

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
	RequestId      *string        `json:"requestId"`
	ContextStack   *string        `json:"contextStack"`
	Name           *string        `json:"name"`
	Description    *string        `json:"description"`
	AcquireScript  *ScriptSetting `json:"acquireScript"`
	OverflowScript *ScriptSetting `json:"overflowScript"`
	ConsumeScript  *ScriptSetting `json:"consumeScript"`
	LogSetting     *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		AcquireScript:  NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript: NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:  NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		LogSetting:     NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":           p.Name,
		"description":    p.Description,
		"acquireScript":  p.AcquireScript.ToDict(),
		"overflowScript": p.OverflowScript.ToDict(),
		"consumeScript":  p.ConsumeScript.ToDict(),
		"logSetting":     p.LogSetting.ToDict(),
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
	RequestId      *string        `json:"requestId"`
	ContextStack   *string        `json:"contextStack"`
	NamespaceName  *string        `json:"namespaceName"`
	Description    *string        `json:"description"`
	AcquireScript  *ScriptSetting `json:"acquireScript"`
	OverflowScript *ScriptSetting `json:"overflowScript"`
	ConsumeScript  *ScriptSetting `json:"consumeScript"`
	LogSetting     *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Description:    core.CastString(data["description"]),
		AcquireScript:  NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript: NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:  NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		LogSetting:     NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"description":    p.Description,
		"acquireScript":  p.AcquireScript.ToDict(),
		"overflowScript": p.OverflowScript.ToDict(),
		"consumeScript":  p.ConsumeScript.ToDict(),
		"logSetting":     p.LogSetting.ToDict(),
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

type DescribeInventoryModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeInventoryModelMastersRequestFromJson(data string) DescribeInventoryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryModelMastersRequestFromDict(dict)
}

func NewDescribeInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeInventoryModelMastersRequest {
	return DescribeInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInventoryModelMastersRequest) Pointer() *DescribeInventoryModelMastersRequest {
	return &p
}

type CreateInventoryModelMasterRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	Name                  *string `json:"name"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
}

func NewCreateInventoryModelMasterRequestFromJson(data string) CreateInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateInventoryModelMasterRequestFromDict(dict)
}

func NewCreateInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateInventoryModelMasterRequest {
	return CreateInventoryModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
	}
}

func (p CreateInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"name":                  p.Name,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
	}
}

func (p CreateInventoryModelMasterRequest) Pointer() *CreateInventoryModelMasterRequest {
	return &p
}

type GetInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetInventoryModelMasterRequestFromJson(data string) GetInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryModelMasterRequestFromDict(dict)
}

func NewGetInventoryModelMasterRequestFromDict(data map[string]interface{}) GetInventoryModelMasterRequest {
	return GetInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetInventoryModelMasterRequest) Pointer() *GetInventoryModelMasterRequest {
	return &p
}

type UpdateInventoryModelMasterRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	InventoryName         *string `json:"inventoryName"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
}

func NewUpdateInventoryModelMasterRequestFromJson(data string) UpdateInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateInventoryModelMasterRequestFromDict(dict)
}

func NewUpdateInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateInventoryModelMasterRequest {
	return UpdateInventoryModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		InventoryName:         core.CastString(data["inventoryName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
	}
}

func (p UpdateInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"inventoryName":         p.InventoryName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
	}
}

func (p UpdateInventoryModelMasterRequest) Pointer() *UpdateInventoryModelMasterRequest {
	return &p
}

type DeleteInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDeleteInventoryModelMasterRequestFromJson(data string) DeleteInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInventoryModelMasterRequestFromDict(dict)
}

func NewDeleteInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteInventoryModelMasterRequest {
	return DeleteInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteInventoryModelMasterRequest) Pointer() *DeleteInventoryModelMasterRequest {
	return &p
}

type DescribeInventoryModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeInventoryModelsRequestFromJson(data string) DescribeInventoryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryModelsRequestFromDict(dict)
}

func NewDescribeInventoryModelsRequestFromDict(data map[string]interface{}) DescribeInventoryModelsRequest {
	return DescribeInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeInventoryModelsRequest) Pointer() *DescribeInventoryModelsRequest {
	return &p
}

type GetInventoryModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetInventoryModelRequestFromJson(data string) GetInventoryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryModelRequestFromDict(dict)
}

func NewGetInventoryModelRequestFromDict(data map[string]interface{}) GetInventoryModelRequest {
	return GetInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetInventoryModelRequest) Pointer() *GetInventoryModelRequest {
	return &p
}

type DescribeItemModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeItemModelMastersRequestFromJson(data string) DescribeItemModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemModelMastersRequestFromDict(dict)
}

func NewDescribeItemModelMastersRequestFromDict(data map[string]interface{}) DescribeItemModelMastersRequest {
	return DescribeItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeItemModelMastersRequest) Pointer() *DescribeItemModelMastersRequest {
	return &p
}

type CreateItemModelMasterRequest struct {
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func NewCreateItemModelMasterRequestFromJson(data string) CreateItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateItemModelMasterRequestFromDict(dict)
}

func NewCreateItemModelMasterRequestFromDict(data map[string]interface{}) CreateItemModelMasterRequest {
	return CreateItemModelMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		InventoryName:       core.CastString(data["inventoryName"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
	}
}

func (p CreateItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"inventoryName":       p.InventoryName,
		"name":                p.Name,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
	}
}

func (p CreateItemModelMasterRequest) Pointer() *CreateItemModelMasterRequest {
	return &p
}

type GetItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetItemModelMasterRequestFromJson(data string) GetItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemModelMasterRequestFromDict(dict)
}

func NewGetItemModelMasterRequestFromDict(data map[string]interface{}) GetItemModelMasterRequest {
	return GetItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetItemModelMasterRequest) Pointer() *GetItemModelMasterRequest {
	return &p
}

type UpdateItemModelMasterRequest struct {
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	ItemName            *string `json:"itemName"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func NewUpdateItemModelMasterRequestFromJson(data string) UpdateItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateItemModelMasterRequestFromDict(dict)
}

func NewUpdateItemModelMasterRequestFromDict(data map[string]interface{}) UpdateItemModelMasterRequest {
	return UpdateItemModelMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		InventoryName:       core.CastString(data["inventoryName"]),
		ItemName:            core.CastString(data["itemName"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
	}
}

func (p UpdateItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"inventoryName":       p.InventoryName,
		"itemName":            p.ItemName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
	}
}

func (p UpdateItemModelMasterRequest) Pointer() *UpdateItemModelMasterRequest {
	return &p
}

type DeleteItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewDeleteItemModelMasterRequestFromJson(data string) DeleteItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteItemModelMasterRequestFromDict(dict)
}

func NewDeleteItemModelMasterRequestFromDict(data map[string]interface{}) DeleteItemModelMasterRequest {
	return DeleteItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteItemModelMasterRequest) Pointer() *DeleteItemModelMasterRequest {
	return &p
}

type DescribeItemModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDescribeItemModelsRequestFromJson(data string) DescribeItemModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemModelsRequestFromDict(dict)
}

func NewDescribeItemModelsRequestFromDict(data map[string]interface{}) DescribeItemModelsRequest {
	return DescribeItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeItemModelsRequest) Pointer() *DescribeItemModelsRequest {
	return &p
}

type GetItemModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetItemModelRequestFromJson(data string) GetItemModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemModelRequestFromDict(dict)
}

func NewGetItemModelRequestFromDict(data map[string]interface{}) GetItemModelRequest {
	return GetItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetItemModelRequest) Pointer() *GetItemModelRequest {
	return &p
}

type DescribeSimpleInventoryModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSimpleInventoryModelMastersRequestFromJson(data string) DescribeSimpleInventoryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleInventoryModelMastersRequestFromDict(dict)
}

func NewDescribeSimpleInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeSimpleInventoryModelMastersRequest {
	return DescribeSimpleInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleInventoryModelMastersRequest) Pointer() *DescribeSimpleInventoryModelMastersRequest {
	return &p
}

type CreateSimpleInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateSimpleInventoryModelMasterRequestFromJson(data string) CreateSimpleInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSimpleInventoryModelMasterRequestFromDict(dict)
}

func NewCreateSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateSimpleInventoryModelMasterRequest {
	return CreateSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateSimpleInventoryModelMasterRequest) Pointer() *CreateSimpleInventoryModelMasterRequest {
	return &p
}

type GetSimpleInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetSimpleInventoryModelMasterRequestFromJson(data string) GetSimpleInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleInventoryModelMasterRequestFromDict(dict)
}

func NewGetSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) GetSimpleInventoryModelMasterRequest {
	return GetSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetSimpleInventoryModelMasterRequest) Pointer() *GetSimpleInventoryModelMasterRequest {
	return &p
}

type UpdateSimpleInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewUpdateSimpleInventoryModelMasterRequestFromJson(data string) UpdateSimpleInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSimpleInventoryModelMasterRequestFromDict(dict)
}

func NewUpdateSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateSimpleInventoryModelMasterRequest {
	return UpdateSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateSimpleInventoryModelMasterRequest) Pointer() *UpdateSimpleInventoryModelMasterRequest {
	return &p
}

type DeleteSimpleInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDeleteSimpleInventoryModelMasterRequestFromJson(data string) DeleteSimpleInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleInventoryModelMasterRequestFromDict(dict)
}

func NewDeleteSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteSimpleInventoryModelMasterRequest {
	return DeleteSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteSimpleInventoryModelMasterRequest) Pointer() *DeleteSimpleInventoryModelMasterRequest {
	return &p
}

type DescribeSimpleInventoryModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeSimpleInventoryModelsRequestFromJson(data string) DescribeSimpleInventoryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleInventoryModelsRequestFromDict(dict)
}

func NewDescribeSimpleInventoryModelsRequestFromDict(data map[string]interface{}) DescribeSimpleInventoryModelsRequest {
	return DescribeSimpleInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeSimpleInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeSimpleInventoryModelsRequest) Pointer() *DescribeSimpleInventoryModelsRequest {
	return &p
}

type GetSimpleInventoryModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetSimpleInventoryModelRequestFromJson(data string) GetSimpleInventoryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleInventoryModelRequestFromDict(dict)
}

func NewGetSimpleInventoryModelRequestFromDict(data map[string]interface{}) GetSimpleInventoryModelRequest {
	return GetSimpleInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetSimpleInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetSimpleInventoryModelRequest) Pointer() *GetSimpleInventoryModelRequest {
	return &p
}

type DescribeSimpleItemModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSimpleItemModelMastersRequestFromJson(data string) DescribeSimpleItemModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemModelMastersRequestFromDict(dict)
}

func NewDescribeSimpleItemModelMastersRequestFromDict(data map[string]interface{}) DescribeSimpleItemModelMastersRequest {
	return DescribeSimpleItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleItemModelMastersRequest) Pointer() *DescribeSimpleItemModelMastersRequest {
	return &p
}

type CreateSimpleItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateSimpleItemModelMasterRequestFromJson(data string) CreateSimpleItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSimpleItemModelMasterRequestFromDict(dict)
}

func NewCreateSimpleItemModelMasterRequestFromDict(data map[string]interface{}) CreateSimpleItemModelMasterRequest {
	return CreateSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateSimpleItemModelMasterRequest) Pointer() *CreateSimpleItemModelMasterRequest {
	return &p
}

type GetSimpleItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetSimpleItemModelMasterRequestFromJson(data string) GetSimpleItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemModelMasterRequestFromDict(dict)
}

func NewGetSimpleItemModelMasterRequestFromDict(data map[string]interface{}) GetSimpleItemModelMasterRequest {
	return GetSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemModelMasterRequest) Pointer() *GetSimpleItemModelMasterRequest {
	return &p
}

type UpdateSimpleItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewUpdateSimpleItemModelMasterRequestFromJson(data string) UpdateSimpleItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSimpleItemModelMasterRequestFromDict(dict)
}

func NewUpdateSimpleItemModelMasterRequestFromDict(data map[string]interface{}) UpdateSimpleItemModelMasterRequest {
	return UpdateSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateSimpleItemModelMasterRequest) Pointer() *UpdateSimpleItemModelMasterRequest {
	return &p
}

type DeleteSimpleItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewDeleteSimpleItemModelMasterRequestFromJson(data string) DeleteSimpleItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleItemModelMasterRequestFromDict(dict)
}

func NewDeleteSimpleItemModelMasterRequestFromDict(data map[string]interface{}) DeleteSimpleItemModelMasterRequest {
	return DeleteSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteSimpleItemModelMasterRequest) Pointer() *DeleteSimpleItemModelMasterRequest {
	return &p
}

type DescribeSimpleItemModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDescribeSimpleItemModelsRequestFromJson(data string) DescribeSimpleItemModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemModelsRequestFromDict(dict)
}

func NewDescribeSimpleItemModelsRequestFromDict(data map[string]interface{}) DescribeSimpleItemModelsRequest {
	return DescribeSimpleItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeSimpleItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeSimpleItemModelsRequest) Pointer() *DescribeSimpleItemModelsRequest {
	return &p
}

type GetSimpleItemModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetSimpleItemModelRequestFromJson(data string) GetSimpleItemModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemModelRequestFromDict(dict)
}

func NewGetSimpleItemModelRequestFromDict(data map[string]interface{}) GetSimpleItemModelRequest {
	return GetSimpleItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemModelRequest) Pointer() *GetSimpleItemModelRequest {
	return &p
}

type DescribeBigInventoryModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBigInventoryModelMastersRequestFromJson(data string) DescribeBigInventoryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigInventoryModelMastersRequestFromDict(dict)
}

func NewDescribeBigInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeBigInventoryModelMastersRequest {
	return DescribeBigInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigInventoryModelMastersRequest) Pointer() *DescribeBigInventoryModelMastersRequest {
	return &p
}

type CreateBigInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateBigInventoryModelMasterRequestFromJson(data string) CreateBigInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBigInventoryModelMasterRequestFromDict(dict)
}

func NewCreateBigInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateBigInventoryModelMasterRequest {
	return CreateBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateBigInventoryModelMasterRequest) Pointer() *CreateBigInventoryModelMasterRequest {
	return &p
}

type GetBigInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetBigInventoryModelMasterRequestFromJson(data string) GetBigInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigInventoryModelMasterRequestFromDict(dict)
}

func NewGetBigInventoryModelMasterRequestFromDict(data map[string]interface{}) GetBigInventoryModelMasterRequest {
	return GetBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetBigInventoryModelMasterRequest) Pointer() *GetBigInventoryModelMasterRequest {
	return &p
}

type UpdateBigInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewUpdateBigInventoryModelMasterRequestFromJson(data string) UpdateBigInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBigInventoryModelMasterRequestFromDict(dict)
}

func NewUpdateBigInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateBigInventoryModelMasterRequest {
	return UpdateBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateBigInventoryModelMasterRequest) Pointer() *UpdateBigInventoryModelMasterRequest {
	return &p
}

type DeleteBigInventoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDeleteBigInventoryModelMasterRequestFromJson(data string) DeleteBigInventoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigInventoryModelMasterRequestFromDict(dict)
}

func NewDeleteBigInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteBigInventoryModelMasterRequest {
	return DeleteBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteBigInventoryModelMasterRequest) Pointer() *DeleteBigInventoryModelMasterRequest {
	return &p
}

type DescribeBigInventoryModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeBigInventoryModelsRequestFromJson(data string) DescribeBigInventoryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigInventoryModelsRequestFromDict(dict)
}

func NewDescribeBigInventoryModelsRequestFromDict(data map[string]interface{}) DescribeBigInventoryModelsRequest {
	return DescribeBigInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeBigInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeBigInventoryModelsRequest) Pointer() *DescribeBigInventoryModelsRequest {
	return &p
}

type GetBigInventoryModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewGetBigInventoryModelRequestFromJson(data string) GetBigInventoryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigInventoryModelRequestFromDict(dict)
}

func NewGetBigInventoryModelRequestFromDict(data map[string]interface{}) GetBigInventoryModelRequest {
	return GetBigInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetBigInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetBigInventoryModelRequest) Pointer() *GetBigInventoryModelRequest {
	return &p
}

type DescribeBigItemModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBigItemModelMastersRequestFromJson(data string) DescribeBigItemModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemModelMastersRequestFromDict(dict)
}

func NewDescribeBigItemModelMastersRequestFromDict(data map[string]interface{}) DescribeBigItemModelMastersRequest {
	return DescribeBigItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigItemModelMastersRequest) Pointer() *DescribeBigItemModelMastersRequest {
	return &p
}

type CreateBigItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateBigItemModelMasterRequestFromJson(data string) CreateBigItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBigItemModelMasterRequestFromDict(dict)
}

func NewCreateBigItemModelMasterRequestFromDict(data map[string]interface{}) CreateBigItemModelMasterRequest {
	return CreateBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateBigItemModelMasterRequest) Pointer() *CreateBigItemModelMasterRequest {
	return &p
}

type GetBigItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetBigItemModelMasterRequestFromJson(data string) GetBigItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemModelMasterRequestFromDict(dict)
}

func NewGetBigItemModelMasterRequestFromDict(data map[string]interface{}) GetBigItemModelMasterRequest {
	return GetBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemModelMasterRequest) Pointer() *GetBigItemModelMasterRequest {
	return &p
}

type UpdateBigItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewUpdateBigItemModelMasterRequestFromJson(data string) UpdateBigItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBigItemModelMasterRequestFromDict(dict)
}

func NewUpdateBigItemModelMasterRequestFromDict(data map[string]interface{}) UpdateBigItemModelMasterRequest {
	return UpdateBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateBigItemModelMasterRequest) Pointer() *UpdateBigItemModelMasterRequest {
	return &p
}

type DeleteBigItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewDeleteBigItemModelMasterRequestFromJson(data string) DeleteBigItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigItemModelMasterRequestFromDict(dict)
}

func NewDeleteBigItemModelMasterRequestFromDict(data map[string]interface{}) DeleteBigItemModelMasterRequest {
	return DeleteBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteBigItemModelMasterRequest) Pointer() *DeleteBigItemModelMasterRequest {
	return &p
}

type DescribeBigItemModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
}

func NewDescribeBigItemModelsRequestFromJson(data string) DescribeBigItemModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemModelsRequestFromDict(dict)
}

func NewDescribeBigItemModelsRequestFromDict(data map[string]interface{}) DescribeBigItemModelsRequest {
	return DescribeBigItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeBigItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeBigItemModelsRequest) Pointer() *DescribeBigItemModelsRequest {
	return &p
}

type GetBigItemModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	ItemName      *string `json:"itemName"`
}

func NewGetBigItemModelRequestFromJson(data string) GetBigItemModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemModelRequestFromDict(dict)
}

func NewGetBigItemModelRequestFromDict(data map[string]interface{}) GetBigItemModelRequest {
	return GetBigItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemModelRequest) Pointer() *GetBigItemModelRequest {
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

type GetCurrentItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentItemModelMasterRequestFromJson(data string) GetCurrentItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentItemModelMasterRequestFromDict(dict)
}

func NewGetCurrentItemModelMasterRequestFromDict(data map[string]interface{}) GetCurrentItemModelMasterRequest {
	return GetCurrentItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentItemModelMasterRequest) Pointer() *GetCurrentItemModelMasterRequest {
	return &p
}

type UpdateCurrentItemModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentItemModelMasterRequestFromJson(data string) UpdateCurrentItemModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentItemModelMasterRequestFromDict(dict)
}

func NewUpdateCurrentItemModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterRequest {
	return UpdateCurrentItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentItemModelMasterRequest) Pointer() *UpdateCurrentItemModelMasterRequest {
	return &p
}

type UpdateCurrentItemModelMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromJson(data string) UpdateCurrentItemModelMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentItemModelMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterFromGitHubRequest {
	return UpdateCurrentItemModelMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) Pointer() *UpdateCurrentItemModelMasterFromGitHubRequest {
	return &p
}

type DescribeInventoriesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeInventoriesRequestFromJson(data string) DescribeInventoriesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoriesRequestFromDict(dict)
}

func NewDescribeInventoriesRequestFromDict(data map[string]interface{}) DescribeInventoriesRequest {
	return DescribeInventoriesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInventoriesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInventoriesRequest) Pointer() *DescribeInventoriesRequest {
	return &p
}

type DescribeInventoriesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeInventoriesByUserIdRequestFromJson(data string) DescribeInventoriesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoriesByUserIdRequestFromDict(dict)
}

func NewDescribeInventoriesByUserIdRequestFromDict(data map[string]interface{}) DescribeInventoriesByUserIdRequest {
	return DescribeInventoriesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInventoriesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInventoriesByUserIdRequest) Pointer() *DescribeInventoriesByUserIdRequest {
	return &p
}

type GetInventoryRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
}

func NewGetInventoryRequestFromJson(data string) GetInventoryRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryRequestFromDict(dict)
}

func NewGetInventoryRequestFromDict(data map[string]interface{}) GetInventoryRequest {
	return GetInventoryRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetInventoryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetInventoryRequest) Pointer() *GetInventoryRequest {
	return &p
}

type GetInventoryByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
}

func NewGetInventoryByUserIdRequestFromJson(data string) GetInventoryByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryByUserIdRequestFromDict(dict)
}

func NewGetInventoryByUserIdRequestFromDict(data map[string]interface{}) GetInventoryByUserIdRequest {
	return GetInventoryByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetInventoryByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
	}
}

func (p GetInventoryByUserIdRequest) Pointer() *GetInventoryByUserIdRequest {
	return &p
}

type AddCapacityByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	AddCapacityValue   *int32  `json:"addCapacityValue"`
}

func NewAddCapacityByUserIdRequestFromJson(data string) AddCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddCapacityByUserIdRequestFromDict(dict)
}

func NewAddCapacityByUserIdRequestFromDict(data map[string]interface{}) AddCapacityByUserIdRequest {
	return AddCapacityByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		UserId:           core.CastString(data["userId"]),
		AddCapacityValue: core.CastInt32(data["addCapacityValue"]),
	}
}

func (p AddCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"userId":           p.UserId,
		"addCapacityValue": p.AddCapacityValue,
	}
}

func (p AddCapacityByUserIdRequest) Pointer() *AddCapacityByUserIdRequest {
	return &p
}

type SetCapacityByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	NewCapacityValue   *int32  `json:"newCapacityValue"`
}

func NewSetCapacityByUserIdRequestFromJson(data string) SetCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetCapacityByUserIdRequestFromDict(dict)
}

func NewSetCapacityByUserIdRequestFromDict(data map[string]interface{}) SetCapacityByUserIdRequest {
	return SetCapacityByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		UserId:           core.CastString(data["userId"]),
		NewCapacityValue: core.CastInt32(data["newCapacityValue"]),
	}
}

func (p SetCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"userId":           p.UserId,
		"newCapacityValue": p.NewCapacityValue,
	}
}

func (p SetCapacityByUserIdRequest) Pointer() *SetCapacityByUserIdRequest {
	return &p
}

type DeleteInventoryByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
}

func NewDeleteInventoryByUserIdRequestFromJson(data string) DeleteInventoryByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInventoryByUserIdRequestFromDict(dict)
}

func NewDeleteInventoryByUserIdRequestFromDict(data map[string]interface{}) DeleteInventoryByUserIdRequest {
	return DeleteInventoryByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteInventoryByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
	}
}

func (p DeleteInventoryByUserIdRequest) Pointer() *DeleteInventoryByUserIdRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityRequest struct {
	RequestId                   *string `json:"requestId"`
	ContextStack                *string `json:"contextStack"`
	DuplicationAvoider          *string `json:"duplicationAvoider"`
	NamespaceName               *string `json:"namespaceName"`
	AccessToken                 *string `json:"accessToken"`
	InventoryName               *string `json:"inventoryName"`
	VerifyType                  *string `json:"verifyType"`
	CurrentInventoryMaxCapacity *int32  `json:"currentInventoryMaxCapacity"`
}

func NewVerifyInventoryCurrentMaxCapacityRequestFromJson(data string) VerifyInventoryCurrentMaxCapacityRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityRequestFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityRequest {
	return VerifyInventoryCurrentMaxCapacityRequest{
		NamespaceName:               core.CastString(data["namespaceName"]),
		AccessToken:                 core.CastString(data["accessToken"]),
		InventoryName:               core.CastString(data["inventoryName"]),
		VerifyType:                  core.CastString(data["verifyType"]),
		CurrentInventoryMaxCapacity: core.CastInt32(data["currentInventoryMaxCapacity"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":               p.NamespaceName,
		"accessToken":                 p.AccessToken,
		"inventoryName":               p.InventoryName,
		"verifyType":                  p.VerifyType,
		"currentInventoryMaxCapacity": p.CurrentInventoryMaxCapacity,
	}
}

func (p VerifyInventoryCurrentMaxCapacityRequest) Pointer() *VerifyInventoryCurrentMaxCapacityRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByUserIdRequest struct {
	RequestId                   *string `json:"requestId"`
	ContextStack                *string `json:"contextStack"`
	DuplicationAvoider          *string `json:"duplicationAvoider"`
	NamespaceName               *string `json:"namespaceName"`
	UserId                      *string `json:"userId"`
	InventoryName               *string `json:"inventoryName"`
	VerifyType                  *string `json:"verifyType"`
	CurrentInventoryMaxCapacity *int32  `json:"currentInventoryMaxCapacity"`
}

func NewVerifyInventoryCurrentMaxCapacityByUserIdRequestFromJson(data string) VerifyInventoryCurrentMaxCapacityByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityByUserIdRequestFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityByUserIdRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByUserIdRequest {
	return VerifyInventoryCurrentMaxCapacityByUserIdRequest{
		NamespaceName:               core.CastString(data["namespaceName"]),
		UserId:                      core.CastString(data["userId"]),
		InventoryName:               core.CastString(data["inventoryName"]),
		VerifyType:                  core.CastString(data["verifyType"]),
		CurrentInventoryMaxCapacity: core.CastInt32(data["currentInventoryMaxCapacity"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":               p.NamespaceName,
		"userId":                      p.UserId,
		"inventoryName":               p.InventoryName,
		"verifyType":                  p.VerifyType,
		"currentInventoryMaxCapacity": p.CurrentInventoryMaxCapacity,
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdRequest) Pointer() *VerifyInventoryCurrentMaxCapacityByUserIdRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskRequestFromJson(data string) VerifyInventoryCurrentMaxCapacityByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityByStampTaskRequestFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByStampTaskRequest {
	return VerifyInventoryCurrentMaxCapacityByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyInventoryCurrentMaxCapacityByStampTaskRequest) Pointer() *VerifyInventoryCurrentMaxCapacityByStampTaskRequest {
	return &p
}

type AddCapacityByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
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

type SetCapacityByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
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

type DescribeItemSetsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeItemSetsRequestFromJson(data string) DescribeItemSetsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemSetsRequestFromDict(dict)
}

func NewDescribeItemSetsRequestFromDict(data map[string]interface{}) DescribeItemSetsRequest {
	return DescribeItemSetsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeItemSetsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeItemSetsRequest) Pointer() *DescribeItemSetsRequest {
	return &p
}

type DescribeItemSetsByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeItemSetsByUserIdRequestFromJson(data string) DescribeItemSetsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemSetsByUserIdRequestFromDict(dict)
}

func NewDescribeItemSetsByUserIdRequestFromDict(data map[string]interface{}) DescribeItemSetsByUserIdRequest {
	return DescribeItemSetsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeItemSetsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeItemSetsByUserIdRequest) Pointer() *DescribeItemSetsByUserIdRequest {
	return &p
}

type GetItemSetRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
}

func NewGetItemSetRequestFromJson(data string) GetItemSetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemSetRequestFromDict(dict)
}

func NewGetItemSetRequestFromDict(data map[string]interface{}) GetItemSetRequest {
	return GetItemSetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p GetItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p GetItemSetRequest) Pointer() *GetItemSetRequest {
	return &p
}

type GetItemSetByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
}

func NewGetItemSetByUserIdRequestFromJson(data string) GetItemSetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemSetByUserIdRequestFromDict(dict)
}

func NewGetItemSetByUserIdRequestFromDict(data map[string]interface{}) GetItemSetByUserIdRequest {
	return GetItemSetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p GetItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p GetItemSetByUserIdRequest) Pointer() *GetItemSetByUserIdRequest {
	return &p
}

type GetItemWithSignatureRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
	KeyId         *string `json:"keyId"`
}

func NewGetItemWithSignatureRequestFromJson(data string) GetItemWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemWithSignatureRequestFromDict(dict)
}

func NewGetItemWithSignatureRequestFromDict(data map[string]interface{}) GetItemWithSignatureRequest {
	return GetItemWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetItemWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"keyId":         p.KeyId,
	}
}

func (p GetItemWithSignatureRequest) Pointer() *GetItemWithSignatureRequest {
	return &p
}

type GetItemWithSignatureByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
	KeyId         *string `json:"keyId"`
}

func NewGetItemWithSignatureByUserIdRequestFromJson(data string) GetItemWithSignatureByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetItemWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetItemWithSignatureByUserIdRequest {
	return GetItemWithSignatureByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetItemWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"keyId":         p.KeyId,
	}
}

func (p GetItemWithSignatureByUserIdRequest) Pointer() *GetItemWithSignatureByUserIdRequest {
	return &p
}

type AcquireItemSetByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	UserId             *string `json:"userId"`
	AcquireCount       *int64  `json:"acquireCount"`
	ExpiresAt          *int64  `json:"expiresAt"`
	CreateNewItemSet   *bool   `json:"createNewItemSet"`
	ItemSetName        *string `json:"itemSetName"`
}

func NewAcquireItemSetByUserIdRequestFromJson(data string) AcquireItemSetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetByUserIdRequestFromDict(dict)
}

func NewAcquireItemSetByUserIdRequestFromDict(data map[string]interface{}) AcquireItemSetByUserIdRequest {
	return AcquireItemSetByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		ItemName:         core.CastString(data["itemName"]),
		UserId:           core.CastString(data["userId"]),
		AcquireCount:     core.CastInt64(data["acquireCount"]),
		ExpiresAt:        core.CastInt64(data["expiresAt"]),
		CreateNewItemSet: core.CastBool(data["createNewItemSet"]),
		ItemSetName:      core.CastString(data["itemSetName"]),
	}
}

func (p AcquireItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"itemName":         p.ItemName,
		"userId":           p.UserId,
		"acquireCount":     p.AcquireCount,
		"expiresAt":        p.ExpiresAt,
		"createNewItemSet": p.CreateNewItemSet,
		"itemSetName":      p.ItemSetName,
	}
}

func (p AcquireItemSetByUserIdRequest) Pointer() *AcquireItemSetByUserIdRequest {
	return &p
}

type ConsumeItemSetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *int64  `json:"consumeCount"`
	ItemSetName        *string `json:"itemSetName"`
}

func NewConsumeItemSetRequestFromJson(data string) ConsumeItemSetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetRequestFromDict(dict)
}

func NewConsumeItemSetRequestFromDict(data map[string]interface{}) ConsumeItemSetRequest {
	return ConsumeItemSetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastInt64(data["consumeCount"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p ConsumeItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
		"itemSetName":   p.ItemSetName,
	}
}

func (p ConsumeItemSetRequest) Pointer() *ConsumeItemSetRequest {
	return &p
}

type ConsumeItemSetByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *int64  `json:"consumeCount"`
	ItemSetName        *string `json:"itemSetName"`
}

func NewConsumeItemSetByUserIdRequestFromJson(data string) ConsumeItemSetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetByUserIdRequestFromDict(dict)
}

func NewConsumeItemSetByUserIdRequestFromDict(data map[string]interface{}) ConsumeItemSetByUserIdRequest {
	return ConsumeItemSetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastInt64(data["consumeCount"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p ConsumeItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
		"itemSetName":   p.ItemSetName,
	}
}

func (p ConsumeItemSetByUserIdRequest) Pointer() *ConsumeItemSetByUserIdRequest {
	return &p
}

type DeleteItemSetByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
}

func NewDeleteItemSetByUserIdRequestFromJson(data string) DeleteItemSetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteItemSetByUserIdRequestFromDict(dict)
}

func NewDeleteItemSetByUserIdRequestFromDict(data map[string]interface{}) DeleteItemSetByUserIdRequest {
	return DeleteItemSetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p DeleteItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p DeleteItemSetByUserIdRequest) Pointer() *DeleteItemSetByUserIdRequest {
	return &p
}

type VerifyItemSetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	ItemSetName        *string `json:"itemSetName"`
	Count              *int64  `json:"count"`
}

func NewVerifyItemSetRequestFromJson(data string) VerifyItemSetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetRequestFromDict(dict)
}

func NewVerifyItemSetRequestFromDict(data map[string]interface{}) VerifyItemSetRequest {
	return VerifyItemSetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p VerifyItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"itemSetName":   p.ItemSetName,
		"count":         p.Count,
	}
}

func (p VerifyItemSetRequest) Pointer() *VerifyItemSetRequest {
	return &p
}

type VerifyItemSetByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	ItemSetName        *string `json:"itemSetName"`
	Count              *int64  `json:"count"`
}

func NewVerifyItemSetByUserIdRequestFromJson(data string) VerifyItemSetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetByUserIdRequestFromDict(dict)
}

func NewVerifyItemSetByUserIdRequestFromDict(data map[string]interface{}) VerifyItemSetByUserIdRequest {
	return VerifyItemSetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p VerifyItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"itemSetName":   p.ItemSetName,
		"count":         p.Count,
	}
}

func (p VerifyItemSetByUserIdRequest) Pointer() *VerifyItemSetByUserIdRequest {
	return &p
}

type AcquireItemSetByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquireItemSetByStampSheetRequestFromJson(data string) AcquireItemSetByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetByStampSheetRequestFromDict(dict)
}

func NewAcquireItemSetByStampSheetRequestFromDict(data map[string]interface{}) AcquireItemSetByStampSheetRequest {
	return AcquireItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireItemSetByStampSheetRequest) Pointer() *AcquireItemSetByStampSheetRequest {
	return &p
}

type ConsumeItemSetByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewConsumeItemSetByStampTaskRequestFromJson(data string) ConsumeItemSetByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetByStampTaskRequestFromDict(dict)
}

func NewConsumeItemSetByStampTaskRequestFromDict(data map[string]interface{}) ConsumeItemSetByStampTaskRequest {
	return ConsumeItemSetByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeItemSetByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeItemSetByStampTaskRequest) Pointer() *ConsumeItemSetByStampTaskRequest {
	return &p
}

type VerifyItemSetByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyItemSetByStampTaskRequestFromJson(data string) VerifyItemSetByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetByStampTaskRequestFromDict(dict)
}

func NewVerifyItemSetByStampTaskRequestFromDict(data map[string]interface{}) VerifyItemSetByStampTaskRequest {
	return VerifyItemSetByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyItemSetByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyItemSetByStampTaskRequest) Pointer() *VerifyItemSetByStampTaskRequest {
	return &p
}

type DescribeReferenceOfRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
}

func NewDescribeReferenceOfRequestFromJson(data string) DescribeReferenceOfRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReferenceOfRequestFromDict(dict)
}

func NewDescribeReferenceOfRequestFromDict(data map[string]interface{}) DescribeReferenceOfRequest {
	return DescribeReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p DescribeReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p DescribeReferenceOfRequest) Pointer() *DescribeReferenceOfRequest {
	return &p
}

type DescribeReferenceOfByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
}

func NewDescribeReferenceOfByUserIdRequestFromJson(data string) DescribeReferenceOfByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReferenceOfByUserIdRequestFromDict(dict)
}

func NewDescribeReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DescribeReferenceOfByUserIdRequest {
	return DescribeReferenceOfByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p DescribeReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p DescribeReferenceOfByUserIdRequest) Pointer() *DescribeReferenceOfByUserIdRequest {
	return &p
}

type GetReferenceOfRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
	ReferenceOf   *string `json:"referenceOf"`
}

func NewGetReferenceOfRequestFromJson(data string) GetReferenceOfRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReferenceOfRequestFromDict(dict)
}

func NewGetReferenceOfRequestFromDict(data map[string]interface{}) GetReferenceOfRequest {
	return GetReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p GetReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p GetReferenceOfRequest) Pointer() *GetReferenceOfRequest {
	return &p
}

type GetReferenceOfByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
	ItemSetName   *string `json:"itemSetName"`
	ReferenceOf   *string `json:"referenceOf"`
}

func NewGetReferenceOfByUserIdRequestFromJson(data string) GetReferenceOfByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReferenceOfByUserIdRequestFromDict(dict)
}

func NewGetReferenceOfByUserIdRequestFromDict(data map[string]interface{}) GetReferenceOfByUserIdRequest {
	return GetReferenceOfByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p GetReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p GetReferenceOfByUserIdRequest) Pointer() *GetReferenceOfByUserIdRequest {
	return &p
}

type VerifyReferenceOfRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	VerifyType         *string `json:"verifyType"`
}

func NewVerifyReferenceOfRequestFromJson(data string) VerifyReferenceOfRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfRequestFromDict(dict)
}

func NewVerifyReferenceOfRequestFromDict(data map[string]interface{}) VerifyReferenceOfRequest {
	return VerifyReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
		VerifyType:    core.CastString(data["verifyType"]),
	}
}

func (p VerifyReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
		"verifyType":    p.VerifyType,
	}
}

func (p VerifyReferenceOfRequest) Pointer() *VerifyReferenceOfRequest {
	return &p
}

type VerifyReferenceOfByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	VerifyType         *string `json:"verifyType"`
}

func NewVerifyReferenceOfByUserIdRequestFromJson(data string) VerifyReferenceOfByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfByUserIdRequestFromDict(dict)
}

func NewVerifyReferenceOfByUserIdRequestFromDict(data map[string]interface{}) VerifyReferenceOfByUserIdRequest {
	return VerifyReferenceOfByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
		VerifyType:    core.CastString(data["verifyType"]),
	}
}

func (p VerifyReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
		"verifyType":    p.VerifyType,
	}
}

func (p VerifyReferenceOfByUserIdRequest) Pointer() *VerifyReferenceOfByUserIdRequest {
	return &p
}

type AddReferenceOfRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func NewAddReferenceOfRequestFromJson(data string) AddReferenceOfRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfRequestFromDict(dict)
}

func NewAddReferenceOfRequestFromDict(data map[string]interface{}) AddReferenceOfRequest {
	return AddReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p AddReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p AddReferenceOfRequest) Pointer() *AddReferenceOfRequest {
	return &p
}

type AddReferenceOfByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func NewAddReferenceOfByUserIdRequestFromJson(data string) AddReferenceOfByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfByUserIdRequestFromDict(dict)
}

func NewAddReferenceOfByUserIdRequestFromDict(data map[string]interface{}) AddReferenceOfByUserIdRequest {
	return AddReferenceOfByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p AddReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p AddReferenceOfByUserIdRequest) Pointer() *AddReferenceOfByUserIdRequest {
	return &p
}

type DeleteReferenceOfRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func NewDeleteReferenceOfRequestFromJson(data string) DeleteReferenceOfRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfRequestFromDict(dict)
}

func NewDeleteReferenceOfRequestFromDict(data map[string]interface{}) DeleteReferenceOfRequest {
	return DeleteReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p DeleteReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p DeleteReferenceOfRequest) Pointer() *DeleteReferenceOfRequest {
	return &p
}

type DeleteReferenceOfByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func NewDeleteReferenceOfByUserIdRequestFromJson(data string) DeleteReferenceOfByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfByUserIdRequestFromDict(dict)
}

func NewDeleteReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DeleteReferenceOfByUserIdRequest {
	return DeleteReferenceOfByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p DeleteReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p DeleteReferenceOfByUserIdRequest) Pointer() *DeleteReferenceOfByUserIdRequest {
	return &p
}

type AddReferenceOfItemSetByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAddReferenceOfItemSetByStampSheetRequestFromJson(data string) AddReferenceOfItemSetByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfItemSetByStampSheetRequestFromDict(dict)
}

func NewAddReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) AddReferenceOfItemSetByStampSheetRequest {
	return AddReferenceOfItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddReferenceOfItemSetByStampSheetRequest) Pointer() *AddReferenceOfItemSetByStampSheetRequest {
	return &p
}

type DeleteReferenceOfItemSetByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromJson(data string) DeleteReferenceOfItemSetByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfItemSetByStampSheetRequestFromDict(dict)
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) DeleteReferenceOfItemSetByStampSheetRequest {
	return DeleteReferenceOfItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) Pointer() *DeleteReferenceOfItemSetByStampSheetRequest {
	return &p
}

type VerifyReferenceOfByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyReferenceOfByStampTaskRequestFromJson(data string) VerifyReferenceOfByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfByStampTaskRequestFromDict(dict)
}

func NewVerifyReferenceOfByStampTaskRequestFromDict(data map[string]interface{}) VerifyReferenceOfByStampTaskRequest {
	return VerifyReferenceOfByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyReferenceOfByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyReferenceOfByStampTaskRequest) Pointer() *VerifyReferenceOfByStampTaskRequest {
	return &p
}

type DescribeSimpleItemsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSimpleItemsRequestFromJson(data string) DescribeSimpleItemsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemsRequestFromDict(dict)
}

func NewDescribeSimpleItemsRequestFromDict(data map[string]interface{}) DescribeSimpleItemsRequest {
	return DescribeSimpleItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleItemsRequest) Pointer() *DescribeSimpleItemsRequest {
	return &p
}

type DescribeSimpleItemsByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSimpleItemsByUserIdRequestFromJson(data string) DescribeSimpleItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemsByUserIdRequestFromDict(dict)
}

func NewDescribeSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) DescribeSimpleItemsByUserIdRequest {
	return DescribeSimpleItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleItemsByUserIdRequest) Pointer() *DescribeSimpleItemsByUserIdRequest {
	return &p
}

type GetSimpleItemRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
}

func NewGetSimpleItemRequestFromJson(data string) GetSimpleItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemRequestFromDict(dict)
}

func NewGetSimpleItemRequestFromDict(data map[string]interface{}) GetSimpleItemRequest {
	return GetSimpleItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemRequest) Pointer() *GetSimpleItemRequest {
	return &p
}

type GetSimpleItemByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
}

func NewGetSimpleItemByUserIdRequestFromJson(data string) GetSimpleItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemByUserIdRequestFromDict(dict)
}

func NewGetSimpleItemByUserIdRequestFromDict(data map[string]interface{}) GetSimpleItemByUserIdRequest {
	return GetSimpleItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemByUserIdRequest) Pointer() *GetSimpleItemByUserIdRequest {
	return &p
}

type GetSimpleItemWithSignatureRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
	KeyId         *string `json:"keyId"`
}

func NewGetSimpleItemWithSignatureRequestFromJson(data string) GetSimpleItemWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemWithSignatureRequestFromDict(dict)
}

func NewGetSimpleItemWithSignatureRequestFromDict(data map[string]interface{}) GetSimpleItemWithSignatureRequest {
	return GetSimpleItemWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetSimpleItemWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"keyId":         p.KeyId,
	}
}

func (p GetSimpleItemWithSignatureRequest) Pointer() *GetSimpleItemWithSignatureRequest {
	return &p
}

type GetSimpleItemWithSignatureByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
	KeyId         *string `json:"keyId"`
}

func NewGetSimpleItemWithSignatureByUserIdRequestFromJson(data string) GetSimpleItemWithSignatureByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetSimpleItemWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetSimpleItemWithSignatureByUserIdRequest {
	return GetSimpleItemWithSignatureByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetSimpleItemWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"keyId":         p.KeyId,
	}
}

func (p GetSimpleItemWithSignatureByUserIdRequest) Pointer() *GetSimpleItemWithSignatureByUserIdRequest {
	return &p
}

type AcquireSimpleItemsByUserIdRequest struct {
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	UserId             *string        `json:"userId"`
	AcquireCounts      []AcquireCount `json:"acquireCounts"`
}

func NewAcquireSimpleItemsByUserIdRequestFromJson(data string) AcquireSimpleItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireSimpleItemsByUserIdRequestFromDict(dict)
}

func NewAcquireSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) AcquireSimpleItemsByUserIdRequest {
	return AcquireSimpleItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		AcquireCounts: CastAcquireCounts(core.CastArray(data["acquireCounts"])),
	}
}

func (p AcquireSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"acquireCounts": CastAcquireCountsFromDict(
			p.AcquireCounts,
		),
	}
}

func (p AcquireSimpleItemsByUserIdRequest) Pointer() *AcquireSimpleItemsByUserIdRequest {
	return &p
}

type ConsumeSimpleItemsRequest struct {
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	AccessToken        *string        `json:"accessToken"`
	ConsumeCounts      []ConsumeCount `json:"consumeCounts"`
}

func NewConsumeSimpleItemsRequestFromJson(data string) ConsumeSimpleItemsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsRequestFromDict(dict)
}

func NewConsumeSimpleItemsRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsRequest {
	return ConsumeSimpleItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ConsumeCounts: CastConsumeCounts(core.CastArray(data["consumeCounts"])),
	}
}

func (p ConsumeSimpleItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"consumeCounts": CastConsumeCountsFromDict(
			p.ConsumeCounts,
		),
	}
}

func (p ConsumeSimpleItemsRequest) Pointer() *ConsumeSimpleItemsRequest {
	return &p
}

type ConsumeSimpleItemsByUserIdRequest struct {
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	UserId             *string        `json:"userId"`
	ConsumeCounts      []ConsumeCount `json:"consumeCounts"`
}

func NewConsumeSimpleItemsByUserIdRequestFromJson(data string) ConsumeSimpleItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsByUserIdRequestFromDict(dict)
}

func NewConsumeSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsByUserIdRequest {
	return ConsumeSimpleItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ConsumeCounts: CastConsumeCounts(core.CastArray(data["consumeCounts"])),
	}
}

func (p ConsumeSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"consumeCounts": CastConsumeCountsFromDict(
			p.ConsumeCounts,
		),
	}
}

func (p ConsumeSimpleItemsByUserIdRequest) Pointer() *ConsumeSimpleItemsByUserIdRequest {
	return &p
}

type SetSimpleItemsByUserIdRequest struct {
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	InventoryName      *string     `json:"inventoryName"`
	UserId             *string     `json:"userId"`
	Counts             []HeldCount `json:"counts"`
}

func NewSetSimpleItemsByUserIdRequestFromJson(data string) SetSimpleItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetSimpleItemsByUserIdRequestFromDict(dict)
}

func NewSetSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) SetSimpleItemsByUserIdRequest {
	return SetSimpleItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		Counts:        CastHeldCounts(core.CastArray(data["counts"])),
	}
}

func (p SetSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"counts": CastHeldCountsFromDict(
			p.Counts,
		),
	}
}

func (p SetSimpleItemsByUserIdRequest) Pointer() *SetSimpleItemsByUserIdRequest {
	return &p
}

type DeleteSimpleItemsByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
}

func NewDeleteSimpleItemsByUserIdRequestFromJson(data string) DeleteSimpleItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleItemsByUserIdRequestFromDict(dict)
}

func NewDeleteSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) DeleteSimpleItemsByUserIdRequest {
	return DeleteSimpleItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
	}
}

func (p DeleteSimpleItemsByUserIdRequest) Pointer() *DeleteSimpleItemsByUserIdRequest {
	return &p
}

type VerifySimpleItemRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	Count              *int64  `json:"count"`
}

func NewVerifySimpleItemRequestFromJson(data string) VerifySimpleItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemRequestFromDict(dict)
}

func NewVerifySimpleItemRequestFromDict(data map[string]interface{}) VerifySimpleItemRequest {
	return VerifySimpleItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p VerifySimpleItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifySimpleItemRequest) Pointer() *VerifySimpleItemRequest {
	return &p
}

type VerifySimpleItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	Count              *int64  `json:"count"`
}

func NewVerifySimpleItemByUserIdRequestFromJson(data string) VerifySimpleItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemByUserIdRequestFromDict(dict)
}

func NewVerifySimpleItemByUserIdRequestFromDict(data map[string]interface{}) VerifySimpleItemByUserIdRequest {
	return VerifySimpleItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p VerifySimpleItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifySimpleItemByUserIdRequest) Pointer() *VerifySimpleItemByUserIdRequest {
	return &p
}

type AcquireSimpleItemsByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquireSimpleItemsByStampSheetRequestFromJson(data string) AcquireSimpleItemsByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireSimpleItemsByStampSheetRequestFromDict(dict)
}

func NewAcquireSimpleItemsByStampSheetRequestFromDict(data map[string]interface{}) AcquireSimpleItemsByStampSheetRequest {
	return AcquireSimpleItemsByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireSimpleItemsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireSimpleItemsByStampSheetRequest) Pointer() *AcquireSimpleItemsByStampSheetRequest {
	return &p
}

type ConsumeSimpleItemsByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewConsumeSimpleItemsByStampTaskRequestFromJson(data string) ConsumeSimpleItemsByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsByStampTaskRequestFromDict(dict)
}

func NewConsumeSimpleItemsByStampTaskRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsByStampTaskRequest {
	return ConsumeSimpleItemsByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeSimpleItemsByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeSimpleItemsByStampTaskRequest) Pointer() *ConsumeSimpleItemsByStampTaskRequest {
	return &p
}

type SetSimpleItemsByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetSimpleItemsByStampSheetRequestFromJson(data string) SetSimpleItemsByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetSimpleItemsByStampSheetRequestFromDict(dict)
}

func NewSetSimpleItemsByStampSheetRequestFromDict(data map[string]interface{}) SetSimpleItemsByStampSheetRequest {
	return SetSimpleItemsByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetSimpleItemsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetSimpleItemsByStampSheetRequest) Pointer() *SetSimpleItemsByStampSheetRequest {
	return &p
}

type VerifySimpleItemByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifySimpleItemByStampTaskRequestFromJson(data string) VerifySimpleItemByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemByStampTaskRequestFromDict(dict)
}

func NewVerifySimpleItemByStampTaskRequestFromDict(data map[string]interface{}) VerifySimpleItemByStampTaskRequest {
	return VerifySimpleItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifySimpleItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifySimpleItemByStampTaskRequest) Pointer() *VerifySimpleItemByStampTaskRequest {
	return &p
}

type DescribeBigItemsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBigItemsRequestFromJson(data string) DescribeBigItemsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemsRequestFromDict(dict)
}

func NewDescribeBigItemsRequestFromDict(data map[string]interface{}) DescribeBigItemsRequest {
	return DescribeBigItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigItemsRequest) Pointer() *DescribeBigItemsRequest {
	return &p
}

type DescribeBigItemsByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBigItemsByUserIdRequestFromJson(data string) DescribeBigItemsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemsByUserIdRequestFromDict(dict)
}

func NewDescribeBigItemsByUserIdRequestFromDict(data map[string]interface{}) DescribeBigItemsByUserIdRequest {
	return DescribeBigItemsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigItemsByUserIdRequest) Pointer() *DescribeBigItemsByUserIdRequest {
	return &p
}

type GetBigItemRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	AccessToken   *string `json:"accessToken"`
	ItemName      *string `json:"itemName"`
}

func NewGetBigItemRequestFromJson(data string) GetBigItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemRequestFromDict(dict)
}

func NewGetBigItemRequestFromDict(data map[string]interface{}) GetBigItemRequest {
	return GetBigItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemRequest) Pointer() *GetBigItemRequest {
	return &p
}

type GetBigItemByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	InventoryName *string `json:"inventoryName"`
	UserId        *string `json:"userId"`
	ItemName      *string `json:"itemName"`
}

func NewGetBigItemByUserIdRequestFromJson(data string) GetBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemByUserIdRequestFromDict(dict)
}

func NewGetBigItemByUserIdRequestFromDict(data map[string]interface{}) GetBigItemByUserIdRequest {
	return GetBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemByUserIdRequest) Pointer() *GetBigItemByUserIdRequest {
	return &p
}

type AcquireBigItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	AcquireCount       *string `json:"acquireCount"`
}

func NewAcquireBigItemByUserIdRequestFromJson(data string) AcquireBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireBigItemByUserIdRequestFromDict(dict)
}

func NewAcquireBigItemByUserIdRequestFromDict(data map[string]interface{}) AcquireBigItemByUserIdRequest {
	return AcquireBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		AcquireCount:  core.CastString(data["acquireCount"]),
	}
}

func (p AcquireBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"acquireCount":  p.AcquireCount,
	}
}

func (p AcquireBigItemByUserIdRequest) Pointer() *AcquireBigItemByUserIdRequest {
	return &p
}

type ConsumeBigItemRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *string `json:"consumeCount"`
}

func NewConsumeBigItemRequestFromJson(data string) ConsumeBigItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemRequestFromDict(dict)
}

func NewConsumeBigItemRequestFromDict(data map[string]interface{}) ConsumeBigItemRequest {
	return ConsumeBigItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastString(data["consumeCount"]),
	}
}

func (p ConsumeBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
	}
}

func (p ConsumeBigItemRequest) Pointer() *ConsumeBigItemRequest {
	return &p
}

type ConsumeBigItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *string `json:"consumeCount"`
}

func NewConsumeBigItemByUserIdRequestFromJson(data string) ConsumeBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemByUserIdRequestFromDict(dict)
}

func NewConsumeBigItemByUserIdRequestFromDict(data map[string]interface{}) ConsumeBigItemByUserIdRequest {
	return ConsumeBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastString(data["consumeCount"]),
	}
}

func (p ConsumeBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
	}
}

func (p ConsumeBigItemByUserIdRequest) Pointer() *ConsumeBigItemByUserIdRequest {
	return &p
}

type SetBigItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	Count              *string `json:"count"`
}

func NewSetBigItemByUserIdRequestFromJson(data string) SetBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBigItemByUserIdRequestFromDict(dict)
}

func NewSetBigItemByUserIdRequestFromDict(data map[string]interface{}) SetBigItemByUserIdRequest {
	return SetBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		Count:         core.CastString(data["count"]),
	}
}

func (p SetBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"count":         p.Count,
	}
}

func (p SetBigItemByUserIdRequest) Pointer() *SetBigItemByUserIdRequest {
	return &p
}

type DeleteBigItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
}

func NewDeleteBigItemByUserIdRequestFromJson(data string) DeleteBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigItemByUserIdRequestFromDict(dict)
}

func NewDeleteBigItemByUserIdRequestFromDict(data map[string]interface{}) DeleteBigItemByUserIdRequest {
	return DeleteBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
	}
}

func (p DeleteBigItemByUserIdRequest) Pointer() *DeleteBigItemByUserIdRequest {
	return &p
}

type VerifyBigItemRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	Count              *string `json:"count"`
}

func NewVerifyBigItemRequestFromJson(data string) VerifyBigItemRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemRequestFromDict(dict)
}

func NewVerifyBigItemRequestFromDict(data map[string]interface{}) VerifyBigItemRequest {
	return VerifyBigItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastString(data["count"]),
	}
}

func (p VerifyBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifyBigItemRequest) Pointer() *VerifyBigItemRequest {
	return &p
}

type VerifyBigItemByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	VerifyType         *string `json:"verifyType"`
	Count              *string `json:"count"`
}

func NewVerifyBigItemByUserIdRequestFromJson(data string) VerifyBigItemByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemByUserIdRequestFromDict(dict)
}

func NewVerifyBigItemByUserIdRequestFromDict(data map[string]interface{}) VerifyBigItemByUserIdRequest {
	return VerifyBigItemByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastString(data["count"]),
	}
}

func (p VerifyBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifyBigItemByUserIdRequest) Pointer() *VerifyBigItemByUserIdRequest {
	return &p
}

type AcquireBigItemByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquireBigItemByStampSheetRequestFromJson(data string) AcquireBigItemByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireBigItemByStampSheetRequestFromDict(dict)
}

func NewAcquireBigItemByStampSheetRequestFromDict(data map[string]interface{}) AcquireBigItemByStampSheetRequest {
	return AcquireBigItemByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireBigItemByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireBigItemByStampSheetRequest) Pointer() *AcquireBigItemByStampSheetRequest {
	return &p
}

type ConsumeBigItemByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewConsumeBigItemByStampTaskRequestFromJson(data string) ConsumeBigItemByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemByStampTaskRequestFromDict(dict)
}

func NewConsumeBigItemByStampTaskRequestFromDict(data map[string]interface{}) ConsumeBigItemByStampTaskRequest {
	return ConsumeBigItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeBigItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeBigItemByStampTaskRequest) Pointer() *ConsumeBigItemByStampTaskRequest {
	return &p
}

type SetBigItemByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewSetBigItemByStampSheetRequestFromJson(data string) SetBigItemByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBigItemByStampSheetRequestFromDict(dict)
}

func NewSetBigItemByStampSheetRequestFromDict(data map[string]interface{}) SetBigItemByStampSheetRequest {
	return SetBigItemByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetBigItemByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetBigItemByStampSheetRequest) Pointer() *SetBigItemByStampSheetRequest {
	return &p
}

type VerifyBigItemByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyBigItemByStampTaskRequestFromJson(data string) VerifyBigItemByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemByStampTaskRequestFromDict(dict)
}

func NewVerifyBigItemByStampTaskRequestFromDict(data map[string]interface{}) VerifyBigItemByStampTaskRequest {
	return VerifyBigItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyBigItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyBigItemByStampTaskRequest) Pointer() *VerifyBigItemByStampTaskRequest {
	return &p
}
