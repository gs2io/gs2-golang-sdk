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

package inbox

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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	TransactionSetting         *TransactionSetting  `json:"transactionSetting"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		IsAutomaticDeletingEnabled: core.CastBool(data["isAutomaticDeletingEnabled"]),
		TransactionSetting:         NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ReceiveMessageScript:       NewScriptSettingFromDict(core.CastMap(data["receiveMessageScript"])).Pointer(),
		ReadMessageScript:          NewScriptSettingFromDict(core.CastMap(data["readMessageScript"])).Pointer(),
		DeleteMessageScript:        NewScriptSettingFromDict(core.CastMap(data["deleteMessageScript"])).Pointer(),
		ReceiveNotification:        NewNotificationSettingFromDict(core.CastMap(data["receiveNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:           core.CastString(data["queueNamespaceId"]),
		KeyId:                      core.CastString(data["keyId"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                       p.Name,
		"description":                p.Description,
		"isAutomaticDeletingEnabled": p.IsAutomaticDeletingEnabled,
		"transactionSetting":         p.TransactionSetting.ToDict(),
		"receiveMessageScript":       p.ReceiveMessageScript.ToDict(),
		"readMessageScript":          p.ReadMessageScript.ToDict(),
		"deleteMessageScript":        p.DeleteMessageScript.ToDict(),
		"receiveNotification":        p.ReceiveNotification.ToDict(),
		"logSetting":                 p.LogSetting.ToDict(),
		"queueNamespaceId":           p.QueueNamespaceId,
		"keyId":                      p.KeyId,
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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	NamespaceName              *string              `json:"namespaceName"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	TransactionSetting         *TransactionSetting  `json:"transactionSetting"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		Description:                core.CastString(data["description"]),
		IsAutomaticDeletingEnabled: core.CastBool(data["isAutomaticDeletingEnabled"]),
		TransactionSetting:         NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ReceiveMessageScript:       NewScriptSettingFromDict(core.CastMap(data["receiveMessageScript"])).Pointer(),
		ReadMessageScript:          NewScriptSettingFromDict(core.CastMap(data["readMessageScript"])).Pointer(),
		DeleteMessageScript:        NewScriptSettingFromDict(core.CastMap(data["deleteMessageScript"])).Pointer(),
		ReceiveNotification:        NewNotificationSettingFromDict(core.CastMap(data["receiveNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:           core.CastString(data["queueNamespaceId"]),
		KeyId:                      core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":              p.NamespaceName,
		"description":                p.Description,
		"isAutomaticDeletingEnabled": p.IsAutomaticDeletingEnabled,
		"transactionSetting":         p.TransactionSetting.ToDict(),
		"receiveMessageScript":       p.ReceiveMessageScript.ToDict(),
		"readMessageScript":          p.ReadMessageScript.ToDict(),
		"deleteMessageScript":        p.DeleteMessageScript.ToDict(),
		"receiveNotification":        p.ReceiveNotification.ToDict(),
		"logSetting":                 p.LogSetting.ToDict(),
		"queueNamespaceId":           p.QueueNamespaceId,
		"keyId":                      p.KeyId,
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

type DescribeMessagesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	IsRead          *bool   `json:"isRead"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMessagesRequestFromJson(data string) DescribeMessagesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesRequestFromDict(dict)
}

func NewDescribeMessagesRequestFromDict(data map[string]interface{}) DescribeMessagesRequest {
	return DescribeMessagesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		IsRead:        core.CastBool(data["isRead"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMessagesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"isRead":        p.IsRead,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMessagesRequest) Pointer() *DescribeMessagesRequest {
	return &p
}

type DescribeMessagesByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	IsRead          *bool   `json:"isRead"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeMessagesByUserIdRequestFromJson(data string) DescribeMessagesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesByUserIdRequestFromDict(dict)
}

func NewDescribeMessagesByUserIdRequestFromDict(data map[string]interface{}) DescribeMessagesByUserIdRequest {
	return DescribeMessagesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		IsRead:          core.CastBool(data["isRead"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeMessagesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"isRead":          p.IsRead,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeMessagesByUserIdRequest) Pointer() *DescribeMessagesByUserIdRequest {
	return &p
}

type SendMessageByUserIdRequest struct {
	SourceRequestId    *string         `json:"sourceRequestId"`
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresAt          *int64          `json:"expiresAt"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	TimeOffsetToken    *string         `json:"timeOffsetToken"`
}

func NewSendMessageByUserIdRequestFromJson(data string) SendMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendMessageByUserIdRequestFromDict(dict)
}

func NewSendMessageByUserIdRequestFromDict(data map[string]interface{}) SendMessageByUserIdRequest {
	return SendMessageByUserIdRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		UserId:             core.CastString(data["userId"]),
		Metadata:           core.CastString(data["metadata"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
		ExpiresTimeSpan:    NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer(),
		TimeOffsetToken:    core.CastString(data["timeOffsetToken"]),
	}
}

func (p SendMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"metadata":      p.Metadata,
		"readAcquireActions": CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		),
		"expiresAt":       p.ExpiresAt,
		"expiresTimeSpan": p.ExpiresTimeSpan.ToDict(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SendMessageByUserIdRequest) Pointer() *SendMessageByUserIdRequest {
	return &p
}

type GetMessageRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	MessageName     *string `json:"messageName"`
}

func NewGetMessageRequestFromJson(data string) GetMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageRequestFromDict(dict)
}

func NewGetMessageRequestFromDict(data map[string]interface{}) GetMessageRequest {
	return GetMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p GetMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"messageName":   p.MessageName,
	}
}

func (p GetMessageRequest) Pointer() *GetMessageRequest {
	return &p
}

type GetMessageByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	MessageName     *string `json:"messageName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetMessageByUserIdRequestFromJson(data string) GetMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageByUserIdRequestFromDict(dict)
}

func NewGetMessageByUserIdRequestFromDict(data map[string]interface{}) GetMessageByUserIdRequest {
	return GetMessageByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MessageName:     core.CastString(data["messageName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"messageName":     p.MessageName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetMessageByUserIdRequest) Pointer() *GetMessageByUserIdRequest {
	return &p
}

type ReceiveGlobalMessageRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
}

func NewReceiveGlobalMessageRequestFromJson(data string) ReceiveGlobalMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalMessageRequestFromDict(dict)
}

func NewReceiveGlobalMessageRequestFromDict(data map[string]interface{}) ReceiveGlobalMessageRequest {
	return ReceiveGlobalMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p ReceiveGlobalMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p ReceiveGlobalMessageRequest) Pointer() *ReceiveGlobalMessageRequest {
	return &p
}

type ReceiveGlobalMessageByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewReceiveGlobalMessageByUserIdRequestFromJson(data string) ReceiveGlobalMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalMessageByUserIdRequestFromDict(dict)
}

func NewReceiveGlobalMessageByUserIdRequestFromDict(data map[string]interface{}) ReceiveGlobalMessageByUserIdRequest {
	return ReceiveGlobalMessageByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ReceiveGlobalMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ReceiveGlobalMessageByUserIdRequest) Pointer() *ReceiveGlobalMessageByUserIdRequest {
	return &p
}

type OpenMessageRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MessageName        *string `json:"messageName"`
}

func NewOpenMessageRequestFromJson(data string) OpenMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenMessageRequestFromDict(dict)
}

func NewOpenMessageRequestFromDict(data map[string]interface{}) OpenMessageRequest {
	return OpenMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p OpenMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"messageName":   p.MessageName,
	}
}

func (p OpenMessageRequest) Pointer() *OpenMessageRequest {
	return &p
}

type OpenMessageByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MessageName        *string `json:"messageName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewOpenMessageByUserIdRequestFromJson(data string) OpenMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenMessageByUserIdRequestFromDict(dict)
}

func NewOpenMessageByUserIdRequestFromDict(data map[string]interface{}) OpenMessageByUserIdRequest {
	return OpenMessageByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MessageName:     core.CastString(data["messageName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p OpenMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"messageName":     p.MessageName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p OpenMessageByUserIdRequest) Pointer() *OpenMessageByUserIdRequest {
	return &p
}

type ReadMessageRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	MessageName        *string  `json:"messageName"`
	Config             []Config `json:"config"`
}

func NewReadMessageRequestFromJson(data string) ReadMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReadMessageRequestFromDict(dict)
}

func NewReadMessageRequestFromDict(data map[string]interface{}) ReadMessageRequest {
	return ReadMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MessageName:   core.CastString(data["messageName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReadMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"messageName":   p.MessageName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReadMessageRequest) Pointer() *ReadMessageRequest {
	return &p
}

type ReadMessageByUserIdRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	MessageName        *string  `json:"messageName"`
	Config             []Config `json:"config"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
}

func NewReadMessageByUserIdRequestFromJson(data string) ReadMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReadMessageByUserIdRequestFromDict(dict)
}

func NewReadMessageByUserIdRequestFromDict(data map[string]interface{}) ReadMessageByUserIdRequest {
	return ReadMessageByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MessageName:     core.CastString(data["messageName"]),
		Config:          CastConfigs(core.CastArray(data["config"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ReadMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"messageName":   p.MessageName,
		"config": CastConfigsFromDict(
			p.Config,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ReadMessageByUserIdRequest) Pointer() *ReadMessageByUserIdRequest {
	return &p
}

type DeleteMessageRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	MessageName        *string `json:"messageName"`
}

func NewDeleteMessageRequestFromJson(data string) DeleteMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageRequestFromDict(dict)
}

func NewDeleteMessageRequestFromDict(data map[string]interface{}) DeleteMessageRequest {
	return DeleteMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p DeleteMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"messageName":   p.MessageName,
	}
}

func (p DeleteMessageRequest) Pointer() *DeleteMessageRequest {
	return &p
}

type DeleteMessageByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MessageName        *string `json:"messageName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteMessageByUserIdRequestFromJson(data string) DeleteMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageByUserIdRequestFromDict(dict)
}

func NewDeleteMessageByUserIdRequestFromDict(data map[string]interface{}) DeleteMessageByUserIdRequest {
	return DeleteMessageByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		MessageName:     core.CastString(data["messageName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"messageName":     p.MessageName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteMessageByUserIdRequest) Pointer() *DeleteMessageByUserIdRequest {
	return &p
}

type SendByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewSendByStampSheetRequestFromJson(data string) SendByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendByStampSheetRequestFromDict(dict)
}

func NewSendByStampSheetRequestFromDict(data map[string]interface{}) SendByStampSheetRequest {
	return SendByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SendByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SendByStampSheetRequest) Pointer() *SendByStampSheetRequest {
	return &p
}

type OpenByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewOpenByStampTaskRequestFromJson(data string) OpenByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenByStampTaskRequestFromDict(dict)
}

func NewOpenByStampTaskRequestFromDict(data map[string]interface{}) OpenByStampTaskRequest {
	return OpenByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p OpenByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p OpenByStampTaskRequest) Pointer() *OpenByStampTaskRequest {
	return &p
}

type DeleteMessageByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewDeleteMessageByStampTaskRequestFromJson(data string) DeleteMessageByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageByStampTaskRequestFromDict(dict)
}

func NewDeleteMessageByStampTaskRequestFromDict(data map[string]interface{}) DeleteMessageByStampTaskRequest {
	return DeleteMessageByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DeleteMessageByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DeleteMessageByStampTaskRequest) Pointer() *DeleteMessageByStampTaskRequest {
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

type GetCurrentMessageMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentMessageMasterRequestFromJson(data string) GetCurrentMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentMessageMasterRequestFromDict(dict)
}

func NewGetCurrentMessageMasterRequestFromDict(data map[string]interface{}) GetCurrentMessageMasterRequest {
	return GetCurrentMessageMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentMessageMasterRequest) Pointer() *GetCurrentMessageMasterRequest {
	return &p
}

type UpdateCurrentMessageMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentMessageMasterRequestFromJson(data string) UpdateCurrentMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMessageMasterRequestFromDict(dict)
}

func NewUpdateCurrentMessageMasterRequestFromDict(data map[string]interface{}) UpdateCurrentMessageMasterRequest {
	return UpdateCurrentMessageMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentMessageMasterRequest) Pointer() *UpdateCurrentMessageMasterRequest {
	return &p
}

type UpdateCurrentMessageMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentMessageMasterFromGitHubRequestFromJson(data string) UpdateCurrentMessageMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMessageMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentMessageMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentMessageMasterFromGitHubRequest {
	return UpdateCurrentMessageMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentMessageMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentMessageMasterFromGitHubRequest) Pointer() *UpdateCurrentMessageMasterFromGitHubRequest {
	return &p
}

type DescribeGlobalMessageMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeGlobalMessageMastersRequestFromJson(data string) DescribeGlobalMessageMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalMessageMastersRequestFromDict(dict)
}

func NewDescribeGlobalMessageMastersRequestFromDict(data map[string]interface{}) DescribeGlobalMessageMastersRequest {
	return DescribeGlobalMessageMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeGlobalMessageMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGlobalMessageMastersRequest) Pointer() *DescribeGlobalMessageMastersRequest {
	return &p
}

type CreateGlobalMessageMasterRequest struct {
	SourceRequestId    *string         `json:"sourceRequestId"`
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	NamespaceName      *string         `json:"namespaceName"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	ExpiresAt          *int64          `json:"expiresAt"`
}

func NewCreateGlobalMessageMasterRequestFromJson(data string) CreateGlobalMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalMessageMasterRequestFromDict(dict)
}

func NewCreateGlobalMessageMasterRequestFromDict(data map[string]interface{}) CreateGlobalMessageMasterRequest {
	return CreateGlobalMessageMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ExpiresTimeSpan:    NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer(),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
	}
}

func (p CreateGlobalMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"metadata":      p.Metadata,
		"readAcquireActions": CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		),
		"expiresTimeSpan": p.ExpiresTimeSpan.ToDict(),
		"expiresAt":       p.ExpiresAt,
	}
}

func (p CreateGlobalMessageMasterRequest) Pointer() *CreateGlobalMessageMasterRequest {
	return &p
}

type GetGlobalMessageMasterRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	GlobalMessageName *string `json:"globalMessageName"`
}

func NewGetGlobalMessageMasterRequestFromJson(data string) GetGlobalMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalMessageMasterRequestFromDict(dict)
}

func NewGetGlobalMessageMasterRequestFromDict(data map[string]interface{}) GetGlobalMessageMasterRequest {
	return GetGlobalMessageMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		GlobalMessageName: core.CastString(data["globalMessageName"]),
	}
}

func (p GetGlobalMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"globalMessageName": p.GlobalMessageName,
	}
}

func (p GetGlobalMessageMasterRequest) Pointer() *GetGlobalMessageMasterRequest {
	return &p
}

type UpdateGlobalMessageMasterRequest struct {
	SourceRequestId    *string         `json:"sourceRequestId"`
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	NamespaceName      *string         `json:"namespaceName"`
	GlobalMessageName  *string         `json:"globalMessageName"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	ExpiresAt          *int64          `json:"expiresAt"`
}

func NewUpdateGlobalMessageMasterRequestFromJson(data string) UpdateGlobalMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGlobalMessageMasterRequestFromDict(dict)
}

func NewUpdateGlobalMessageMasterRequestFromDict(data map[string]interface{}) UpdateGlobalMessageMasterRequest {
	return UpdateGlobalMessageMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		GlobalMessageName:  core.CastString(data["globalMessageName"]),
		Metadata:           core.CastString(data["metadata"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ExpiresTimeSpan:    NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer(),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
	}
}

func (p UpdateGlobalMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"globalMessageName": p.GlobalMessageName,
		"metadata":          p.Metadata,
		"readAcquireActions": CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		),
		"expiresTimeSpan": p.ExpiresTimeSpan.ToDict(),
		"expiresAt":       p.ExpiresAt,
	}
}

func (p UpdateGlobalMessageMasterRequest) Pointer() *UpdateGlobalMessageMasterRequest {
	return &p
}

type DeleteGlobalMessageMasterRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	GlobalMessageName *string `json:"globalMessageName"`
}

func NewDeleteGlobalMessageMasterRequestFromJson(data string) DeleteGlobalMessageMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGlobalMessageMasterRequestFromDict(dict)
}

func NewDeleteGlobalMessageMasterRequestFromDict(data map[string]interface{}) DeleteGlobalMessageMasterRequest {
	return DeleteGlobalMessageMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		GlobalMessageName: core.CastString(data["globalMessageName"]),
	}
}

func (p DeleteGlobalMessageMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"globalMessageName": p.GlobalMessageName,
	}
}

func (p DeleteGlobalMessageMasterRequest) Pointer() *DeleteGlobalMessageMasterRequest {
	return &p
}

type DescribeGlobalMessagesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeGlobalMessagesRequestFromJson(data string) DescribeGlobalMessagesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalMessagesRequestFromDict(dict)
}

func NewDescribeGlobalMessagesRequestFromDict(data map[string]interface{}) DescribeGlobalMessagesRequest {
	return DescribeGlobalMessagesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeGlobalMessagesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGlobalMessagesRequest) Pointer() *DescribeGlobalMessagesRequest {
	return &p
}

type GetGlobalMessageRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	GlobalMessageName *string `json:"globalMessageName"`
}

func NewGetGlobalMessageRequestFromJson(data string) GetGlobalMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalMessageRequestFromDict(dict)
}

func NewGetGlobalMessageRequestFromDict(data map[string]interface{}) GetGlobalMessageRequest {
	return GetGlobalMessageRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		GlobalMessageName: core.CastString(data["globalMessageName"]),
	}
}

func (p GetGlobalMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"globalMessageName": p.GlobalMessageName,
	}
}

func (p GetGlobalMessageRequest) Pointer() *GetGlobalMessageRequest {
	return &p
}

type GetReceivedByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetReceivedByUserIdRequestFromJson(data string) GetReceivedByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceivedByUserIdRequestFromDict(dict)
}

func NewGetReceivedByUserIdRequestFromDict(data map[string]interface{}) GetReceivedByUserIdRequest {
	return GetReceivedByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetReceivedByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetReceivedByUserIdRequest) Pointer() *GetReceivedByUserIdRequest {
	return &p
}

type UpdateReceivedByUserIdRequest struct {
	SourceRequestId            *string   `json:"sourceRequestId"`
	RequestId                  *string   `json:"requestId"`
	ContextStack               *string   `json:"contextStack"`
	DuplicationAvoider         *string   `json:"duplicationAvoider"`
	NamespaceName              *string   `json:"namespaceName"`
	UserId                     *string   `json:"userId"`
	ReceivedGlobalMessageNames []*string `json:"receivedGlobalMessageNames"`
	TimeOffsetToken            *string   `json:"timeOffsetToken"`
}

func NewUpdateReceivedByUserIdRequestFromJson(data string) UpdateReceivedByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateReceivedByUserIdRequestFromDict(dict)
}

func NewUpdateReceivedByUserIdRequestFromDict(data map[string]interface{}) UpdateReceivedByUserIdRequest {
	return UpdateReceivedByUserIdRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		UserId:                     core.CastString(data["userId"]),
		ReceivedGlobalMessageNames: core.CastStrings(core.CastArray(data["receivedGlobalMessageNames"])),
		TimeOffsetToken:            core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateReceivedByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"receivedGlobalMessageNames": core.CastStringsFromDict(
			p.ReceivedGlobalMessageNames,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateReceivedByUserIdRequest) Pointer() *UpdateReceivedByUserIdRequest {
	return &p
}

type DeleteReceivedByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteReceivedByUserIdRequestFromJson(data string) DeleteReceivedByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceivedByUserIdRequestFromDict(dict)
}

func NewDeleteReceivedByUserIdRequestFromDict(data map[string]interface{}) DeleteReceivedByUserIdRequest {
	return DeleteReceivedByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteReceivedByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteReceivedByUserIdRequest) Pointer() *DeleteReceivedByUserIdRequest {
	return &p
}
