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

package limit

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
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeCountersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeCountersByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetCounterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	CounterName *string	`json:"counterName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetCounterByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	UserId *string	`json:"userId"`
	CounterName *string	`json:"counterName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountUpRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	CounterName *string	`json:"counterName"`
	CountUpValue *int32	`json:"countUpValue"`
	MaxValue *int32	`json:"maxValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CountUpByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	CounterName *string	`json:"counterName"`
	UserId *string	`json:"userId"`
	CountUpValue *int32	`json:"countUpValue"`
	MaxValue *int32	`json:"maxValue"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteCounterByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	UserId *string	`json:"userId"`
	CounterName *string	`json:"counterName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountUpByStampTaskRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampTask *string	`json:"stampTask"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DeleteByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DescribeLimitModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateLimitModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ResetType *string	`json:"resetType"`
	ResetDayOfMonth *int32	`json:"resetDayOfMonth"`
	ResetDayOfWeek *string	`json:"resetDayOfWeek"`
	ResetHour *int32	`json:"resetHour"`
}

type GetLimitModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
}

type UpdateLimitModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
	ResetType *string	`json:"resetType"`
	ResetDayOfMonth *int32	`json:"resetDayOfMonth"`
	ResetDayOfWeek *string	`json:"resetDayOfWeek"`
	ResetHour *int32	`json:"resetHour"`
}

type DeleteLimitModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetCurrentLimitMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateCurrentLimitMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Settings *string	`json:"settings"`
}

type UpdateCurrentLimitMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeLimitModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetLimitModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	LimitName *string	`json:"limitName"`
}
