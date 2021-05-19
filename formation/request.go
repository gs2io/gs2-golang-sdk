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

package formation

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
	UpdateMoldScript *ScriptSetting	`json:"updateMoldScript"`
	UpdateFormScript *ScriptSetting	`json:"updateFormScript"`
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
	UpdateMoldScript *ScriptSetting	`json:"updateMoldScript"`
	UpdateFormScript *ScriptSetting	`json:"updateFormScript"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeFormModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateFormModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	Slots []SlotModel	`json:"slots"`
}

type GetFormModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	FormModelName *string	`json:"formModelName"`
}

type UpdateFormModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	FormModelName *string	`json:"formModelName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	Slots []SlotModel	`json:"slots"`
}

type DeleteFormModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	FormModelName *string	`json:"formModelName"`
}

type DescribeMoldModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetMoldModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
}

type DescribeMoldModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateMoldModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	FormModelName *string	`json:"formModelName"`
	InitialMaxCapacity *int32	`json:"initialMaxCapacity"`
	MaxCapacity *int32	`json:"maxCapacity"`
}

type GetMoldModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
}

type UpdateMoldModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	FormModelName *string	`json:"formModelName"`
	InitialMaxCapacity *int32	`json:"initialMaxCapacity"`
	MaxCapacity *int32	`json:"maxCapacity"`
}

type DeleteMoldModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetCurrentFormMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateCurrentFormMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Settings *string	`json:"settings"`
}

type UpdateCurrentFormMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeMoldsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeMoldsByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetMoldRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetMoldByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetMoldCapacityByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Capacity *int32	`json:"capacity"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddMoldCapacityByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Capacity *int32	`json:"capacity"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteMoldRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteMoldByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
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

type DescribeFormsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeFormsByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetFormRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetFormByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetFormWithSignatureRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetFormWithSignatureByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetFormByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	Slots []Slot	`json:"slots"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type SetFormWithSignatureRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	Slots []SlotWithSignature	`json:"slots"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type AcquireActionsToFormPropertiesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	AcquireAction *AcquireAction	`json:"acquireAction"`
	QueueNamespaceId *string	`json:"queueNamespaceId"`
	KeyId *string	`json:"keyId"`
	Config []AcquireActionConfig	`json:"config"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteFormRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteFormByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	MoldName *string	`json:"moldName"`
	Index *int32	`json:"index"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AcquireActionToFormPropertiesByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
