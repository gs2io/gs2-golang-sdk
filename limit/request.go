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
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	Name               *string     `json:"name"`
	Description        *string     `json:"description"`
	LogSetting         *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
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
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	Description        *string     `json:"description"`
	LogSetting         *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Description:   core.CastString(data["description"]),
		LogSetting:    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"logSetting":    p.LogSetting.ToDict(),
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

type DescribeCountersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	LimitName          *string `json:"limitName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCountersRequestFromDict(data map[string]interface{}) DescribeCountersRequest {
	return DescribeCountersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		LimitName:     core.CastString(data["limitName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCountersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"limitName":     p.LimitName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCountersRequest) Pointer() *DescribeCountersRequest {
	return &p
}

type DescribeCountersByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	LimitName          *string `json:"limitName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCountersByUserIdRequestFromDict(data map[string]interface{}) DescribeCountersByUserIdRequest {
	return DescribeCountersByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		LimitName:     core.CastString(data["limitName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCountersByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"limitName":     p.LimitName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCountersByUserIdRequest) Pointer() *DescribeCountersByUserIdRequest {
	return &p
}

type GetCounterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	AccessToken        *string `json:"accessToken"`
	CounterName        *string `json:"counterName"`
}

func NewGetCounterRequestFromDict(data map[string]interface{}) GetCounterRequest {
	return GetCounterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p GetCounterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
		"accessToken":   p.AccessToken,
		"counterName":   p.CounterName,
	}
}

func (p GetCounterRequest) Pointer() *GetCounterRequest {
	return &p
}

type GetCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	UserId             *string `json:"userId"`
	CounterName        *string `json:"counterName"`
}

func NewGetCounterByUserIdRequestFromDict(data map[string]interface{}) GetCounterByUserIdRequest {
	return GetCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
		UserId:        core.CastString(data["userId"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p GetCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
		"userId":        p.UserId,
		"counterName":   p.CounterName,
	}
}

func (p GetCounterByUserIdRequest) Pointer() *GetCounterByUserIdRequest {
	return &p
}

type CountUpRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	CounterName        *string `json:"counterName"`
	AccessToken        *string `json:"accessToken"`
	CountUpValue       *int32  `json:"countUpValue"`
	MaxValue           *int32  `json:"maxValue"`
}

func NewCountUpRequestFromDict(data map[string]interface{}) CountUpRequest {
	return CountUpRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
		CounterName:   core.CastString(data["counterName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		CountUpValue:  core.CastInt32(data["countUpValue"]),
		MaxValue:      core.CastInt32(data["maxValue"]),
	}
}

func (p CountUpRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
		"counterName":   p.CounterName,
		"accessToken":   p.AccessToken,
		"countUpValue":  p.CountUpValue,
		"maxValue":      p.MaxValue,
	}
}

func (p CountUpRequest) Pointer() *CountUpRequest {
	return &p
}

type CountUpByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
	CountUpValue       *int32  `json:"countUpValue"`
	MaxValue           *int32  `json:"maxValue"`
}

func NewCountUpByUserIdRequestFromDict(data map[string]interface{}) CountUpByUserIdRequest {
	return CountUpByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
		CounterName:   core.CastString(data["counterName"]),
		UserId:        core.CastString(data["userId"]),
		CountUpValue:  core.CastInt32(data["countUpValue"]),
		MaxValue:      core.CastInt32(data["maxValue"]),
	}
}

func (p CountUpByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
		"counterName":   p.CounterName,
		"userId":        p.UserId,
		"countUpValue":  p.CountUpValue,
		"maxValue":      p.MaxValue,
	}
}

func (p CountUpByUserIdRequest) Pointer() *CountUpByUserIdRequest {
	return &p
}

type DeleteCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	UserId             *string `json:"userId"`
	CounterName        *string `json:"counterName"`
}

func NewDeleteCounterByUserIdRequestFromDict(data map[string]interface{}) DeleteCounterByUserIdRequest {
	return DeleteCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
		UserId:        core.CastString(data["userId"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p DeleteCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
		"userId":        p.UserId,
		"counterName":   p.CounterName,
	}
}

func (p DeleteCounterByUserIdRequest) Pointer() *DeleteCounterByUserIdRequest {
	return &p
}

type CountUpByStampTaskRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampTask          *string `json:"stampTask"`
	KeyId              *string `json:"keyId"`
}

func NewCountUpByStampTaskRequestFromDict(data map[string]interface{}) CountUpByStampTaskRequest {
	return CountUpByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p CountUpByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p CountUpByStampTaskRequest) Pointer() *CountUpByStampTaskRequest {
	return &p
}

type DeleteByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewDeleteByStampSheetRequestFromDict(data map[string]interface{}) DeleteByStampSheetRequest {
	return DeleteByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DeleteByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DeleteByStampSheetRequest) Pointer() *DeleteByStampSheetRequest {
	return &p
}

type DescribeLimitModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeLimitModelMastersRequestFromDict(data map[string]interface{}) DescribeLimitModelMastersRequest {
	return DescribeLimitModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeLimitModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeLimitModelMastersRequest) Pointer() *DescribeLimitModelMastersRequest {
	return &p
}

type CreateLimitModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	ResetType          *string `json:"resetType"`
	ResetDayOfMonth    *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek     *string `json:"resetDayOfWeek"`
	ResetHour          *int32  `json:"resetHour"`
}

func NewCreateLimitModelMasterRequestFromDict(data map[string]interface{}) CreateLimitModelMasterRequest {
	return CreateLimitModelMasterRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Name:            core.CastString(data["name"]),
		Description:     core.CastString(data["description"]),
		Metadata:        core.CastString(data["metadata"]),
		ResetType:       core.CastString(data["resetType"]),
		ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:  core.CastString(data["resetDayOfWeek"]),
		ResetHour:       core.CastInt32(data["resetHour"]),
	}
}

func (p CreateLimitModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"name":            p.Name,
		"description":     p.Description,
		"metadata":        p.Metadata,
		"resetType":       p.ResetType,
		"resetDayOfMonth": p.ResetDayOfMonth,
		"resetDayOfWeek":  p.ResetDayOfWeek,
		"resetHour":       p.ResetHour,
	}
}

func (p CreateLimitModelMasterRequest) Pointer() *CreateLimitModelMasterRequest {
	return &p
}

type GetLimitModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
}

func NewGetLimitModelMasterRequestFromDict(data map[string]interface{}) GetLimitModelMasterRequest {
	return GetLimitModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
	}
}

func (p GetLimitModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
	}
}

func (p GetLimitModelMasterRequest) Pointer() *GetLimitModelMasterRequest {
	return &p
}

type UpdateLimitModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	ResetType          *string `json:"resetType"`
	ResetDayOfMonth    *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek     *string `json:"resetDayOfWeek"`
	ResetHour          *int32  `json:"resetHour"`
}

func NewUpdateLimitModelMasterRequestFromDict(data map[string]interface{}) UpdateLimitModelMasterRequest {
	return UpdateLimitModelMasterRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		LimitName:       core.CastString(data["limitName"]),
		Description:     core.CastString(data["description"]),
		Metadata:        core.CastString(data["metadata"]),
		ResetType:       core.CastString(data["resetType"]),
		ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:  core.CastString(data["resetDayOfWeek"]),
		ResetHour:       core.CastInt32(data["resetHour"]),
	}
}

func (p UpdateLimitModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"limitName":       p.LimitName,
		"description":     p.Description,
		"metadata":        p.Metadata,
		"resetType":       p.ResetType,
		"resetDayOfMonth": p.ResetDayOfMonth,
		"resetDayOfWeek":  p.ResetDayOfWeek,
		"resetHour":       p.ResetHour,
	}
}

func (p UpdateLimitModelMasterRequest) Pointer() *UpdateLimitModelMasterRequest {
	return &p
}

type DeleteLimitModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
}

func NewDeleteLimitModelMasterRequestFromDict(data map[string]interface{}) DeleteLimitModelMasterRequest {
	return DeleteLimitModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
	}
}

func (p DeleteLimitModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
	}
}

func (p DeleteLimitModelMasterRequest) Pointer() *DeleteLimitModelMasterRequest {
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

type GetCurrentLimitMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetCurrentLimitMasterRequestFromDict(data map[string]interface{}) GetCurrentLimitMasterRequest {
	return GetCurrentLimitMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentLimitMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentLimitMasterRequest) Pointer() *GetCurrentLimitMasterRequest {
	return &p
}

type UpdateCurrentLimitMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
}

func NewUpdateCurrentLimitMasterRequestFromDict(data map[string]interface{}) UpdateCurrentLimitMasterRequest {
	return UpdateCurrentLimitMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentLimitMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentLimitMasterRequest) Pointer() *UpdateCurrentLimitMasterRequest {
	return &p
}

type UpdateCurrentLimitMasterFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentLimitMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentLimitMasterFromGitHubRequest {
	return UpdateCurrentLimitMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentLimitMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentLimitMasterFromGitHubRequest) Pointer() *UpdateCurrentLimitMasterFromGitHubRequest {
	return &p
}

type DescribeLimitModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeLimitModelsRequestFromDict(data map[string]interface{}) DescribeLimitModelsRequest {
	return DescribeLimitModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeLimitModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeLimitModelsRequest) Pointer() *DescribeLimitModelsRequest {
	return &p
}

type GetLimitModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
}

func NewGetLimitModelRequestFromDict(data map[string]interface{}) GetLimitModelRequest {
	return GetLimitModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
	}
}

func (p GetLimitModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
	}
}

func (p GetLimitModelRequest) Pointer() *GetLimitModelRequest {
	return &p
}
