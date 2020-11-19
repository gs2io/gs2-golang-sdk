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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	Name                   *string            `json:"name"`
	Description            *string            `json:"description"`
	ExperienceCapScriptId  *string            `json:"experienceCapScriptId"`
	ChangeExperienceScript *ScriptSetting     `json:"changeExperienceScript"`
	ChangeRankScript       *ScriptSetting     `json:"changeRankScript"`
	ChangeRankCapScript    *ScriptSetting     `json:"changeRankCapScript"`
	LogSetting             *LogSetting        `json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Description            *string            `json:"description"`
	ExperienceCapScriptId  *string            `json:"experienceCapScriptId"`
	ChangeExperienceScript *ScriptSetting     `json:"changeExperienceScript"`
	ChangeRankScript       *ScriptSetting     `json:"changeRankScript"`
	ChangeRankCapScript    *ScriptSetting     `json:"changeRankCapScript"`
	LogSetting             *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeExperienceModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateExperienceModelMasterRequest struct {
	RequestId         *core.RequestId    `json:"requestId"`
	ContextStack      *core.ContextStack `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	Name              *string            `json:"name"`
	Description       *string            `json:"description"`
	Metadata          *string            `json:"metadata"`
	DefaultExperience *int64             `json:"defaultExperience"`
	DefaultRankCap    *int64             `json:"defaultRankCap"`
	MaxRankCap        *int64             `json:"maxRankCap"`
	RankThresholdId   *string            `json:"rankThresholdId"`
}

type GetExperienceModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	ExperienceName *string            `json:"experienceName"`
}

type UpdateExperienceModelMasterRequest struct {
	RequestId         *core.RequestId    `json:"requestId"`
	ContextStack      *core.ContextStack `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	ExperienceName    *string            `json:"experienceName"`
	Description       *string            `json:"description"`
	Metadata          *string            `json:"metadata"`
	DefaultExperience *int64             `json:"defaultExperience"`
	DefaultRankCap    *int64             `json:"defaultRankCap"`
	MaxRankCap        *int64             `json:"maxRankCap"`
	RankThresholdId   *string            `json:"rankThresholdId"`
}

type DeleteExperienceModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	ExperienceName *string            `json:"experienceName"`
}

type DescribeExperienceModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetExperienceModelRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	ExperienceName *string            `json:"experienceName"`
}

type DescribeThresholdMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateThresholdMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Name          *string            `json:"name"`
	Description   *string            `json:"description"`
	Metadata      *string            `json:"metadata"`
	Values        *[]int64           `json:"values"`
}

type GetThresholdMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	ThresholdName *string            `json:"thresholdName"`
}

type UpdateThresholdMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	ThresholdName *string            `json:"thresholdName"`
	Description   *string            `json:"description"`
	Metadata      *string            `json:"metadata"`
	Values        *[]int64           `json:"values"`
}

type DeleteThresholdMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	ThresholdName *string            `json:"thresholdName"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentExperienceMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentExperienceMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentExperienceMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeStatusesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ExperienceName         *string            `json:"experienceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeStatusesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ExperienceName         *string            `json:"experienceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetStatusRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetStatusByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetStatusWithSignatureRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type AddExperienceByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	ExperienceValue        *int64             `json:"experienceValue"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SetExperienceByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	ExperienceValue        *int64             `json:"experienceValue"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AddRankCapByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	RankCapValue           *int64             `json:"rankCapValue"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SetRankCapByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	RankCapValue           *int64             `json:"rankCapValue"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteStatusByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	ExperienceName         *string            `json:"experienceName"`
	PropertyId             *string            `json:"propertyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AddExperienceByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AddRankCapByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SetRankCapByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
