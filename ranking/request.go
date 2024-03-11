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

package ranking

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
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
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	Name            *string     `json:"name"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeCategoryModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeCategoryModelsRequestFromJson(data string) DescribeCategoryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelsRequestFromDict(dict)
}

func NewDescribeCategoryModelsRequestFromDict(data map[string]interface{}) DescribeCategoryModelsRequest {
	return DescribeCategoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeCategoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeCategoryModelsRequest) Pointer() *DescribeCategoryModelsRequest {
	return &p
}

type GetCategoryModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
}

func NewGetCategoryModelRequestFromJson(data string) GetCategoryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelRequestFromDict(dict)
}

func NewGetCategoryModelRequestFromDict(data map[string]interface{}) GetCategoryModelRequest {
	return GetCategoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetCategoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p GetCategoryModelRequest) Pointer() *GetCategoryModelRequest {
	return &p
}

type DescribeCategoryModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCategoryModelMastersRequestFromJson(data string) DescribeCategoryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelMastersRequestFromDict(dict)
}

func NewDescribeCategoryModelMastersRequestFromDict(data map[string]interface{}) DescribeCategoryModelMastersRequest {
	return DescribeCategoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCategoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCategoryModelMastersRequest) Pointer() *DescribeCategoryModelMastersRequest {
	return &p
}

type CreateCategoryModelMasterRequest struct {
	SourceRequestId            *string   `json:"sourceRequestId"`
	RequestId                  *string   `json:"requestId"`
	ContextStack               *string   `json:"contextStack"`
	NamespaceName              *string   `json:"namespaceName"`
	Name                       *string   `json:"name"`
	Description                *string   `json:"description"`
	Metadata                   *string   `json:"metadata"`
	MinimumValue               *int64    `json:"minimumValue"`
	MaximumValue               *int64    `json:"maximumValue"`
	OrderDirection             *string   `json:"orderDirection"`
	Scope                      *string   `json:"scope"`
	UniqueByUserId             *bool     `json:"uniqueByUserId"`
	Sum                        *bool     `json:"sum"`
	CalculateFixedTimingHour   *int32    `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32    `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes   *int32    `json:"calculateIntervalMinutes"`
	AdditionalScopes           []Scope   `json:"additionalScopes"`
	EntryPeriodEventId         *string   `json:"entryPeriodEventId"`
	AccessPeriodEventId        *string   `json:"accessPeriodEventId"`
	IgnoreUserIds              []*string `json:"ignoreUserIds"`
	Generation                 *string   `json:"generation"`
}

func NewCreateCategoryModelMasterRequestFromJson(data string) CreateCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCategoryModelMasterRequestFromDict(dict)
}

func NewCreateCategoryModelMasterRequestFromDict(data map[string]interface{}) CreateCategoryModelMasterRequest {
	return CreateCategoryModelMasterRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		Sum:                        core.CastBool(data["sum"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		AdditionalScopes:           CastScopes(core.CastArray(data["additionalScopes"])),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		IgnoreUserIds:              core.CastStrings(core.CastArray(data["ignoreUserIds"])),
		Generation:                 core.CastString(data["generation"]),
	}
}

func (p CreateCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":              p.NamespaceName,
		"name":                       p.Name,
		"description":                p.Description,
		"metadata":                   p.Metadata,
		"minimumValue":               p.MinimumValue,
		"maximumValue":               p.MaximumValue,
		"orderDirection":             p.OrderDirection,
		"scope":                      p.Scope,
		"uniqueByUserId":             p.UniqueByUserId,
		"sum":                        p.Sum,
		"calculateFixedTimingHour":   p.CalculateFixedTimingHour,
		"calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
		"calculateIntervalMinutes":   p.CalculateIntervalMinutes,
		"additionalScopes": CastScopesFromDict(
			p.AdditionalScopes,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
		"ignoreUserIds": core.CastStringsFromDict(
			p.IgnoreUserIds,
		),
		"generation": p.Generation,
	}
}

func (p CreateCategoryModelMasterRequest) Pointer() *CreateCategoryModelMasterRequest {
	return &p
}

type GetCategoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
}

func NewGetCategoryModelMasterRequestFromJson(data string) GetCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelMasterRequestFromDict(dict)
}

func NewGetCategoryModelMasterRequestFromDict(data map[string]interface{}) GetCategoryModelMasterRequest {
	return GetCategoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p GetCategoryModelMasterRequest) Pointer() *GetCategoryModelMasterRequest {
	return &p
}

type UpdateCategoryModelMasterRequest struct {
	SourceRequestId            *string   `json:"sourceRequestId"`
	RequestId                  *string   `json:"requestId"`
	ContextStack               *string   `json:"contextStack"`
	NamespaceName              *string   `json:"namespaceName"`
	CategoryName               *string   `json:"categoryName"`
	Description                *string   `json:"description"`
	Metadata                   *string   `json:"metadata"`
	MinimumValue               *int64    `json:"minimumValue"`
	MaximumValue               *int64    `json:"maximumValue"`
	OrderDirection             *string   `json:"orderDirection"`
	Scope                      *string   `json:"scope"`
	UniqueByUserId             *bool     `json:"uniqueByUserId"`
	Sum                        *bool     `json:"sum"`
	CalculateFixedTimingHour   *int32    `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32    `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes   *int32    `json:"calculateIntervalMinutes"`
	AdditionalScopes           []Scope   `json:"additionalScopes"`
	EntryPeriodEventId         *string   `json:"entryPeriodEventId"`
	AccessPeriodEventId        *string   `json:"accessPeriodEventId"`
	IgnoreUserIds              []*string `json:"ignoreUserIds"`
	Generation                 *string   `json:"generation"`
}

func NewUpdateCategoryModelMasterRequestFromJson(data string) UpdateCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCategoryModelMasterRequestFromDict(dict)
}

func NewUpdateCategoryModelMasterRequestFromDict(data map[string]interface{}) UpdateCategoryModelMasterRequest {
	return UpdateCategoryModelMasterRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		CategoryName:               core.CastString(data["categoryName"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		Sum:                        core.CastBool(data["sum"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		AdditionalScopes:           CastScopes(core.CastArray(data["additionalScopes"])),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		IgnoreUserIds:              core.CastStrings(core.CastArray(data["ignoreUserIds"])),
		Generation:                 core.CastString(data["generation"]),
	}
}

func (p UpdateCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":              p.NamespaceName,
		"categoryName":               p.CategoryName,
		"description":                p.Description,
		"metadata":                   p.Metadata,
		"minimumValue":               p.MinimumValue,
		"maximumValue":               p.MaximumValue,
		"orderDirection":             p.OrderDirection,
		"scope":                      p.Scope,
		"uniqueByUserId":             p.UniqueByUserId,
		"sum":                        p.Sum,
		"calculateFixedTimingHour":   p.CalculateFixedTimingHour,
		"calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
		"calculateIntervalMinutes":   p.CalculateIntervalMinutes,
		"additionalScopes": CastScopesFromDict(
			p.AdditionalScopes,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
		"ignoreUserIds": core.CastStringsFromDict(
			p.IgnoreUserIds,
		),
		"generation": p.Generation,
	}
}

func (p UpdateCategoryModelMasterRequest) Pointer() *UpdateCategoryModelMasterRequest {
	return &p
}

type DeleteCategoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
}

func NewDeleteCategoryModelMasterRequestFromJson(data string) DeleteCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCategoryModelMasterRequestFromDict(dict)
}

func NewDeleteCategoryModelMasterRequestFromDict(data map[string]interface{}) DeleteCategoryModelMasterRequest {
	return DeleteCategoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p DeleteCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p DeleteCategoryModelMasterRequest) Pointer() *DeleteCategoryModelMasterRequest {
	return &p
}

type SubscribeRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
}

func NewSubscribeRequestFromJson(data string) SubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeRequestFromDict(dict)
}

func NewSubscribeRequestFromDict(data map[string]interface{}) SubscribeRequest {
	return SubscribeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TargetUserId:  core.CastString(data["targetUserId"]),
	}
}

func (p SubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p SubscribeRequest) Pointer() *SubscribeRequest {
	return &p
}

type SubscribeByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	UserId             *string `json:"userId"`
	TargetUserId       *string `json:"targetUserId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewSubscribeByUserIdRequestFromJson(data string) SubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeByUserIdRequestFromDict(dict)
}

func NewSubscribeByUserIdRequestFromDict(data map[string]interface{}) SubscribeByUserIdRequest {
	return SubscribeByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		TargetUserId:    core.CastString(data["targetUserId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SubscribeByUserIdRequest) Pointer() *SubscribeByUserIdRequest {
	return &p
}

type DescribeScoresRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	AccessToken     *string `json:"accessToken"`
	ScorerUserId    *string `json:"scorerUserId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeScoresRequestFromJson(data string) DescribeScoresRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScoresRequestFromDict(dict)
}

func NewDescribeScoresRequestFromDict(data map[string]interface{}) DescribeScoresRequest {
	return DescribeScoresRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ScorerUserId:  core.CastString(data["scorerUserId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeScoresRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"scorerUserId":  p.ScorerUserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeScoresRequest) Pointer() *DescribeScoresRequest {
	return &p
}

type DescribeScoresByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	UserId          *string `json:"userId"`
	ScorerUserId    *string `json:"scorerUserId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeScoresByUserIdRequestFromJson(data string) DescribeScoresByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScoresByUserIdRequestFromDict(dict)
}

func NewDescribeScoresByUserIdRequestFromDict(data map[string]interface{}) DescribeScoresByUserIdRequest {
	return DescribeScoresByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		ScorerUserId:    core.CastString(data["scorerUserId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeScoresByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"scorerUserId":    p.ScorerUserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeScoresByUserIdRequest) Pointer() *DescribeScoresByUserIdRequest {
	return &p
}

type GetScoreRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	AccessToken     *string `json:"accessToken"`
	ScorerUserId    *string `json:"scorerUserId"`
	UniqueId        *string `json:"uniqueId"`
}

func NewGetScoreRequestFromJson(data string) GetScoreRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScoreRequestFromDict(dict)
}

func NewGetScoreRequestFromDict(data map[string]interface{}) GetScoreRequest {
	return GetScoreRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ScorerUserId:  core.CastString(data["scorerUserId"]),
		UniqueId:      core.CastString(data["uniqueId"]),
	}
}

func (p GetScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"scorerUserId":  p.ScorerUserId,
		"uniqueId":      p.UniqueId,
	}
}

func (p GetScoreRequest) Pointer() *GetScoreRequest {
	return &p
}

type GetScoreByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	UserId          *string `json:"userId"`
	ScorerUserId    *string `json:"scorerUserId"`
	UniqueId        *string `json:"uniqueId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetScoreByUserIdRequestFromJson(data string) GetScoreByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScoreByUserIdRequestFromDict(dict)
}

func NewGetScoreByUserIdRequestFromDict(data map[string]interface{}) GetScoreByUserIdRequest {
	return GetScoreByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		ScorerUserId:    core.CastString(data["scorerUserId"]),
		UniqueId:        core.CastString(data["uniqueId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"scorerUserId":    p.ScorerUserId,
		"uniqueId":        p.UniqueId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetScoreByUserIdRequest) Pointer() *GetScoreByUserIdRequest {
	return &p
}

type DescribeRankingsRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	AccessToken         *string `json:"accessToken"`
	AdditionalScopeName *string `json:"additionalScopeName"`
	StartIndex          *int64  `json:"startIndex"`
	PageToken           *string `json:"pageToken"`
	Limit               *int32  `json:"limit"`
}

func NewDescribeRankingsRequestFromJson(data string) DescribeRankingsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingsRequestFromDict(dict)
}

func NewDescribeRankingsRequestFromDict(data map[string]interface{}) DescribeRankingsRequest {
	return DescribeRankingsRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		AccessToken:         core.CastString(data["accessToken"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
		StartIndex:          core.CastInt64(data["startIndex"]),
		PageToken:           core.CastString(data["pageToken"]),
		Limit:               core.CastInt32(data["limit"]),
	}
}

func (p DescribeRankingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"accessToken":         p.AccessToken,
		"additionalScopeName": p.AdditionalScopeName,
		"startIndex":          p.StartIndex,
		"pageToken":           p.PageToken,
		"limit":               p.Limit,
	}
}

func (p DescribeRankingsRequest) Pointer() *DescribeRankingsRequest {
	return &p
}

type DescribeRankingssByUserIdRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	UserId              *string `json:"userId"`
	AdditionalScopeName *string `json:"additionalScopeName"`
	StartIndex          *int64  `json:"startIndex"`
	PageToken           *string `json:"pageToken"`
	Limit               *int32  `json:"limit"`
	TimeOffsetToken     *string `json:"timeOffsetToken"`
}

func NewDescribeRankingssByUserIdRequestFromJson(data string) DescribeRankingssByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingssByUserIdRequestFromDict(dict)
}

func NewDescribeRankingssByUserIdRequestFromDict(data map[string]interface{}) DescribeRankingssByUserIdRequest {
	return DescribeRankingssByUserIdRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		UserId:              core.CastString(data["userId"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
		StartIndex:          core.CastInt64(data["startIndex"]),
		PageToken:           core.CastString(data["pageToken"]),
		Limit:               core.CastInt32(data["limit"]),
		TimeOffsetToken:     core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeRankingssByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"userId":              p.UserId,
		"additionalScopeName": p.AdditionalScopeName,
		"startIndex":          p.StartIndex,
		"pageToken":           p.PageToken,
		"limit":               p.Limit,
		"timeOffsetToken":     p.TimeOffsetToken,
	}
}

func (p DescribeRankingssByUserIdRequest) Pointer() *DescribeRankingssByUserIdRequest {
	return &p
}

type DescribeNearRankingsRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	AdditionalScopeName *string `json:"additionalScopeName"`
	Score               *int64  `json:"score"`
}

func NewDescribeNearRankingsRequestFromJson(data string) DescribeNearRankingsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNearRankingsRequestFromDict(dict)
}

func NewDescribeNearRankingsRequestFromDict(data map[string]interface{}) DescribeNearRankingsRequest {
	return DescribeNearRankingsRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
		Score:               core.CastInt64(data["score"]),
	}
}

func (p DescribeNearRankingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"additionalScopeName": p.AdditionalScopeName,
		"score":               p.Score,
	}
}

func (p DescribeNearRankingsRequest) Pointer() *DescribeNearRankingsRequest {
	return &p
}

type GetRankingRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	AccessToken         *string `json:"accessToken"`
	ScorerUserId        *string `json:"scorerUserId"`
	UniqueId            *string `json:"uniqueId"`
	AdditionalScopeName *string `json:"additionalScopeName"`
}

func NewGetRankingRequestFromJson(data string) GetRankingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingRequestFromDict(dict)
}

func NewGetRankingRequestFromDict(data map[string]interface{}) GetRankingRequest {
	return GetRankingRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		AccessToken:         core.CastString(data["accessToken"]),
		ScorerUserId:        core.CastString(data["scorerUserId"]),
		UniqueId:            core.CastString(data["uniqueId"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
	}
}

func (p GetRankingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"accessToken":         p.AccessToken,
		"scorerUserId":        p.ScorerUserId,
		"uniqueId":            p.UniqueId,
		"additionalScopeName": p.AdditionalScopeName,
	}
}

func (p GetRankingRequest) Pointer() *GetRankingRequest {
	return &p
}

type GetRankingByUserIdRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	UserId              *string `json:"userId"`
	ScorerUserId        *string `json:"scorerUserId"`
	UniqueId            *string `json:"uniqueId"`
	AdditionalScopeName *string `json:"additionalScopeName"`
	TimeOffsetToken     *string `json:"timeOffsetToken"`
}

func NewGetRankingByUserIdRequestFromJson(data string) GetRankingByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingByUserIdRequestFromDict(dict)
}

func NewGetRankingByUserIdRequestFromDict(data map[string]interface{}) GetRankingByUserIdRequest {
	return GetRankingByUserIdRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		UserId:              core.CastString(data["userId"]),
		ScorerUserId:        core.CastString(data["scorerUserId"]),
		UniqueId:            core.CastString(data["uniqueId"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
		TimeOffsetToken:     core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetRankingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"userId":              p.UserId,
		"scorerUserId":        p.ScorerUserId,
		"uniqueId":            p.UniqueId,
		"additionalScopeName": p.AdditionalScopeName,
		"timeOffsetToken":     p.TimeOffsetToken,
	}
}

func (p GetRankingByUserIdRequest) Pointer() *GetRankingByUserIdRequest {
	return &p
}

type PutScoreRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	AccessToken        *string `json:"accessToken"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
}

func NewPutScoreRequestFromJson(data string) PutScoreRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutScoreRequestFromDict(dict)
}

func NewPutScoreRequestFromDict(data map[string]interface{}) PutScoreRequest {
	return PutScoreRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Score:         core.CastInt64(data["score"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p PutScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"score":         p.Score,
		"metadata":      p.Metadata,
	}
}

func (p PutScoreRequest) Pointer() *PutScoreRequest {
	return &p
}

type PutScoreByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	UserId             *string `json:"userId"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPutScoreByUserIdRequestFromJson(data string) PutScoreByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutScoreByUserIdRequestFromDict(dict)
}

func NewPutScoreByUserIdRequestFromDict(data map[string]interface{}) PutScoreByUserIdRequest {
	return PutScoreByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		Score:           core.CastInt64(data["score"]),
		Metadata:        core.CastString(data["metadata"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PutScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"score":           p.Score,
		"metadata":        p.Metadata,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PutScoreByUserIdRequest) Pointer() *PutScoreByUserIdRequest {
	return &p
}

type CalcRankingRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	CategoryName        *string `json:"categoryName"`
	AdditionalScopeName *string `json:"additionalScopeName"`
}

func NewCalcRankingRequestFromJson(data string) CalcRankingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCalcRankingRequestFromDict(dict)
}

func NewCalcRankingRequestFromDict(data map[string]interface{}) CalcRankingRequest {
	return CalcRankingRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		CategoryName:        core.CastString(data["categoryName"]),
		AdditionalScopeName: core.CastString(data["additionalScopeName"]),
	}
}

func (p CalcRankingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"categoryName":        p.CategoryName,
		"additionalScopeName": p.AdditionalScopeName,
	}
}

func (p CalcRankingRequest) Pointer() *CalcRankingRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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

type GetCurrentRankingMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentRankingMasterRequestFromJson(data string) GetCurrentRankingMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRankingMasterRequestFromDict(dict)
}

func NewGetCurrentRankingMasterRequestFromDict(data map[string]interface{}) GetCurrentRankingMasterRequest {
	return GetCurrentRankingMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentRankingMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentRankingMasterRequest) Pointer() *GetCurrentRankingMasterRequest {
	return &p
}

type UpdateCurrentRankingMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentRankingMasterRequestFromJson(data string) UpdateCurrentRankingMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterRequestFromDict(dict)
}

func NewUpdateCurrentRankingMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterRequest {
	return UpdateCurrentRankingMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentRankingMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentRankingMasterRequest) Pointer() *UpdateCurrentRankingMasterRequest {
	return &p
}

type UpdateCurrentRankingMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRankingMasterFromGitHubRequestFromJson(data string) UpdateCurrentRankingMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentRankingMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubRequest {
	return UpdateCurrentRankingMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubRequest) Pointer() *UpdateCurrentRankingMasterFromGitHubRequest {
	return &p
}

type GetSubscribeRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	AccessToken     *string `json:"accessToken"`
	TargetUserId    *string `json:"targetUserId"`
}

func NewGetSubscribeRequestFromJson(data string) GetSubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRequestFromDict(dict)
}

func NewGetSubscribeRequestFromDict(data map[string]interface{}) GetSubscribeRequest {
	return GetSubscribeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TargetUserId:  core.CastString(data["targetUserId"]),
	}
}

func (p GetSubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p GetSubscribeRequest) Pointer() *GetSubscribeRequest {
	return &p
}

type GetSubscribeByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	UserId          *string `json:"userId"`
	TargetUserId    *string `json:"targetUserId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetSubscribeByUserIdRequestFromJson(data string) GetSubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeByUserIdRequestFromDict(dict)
}

func NewGetSubscribeByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeByUserIdRequest {
	return GetSubscribeByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		TargetUserId:    core.CastString(data["targetUserId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetSubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSubscribeByUserIdRequest) Pointer() *GetSubscribeByUserIdRequest {
	return &p
}

type UnsubscribeRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
}

func NewUnsubscribeRequestFromJson(data string) UnsubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeRequestFromDict(dict)
}

func NewUnsubscribeRequestFromDict(data map[string]interface{}) UnsubscribeRequest {
	return UnsubscribeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TargetUserId:  core.CastString(data["targetUserId"]),
	}
}

func (p UnsubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p UnsubscribeRequest) Pointer() *UnsubscribeRequest {
	return &p
}

type UnsubscribeByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CategoryName       *string `json:"categoryName"`
	UserId             *string `json:"userId"`
	TargetUserId       *string `json:"targetUserId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUnsubscribeByUserIdRequestFromJson(data string) UnsubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeByUserIdRequestFromDict(dict)
}

func NewUnsubscribeByUserIdRequestFromDict(data map[string]interface{}) UnsubscribeByUserIdRequest {
	return UnsubscribeByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		TargetUserId:    core.CastString(data["targetUserId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UnsubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UnsubscribeByUserIdRequest) Pointer() *UnsubscribeByUserIdRequest {
	return &p
}

type DescribeSubscribesByCategoryNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	AccessToken     *string `json:"accessToken"`
}

func NewDescribeSubscribesByCategoryNameRequestFromJson(data string) DescribeSubscribesByCategoryNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByCategoryNameRequestFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameRequestFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameRequest {
	return DescribeSubscribesByCategoryNameRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DescribeSubscribesByCategoryNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeSubscribesByCategoryNameRequest) Pointer() *DescribeSubscribesByCategoryNameRequest {
	return &p
}

type DescribeSubscribesByCategoryNameAndUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeSubscribesByCategoryNameAndUserIdRequestFromJson(data string) DescribeSubscribesByCategoryNameAndUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByCategoryNameAndUserIdRequestFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameAndUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameAndUserIdRequest {
	return DescribeSubscribesByCategoryNameAndUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"categoryName":    p.CategoryName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdRequest) Pointer() *DescribeSubscribesByCategoryNameAndUserIdRequest {
	return &p
}
