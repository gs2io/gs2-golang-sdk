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

package version

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	RequestId                   *string        `json:"requestId"`
	ContextStack                *string        `json:"contextStack"`
	DuplicationAvoider          *string        `json:"duplicationAvoider"`
	Name                        *string        `json:"name"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		AssumeUserId:                core.CastString(data["assumeUserId"]),
		AcceptVersionScript:         NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer(),
		CheckVersionTriggerScriptId: core.CastString(data["checkVersionTriggerScriptId"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                        p.Name,
		"description":                 p.Description,
		"assumeUserId":                p.AssumeUserId,
		"acceptVersionScript":         p.AcceptVersionScript.ToDict(),
		"checkVersionTriggerScriptId": p.CheckVersionTriggerScriptId,
		"logSetting":                  p.LogSetting.ToDict(),
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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
	RequestId                   *string        `json:"requestId"`
	ContextStack                *string        `json:"contextStack"`
	DuplicationAvoider          *string        `json:"duplicationAvoider"`
	NamespaceName               *string        `json:"namespaceName"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:               core.CastString(data["namespaceName"]),
		Description:                 core.CastString(data["description"]),
		AssumeUserId:                core.CastString(data["assumeUserId"]),
		AcceptVersionScript:         NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer(),
		CheckVersionTriggerScriptId: core.CastString(data["checkVersionTriggerScriptId"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":               p.NamespaceName,
		"description":                 p.Description,
		"assumeUserId":                p.AssumeUserId,
		"acceptVersionScript":         p.AcceptVersionScript.ToDict(),
		"checkVersionTriggerScriptId": p.CheckVersionTriggerScriptId,
		"logSetting":                  p.LogSetting.ToDict(),
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

type DescribeVersionModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeVersionModelMastersRequestFromJson(data string) DescribeVersionModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeVersionModelMastersRequestFromDict(dict)
}

func NewDescribeVersionModelMastersRequestFromDict(data map[string]interface{}) DescribeVersionModelMastersRequest {
	return DescribeVersionModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeVersionModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeVersionModelMastersRequest) Pointer() *DescribeVersionModelMastersRequest {
	return &p
}

type CreateVersionModelMasterRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	Name               *string  `json:"name"`
	Description        *string  `json:"description"`
	Metadata           *string  `json:"metadata"`
	WarningVersion     *Version `json:"warningVersion"`
	ErrorVersion       *Version `json:"errorVersion"`
	Scope              *string  `json:"scope"`
	CurrentVersion     *Version `json:"currentVersion"`
	NeedSignature      *bool    `json:"needSignature"`
	SignatureKeyId     *string  `json:"signatureKeyId"`
}

func NewCreateVersionModelMasterRequestFromJson(data string) CreateVersionModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateVersionModelMasterRequestFromDict(dict)
}

func NewCreateVersionModelMasterRequestFromDict(data map[string]interface{}) CreateVersionModelMasterRequest {
	return CreateVersionModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		WarningVersion: NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer(),
		ErrorVersion:   NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer(),
		Scope:          core.CastString(data["scope"]),
		CurrentVersion: NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer(),
		NeedSignature:  core.CastBool(data["needSignature"]),
		SignatureKeyId: core.CastString(data["signatureKeyId"]),
	}
}

func (p CreateVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"name":           p.Name,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"warningVersion": p.WarningVersion.ToDict(),
		"errorVersion":   p.ErrorVersion.ToDict(),
		"scope":          p.Scope,
		"currentVersion": p.CurrentVersion.ToDict(),
		"needSignature":  p.NeedSignature,
		"signatureKeyId": p.SignatureKeyId,
	}
}

func (p CreateVersionModelMasterRequest) Pointer() *CreateVersionModelMasterRequest {
	return &p
}

type GetVersionModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	VersionName        *string `json:"versionName"`
}

func NewGetVersionModelMasterRequestFromJson(data string) GetVersionModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetVersionModelMasterRequestFromDict(dict)
}

func NewGetVersionModelMasterRequestFromDict(data map[string]interface{}) GetVersionModelMasterRequest {
	return GetVersionModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p GetVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p GetVersionModelMasterRequest) Pointer() *GetVersionModelMasterRequest {
	return &p
}

type UpdateVersionModelMasterRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	Description        *string  `json:"description"`
	Metadata           *string  `json:"metadata"`
	WarningVersion     *Version `json:"warningVersion"`
	ErrorVersion       *Version `json:"errorVersion"`
	Scope              *string  `json:"scope"`
	CurrentVersion     *Version `json:"currentVersion"`
	NeedSignature      *bool    `json:"needSignature"`
	SignatureKeyId     *string  `json:"signatureKeyId"`
}

func NewUpdateVersionModelMasterRequestFromJson(data string) UpdateVersionModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateVersionModelMasterRequestFromDict(dict)
}

func NewUpdateVersionModelMasterRequestFromDict(data map[string]interface{}) UpdateVersionModelMasterRequest {
	return UpdateVersionModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		VersionName:    core.CastString(data["versionName"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		WarningVersion: NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer(),
		ErrorVersion:   NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer(),
		Scope:          core.CastString(data["scope"]),
		CurrentVersion: NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer(),
		NeedSignature:  core.CastBool(data["needSignature"]),
		SignatureKeyId: core.CastString(data["signatureKeyId"]),
	}
}

func (p UpdateVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"versionName":    p.VersionName,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"warningVersion": p.WarningVersion.ToDict(),
		"errorVersion":   p.ErrorVersion.ToDict(),
		"scope":          p.Scope,
		"currentVersion": p.CurrentVersion.ToDict(),
		"needSignature":  p.NeedSignature,
		"signatureKeyId": p.SignatureKeyId,
	}
}

func (p UpdateVersionModelMasterRequest) Pointer() *UpdateVersionModelMasterRequest {
	return &p
}

type DeleteVersionModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	VersionName        *string `json:"versionName"`
}

func NewDeleteVersionModelMasterRequestFromJson(data string) DeleteVersionModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteVersionModelMasterRequestFromDict(dict)
}

func NewDeleteVersionModelMasterRequestFromDict(data map[string]interface{}) DeleteVersionModelMasterRequest {
	return DeleteVersionModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p DeleteVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p DeleteVersionModelMasterRequest) Pointer() *DeleteVersionModelMasterRequest {
	return &p
}

type DescribeVersionModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeVersionModelsRequestFromJson(data string) DescribeVersionModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeVersionModelsRequestFromDict(dict)
}

func NewDescribeVersionModelsRequestFromDict(data map[string]interface{}) DescribeVersionModelsRequest {
	return DescribeVersionModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeVersionModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeVersionModelsRequest) Pointer() *DescribeVersionModelsRequest {
	return &p
}

type GetVersionModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	VersionName        *string `json:"versionName"`
}

func NewGetVersionModelRequestFromJson(data string) GetVersionModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetVersionModelRequestFromDict(dict)
}

func NewGetVersionModelRequestFromDict(data map[string]interface{}) GetVersionModelRequest {
	return GetVersionModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p GetVersionModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p GetVersionModelRequest) Pointer() *GetVersionModelRequest {
	return &p
}

type DescribeAcceptVersionsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeAcceptVersionsRequestFromJson(data string) DescribeAcceptVersionsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAcceptVersionsRequestFromDict(dict)
}

func NewDescribeAcceptVersionsRequestFromDict(data map[string]interface{}) DescribeAcceptVersionsRequest {
	return DescribeAcceptVersionsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAcceptVersionsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAcceptVersionsRequest) Pointer() *DescribeAcceptVersionsRequest {
	return &p
}

type DescribeAcceptVersionsByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeAcceptVersionsByUserIdRequestFromJson(data string) DescribeAcceptVersionsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAcceptVersionsByUserIdRequestFromDict(dict)
}

func NewDescribeAcceptVersionsByUserIdRequestFromDict(data map[string]interface{}) DescribeAcceptVersionsByUserIdRequest {
	return DescribeAcceptVersionsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAcceptVersionsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAcceptVersionsByUserIdRequest) Pointer() *DescribeAcceptVersionsByUserIdRequest {
	return &p
}

type AcceptRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	VersionName        *string `json:"versionName"`
	AccessToken        *string `json:"accessToken"`
}

func NewAcceptRequestFromJson(data string) AcceptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestFromDict(dict)
}

func NewAcceptRequestFromDict(data map[string]interface{}) AcceptRequest {
	return AcceptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p AcceptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"accessToken":   p.AccessToken,
	}
}

func (p AcceptRequest) Pointer() *AcceptRequest {
	return &p
}

type AcceptByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	VersionName        *string `json:"versionName"`
	UserId             *string `json:"userId"`
}

func NewAcceptByUserIdRequestFromJson(data string) AcceptByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptByUserIdRequestFromDict(dict)
}

func NewAcceptByUserIdRequestFromDict(data map[string]interface{}) AcceptByUserIdRequest {
	return AcceptByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p AcceptByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"userId":        p.UserId,
	}
}

func (p AcceptByUserIdRequest) Pointer() *AcceptByUserIdRequest {
	return &p
}

type GetAcceptVersionRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	VersionName        *string `json:"versionName"`
}

func NewGetAcceptVersionRequestFromJson(data string) GetAcceptVersionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAcceptVersionRequestFromDict(dict)
}

func NewGetAcceptVersionRequestFromDict(data map[string]interface{}) GetAcceptVersionRequest {
	return GetAcceptVersionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p GetAcceptVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"versionName":   p.VersionName,
	}
}

func (p GetAcceptVersionRequest) Pointer() *GetAcceptVersionRequest {
	return &p
}

type GetAcceptVersionByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	VersionName        *string `json:"versionName"`
}

func NewGetAcceptVersionByUserIdRequestFromJson(data string) GetAcceptVersionByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAcceptVersionByUserIdRequestFromDict(dict)
}

func NewGetAcceptVersionByUserIdRequestFromDict(data map[string]interface{}) GetAcceptVersionByUserIdRequest {
	return GetAcceptVersionByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p GetAcceptVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"versionName":   p.VersionName,
	}
}

func (p GetAcceptVersionByUserIdRequest) Pointer() *GetAcceptVersionByUserIdRequest {
	return &p
}

type DeleteAcceptVersionRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	VersionName        *string `json:"versionName"`
}

func NewDeleteAcceptVersionRequestFromJson(data string) DeleteAcceptVersionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAcceptVersionRequestFromDict(dict)
}

func NewDeleteAcceptVersionRequestFromDict(data map[string]interface{}) DeleteAcceptVersionRequest {
	return DeleteAcceptVersionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p DeleteAcceptVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"versionName":   p.VersionName,
	}
}

func (p DeleteAcceptVersionRequest) Pointer() *DeleteAcceptVersionRequest {
	return &p
}

type DeleteAcceptVersionByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	VersionName        *string `json:"versionName"`
}

func NewDeleteAcceptVersionByUserIdRequestFromJson(data string) DeleteAcceptVersionByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAcceptVersionByUserIdRequestFromDict(dict)
}

func NewDeleteAcceptVersionByUserIdRequestFromDict(data map[string]interface{}) DeleteAcceptVersionByUserIdRequest {
	return DeleteAcceptVersionByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		VersionName:   core.CastString(data["versionName"]),
	}
}

func (p DeleteAcceptVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"versionName":   p.VersionName,
	}
}

func (p DeleteAcceptVersionByUserIdRequest) Pointer() *DeleteAcceptVersionByUserIdRequest {
	return &p
}

type CheckVersionRequest struct {
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	AccessToken        *string         `json:"accessToken"`
	TargetVersions     []TargetVersion `json:"targetVersions"`
}

func NewCheckVersionRequestFromJson(data string) CheckVersionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckVersionRequestFromDict(dict)
}

func NewCheckVersionRequestFromDict(data map[string]interface{}) CheckVersionRequest {
	return CheckVersionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		TargetVersions: CastTargetVersions(core.CastArray(data["targetVersions"])),
	}
}

func (p CheckVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"targetVersions": CastTargetVersionsFromDict(
			p.TargetVersions,
		),
	}
}

func (p CheckVersionRequest) Pointer() *CheckVersionRequest {
	return &p
}

type CheckVersionByUserIdRequest struct {
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	TargetVersions     []TargetVersion `json:"targetVersions"`
}

func NewCheckVersionByUserIdRequestFromJson(data string) CheckVersionByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckVersionByUserIdRequestFromDict(dict)
}

func NewCheckVersionByUserIdRequestFromDict(data map[string]interface{}) CheckVersionByUserIdRequest {
	return CheckVersionByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		TargetVersions: CastTargetVersions(core.CastArray(data["targetVersions"])),
	}
}

func (p CheckVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"targetVersions": CastTargetVersionsFromDict(
			p.TargetVersions,
		),
	}
}

func (p CheckVersionByUserIdRequest) Pointer() *CheckVersionByUserIdRequest {
	return &p
}

type CalculateSignatureRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	Version            *Version `json:"version"`
}

func NewCalculateSignatureRequestFromJson(data string) CalculateSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCalculateSignatureRequestFromDict(dict)
}

func NewCalculateSignatureRequestFromDict(data map[string]interface{}) CalculateSignatureRequest {
	return CalculateSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
		Version:       NewVersionFromDict(core.CastMap(data["version"])).Pointer(),
	}
}

func (p CalculateSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"version":       p.Version.ToDict(),
	}
}

func (p CalculateSignatureRequest) Pointer() *CalculateSignatureRequest {
	return &p
}

type ExportMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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

type GetCurrentVersionMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetCurrentVersionMasterRequestFromJson(data string) GetCurrentVersionMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentVersionMasterRequestFromDict(dict)
}

func NewGetCurrentVersionMasterRequestFromDict(data map[string]interface{}) GetCurrentVersionMasterRequest {
	return GetCurrentVersionMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentVersionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentVersionMasterRequest) Pointer() *GetCurrentVersionMasterRequest {
	return &p
}

type UpdateCurrentVersionMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
}

func NewUpdateCurrentVersionMasterRequestFromJson(data string) UpdateCurrentVersionMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentVersionMasterRequestFromDict(dict)
}

func NewUpdateCurrentVersionMasterRequestFromDict(data map[string]interface{}) UpdateCurrentVersionMasterRequest {
	return UpdateCurrentVersionMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentVersionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentVersionMasterRequest) Pointer() *UpdateCurrentVersionMasterRequest {
	return &p
}

type UpdateCurrentVersionMasterFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentVersionMasterFromGitHubRequestFromJson(data string) UpdateCurrentVersionMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentVersionMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentVersionMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentVersionMasterFromGitHubRequest {
	return UpdateCurrentVersionMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentVersionMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentVersionMasterFromGitHubRequest) Pointer() *UpdateCurrentVersionMasterFromGitHubRequest {
	return &p
}
