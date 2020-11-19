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

package enhance

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
	RequestId           *core.RequestId    `json:"requestId"`
	ContextStack        *core.ContextStack `json:"contextStack"`
	Name                *string            `json:"name"`
	Description         *string            `json:"description"`
	EnableDirectEnhance *bool              `json:"enableDirectEnhance"`
	QueueNamespaceId    *string            `json:"queueNamespaceId"`
	KeyId               *string            `json:"keyId"`
	LogSetting          *LogSetting        `json:"logSetting"`
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
	RequestId           *core.RequestId    `json:"requestId"`
	ContextStack        *core.ContextStack `json:"contextStack"`
	NamespaceName       *string            `json:"namespaceName"`
	Description         *string            `json:"description"`
	EnableDirectEnhance *bool              `json:"enableDirectEnhance"`
	QueueNamespaceId    *string            `json:"queueNamespaceId"`
	KeyId               *string            `json:"keyId"`
	LogSetting          *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeRateModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetRateModelRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RateName      *string            `json:"rateName"`
}

type DescribeRateModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateRateModelMasterRequest struct {
	RequestId                  *core.RequestId    `json:"requestId"`
	ContextStack               *core.ContextStack `json:"contextStack"`
	NamespaceName              *string            `json:"namespaceName"`
	Name                       *string            `json:"name"`
	Description                *string            `json:"description"`
	Metadata                   *string            `json:"metadata"`
	TargetInventoryModelId     *string            `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string            `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string            `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy *[]string          `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string            `json:"experienceModelId"`
	BonusRates                 *[]*BonusRate      `json:"bonusRates"`
}

type GetRateModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RateName      *string            `json:"rateName"`
}

type UpdateRateModelMasterRequest struct {
	RequestId                  *core.RequestId    `json:"requestId"`
	ContextStack               *core.ContextStack `json:"contextStack"`
	NamespaceName              *string            `json:"namespaceName"`
	RateName                   *string            `json:"rateName"`
	Description                *string            `json:"description"`
	Metadata                   *string            `json:"metadata"`
	TargetInventoryModelId     *string            `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string            `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string            `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy *[]string          `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string            `json:"experienceModelId"`
	BonusRates                 *[]*BonusRate      `json:"bonusRates"`
}

type DeleteRateModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RateName      *string            `json:"rateName"`
}

type DirectEnhanceRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	TargetItemSetId        *string            `json:"targetItemSetId"`
	Materials              *[]*Material       `json:"materials"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DirectEnhanceByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	UserId                 *string            `json:"userId"`
	TargetItemSetId        *string            `json:"targetItemSetId"`
	Materials              *[]*Material       `json:"materials"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DirectEnhanceByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeProgressesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	TargetItemSetId        *string            `json:"targetItemSetId"`
	Materials              *[]*Material       `json:"materials"`
	Force                  *bool              `json:"force"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetProgressRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type StartRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	TargetItemSetId        *string            `json:"targetItemSetId"`
	Materials              *[]*Material       `json:"materials"`
	Force                  *bool              `json:"force"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type StartByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	TargetItemSetId        *string            `json:"targetItemSetId"`
	Materials              *[]*Material       `json:"materials"`
	UserId                 *string            `json:"userId"`
	Force                  *bool              `json:"force"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type EndRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type EndByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteProgressRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateProgressByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteProgressByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentRateMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentRateMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentRateMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}
