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

package schedule

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

type DescribeEventMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeEventMastersRequestFromJson(data string) DescribeEventMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventMastersRequestFromDict(dict)
}

func NewDescribeEventMastersRequestFromDict(data map[string]interface{}) DescribeEventMastersRequest {
	return DescribeEventMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeEventMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeEventMastersRequest) Pointer() *DescribeEventMastersRequest {
	return &p
}

type CreateEventMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	Name                  *string `json:"name"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	ScheduleType          *string `json:"scheduleType"`
	AbsoluteBegin         *int64  `json:"absoluteBegin"`
	AbsoluteEnd           *int64  `json:"absoluteEnd"`
	RepeatType            *string `json:"repeatType"`
	RepeatBeginDayOfMonth *int32  `json:"repeatBeginDayOfMonth"`
	RepeatEndDayOfMonth   *int32  `json:"repeatEndDayOfMonth"`
	RepeatBeginDayOfWeek  *string `json:"repeatBeginDayOfWeek"`
	RepeatEndDayOfWeek    *string `json:"repeatEndDayOfWeek"`
	RepeatBeginHour       *int32  `json:"repeatBeginHour"`
	RepeatEndHour         *int32  `json:"repeatEndHour"`
	RelativeTriggerName   *string `json:"relativeTriggerName"`
}

func NewCreateEventMasterRequestFromJson(data string) CreateEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateEventMasterRequestFromDict(dict)
}

func NewCreateEventMasterRequestFromDict(data map[string]interface{}) CreateEventMasterRequest {
	return CreateEventMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ScheduleType:          core.CastString(data["scheduleType"]),
		AbsoluteBegin:         core.CastInt64(data["absoluteBegin"]),
		AbsoluteEnd:           core.CastInt64(data["absoluteEnd"]),
		RepeatType:            core.CastString(data["repeatType"]),
		RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
		RepeatEndDayOfMonth:   core.CastInt32(data["repeatEndDayOfMonth"]),
		RepeatBeginDayOfWeek:  core.CastString(data["repeatBeginDayOfWeek"]),
		RepeatEndDayOfWeek:    core.CastString(data["repeatEndDayOfWeek"]),
		RepeatBeginHour:       core.CastInt32(data["repeatBeginHour"]),
		RepeatEndHour:         core.CastInt32(data["repeatEndHour"]),
		RelativeTriggerName:   core.CastString(data["relativeTriggerName"]),
	}
}

func (p CreateEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"name":                  p.Name,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"scheduleType":          p.ScheduleType,
		"absoluteBegin":         p.AbsoluteBegin,
		"absoluteEnd":           p.AbsoluteEnd,
		"repeatType":            p.RepeatType,
		"repeatBeginDayOfMonth": p.RepeatBeginDayOfMonth,
		"repeatEndDayOfMonth":   p.RepeatEndDayOfMonth,
		"repeatBeginDayOfWeek":  p.RepeatBeginDayOfWeek,
		"repeatEndDayOfWeek":    p.RepeatEndDayOfWeek,
		"repeatBeginHour":       p.RepeatBeginHour,
		"repeatEndHour":         p.RepeatEndHour,
		"relativeTriggerName":   p.RelativeTriggerName,
	}
}

func (p CreateEventMasterRequest) Pointer() *CreateEventMasterRequest {
	return &p
}

type GetEventMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	EventName       *string `json:"eventName"`
}

func NewGetEventMasterRequestFromJson(data string) GetEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventMasterRequestFromDict(dict)
}

func NewGetEventMasterRequestFromDict(data map[string]interface{}) GetEventMasterRequest {
	return GetEventMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EventName:     core.CastString(data["eventName"]),
	}
}

func (p GetEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"eventName":     p.EventName,
	}
}

func (p GetEventMasterRequest) Pointer() *GetEventMasterRequest {
	return &p
}

type UpdateEventMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	EventName             *string `json:"eventName"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	ScheduleType          *string `json:"scheduleType"`
	AbsoluteBegin         *int64  `json:"absoluteBegin"`
	AbsoluteEnd           *int64  `json:"absoluteEnd"`
	RepeatType            *string `json:"repeatType"`
	RepeatBeginDayOfMonth *int32  `json:"repeatBeginDayOfMonth"`
	RepeatEndDayOfMonth   *int32  `json:"repeatEndDayOfMonth"`
	RepeatBeginDayOfWeek  *string `json:"repeatBeginDayOfWeek"`
	RepeatEndDayOfWeek    *string `json:"repeatEndDayOfWeek"`
	RepeatBeginHour       *int32  `json:"repeatBeginHour"`
	RepeatEndHour         *int32  `json:"repeatEndHour"`
	RelativeTriggerName   *string `json:"relativeTriggerName"`
}

func NewUpdateEventMasterRequestFromJson(data string) UpdateEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateEventMasterRequestFromDict(dict)
}

func NewUpdateEventMasterRequestFromDict(data map[string]interface{}) UpdateEventMasterRequest {
	return UpdateEventMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		EventName:             core.CastString(data["eventName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ScheduleType:          core.CastString(data["scheduleType"]),
		AbsoluteBegin:         core.CastInt64(data["absoluteBegin"]),
		AbsoluteEnd:           core.CastInt64(data["absoluteEnd"]),
		RepeatType:            core.CastString(data["repeatType"]),
		RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
		RepeatEndDayOfMonth:   core.CastInt32(data["repeatEndDayOfMonth"]),
		RepeatBeginDayOfWeek:  core.CastString(data["repeatBeginDayOfWeek"]),
		RepeatEndDayOfWeek:    core.CastString(data["repeatEndDayOfWeek"]),
		RepeatBeginHour:       core.CastInt32(data["repeatBeginHour"]),
		RepeatEndHour:         core.CastInt32(data["repeatEndHour"]),
		RelativeTriggerName:   core.CastString(data["relativeTriggerName"]),
	}
}

func (p UpdateEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"eventName":             p.EventName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"scheduleType":          p.ScheduleType,
		"absoluteBegin":         p.AbsoluteBegin,
		"absoluteEnd":           p.AbsoluteEnd,
		"repeatType":            p.RepeatType,
		"repeatBeginDayOfMonth": p.RepeatBeginDayOfMonth,
		"repeatEndDayOfMonth":   p.RepeatEndDayOfMonth,
		"repeatBeginDayOfWeek":  p.RepeatBeginDayOfWeek,
		"repeatEndDayOfWeek":    p.RepeatEndDayOfWeek,
		"repeatBeginHour":       p.RepeatBeginHour,
		"repeatEndHour":         p.RepeatEndHour,
		"relativeTriggerName":   p.RelativeTriggerName,
	}
}

func (p UpdateEventMasterRequest) Pointer() *UpdateEventMasterRequest {
	return &p
}

type DeleteEventMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	EventName       *string `json:"eventName"`
}

func NewDeleteEventMasterRequestFromJson(data string) DeleteEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEventMasterRequestFromDict(dict)
}

func NewDeleteEventMasterRequestFromDict(data map[string]interface{}) DeleteEventMasterRequest {
	return DeleteEventMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EventName:     core.CastString(data["eventName"]),
	}
}

func (p DeleteEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"eventName":     p.EventName,
	}
}

func (p DeleteEventMasterRequest) Pointer() *DeleteEventMasterRequest {
	return &p
}

type DescribeTriggersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeTriggersRequestFromJson(data string) DescribeTriggersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTriggersRequestFromDict(dict)
}

func NewDescribeTriggersRequestFromDict(data map[string]interface{}) DescribeTriggersRequest {
	return DescribeTriggersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeTriggersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeTriggersRequest) Pointer() *DescribeTriggersRequest {
	return &p
}

type DescribeTriggersByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeTriggersByUserIdRequestFromJson(data string) DescribeTriggersByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTriggersByUserIdRequestFromDict(dict)
}

func NewDescribeTriggersByUserIdRequestFromDict(data map[string]interface{}) DescribeTriggersByUserIdRequest {
	return DescribeTriggersByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeTriggersByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeTriggersByUserIdRequest) Pointer() *DescribeTriggersByUserIdRequest {
	return &p
}

type GetTriggerRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	TriggerName     *string `json:"triggerName"`
}

func NewGetTriggerRequestFromJson(data string) GetTriggerRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTriggerRequestFromDict(dict)
}

func NewGetTriggerRequestFromDict(data map[string]interface{}) GetTriggerRequest {
	return GetTriggerRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TriggerName:   core.CastString(data["triggerName"]),
	}
}

func (p GetTriggerRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"triggerName":   p.TriggerName,
	}
}

func (p GetTriggerRequest) Pointer() *GetTriggerRequest {
	return &p
}

type GetTriggerByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TriggerName     *string `json:"triggerName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetTriggerByUserIdRequestFromJson(data string) GetTriggerByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTriggerByUserIdRequestFromDict(dict)
}

func NewGetTriggerByUserIdRequestFromDict(data map[string]interface{}) GetTriggerByUserIdRequest {
	return GetTriggerByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TriggerName:     core.CastString(data["triggerName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetTriggerByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"triggerName":     p.TriggerName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetTriggerByUserIdRequest) Pointer() *GetTriggerByUserIdRequest {
	return &p
}

type TriggerByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	TriggerName        *string `json:"triggerName"`
	UserId             *string `json:"userId"`
	TriggerStrategy    *string `json:"triggerStrategy"`
	Ttl                *int32  `json:"ttl"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewTriggerByUserIdRequestFromJson(data string) TriggerByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTriggerByUserIdRequestFromDict(dict)
}

func NewTriggerByUserIdRequestFromDict(data map[string]interface{}) TriggerByUserIdRequest {
	return TriggerByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		TriggerName:     core.CastString(data["triggerName"]),
		UserId:          core.CastString(data["userId"]),
		TriggerStrategy: core.CastString(data["triggerStrategy"]),
		Ttl:             core.CastInt32(data["ttl"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p TriggerByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"triggerName":     p.TriggerName,
		"userId":          p.UserId,
		"triggerStrategy": p.TriggerStrategy,
		"ttl":             p.Ttl,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p TriggerByUserIdRequest) Pointer() *TriggerByUserIdRequest {
	return &p
}

type TriggerByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewTriggerByStampSheetRequestFromJson(data string) TriggerByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTriggerByStampSheetRequestFromDict(dict)
}

func NewTriggerByStampSheetRequestFromDict(data map[string]interface{}) TriggerByStampSheetRequest {
	return TriggerByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p TriggerByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p TriggerByStampSheetRequest) Pointer() *TriggerByStampSheetRequest {
	return &p
}

type DeleteTriggerRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	TriggerName        *string `json:"triggerName"`
}

func NewDeleteTriggerRequestFromJson(data string) DeleteTriggerRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerRequestFromDict(dict)
}

func NewDeleteTriggerRequestFromDict(data map[string]interface{}) DeleteTriggerRequest {
	return DeleteTriggerRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TriggerName:   core.CastString(data["triggerName"]),
	}
}

func (p DeleteTriggerRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"triggerName":   p.TriggerName,
	}
}

func (p DeleteTriggerRequest) Pointer() *DeleteTriggerRequest {
	return &p
}

type DeleteTriggerByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TriggerName        *string `json:"triggerName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteTriggerByUserIdRequestFromJson(data string) DeleteTriggerByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerByUserIdRequestFromDict(dict)
}

func NewDeleteTriggerByUserIdRequestFromDict(data map[string]interface{}) DeleteTriggerByUserIdRequest {
	return DeleteTriggerByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TriggerName:     core.CastString(data["triggerName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteTriggerByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"triggerName":     p.TriggerName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteTriggerByUserIdRequest) Pointer() *DeleteTriggerByUserIdRequest {
	return &p
}

type DeleteTriggerByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewDeleteTriggerByStampTaskRequestFromJson(data string) DeleteTriggerByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerByStampTaskRequestFromDict(dict)
}

func NewDeleteTriggerByStampTaskRequestFromDict(data map[string]interface{}) DeleteTriggerByStampTaskRequest {
	return DeleteTriggerByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DeleteTriggerByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DeleteTriggerByStampTaskRequest) Pointer() *DeleteTriggerByStampTaskRequest {
	return &p
}

type DescribeEventsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
}

func NewDescribeEventsRequestFromJson(data string) DescribeEventsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsRequestFromDict(dict)
}

func NewDescribeEventsRequestFromDict(data map[string]interface{}) DescribeEventsRequest {
	return DescribeEventsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DescribeEventsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeEventsRequest) Pointer() *DescribeEventsRequest {
	return &p
}

type DescribeEventsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeEventsByUserIdRequestFromJson(data string) DescribeEventsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsByUserIdRequestFromDict(dict)
}

func NewDescribeEventsByUserIdRequestFromDict(data map[string]interface{}) DescribeEventsByUserIdRequest {
	return DescribeEventsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeEventsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeEventsByUserIdRequest) Pointer() *DescribeEventsByUserIdRequest {
	return &p
}

type DescribeRawEventsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeRawEventsRequestFromJson(data string) DescribeRawEventsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRawEventsRequestFromDict(dict)
}

func NewDescribeRawEventsRequestFromDict(data map[string]interface{}) DescribeRawEventsRequest {
	return DescribeRawEventsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeRawEventsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRawEventsRequest) Pointer() *DescribeRawEventsRequest {
	return &p
}

type GetEventRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	EventName       *string `json:"eventName"`
	AccessToken     *string `json:"accessToken"`
	IsInSchedule    *bool   `json:"isInSchedule"`
}

func NewGetEventRequestFromJson(data string) GetEventRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventRequestFromDict(dict)
}

func NewGetEventRequestFromDict(data map[string]interface{}) GetEventRequest {
	return GetEventRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EventName:     core.CastString(data["eventName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		IsInSchedule:  core.CastBool(data["isInSchedule"]),
	}
}

func (p GetEventRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"eventName":     p.EventName,
		"accessToken":   p.AccessToken,
		"isInSchedule":  p.IsInSchedule,
	}
}

func (p GetEventRequest) Pointer() *GetEventRequest {
	return &p
}

type GetEventByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	EventName       *string `json:"eventName"`
	UserId          *string `json:"userId"`
	IsInSchedule    *bool   `json:"isInSchedule"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetEventByUserIdRequestFromJson(data string) GetEventByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventByUserIdRequestFromDict(dict)
}

func NewGetEventByUserIdRequestFromDict(data map[string]interface{}) GetEventByUserIdRequest {
	return GetEventByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		EventName:       core.CastString(data["eventName"]),
		UserId:          core.CastString(data["userId"]),
		IsInSchedule:    core.CastBool(data["isInSchedule"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetEventByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"eventName":       p.EventName,
		"userId":          p.UserId,
		"isInSchedule":    p.IsInSchedule,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetEventByUserIdRequest) Pointer() *GetEventByUserIdRequest {
	return &p
}

type GetRawEventRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	EventName       *string `json:"eventName"`
}

func NewGetRawEventRequestFromJson(data string) GetRawEventRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRawEventRequestFromDict(dict)
}

func NewGetRawEventRequestFromDict(data map[string]interface{}) GetRawEventRequest {
	return GetRawEventRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EventName:     core.CastString(data["eventName"]),
	}
}

func (p GetRawEventRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"eventName":     p.EventName,
	}
}

func (p GetRawEventRequest) Pointer() *GetRawEventRequest {
	return &p
}

type VerifyEventRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	EventName          *string `json:"eventName"`
	VerifyType         *string `json:"verifyType"`
}

func NewVerifyEventRequestFromJson(data string) VerifyEventRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventRequestFromDict(dict)
}

func NewVerifyEventRequestFromDict(data map[string]interface{}) VerifyEventRequest {
	return VerifyEventRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		EventName:     core.CastString(data["eventName"]),
		VerifyType:    core.CastString(data["verifyType"]),
	}
}

func (p VerifyEventRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"eventName":     p.EventName,
		"verifyType":    p.VerifyType,
	}
}

func (p VerifyEventRequest) Pointer() *VerifyEventRequest {
	return &p
}

type VerifyEventByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	EventName          *string `json:"eventName"`
	VerifyType         *string `json:"verifyType"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewVerifyEventByUserIdRequestFromJson(data string) VerifyEventByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventByUserIdRequestFromDict(dict)
}

func NewVerifyEventByUserIdRequestFromDict(data map[string]interface{}) VerifyEventByUserIdRequest {
	return VerifyEventByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		EventName:       core.CastString(data["eventName"]),
		VerifyType:      core.CastString(data["verifyType"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifyEventByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"eventName":       p.EventName,
		"verifyType":      p.VerifyType,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p VerifyEventByUserIdRequest) Pointer() *VerifyEventByUserIdRequest {
	return &p
}

type VerifyEventByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewVerifyEventByStampTaskRequestFromJson(data string) VerifyEventByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventByStampTaskRequestFromDict(dict)
}

func NewVerifyEventByStampTaskRequestFromDict(data map[string]interface{}) VerifyEventByStampTaskRequest {
	return VerifyEventByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyEventByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyEventByStampTaskRequest) Pointer() *VerifyEventByStampTaskRequest {
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

type GetCurrentEventMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentEventMasterRequestFromJson(data string) GetCurrentEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentEventMasterRequestFromDict(dict)
}

func NewGetCurrentEventMasterRequestFromDict(data map[string]interface{}) GetCurrentEventMasterRequest {
	return GetCurrentEventMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentEventMasterRequest) Pointer() *GetCurrentEventMasterRequest {
	return &p
}

type UpdateCurrentEventMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentEventMasterRequestFromJson(data string) UpdateCurrentEventMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEventMasterRequestFromDict(dict)
}

func NewUpdateCurrentEventMasterRequestFromDict(data map[string]interface{}) UpdateCurrentEventMasterRequest {
	return UpdateCurrentEventMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentEventMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentEventMasterRequest) Pointer() *UpdateCurrentEventMasterRequest {
	return &p
}

type UpdateCurrentEventMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentEventMasterFromGitHubRequestFromJson(data string) UpdateCurrentEventMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEventMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentEventMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentEventMasterFromGitHubRequest {
	return UpdateCurrentEventMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentEventMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentEventMasterFromGitHubRequest) Pointer() *UpdateCurrentEventMasterFromGitHubRequest {
	return &p
}
