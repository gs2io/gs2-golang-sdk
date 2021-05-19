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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	AcquireScript *ScriptSetting	`json:"acquireScript"`
	OverflowScript *ScriptSetting	`json:"overflowScript"`
	ConsumeScript *ScriptSetting	`json:"consumeScript"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Description *string	`json:"description"`
	AcquireScript *ScriptSetting	`json:"acquireScript"`
	OverflowScript *ScriptSetting	`json:"overflowScript"`
	ConsumeScript *ScriptSetting	`json:"consumeScript"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeInventoryModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateInventoryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	InitialCapacity *int32	`json:"initialCapacity"`
	MaxCapacity *int32	`json:"maxCapacity"`
	ProtectReferencedItem *bool	`json:"protectReferencedItem"`
}

type GetInventoryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
}

type UpdateInventoryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	InitialCapacity *int32	`json:"initialCapacity"`
	MaxCapacity *int32	`json:"maxCapacity"`
	ProtectReferencedItem *bool	`json:"protectReferencedItem"`
}

type DeleteInventoryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
}

type DescribeInventoryModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetInventoryModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
}

type DescribeItemModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	StackingLimit *int64	`json:"stackingLimit"`
	AllowMultipleStacks *bool	`json:"allowMultipleStacks"`
	SortValue *int32	`json:"sortValue"`
}

type GetItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
}

type UpdateItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	StackingLimit *int64	`json:"stackingLimit"`
	AllowMultipleStacks *bool	`json:"allowMultipleStacks"`
	SortValue *int32	`json:"sortValue"`
}

type DeleteItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
}

type DescribeItemModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
}

type GetItemModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetCurrentItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateCurrentItemModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Settings *string	`json:"settings"`
}

type UpdateCurrentItemModelMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeInventoriesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeInventoriesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetInventoryRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetInventoryByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddCapacityByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	AddCapacityValue *int32	`json:"addCapacityValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetCapacityByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	NewCapacityValue *int32	`json:"newCapacityValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteInventoryByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddCapacityByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetCapacityByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DescribeItemSetsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeItemSetsByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetItemSetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetItemSetByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetItemWithSignatureRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetItemWithSignatureByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AcquireItemSetByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	UserId *string	`json:"userId"`
	AcquireCount *int64	`json:"acquireCount"`
	ExpiresAt *int64	`json:"expiresAt"`
	CreateNewItemSet *bool	`json:"createNewItemSet"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ConsumeItemSetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ConsumeCount *int64	`json:"consumeCount"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type ConsumeItemSetByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ConsumeCount *int64	`json:"consumeCount"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DescribeReferenceOfRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeReferenceOfByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetReferenceOfRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetReferenceOfByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type VerifyReferenceOfRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	VerifyType *string	`json:"verifyType"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type VerifyReferenceOfByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	VerifyType *string	`json:"verifyType"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddReferenceOfRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type AddReferenceOfByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteReferenceOfRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteReferenceOfByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	ReferenceOf *string	`json:"referenceOf"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteItemSetByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	InventoryName *string	`json:"inventoryName"`
	UserId *string	`json:"userId"`
	ItemName *string	`json:"itemName"`
	ItemSetName *string	`json:"itemSetName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AcquireItemSetByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddReferenceOfItemSetByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteReferenceOfItemSetByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ConsumeItemSetByStampTaskRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampTask *string	`json:"stampTask"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type VerifyReferenceOfByStampTaskRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampTask *string	`json:"stampTask"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
