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

package stamina

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
	OverflowTriggerScriptId *string	`json:"overflowTriggerScriptId"`
	OverflowTriggerNamespaceId *string	`json:"overflowTriggerNamespaceId"`
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
	OverflowTriggerScriptId *string	`json:"overflowTriggerScriptId"`
	OverflowTriggerNamespaceId *string	`json:"overflowTriggerNamespaceId"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeStaminaModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateStaminaModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	RecoverIntervalMinutes *int32	`json:"recoverIntervalMinutes"`
	RecoverValue *int32	`json:"recoverValue"`
	InitialCapacity *int32	`json:"initialCapacity"`
	IsOverflow *bool	`json:"isOverflow"`
	MaxCapacity *int32	`json:"maxCapacity"`
	MaxStaminaTableName *string	`json:"maxStaminaTableName"`
	RecoverIntervalTableName *string	`json:"recoverIntervalTableName"`
	RecoverValueTableName *string	`json:"recoverValueTableName"`
}

type GetStaminaModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
}

type UpdateStaminaModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	RecoverIntervalMinutes *int32	`json:"recoverIntervalMinutes"`
	RecoverValue *int32	`json:"recoverValue"`
	InitialCapacity *int32	`json:"initialCapacity"`
	IsOverflow *bool	`json:"isOverflow"`
	MaxCapacity *int32	`json:"maxCapacity"`
	MaxStaminaTableName *string	`json:"maxStaminaTableName"`
	RecoverIntervalTableName *string	`json:"recoverIntervalTableName"`
	RecoverValueTableName *string	`json:"recoverValueTableName"`
}

type DeleteStaminaModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
}

type DescribeMaxStaminaTableMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateMaxStaminaTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type GetMaxStaminaTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MaxStaminaTableName *string	`json:"maxStaminaTableName"`
}

type UpdateMaxStaminaTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MaxStaminaTableName *string	`json:"maxStaminaTableName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type DeleteMaxStaminaTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MaxStaminaTableName *string	`json:"maxStaminaTableName"`
}

type DescribeRecoverIntervalTableMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateRecoverIntervalTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type GetRecoverIntervalTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverIntervalTableName *string	`json:"recoverIntervalTableName"`
}

type UpdateRecoverIntervalTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverIntervalTableName *string	`json:"recoverIntervalTableName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type DeleteRecoverIntervalTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverIntervalTableName *string	`json:"recoverIntervalTableName"`
}

type DescribeRecoverValueTableMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateRecoverValueTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type GetRecoverValueTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverValueTableName *string	`json:"recoverValueTableName"`
}

type UpdateRecoverValueTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverValueTableName *string	`json:"recoverValueTableName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ExperienceModelId *string	`json:"experienceModelId"`
	Values []int32	`json:"values"`
}

type DeleteRecoverValueTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RecoverValueTableName *string	`json:"recoverValueTableName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetCurrentStaminaMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateCurrentStaminaMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Settings *string	`json:"settings"`
}

type UpdateCurrentStaminaMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeStaminaModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetStaminaModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
}

type DescribeStaminasRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeStaminasByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetStaminaRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetStaminaByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type UpdateStaminaByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	Value *int32	`json:"value"`
	MaxValue *int32	`json:"maxValue"`
	RecoverIntervalMinutes *int32	`json:"recoverIntervalMinutes"`
	RecoverValue *int32	`json:"recoverValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ConsumeStaminaRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	ConsumeValue *int32	`json:"consumeValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type ConsumeStaminaByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	ConsumeValue *int32	`json:"consumeValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type RecoverStaminaByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	RecoverValue *int32	`json:"recoverValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type RaiseMaxValueByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	RaiseValue *int32	`json:"raiseValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetMaxValueByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	MaxValue *int32	`json:"maxValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetRecoverIntervalByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	RecoverIntervalMinutes *int32	`json:"recoverIntervalMinutes"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetRecoverValueByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	RecoverValue *int32	`json:"recoverValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetMaxValueByStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	KeyId *string	`json:"keyId"`
	SignedStatusBody *string	`json:"signedStatusBody"`
	SignedStatusSignature *string	`json:"signedStatusSignature"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type SetRecoverIntervalByStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	KeyId *string	`json:"keyId"`
	SignedStatusBody *string	`json:"signedStatusBody"`
	SignedStatusSignature *string	`json:"signedStatusSignature"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type SetRecoverValueByStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	KeyId *string	`json:"keyId"`
	SignedStatusBody *string	`json:"signedStatusBody"`
	SignedStatusSignature *string	`json:"signedStatusSignature"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteStaminaByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	StaminaName *string	`json:"staminaName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type RecoverStaminaByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type RaiseMaxValueByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetMaxValueByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetRecoverIntervalByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetRecoverValueByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ConsumeStaminaByStampTaskRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampTask *string	`json:"stampTask"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
