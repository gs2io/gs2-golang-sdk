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

import "github.com/gs2io/gs2-golang-sdk/core"

type DescribeNamespacesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	Name               *string        `json:"name"`
	Description        *string        `json:"description"`
	AcquireScript      *ScriptSetting `json:"acquireScript"`
	OverflowScript     *ScriptSetting `json:"overflowScript"`
	ConsumeScript      *ScriptSetting `json:"consumeScript"`
	LogSetting         *LogSetting    `json:"logSetting"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	Description        *string        `json:"description"`
	AcquireScript      *ScriptSetting `json:"acquireScript"`
	OverflowScript     *ScriptSetting `json:"overflowScript"`
	ConsumeScript      *ScriptSetting `json:"consumeScript"`
	LogSetting         *LogSetting    `json:"logSetting"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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

type DescribeInventoryModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	Name                  *string `json:"name"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
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
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	InventoryName         *string `json:"inventoryName"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
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
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	ItemName            *string `json:"itemName"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
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

type ExportMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
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
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
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

type AddCapacityByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	KeyId              *string `json:"keyId"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	KeyId              *string `json:"keyId"`
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

type DescribeReferenceOfRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
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

type AcquireItemSetByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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

type AddReferenceOfItemSetByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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

type ConsumeItemSetByStampTaskRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampTask          *string `json:"stampTask"`
	KeyId              *string `json:"keyId"`
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

type VerifyReferenceOfByStampTaskRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampTask          *string `json:"stampTask"`
	KeyId              *string `json:"keyId"`
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
