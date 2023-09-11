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

package skillTree

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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ReleaseScript      *ScriptSetting      `json:"releaseScript"`
	RestrainScript     *ScriptSetting      `json:"restrainScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ReleaseScript:      NewScriptSettingFromDict(core.CastMap(data["releaseScript"])).Pointer(),
		RestrainScript:     NewScriptSettingFromDict(core.CastMap(data["restrainScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"releaseScript":      p.ReleaseScript.ToDict(),
		"restrainScript":     p.RestrainScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ReleaseScript      *ScriptSetting      `json:"releaseScript"`
	RestrainScript     *ScriptSetting      `json:"restrainScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ReleaseScript:      NewScriptSettingFromDict(core.CastMap(data["releaseScript"])).Pointer(),
		RestrainScript:     NewScriptSettingFromDict(core.CastMap(data["restrainScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"releaseScript":      p.ReleaseScript.ToDict(),
		"restrainScript":     p.RestrainScript.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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

type DescribeStatusModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeStatusModelsRequestFromJson(data string) DescribeStatusModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusModelsRequestFromDict(dict)
}

func NewDescribeStatusModelsRequestFromDict(data map[string]interface{}) DescribeStatusModelsRequest {
	return DescribeStatusModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeStatusModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStatusModelsRequest) Pointer() *DescribeStatusModelsRequest {
	return &p
}

type GetStatusModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewGetStatusModelRequestFromJson(data string) GetStatusModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusModelRequestFromDict(dict)
}

func NewGetStatusModelRequestFromDict(data map[string]interface{}) GetStatusModelRequest {
	return GetStatusModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p GetStatusModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p GetStatusModelRequest) Pointer() *GetStatusModelRequest {
	return &p
}

type DescribeStatusModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStatusModelMastersRequestFromJson(data string) DescribeStatusModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusModelMastersRequestFromDict(dict)
}

func NewDescribeStatusModelMastersRequestFromDict(data map[string]interface{}) DescribeStatusModelMastersRequest {
	return DescribeStatusModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusModelMastersRequest) Pointer() *DescribeStatusModelMastersRequest {
	return &p
}

type CreateStatusModelMasterRequest struct {
	RequestId             *string         `json:"requestId"`
	ContextStack          *string         `json:"contextStack"`
	NamespaceName         *string         `json:"namespaceName"`
	Name                  *string         `json:"name"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
}

func NewCreateStatusModelMasterRequestFromJson(data string) CreateStatusModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStatusModelMasterRequestFromDict(dict)
}

func NewCreateStatusModelMasterRequestFromDict(data map[string]interface{}) CreateStatusModelMasterRequest {
	return CreateStatusModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
	}
}

func (p CreateStatusModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"releaseConsumeActions": CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		),
		"restrainReturnRate": p.RestrainReturnRate,
		"premiseNodeNames": core.CastStringsFromDict(
			p.PremiseNodeNames,
		),
	}
}

func (p CreateStatusModelMasterRequest) Pointer() *CreateStatusModelMasterRequest {
	return &p
}

type GetStatusModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewGetStatusModelMasterRequestFromJson(data string) GetStatusModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusModelMasterRequestFromDict(dict)
}

func NewGetStatusModelMasterRequestFromDict(data map[string]interface{}) GetStatusModelMasterRequest {
	return GetStatusModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p GetStatusModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p GetStatusModelMasterRequest) Pointer() *GetStatusModelMasterRequest {
	return &p
}

type UpdateStatusModelMasterRequest struct {
	RequestId             *string         `json:"requestId"`
	ContextStack          *string         `json:"contextStack"`
	NamespaceName         *string         `json:"namespaceName"`
	NodeModelName         *string         `json:"nodeModelName"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
}

func NewUpdateStatusModelMasterRequestFromJson(data string) UpdateStatusModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStatusModelMasterRequestFromDict(dict)
}

func NewUpdateStatusModelMasterRequestFromDict(data map[string]interface{}) UpdateStatusModelMasterRequest {
	return UpdateStatusModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		NodeModelName:         core.CastString(data["nodeModelName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
	}
}

func (p UpdateStatusModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"releaseConsumeActions": CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		),
		"restrainReturnRate": p.RestrainReturnRate,
		"premiseNodeNames": core.CastStringsFromDict(
			p.PremiseNodeNames,
		),
	}
}

func (p UpdateStatusModelMasterRequest) Pointer() *UpdateStatusModelMasterRequest {
	return &p
}

type DeleteStatusModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	NodeModelName *string `json:"nodeModelName"`
}

func NewDeleteStatusModelMasterRequestFromJson(data string) DeleteStatusModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStatusModelMasterRequestFromDict(dict)
}

func NewDeleteStatusModelMasterRequestFromDict(data map[string]interface{}) DeleteStatusModelMasterRequest {
	return DeleteStatusModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		NodeModelName: core.CastString(data["nodeModelName"]),
	}
}

func (p DeleteStatusModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"nodeModelName": p.NodeModelName,
	}
}

func (p DeleteStatusModelMasterRequest) Pointer() *DeleteStatusModelMasterRequest {
	return &p
}

type MarkReleaseByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
}

func NewMarkReleaseByUserIdRequestFromJson(data string) MarkReleaseByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReleaseByUserIdRequestFromDict(dict)
}

func NewMarkReleaseByUserIdRequestFromDict(data map[string]interface{}) MarkReleaseByUserIdRequest {
	return MarkReleaseByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
	}
}

func (p MarkReleaseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
	}
}

func (p MarkReleaseByUserIdRequest) Pointer() *MarkReleaseByUserIdRequest {
	return &p
}

type ReleaseRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewReleaseRequestFromJson(data string) ReleaseRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReleaseRequestFromDict(dict)
}

func NewReleaseRequestFromDict(data map[string]interface{}) ReleaseRequest {
	return ReleaseRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReleaseRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReleaseRequest) Pointer() *ReleaseRequest {
	return &p
}

type ReleaseByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewReleaseByUserIdRequestFromJson(data string) ReleaseByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReleaseByUserIdRequestFromDict(dict)
}

func NewReleaseByUserIdRequestFromDict(data map[string]interface{}) ReleaseByUserIdRequest {
	return ReleaseByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReleaseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReleaseByUserIdRequest) Pointer() *ReleaseByUserIdRequest {
	return &p
}

type MarkRestrainByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
}

func NewMarkRestrainByUserIdRequestFromJson(data string) MarkRestrainByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkRestrainByUserIdRequestFromDict(dict)
}

func NewMarkRestrainByUserIdRequestFromDict(data map[string]interface{}) MarkRestrainByUserIdRequest {
	return MarkRestrainByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
	}
}

func (p MarkRestrainByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
	}
}

func (p MarkRestrainByUserIdRequest) Pointer() *MarkRestrainByUserIdRequest {
	return &p
}

type RestrainRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewRestrainRequestFromJson(data string) RestrainRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRestrainRequestFromDict(dict)
}

func NewRestrainRequestFromDict(data map[string]interface{}) RestrainRequest {
	return RestrainRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RestrainRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RestrainRequest) Pointer() *RestrainRequest {
	return &p
}

type RestrainByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	NodeModelNames     []*string `json:"nodeModelNames"`
	Config             []Config  `json:"config"`
}

func NewRestrainByUserIdRequestFromJson(data string) RestrainByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRestrainByUserIdRequestFromDict(dict)
}

func NewRestrainByUserIdRequestFromDict(data map[string]interface{}) RestrainByUserIdRequest {
	return RestrainByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		NodeModelNames: core.CastStrings(core.CastArray(data["nodeModelNames"])),
		Config:         CastConfigs(core.CastArray(data["config"])),
	}
}

func (p RestrainByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"nodeModelNames": core.CastStringsFromDict(
			p.NodeModelNames,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p RestrainByUserIdRequest) Pointer() *RestrainByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
}

func NewGetStatusRequestFromJson(data string) GetStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusRequestFromDict(dict)
}

func NewGetStatusRequestFromDict(data map[string]interface{}) GetStatusRequest {
	return GetStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
	return &p
}

type GetStatusByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
}

func NewGetStatusByUserIdRequestFromJson(data string) GetStatusByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusByUserIdRequestFromDict(dict)
}

func NewGetStatusByUserIdRequestFromDict(data map[string]interface{}) GetStatusByUserIdRequest {
	return GetStatusByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type ResetRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	Config             []Config `json:"config"`
}

func NewResetRequestFromJson(data string) ResetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetRequestFromDict(dict)
}

func NewResetRequestFromDict(data map[string]interface{}) ResetRequest {
	return ResetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ResetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ResetRequest) Pointer() *ResetRequest {
	return &p
}

type ResetByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	Config             []Config `json:"config"`
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
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ResetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ResetByUserIdRequest) Pointer() *ResetByUserIdRequest {
	return &p
}

type MarkReleaseByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewMarkReleaseByStampSheetRequestFromJson(data string) MarkReleaseByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReleaseByStampSheetRequestFromDict(dict)
}

func NewMarkReleaseByStampSheetRequestFromDict(data map[string]interface{}) MarkReleaseByStampSheetRequest {
	return MarkReleaseByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p MarkReleaseByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p MarkReleaseByStampSheetRequest) Pointer() *MarkReleaseByStampSheetRequest {
	return &p
}

type MarkRestrainByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewMarkRestrainByStampTaskRequestFromJson(data string) MarkRestrainByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkRestrainByStampTaskRequestFromDict(dict)
}

func NewMarkRestrainByStampTaskRequestFromDict(data map[string]interface{}) MarkRestrainByStampTaskRequest {
	return MarkRestrainByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p MarkRestrainByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p MarkRestrainByStampTaskRequest) Pointer() *MarkRestrainByStampTaskRequest {
	return &p
}

type ExportMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
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

type GetCurrentStatusMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentStatusMasterRequestFromJson(data string) GetCurrentStatusMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentStatusMasterRequestFromDict(dict)
}

func NewGetCurrentStatusMasterRequestFromDict(data map[string]interface{}) GetCurrentStatusMasterRequest {
	return GetCurrentStatusMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentStatusMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentStatusMasterRequest) Pointer() *GetCurrentStatusMasterRequest {
	return &p
}

type UpdateCurrentStatusMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentStatusMasterRequestFromJson(data string) UpdateCurrentStatusMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStatusMasterRequestFromDict(dict)
}

func NewUpdateCurrentStatusMasterRequestFromDict(data map[string]interface{}) UpdateCurrentStatusMasterRequest {
	return UpdateCurrentStatusMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentStatusMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentStatusMasterRequest) Pointer() *UpdateCurrentStatusMasterRequest {
	return &p
}

type UpdateCurrentStatusMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentStatusMasterFromGitHubRequestFromJson(data string) UpdateCurrentStatusMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStatusMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentStatusMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentStatusMasterFromGitHubRequest {
	return UpdateCurrentStatusMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentStatusMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentStatusMasterFromGitHubRequest) Pointer() *UpdateCurrentStatusMasterFromGitHubRequest {
	return &p
}
