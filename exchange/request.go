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

package exchange

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
	RequestId            *core.RequestId    `json:"requestId"`
	ContextStack         *core.ContextStack `json:"contextStack"`
	Name                 *string            `json:"name"`
	Description          *string            `json:"description"`
	EnableAwaitExchange  *bool              `json:"enableAwaitExchange"`
	EnableDirectExchange *bool              `json:"enableDirectExchange"`
	QueueNamespaceId     *string            `json:"queueNamespaceId"`
	KeyId                *string            `json:"keyId"`
	LogSetting           *LogSetting        `json:"logSetting"`
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
	RequestId            *core.RequestId    `json:"requestId"`
	ContextStack         *core.ContextStack `json:"contextStack"`
	NamespaceName        *string            `json:"namespaceName"`
	Description          *string            `json:"description"`
	EnableAwaitExchange  *bool              `json:"enableAwaitExchange"`
	EnableDirectExchange *bool              `json:"enableDirectExchange"`
	QueueNamespaceId     *string            `json:"queueNamespaceId"`
	KeyId                *string            `json:"keyId"`
	LogSetting           *LogSetting        `json:"logSetting"`
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
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	Name               *string            `json:"name"`
	Description        *string            `json:"description"`
	Metadata           *string            `json:"metadata"`
	TimingType         *string            `json:"timingType"`
	LockTime           *int32             `json:"lockTime"`
	EnableSkip         *bool              `json:"enableSkip"`
	SkipConsumeActions *[]*ConsumeAction  `json:"skipConsumeActions"`
	AcquireActions     *[]*AcquireAction  `json:"acquireActions"`
	ConsumeActions     *[]*ConsumeAction  `json:"consumeActions"`
}

type GetRateModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RateName      *string            `json:"rateName"`
}

type UpdateRateModelMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	RateName           *string            `json:"rateName"`
	Description        *string            `json:"description"`
	Metadata           *string            `json:"metadata"`
	TimingType         *string            `json:"timingType"`
	LockTime           *int32             `json:"lockTime"`
	EnableSkip         *bool              `json:"enableSkip"`
	SkipConsumeActions *[]*ConsumeAction  `json:"skipConsumeActions"`
	AcquireActions     *[]*AcquireAction  `json:"acquireActions"`
	ConsumeActions     *[]*ConsumeAction  `json:"consumeActions"`
}

type DeleteRateModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RateName      *string            `json:"rateName"`
}

type ExchangeRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	Count                  *int32             `json:"count"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type ExchangeByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	UserId                 *string            `json:"userId"`
	Count                  *int32             `json:"count"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ExchangeByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
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

type CreateAwaitByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	Count                  *int32             `json:"count"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeAwaitsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeAwaitsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetAwaitRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetAwaitByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AcquireRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type AcquireByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AcquireForceByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SkipRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type SkipByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteAwaitRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteAwaitByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RateName               *string            `json:"rateName"`
	AwaitName              *string            `json:"awaitName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateAwaitByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteAwaitByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
