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

package adReward

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
	RequestId               *string              `json:"requestId"`
	ContextStack            *string              `json:"contextStack"`
	Name                    *string              `json:"name"`
	Admob                   *AdMob               `json:"admob"`
	UnityAd                 *UnityAd             `json:"unityAd"`
	Description             *string              `json:"description"`
	ChangePointNotification *NotificationSetting `json:"changePointNotification"`
	LogSetting              *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                    core.CastString(data["name"]),
		Admob:                   NewAdMobFromDict(core.CastMap(data["admob"])).Pointer(),
		UnityAd:                 NewUnityAdFromDict(core.CastMap(data["unityAd"])).Pointer(),
		Description:             core.CastString(data["description"]),
		ChangePointNotification: NewNotificationSettingFromDict(core.CastMap(data["changePointNotification"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                    p.Name,
		"admob":                   p.Admob.ToDict(),
		"unityAd":                 p.UnityAd.ToDict(),
		"description":             p.Description,
		"changePointNotification": p.ChangePointNotification.ToDict(),
		"logSetting":              p.LogSetting.ToDict(),
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
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
	RequestId               *string              `json:"requestId"`
	ContextStack            *string              `json:"contextStack"`
	NamespaceName           *string              `json:"namespaceName"`
	Description             *string              `json:"description"`
	Admob                   *AdMob               `json:"admob"`
	UnityAd                 *UnityAd             `json:"unityAd"`
	ChangePointNotification *NotificationSetting `json:"changePointNotification"`
	LogSetting              *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		Description:             core.CastString(data["description"]),
		Admob:                   NewAdMobFromDict(core.CastMap(data["admob"])).Pointer(),
		UnityAd:                 NewUnityAdFromDict(core.CastMap(data["unityAd"])).Pointer(),
		ChangePointNotification: NewNotificationSettingFromDict(core.CastMap(data["changePointNotification"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"description":             p.Description,
		"admob":                   p.Admob.ToDict(),
		"unityAd":                 p.UnityAd.ToDict(),
		"changePointNotification": p.ChangePointNotification.ToDict(),
		"logSetting":              p.LogSetting.ToDict(),
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

type GetPointRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
}

func NewGetPointRequestFromJson(data string) GetPointRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPointRequestFromDict(dict)
}

func NewGetPointRequestFromDict(data map[string]interface{}) GetPointRequest {
	return GetPointRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetPointRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetPointRequest) Pointer() *GetPointRequest {
	return &p
}

type GetPointByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
}

func NewGetPointByUserIdRequestFromJson(data string) GetPointByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPointByUserIdRequestFromDict(dict)
}

func NewGetPointByUserIdRequestFromDict(data map[string]interface{}) GetPointByUserIdRequest {
	return GetPointByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetPointByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p GetPointByUserIdRequest) Pointer() *GetPointByUserIdRequest {
	return &p
}

type AcquirePointByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Point              *int64  `json:"point"`
}

func NewAcquirePointByUserIdRequestFromJson(data string) AcquirePointByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquirePointByUserIdRequestFromDict(dict)
}

func NewAcquirePointByUserIdRequestFromDict(data map[string]interface{}) AcquirePointByUserIdRequest {
	return AcquirePointByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Point:         core.CastInt64(data["point"]),
	}
}

func (p AcquirePointByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"point":         p.Point,
	}
}

func (p AcquirePointByUserIdRequest) Pointer() *AcquirePointByUserIdRequest {
	return &p
}

type ConsumePointRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Point              *int64  `json:"point"`
}

func NewConsumePointRequestFromJson(data string) ConsumePointRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointRequestFromDict(dict)
}

func NewConsumePointRequestFromDict(data map[string]interface{}) ConsumePointRequest {
	return ConsumePointRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Point:         core.CastInt64(data["point"]),
	}
}

func (p ConsumePointRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"point":         p.Point,
	}
}

func (p ConsumePointRequest) Pointer() *ConsumePointRequest {
	return &p
}

type ConsumePointByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Point              *int64  `json:"point"`
}

func NewConsumePointByUserIdRequestFromJson(data string) ConsumePointByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointByUserIdRequestFromDict(dict)
}

func NewConsumePointByUserIdRequestFromDict(data map[string]interface{}) ConsumePointByUserIdRequest {
	return ConsumePointByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Point:         core.CastInt64(data["point"]),
	}
}

func (p ConsumePointByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"point":         p.Point,
	}
}

func (p ConsumePointByUserIdRequest) Pointer() *ConsumePointByUserIdRequest {
	return &p
}

type DeletePointByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewDeletePointByUserIdRequestFromJson(data string) DeletePointByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePointByUserIdRequestFromDict(dict)
}

func NewDeletePointByUserIdRequestFromDict(data map[string]interface{}) DeletePointByUserIdRequest {
	return DeletePointByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeletePointByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p DeletePointByUserIdRequest) Pointer() *DeletePointByUserIdRequest {
	return &p
}

type ConsumePointByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewConsumePointByStampTaskRequestFromJson(data string) ConsumePointByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointByStampTaskRequestFromDict(dict)
}

func NewConsumePointByStampTaskRequestFromDict(data map[string]interface{}) ConsumePointByStampTaskRequest {
	return ConsumePointByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumePointByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumePointByStampTaskRequest) Pointer() *ConsumePointByStampTaskRequest {
	return &p
}

type AcquirePointByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewAcquirePointByStampSheetRequestFromJson(data string) AcquirePointByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquirePointByStampSheetRequestFromDict(dict)
}

func NewAcquirePointByStampSheetRequestFromDict(data map[string]interface{}) AcquirePointByStampSheetRequest {
	return AcquirePointByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquirePointByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquirePointByStampSheetRequest) Pointer() *AcquirePointByStampSheetRequest {
	return &p
}
