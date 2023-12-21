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
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNamespacesRequestFromDict(dict2)
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
	RequestId    *string     `json:"requestId"`
	ContextStack *string     `json:"contextStack"`
	Name         *string     `json:"name"`
	Description  *string     `json:"description"`
	LogSetting   *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateNamespaceRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceStatusRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceRequestFromDict(dict2)
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Description   *string     `json:"description"`
	LogSetting    *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateNamespaceRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteNamespaceRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDumpUserDataByUserIdRequestFromDict(dict2)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCleanUserDataByUserIdRequestFromDict(dict2)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict2)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewImportUserDataByUserIdRequestFromDict(dict2)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict2)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeCountersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	LimitName     *string `json:"limitName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeCountersRequestFromJson(data string) DescribeCountersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeCountersRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	LimitName     *string `json:"limitName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeCountersByUserIdRequestFromJson(data string) DescribeCountersByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeCountersByUserIdRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LimitName     *string `json:"limitName"`
	AccessToken   *string `json:"accessToken"`
	CounterName   *string `json:"counterName"`
}

func NewGetCounterRequestFromJson(data string) GetCounterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCounterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LimitName     *string `json:"limitName"`
	UserId        *string `json:"userId"`
	CounterName   *string `json:"counterName"`
}

func NewGetCounterByUserIdRequestFromJson(data string) GetCounterByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCounterByUserIdRequestFromDict(dict2)
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

func NewCountUpRequestFromJson(data string) CountUpRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCountUpRequestFromDict(dict2)
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

func NewCountUpByUserIdRequestFromJson(data string) CountUpByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCountUpByUserIdRequestFromDict(dict2)
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

type CountDownByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LimitName          *string `json:"limitName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
	CountDownValue     *int32  `json:"countDownValue"`
}

func NewCountDownByUserIdRequestFromJson(data string) CountDownByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCountDownByUserIdRequestFromDict(dict2)
}

func NewCountDownByUserIdRequestFromDict(data map[string]interface{}) CountDownByUserIdRequest {
	return CountDownByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		LimitName:      core.CastString(data["limitName"]),
		CounterName:    core.CastString(data["counterName"]),
		UserId:         core.CastString(data["userId"]),
		CountDownValue: core.CastInt32(data["countDownValue"]),
	}
}

func (p CountDownByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"limitName":      p.LimitName,
		"counterName":    p.CounterName,
		"userId":         p.UserId,
		"countDownValue": p.CountDownValue,
	}
}

func (p CountDownByUserIdRequest) Pointer() *CountDownByUserIdRequest {
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

func NewDeleteCounterByUserIdRequestFromJson(data string) DeleteCounterByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteCounterByUserIdRequestFromDict(dict2)
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

type VerifyCounterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	LimitName          *string `json:"limitName"`
	CounterName        *string `json:"counterName"`
	VerifyType         *string `json:"verifyType"`
	Count              *int32  `json:"count"`
}

func NewVerifyCounterRequestFromJson(data string) VerifyCounterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyCounterRequestFromDict(dict2)
}

func NewVerifyCounterRequestFromDict(data map[string]interface{}) VerifyCounterRequest {
	return VerifyCounterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		LimitName:     core.CastString(data["limitName"]),
		CounterName:   core.CastString(data["counterName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastInt32(data["count"]),
	}
}

func (p VerifyCounterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"limitName":     p.LimitName,
		"counterName":   p.CounterName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifyCounterRequest) Pointer() *VerifyCounterRequest {
	return &p
}

type VerifyCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	LimitName          *string `json:"limitName"`
	CounterName        *string `json:"counterName"`
	VerifyType         *string `json:"verifyType"`
	Count              *int32  `json:"count"`
}

func NewVerifyCounterByUserIdRequestFromJson(data string) VerifyCounterByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyCounterByUserIdRequestFromDict(dict2)
}

func NewVerifyCounterByUserIdRequestFromDict(data map[string]interface{}) VerifyCounterByUserIdRequest {
	return VerifyCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		LimitName:     core.CastString(data["limitName"]),
		CounterName:   core.CastString(data["counterName"]),
		VerifyType:    core.CastString(data["verifyType"]),
		Count:         core.CastInt32(data["count"]),
	}
}

func (p VerifyCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"limitName":     p.LimitName,
		"counterName":   p.CounterName,
		"verifyType":    p.VerifyType,
		"count":         p.Count,
	}
}

func (p VerifyCounterByUserIdRequest) Pointer() *VerifyCounterByUserIdRequest {
	return &p
}

type CountUpByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewCountUpByStampTaskRequestFromJson(data string) CountUpByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCountUpByStampTaskRequestFromDict(dict2)
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

type CountDownByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewCountDownByStampSheetRequestFromJson(data string) CountDownByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCountDownByStampSheetRequestFromDict(dict2)
}

func NewCountDownByStampSheetRequestFromDict(data map[string]interface{}) CountDownByStampSheetRequest {
	return CountDownByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p CountDownByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p CountDownByStampSheetRequest) Pointer() *CountDownByStampSheetRequest {
	return &p
}

type DeleteByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewDeleteByStampSheetRequestFromJson(data string) DeleteByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteByStampSheetRequestFromDict(dict2)
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

type VerifyCounterByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewVerifyCounterByStampTaskRequestFromJson(data string) VerifyCounterByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewVerifyCounterByStampTaskRequestFromDict(dict2)
}

func NewVerifyCounterByStampTaskRequestFromDict(data map[string]interface{}) VerifyCounterByStampTaskRequest {
	return VerifyCounterByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyCounterByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyCounterByStampTaskRequest) Pointer() *VerifyCounterByStampTaskRequest {
	return &p
}

type DescribeLimitModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeLimitModelMastersRequestFromJson(data string) DescribeLimitModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeLimitModelMastersRequestFromDict(dict2)
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
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	ResetType       *string `json:"resetType"`
	ResetDayOfMonth *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek  *string `json:"resetDayOfWeek"`
	ResetHour       *int32  `json:"resetHour"`
}

func NewCreateLimitModelMasterRequestFromJson(data string) CreateLimitModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateLimitModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LimitName     *string `json:"limitName"`
}

func NewGetLimitModelMasterRequestFromJson(data string) GetLimitModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetLimitModelMasterRequestFromDict(dict2)
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
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	LimitName       *string `json:"limitName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	ResetType       *string `json:"resetType"`
	ResetDayOfMonth *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek  *string `json:"resetDayOfWeek"`
	ResetHour       *int32  `json:"resetHour"`
}

func NewUpdateLimitModelMasterRequestFromJson(data string) UpdateLimitModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateLimitModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LimitName     *string `json:"limitName"`
}

func NewDeleteLimitModelMasterRequestFromJson(data string) DeleteLimitModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteLimitModelMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewExportMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentLimitMasterRequestFromJson(data string) GetCurrentLimitMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentLimitMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentLimitMasterRequestFromJson(data string) UpdateCurrentLimitMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentLimitMasterRequestFromDict(dict2)
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
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentLimitMasterFromGitHubRequestFromJson(data string) UpdateCurrentLimitMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentLimitMasterFromGitHubRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeLimitModelsRequestFromJson(data string) DescribeLimitModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeLimitModelsRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LimitName     *string `json:"limitName"`
}

func NewGetLimitModelRequestFromJson(data string) GetLimitModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetLimitModelRequestFromDict(dict2)
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
