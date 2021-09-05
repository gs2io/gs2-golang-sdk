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

package dictionary

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
	RequestId            *string        `json:"requestId"`
	ContextStack         *string        `json:"contextStack"`
	DuplicationAvoider   *string        `json:"duplicationAvoider"`
	Name                 *string        `json:"name"`
	Description          *string        `json:"description"`
	EntryScript          *ScriptSetting `json:"entryScript"`
	DuplicateEntryScript *ScriptSetting `json:"duplicateEntryScript"`
	LogSetting           *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		EntryScript:          NewScriptSettingFromDict(core.CastMap(data["entryScript"])).Pointer(),
		DuplicateEntryScript: NewScriptSettingFromDict(core.CastMap(data["duplicateEntryScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                 p.Name,
		"description":          p.Description,
		"entryScript":          p.EntryScript.ToDict(),
		"duplicateEntryScript": p.DuplicateEntryScript.ToDict(),
		"logSetting":           p.LogSetting.ToDict(),
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
	RequestId            *string        `json:"requestId"`
	ContextStack         *string        `json:"contextStack"`
	DuplicationAvoider   *string        `json:"duplicationAvoider"`
	NamespaceName        *string        `json:"namespaceName"`
	Description          *string        `json:"description"`
	EntryScript          *ScriptSetting `json:"entryScript"`
	DuplicateEntryScript *ScriptSetting `json:"duplicateEntryScript"`
	LogSetting           *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		Description:          core.CastString(data["description"]),
		EntryScript:          NewScriptSettingFromDict(core.CastMap(data["entryScript"])).Pointer(),
		DuplicateEntryScript: NewScriptSettingFromDict(core.CastMap(data["duplicateEntryScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"description":          p.Description,
		"entryScript":          p.EntryScript.ToDict(),
		"duplicateEntryScript": p.DuplicateEntryScript.ToDict(),
		"logSetting":           p.LogSetting.ToDict(),
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

type DescribeEntryModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeEntryModelsRequestFromJson(data string) DescribeEntryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntryModelsRequestFromDict(dict)
}

func NewDescribeEntryModelsRequestFromDict(data map[string]interface{}) DescribeEntryModelsRequest {
	return DescribeEntryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeEntryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeEntryModelsRequest) Pointer() *DescribeEntryModelsRequest {
	return &p
}

type GetEntryModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	EntryName          *string `json:"entryName"`
}

func NewGetEntryModelRequestFromJson(data string) GetEntryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryModelRequestFromDict(dict)
}

func NewGetEntryModelRequestFromDict(data map[string]interface{}) GetEntryModelRequest {
	return GetEntryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EntryName:     core.CastString(data["entryName"]),
	}
}

func (p GetEntryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"entryName":     p.EntryName,
	}
}

func (p GetEntryModelRequest) Pointer() *GetEntryModelRequest {
	return &p
}

type DescribeEntryModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeEntryModelMastersRequestFromJson(data string) DescribeEntryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntryModelMastersRequestFromDict(dict)
}

func NewDescribeEntryModelMastersRequestFromDict(data map[string]interface{}) DescribeEntryModelMastersRequest {
	return DescribeEntryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeEntryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeEntryModelMastersRequest) Pointer() *DescribeEntryModelMastersRequest {
	return &p
}

type CreateEntryModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
}

func NewCreateEntryModelMasterRequestFromJson(data string) CreateEntryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateEntryModelMasterRequestFromDict(dict)
}

func NewCreateEntryModelMasterRequestFromDict(data map[string]interface{}) CreateEntryModelMasterRequest {
	return CreateEntryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateEntryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateEntryModelMasterRequest) Pointer() *CreateEntryModelMasterRequest {
	return &p
}

type GetEntryModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	EntryName          *string `json:"entryName"`
}

func NewGetEntryModelMasterRequestFromJson(data string) GetEntryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryModelMasterRequestFromDict(dict)
}

func NewGetEntryModelMasterRequestFromDict(data map[string]interface{}) GetEntryModelMasterRequest {
	return GetEntryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EntryName:     core.CastString(data["entryName"]),
	}
}

func (p GetEntryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"entryName":     p.EntryName,
	}
}

func (p GetEntryModelMasterRequest) Pointer() *GetEntryModelMasterRequest {
	return &p
}

type UpdateEntryModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	EntryName          *string `json:"entryName"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
}

func NewUpdateEntryModelMasterRequestFromJson(data string) UpdateEntryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateEntryModelMasterRequestFromDict(dict)
}

func NewUpdateEntryModelMasterRequestFromDict(data map[string]interface{}) UpdateEntryModelMasterRequest {
	return UpdateEntryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EntryName:     core.CastString(data["entryName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateEntryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"entryName":     p.EntryName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateEntryModelMasterRequest) Pointer() *UpdateEntryModelMasterRequest {
	return &p
}

type DeleteEntryModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	EntryName          *string `json:"entryName"`
}

func NewDeleteEntryModelMasterRequestFromJson(data string) DeleteEntryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEntryModelMasterRequestFromDict(dict)
}

func NewDeleteEntryModelMasterRequestFromDict(data map[string]interface{}) DeleteEntryModelMasterRequest {
	return DeleteEntryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		EntryName:     core.CastString(data["entryName"]),
	}
}

func (p DeleteEntryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"entryName":     p.EntryName,
	}
}

func (p DeleteEntryModelMasterRequest) Pointer() *DeleteEntryModelMasterRequest {
	return &p
}

type DescribeEntriesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeEntriesRequestFromJson(data string) DescribeEntriesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntriesRequestFromDict(dict)
}

func NewDescribeEntriesRequestFromDict(data map[string]interface{}) DescribeEntriesRequest {
	return DescribeEntriesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeEntriesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeEntriesRequest) Pointer() *DescribeEntriesRequest {
	return &p
}

type DescribeEntriesByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeEntriesByUserIdRequestFromJson(data string) DescribeEntriesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntriesByUserIdRequestFromDict(dict)
}

func NewDescribeEntriesByUserIdRequestFromDict(data map[string]interface{}) DescribeEntriesByUserIdRequest {
	return DescribeEntriesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeEntriesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeEntriesByUserIdRequest) Pointer() *DescribeEntriesByUserIdRequest {
	return &p
}

type AddEntriesByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	EntryModelNames    []string `json:"entryModelNames"`
}

func NewAddEntriesByUserIdRequestFromJson(data string) AddEntriesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddEntriesByUserIdRequestFromDict(dict)
}

func NewAddEntriesByUserIdRequestFromDict(data map[string]interface{}) AddEntriesByUserIdRequest {
	return AddEntriesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		EntryModelNames: core.CastStrings(core.CastArray(data["entryModelNames"])),
	}
}

func (p AddEntriesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"entryModelNames": core.CastStringsFromDict(
			p.EntryModelNames,
		),
	}
}

func (p AddEntriesByUserIdRequest) Pointer() *AddEntriesByUserIdRequest {
	return &p
}

type GetEntryRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	EntryModelName     *string `json:"entryModelName"`
}

func NewGetEntryRequestFromJson(data string) GetEntryRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryRequestFromDict(dict)
}

func NewGetEntryRequestFromDict(data map[string]interface{}) GetEntryRequest {
	return GetEntryRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		EntryModelName: core.CastString(data["entryModelName"]),
	}
}

func (p GetEntryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"entryModelName": p.EntryModelName,
	}
}

func (p GetEntryRequest) Pointer() *GetEntryRequest {
	return &p
}

type GetEntryByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	EntryModelName     *string `json:"entryModelName"`
}

func NewGetEntryByUserIdRequestFromJson(data string) GetEntryByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryByUserIdRequestFromDict(dict)
}

func NewGetEntryByUserIdRequestFromDict(data map[string]interface{}) GetEntryByUserIdRequest {
	return GetEntryByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		EntryModelName: core.CastString(data["entryModelName"]),
	}
}

func (p GetEntryByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"entryModelName": p.EntryModelName,
	}
}

func (p GetEntryByUserIdRequest) Pointer() *GetEntryByUserIdRequest {
	return &p
}

type GetEntryWithSignatureRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	EntryModelName     *string `json:"entryModelName"`
	KeyId              *string `json:"keyId"`
}

func NewGetEntryWithSignatureRequestFromJson(data string) GetEntryWithSignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryWithSignatureRequestFromDict(dict)
}

func NewGetEntryWithSignatureRequestFromDict(data map[string]interface{}) GetEntryWithSignatureRequest {
	return GetEntryWithSignatureRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		EntryModelName: core.CastString(data["entryModelName"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetEntryWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"entryModelName": p.EntryModelName,
		"keyId":          p.KeyId,
	}
}

func (p GetEntryWithSignatureRequest) Pointer() *GetEntryWithSignatureRequest {
	return &p
}

type GetEntryWithSignatureByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	EntryModelName     *string `json:"entryModelName"`
	KeyId              *string `json:"keyId"`
}

func NewGetEntryWithSignatureByUserIdRequestFromJson(data string) GetEntryWithSignatureByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetEntryWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetEntryWithSignatureByUserIdRequest {
	return GetEntryWithSignatureByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		EntryModelName: core.CastString(data["entryModelName"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetEntryWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"entryModelName": p.EntryModelName,
		"keyId":          p.KeyId,
	}
}

func (p GetEntryWithSignatureByUserIdRequest) Pointer() *GetEntryWithSignatureByUserIdRequest {
	return &p
}

type ResetByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewResetByUserIdRequestFromJson(data string) ResetByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetByUserIdRequestFromDict(dict)
}

func NewResetByUserIdRequestFromDict(data map[string]interface{}) ResetByUserIdRequest {
	return ResetByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p ResetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p ResetByUserIdRequest) Pointer() *ResetByUserIdRequest {
	return &p
}

type AddEntriesByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewAddEntriesByStampSheetRequestFromJson(data string) AddEntriesByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddEntriesByStampSheetRequestFromDict(dict)
}

func NewAddEntriesByStampSheetRequestFromDict(data map[string]interface{}) AddEntriesByStampSheetRequest {
	return AddEntriesByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddEntriesByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddEntriesByStampSheetRequest) Pointer() *AddEntriesByStampSheetRequest {
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

type GetCurrentEntryMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetCurrentEntryMasterRequestFromJson(data string) GetCurrentEntryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentEntryMasterRequestFromDict(dict)
}

func NewGetCurrentEntryMasterRequestFromDict(data map[string]interface{}) GetCurrentEntryMasterRequest {
	return GetCurrentEntryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentEntryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentEntryMasterRequest) Pointer() *GetCurrentEntryMasterRequest {
	return &p
}

type UpdateCurrentEntryMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
}

func NewUpdateCurrentEntryMasterRequestFromJson(data string) UpdateCurrentEntryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEntryMasterRequestFromDict(dict)
}

func NewUpdateCurrentEntryMasterRequestFromDict(data map[string]interface{}) UpdateCurrentEntryMasterRequest {
	return UpdateCurrentEntryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentEntryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentEntryMasterRequest) Pointer() *UpdateCurrentEntryMasterRequest {
	return &p
}

type UpdateCurrentEntryMasterFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentEntryMasterFromGitHubRequestFromJson(data string) UpdateCurrentEntryMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEntryMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentEntryMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentEntryMasterFromGitHubRequest {
	return UpdateCurrentEntryMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentEntryMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentEntryMasterFromGitHubRequest) Pointer() *UpdateCurrentEntryMasterFromGitHubRequest {
	return &p
}
