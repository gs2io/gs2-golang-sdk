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

package lock

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

type DescribeMutexesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMutexesRequestFromJson(data string) DescribeMutexesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMutexesRequestFromDict(dict)
}

func NewDescribeMutexesRequestFromDict(data map[string]interface{}) DescribeMutexesRequest {
	return DescribeMutexesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMutexesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMutexesRequest) Pointer() *DescribeMutexesRequest {
	return &p
}

type DescribeMutexesByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeMutexesByUserIdRequestFromJson(data string) DescribeMutexesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMutexesByUserIdRequestFromDict(dict)
}

func NewDescribeMutexesByUserIdRequestFromDict(data map[string]interface{}) DescribeMutexesByUserIdRequest {
	return DescribeMutexesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeMutexesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeMutexesByUserIdRequest) Pointer() *DescribeMutexesByUserIdRequest {
	return &p
}

type LockRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PropertyId         *string `json:"propertyId"`
	AccessToken        *string `json:"accessToken"`
	TransactionId      *string `json:"transactionId"`
	Ttl                *int64  `json:"ttl"`
}

func NewLockRequestFromJson(data string) LockRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLockRequestFromDict(dict)
}

func NewLockRequestFromDict(data map[string]interface{}) LockRequest {
	return LockRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TransactionId: core.CastString(data["transactionId"]),
		Ttl:           core.CastInt64(data["ttl"]),
	}
}

func (p LockRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"propertyId":    p.PropertyId,
		"accessToken":   p.AccessToken,
		"transactionId": p.TransactionId,
		"ttl":           p.Ttl,
	}
}

func (p LockRequest) Pointer() *LockRequest {
	return &p
}

type LockByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PropertyId         *string `json:"propertyId"`
	UserId             *string `json:"userId"`
	TransactionId      *string `json:"transactionId"`
	Ttl                *int64  `json:"ttl"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewLockByUserIdRequestFromJson(data string) LockByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLockByUserIdRequestFromDict(dict)
}

func NewLockByUserIdRequestFromDict(data map[string]interface{}) LockByUserIdRequest {
	return LockByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		UserId:          core.CastString(data["userId"]),
		TransactionId:   core.CastString(data["transactionId"]),
		Ttl:             core.CastInt64(data["ttl"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p LockByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"propertyId":      p.PropertyId,
		"userId":          p.UserId,
		"transactionId":   p.TransactionId,
		"ttl":             p.Ttl,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p LockByUserIdRequest) Pointer() *LockByUserIdRequest {
	return &p
}

type UnlockRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PropertyId         *string `json:"propertyId"`
	AccessToken        *string `json:"accessToken"`
	TransactionId      *string `json:"transactionId"`
}

func NewUnlockRequestFromJson(data string) UnlockRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockRequestFromDict(dict)
}

func NewUnlockRequestFromDict(data map[string]interface{}) UnlockRequest {
	return UnlockRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PropertyId:    core.CastString(data["propertyId"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p UnlockRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"propertyId":    p.PropertyId,
		"accessToken":   p.AccessToken,
		"transactionId": p.TransactionId,
	}
}

func (p UnlockRequest) Pointer() *UnlockRequest {
	return &p
}

type UnlockByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PropertyId         *string `json:"propertyId"`
	UserId             *string `json:"userId"`
	TransactionId      *string `json:"transactionId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUnlockByUserIdRequestFromJson(data string) UnlockByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockByUserIdRequestFromDict(dict)
}

func NewUnlockByUserIdRequestFromDict(data map[string]interface{}) UnlockByUserIdRequest {
	return UnlockByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		UserId:          core.CastString(data["userId"]),
		TransactionId:   core.CastString(data["transactionId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UnlockByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"propertyId":      p.PropertyId,
		"userId":          p.UserId,
		"transactionId":   p.TransactionId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UnlockByUserIdRequest) Pointer() *UnlockByUserIdRequest {
	return &p
}

type GetMutexRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PropertyId      *string `json:"propertyId"`
}

func NewGetMutexRequestFromJson(data string) GetMutexRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMutexRequestFromDict(dict)
}

func NewGetMutexRequestFromDict(data map[string]interface{}) GetMutexRequest {
	return GetMutexRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PropertyId:    core.CastString(data["propertyId"]),
	}
}

func (p GetMutexRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"propertyId":    p.PropertyId,
	}
}

func (p GetMutexRequest) Pointer() *GetMutexRequest {
	return &p
}

type GetMutexByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PropertyId      *string `json:"propertyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetMutexByUserIdRequestFromJson(data string) GetMutexByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMutexByUserIdRequestFromDict(dict)
}

func NewGetMutexByUserIdRequestFromDict(data map[string]interface{}) GetMutexByUserIdRequest {
	return GetMutexByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PropertyId:      core.CastString(data["propertyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetMutexByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetMutexByUserIdRequest) Pointer() *GetMutexByUserIdRequest {
	return &p
}

type DeleteMutexByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PropertyId         *string `json:"propertyId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteMutexByUserIdRequestFromJson(data string) DeleteMutexByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMutexByUserIdRequestFromDict(dict)
}

func NewDeleteMutexByUserIdRequestFromDict(data map[string]interface{}) DeleteMutexByUserIdRequest {
	return DeleteMutexByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PropertyId:      core.CastString(data["propertyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteMutexByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteMutexByUserIdRequest) Pointer() *DeleteMutexByUserIdRequest {
	return &p
}
