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

package stamina

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
	RequestId             *string        `json:"requestId"`
	ContextStack          *string        `json:"contextStack"`
	DuplicationAvoider    *string        `json:"duplicationAvoider"`
	Name                  *string        `json:"name"`
	Description           *string        `json:"description"`
	OverflowTriggerScript *ScriptSetting `json:"overflowTriggerScript"`
	LogSetting            *LogSetting    `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		OverflowTriggerScript: NewScriptSettingFromDict(core.CastMap(data["overflowTriggerScript"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                  p.Name,
		"description":           p.Description,
		"overflowTriggerScript": p.OverflowTriggerScript.ToDict(),
		"logSetting":            p.LogSetting.ToDict(),
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
	RequestId             *string        `json:"requestId"`
	ContextStack          *string        `json:"contextStack"`
	DuplicationAvoider    *string        `json:"duplicationAvoider"`
	NamespaceName         *string        `json:"namespaceName"`
	Description           *string        `json:"description"`
	OverflowTriggerScript *ScriptSetting `json:"overflowTriggerScript"`
	LogSetting            *LogSetting    `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Description:           core.CastString(data["description"]),
		OverflowTriggerScript: NewScriptSettingFromDict(core.CastMap(data["overflowTriggerScript"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"description":           p.Description,
		"overflowTriggerScript": p.OverflowTriggerScript.ToDict(),
		"logSetting":            p.LogSetting.ToDict(),
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

type DescribeStaminaModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeStaminaModelMastersRequestFromJson(data string) DescribeStaminaModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaModelMastersRequestFromDict(dict)
}

func NewDescribeStaminaModelMastersRequestFromDict(data map[string]interface{}) DescribeStaminaModelMastersRequest {
	return DescribeStaminaModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStaminaModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStaminaModelMastersRequest) Pointer() *DescribeStaminaModelMastersRequest {
	return &p
}

type CreateStaminaModelMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	DuplicationAvoider       *string `json:"duplicationAvoider"`
	NamespaceName            *string `json:"namespaceName"`
	Name                     *string `json:"name"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
}

func NewCreateStaminaModelMasterRequestFromJson(data string) CreateStaminaModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStaminaModelMasterRequestFromDict(dict)
}

func NewCreateStaminaModelMasterRequestFromDict(data map[string]interface{}) CreateStaminaModelMasterRequest {
	return CreateStaminaModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		RecoverIntervalMinutes:   core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:             core.CastInt32(data["recoverValue"]),
		InitialCapacity:          core.CastInt32(data["initialCapacity"]),
		IsOverflow:               core.CastBool(data["isOverflow"]),
		MaxCapacity:              core.CastInt32(data["maxCapacity"]),
		MaxStaminaTableName:      core.CastString(data["maxStaminaTableName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		RecoverValueTableName:    core.CastString(data["recoverValueTableName"]),
	}
}

func (p CreateStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"recoverIntervalMinutes":   p.RecoverIntervalMinutes,
		"recoverValue":             p.RecoverValue,
		"initialCapacity":          p.InitialCapacity,
		"isOverflow":               p.IsOverflow,
		"maxCapacity":              p.MaxCapacity,
		"maxStaminaTableName":      p.MaxStaminaTableName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"recoverValueTableName":    p.RecoverValueTableName,
	}
}

func (p CreateStaminaModelMasterRequest) Pointer() *CreateStaminaModelMasterRequest {
	return &p
}

type GetStaminaModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
}

func NewGetStaminaModelMasterRequestFromJson(data string) GetStaminaModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaModelMasterRequestFromDict(dict)
}

func NewGetStaminaModelMasterRequestFromDict(data map[string]interface{}) GetStaminaModelMasterRequest {
	return GetStaminaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p GetStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p GetStaminaModelMasterRequest) Pointer() *GetStaminaModelMasterRequest {
	return &p
}

type UpdateStaminaModelMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	DuplicationAvoider       *string `json:"duplicationAvoider"`
	NamespaceName            *string `json:"namespaceName"`
	StaminaName              *string `json:"staminaName"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
}

func NewUpdateStaminaModelMasterRequestFromJson(data string) UpdateStaminaModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStaminaModelMasterRequestFromDict(dict)
}

func NewUpdateStaminaModelMasterRequestFromDict(data map[string]interface{}) UpdateStaminaModelMasterRequest {
	return UpdateStaminaModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		StaminaName:              core.CastString(data["staminaName"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		RecoverIntervalMinutes:   core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:             core.CastInt32(data["recoverValue"]),
		InitialCapacity:          core.CastInt32(data["initialCapacity"]),
		IsOverflow:               core.CastBool(data["isOverflow"]),
		MaxCapacity:              core.CastInt32(data["maxCapacity"]),
		MaxStaminaTableName:      core.CastString(data["maxStaminaTableName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		RecoverValueTableName:    core.CastString(data["recoverValueTableName"]),
	}
}

func (p UpdateStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"staminaName":              p.StaminaName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"recoverIntervalMinutes":   p.RecoverIntervalMinutes,
		"recoverValue":             p.RecoverValue,
		"initialCapacity":          p.InitialCapacity,
		"isOverflow":               p.IsOverflow,
		"maxCapacity":              p.MaxCapacity,
		"maxStaminaTableName":      p.MaxStaminaTableName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"recoverValueTableName":    p.RecoverValueTableName,
	}
}

func (p UpdateStaminaModelMasterRequest) Pointer() *UpdateStaminaModelMasterRequest {
	return &p
}

type DeleteStaminaModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
}

func NewDeleteStaminaModelMasterRequestFromJson(data string) DeleteStaminaModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStaminaModelMasterRequestFromDict(dict)
}

func NewDeleteStaminaModelMasterRequestFromDict(data map[string]interface{}) DeleteStaminaModelMasterRequest {
	return DeleteStaminaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p DeleteStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p DeleteStaminaModelMasterRequest) Pointer() *DeleteStaminaModelMasterRequest {
	return &p
}

type DescribeMaxStaminaTableMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeMaxStaminaTableMastersRequestFromJson(data string) DescribeMaxStaminaTableMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMaxStaminaTableMastersRequestFromDict(dict)
}

func NewDescribeMaxStaminaTableMastersRequestFromDict(data map[string]interface{}) DescribeMaxStaminaTableMastersRequest {
	return DescribeMaxStaminaTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMaxStaminaTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMaxStaminaTableMastersRequest) Pointer() *DescribeMaxStaminaTableMastersRequest {
	return &p
}

type CreateMaxStaminaTableMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	ExperienceModelId  *string `json:"experienceModelId"`
	Values             []int32 `json:"values"`
}

func NewCreateMaxStaminaTableMasterRequestFromJson(data string) CreateMaxStaminaTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMaxStaminaTableMasterRequestFromDict(dict)
}

func NewCreateMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) CreateMaxStaminaTableMasterRequest {
	return CreateMaxStaminaTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateMaxStaminaTableMasterRequest) Pointer() *CreateMaxStaminaTableMasterRequest {
	return &p
}

type GetMaxStaminaTableMasterRequest struct {
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
}

func NewGetMaxStaminaTableMasterRequestFromJson(data string) GetMaxStaminaTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMaxStaminaTableMasterRequestFromDict(dict)
}

func NewGetMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) GetMaxStaminaTableMasterRequest {
	return GetMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
	}
}

func (p GetMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
	}
}

func (p GetMaxStaminaTableMasterRequest) Pointer() *GetMaxStaminaTableMasterRequest {
	return &p
}

type UpdateMaxStaminaTableMasterRequest struct {
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	ExperienceModelId   *string `json:"experienceModelId"`
	Values              []int32 `json:"values"`
}

func NewUpdateMaxStaminaTableMasterRequestFromJson(data string) UpdateMaxStaminaTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMaxStaminaTableMasterRequestFromDict(dict)
}

func NewUpdateMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) UpdateMaxStaminaTableMasterRequest {
	return UpdateMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		ExperienceModelId:   core.CastString(data["experienceModelId"]),
		Values:              core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"experienceModelId":   p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateMaxStaminaTableMasterRequest) Pointer() *UpdateMaxStaminaTableMasterRequest {
	return &p
}

type DeleteMaxStaminaTableMasterRequest struct {
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
}

func NewDeleteMaxStaminaTableMasterRequestFromJson(data string) DeleteMaxStaminaTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMaxStaminaTableMasterRequestFromDict(dict)
}

func NewDeleteMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) DeleteMaxStaminaTableMasterRequest {
	return DeleteMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
	}
}

func (p DeleteMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
	}
}

func (p DeleteMaxStaminaTableMasterRequest) Pointer() *DeleteMaxStaminaTableMasterRequest {
	return &p
}

type DescribeRecoverIntervalTableMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeRecoverIntervalTableMastersRequestFromJson(data string) DescribeRecoverIntervalTableMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRecoverIntervalTableMastersRequestFromDict(dict)
}

func NewDescribeRecoverIntervalTableMastersRequestFromDict(data map[string]interface{}) DescribeRecoverIntervalTableMastersRequest {
	return DescribeRecoverIntervalTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRecoverIntervalTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRecoverIntervalTableMastersRequest) Pointer() *DescribeRecoverIntervalTableMastersRequest {
	return &p
}

type CreateRecoverIntervalTableMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	ExperienceModelId  *string `json:"experienceModelId"`
	Values             []int32 `json:"values"`
}

func NewCreateRecoverIntervalTableMasterRequestFromJson(data string) CreateRecoverIntervalTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRecoverIntervalTableMasterRequestFromDict(dict)
}

func NewCreateRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) CreateRecoverIntervalTableMasterRequest {
	return CreateRecoverIntervalTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateRecoverIntervalTableMasterRequest) Pointer() *CreateRecoverIntervalTableMasterRequest {
	return &p
}

type GetRecoverIntervalTableMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	DuplicationAvoider       *string `json:"duplicationAvoider"`
	NamespaceName            *string `json:"namespaceName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
}

func NewGetRecoverIntervalTableMasterRequestFromJson(data string) GetRecoverIntervalTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRecoverIntervalTableMasterRequestFromDict(dict)
}

func NewGetRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) GetRecoverIntervalTableMasterRequest {
	return GetRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
	}
}

func (p GetRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
	}
}

func (p GetRecoverIntervalTableMasterRequest) Pointer() *GetRecoverIntervalTableMasterRequest {
	return &p
}

type UpdateRecoverIntervalTableMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	DuplicationAvoider       *string `json:"duplicationAvoider"`
	NamespaceName            *string `json:"namespaceName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	ExperienceModelId        *string `json:"experienceModelId"`
	Values                   []int32 `json:"values"`
}

func NewUpdateRecoverIntervalTableMasterRequestFromJson(data string) UpdateRecoverIntervalTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRecoverIntervalTableMasterRequestFromDict(dict)
}

func NewUpdateRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) UpdateRecoverIntervalTableMasterRequest {
	return UpdateRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		ExperienceModelId:        core.CastString(data["experienceModelId"]),
		Values:                   core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"experienceModelId":        p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateRecoverIntervalTableMasterRequest) Pointer() *UpdateRecoverIntervalTableMasterRequest {
	return &p
}

type DeleteRecoverIntervalTableMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	DuplicationAvoider       *string `json:"duplicationAvoider"`
	NamespaceName            *string `json:"namespaceName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
}

func NewDeleteRecoverIntervalTableMasterRequestFromJson(data string) DeleteRecoverIntervalTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRecoverIntervalTableMasterRequestFromDict(dict)
}

func NewDeleteRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) DeleteRecoverIntervalTableMasterRequest {
	return DeleteRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
	}
}

func (p DeleteRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
	}
}

func (p DeleteRecoverIntervalTableMasterRequest) Pointer() *DeleteRecoverIntervalTableMasterRequest {
	return &p
}

type DescribeRecoverValueTableMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeRecoverValueTableMastersRequestFromJson(data string) DescribeRecoverValueTableMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRecoverValueTableMastersRequestFromDict(dict)
}

func NewDescribeRecoverValueTableMastersRequestFromDict(data map[string]interface{}) DescribeRecoverValueTableMastersRequest {
	return DescribeRecoverValueTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRecoverValueTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRecoverValueTableMastersRequest) Pointer() *DescribeRecoverValueTableMastersRequest {
	return &p
}

type CreateRecoverValueTableMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	ExperienceModelId  *string `json:"experienceModelId"`
	Values             []int32 `json:"values"`
}

func NewCreateRecoverValueTableMasterRequestFromJson(data string) CreateRecoverValueTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRecoverValueTableMasterRequestFromDict(dict)
}

func NewCreateRecoverValueTableMasterRequestFromDict(data map[string]interface{}) CreateRecoverValueTableMasterRequest {
	return CreateRecoverValueTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateRecoverValueTableMasterRequest) Pointer() *CreateRecoverValueTableMasterRequest {
	return &p
}

type GetRecoverValueTableMasterRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
}

func NewGetRecoverValueTableMasterRequestFromJson(data string) GetRecoverValueTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRecoverValueTableMasterRequestFromDict(dict)
}

func NewGetRecoverValueTableMasterRequestFromDict(data map[string]interface{}) GetRecoverValueTableMasterRequest {
	return GetRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
	}
}

func (p GetRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
	}
}

func (p GetRecoverValueTableMasterRequest) Pointer() *GetRecoverValueTableMasterRequest {
	return &p
}

type UpdateRecoverValueTableMasterRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	ExperienceModelId     *string `json:"experienceModelId"`
	Values                []int32 `json:"values"`
}

func NewUpdateRecoverValueTableMasterRequestFromJson(data string) UpdateRecoverValueTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRecoverValueTableMasterRequestFromDict(dict)
}

func NewUpdateRecoverValueTableMasterRequestFromDict(data map[string]interface{}) UpdateRecoverValueTableMasterRequest {
	return UpdateRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ExperienceModelId:     core.CastString(data["experienceModelId"]),
		Values:                core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"experienceModelId":     p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateRecoverValueTableMasterRequest) Pointer() *UpdateRecoverValueTableMasterRequest {
	return &p
}

type DeleteRecoverValueTableMasterRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
}

func NewDeleteRecoverValueTableMasterRequestFromJson(data string) DeleteRecoverValueTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRecoverValueTableMasterRequestFromDict(dict)
}

func NewDeleteRecoverValueTableMasterRequestFromDict(data map[string]interface{}) DeleteRecoverValueTableMasterRequest {
	return DeleteRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
	}
}

func (p DeleteRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
	}
}

func (p DeleteRecoverValueTableMasterRequest) Pointer() *DeleteRecoverValueTableMasterRequest {
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

type GetCurrentStaminaMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetCurrentStaminaMasterRequestFromJson(data string) GetCurrentStaminaMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentStaminaMasterRequestFromDict(dict)
}

func NewGetCurrentStaminaMasterRequestFromDict(data map[string]interface{}) GetCurrentStaminaMasterRequest {
	return GetCurrentStaminaMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentStaminaMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentStaminaMasterRequest) Pointer() *GetCurrentStaminaMasterRequest {
	return &p
}

type UpdateCurrentStaminaMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
}

func NewUpdateCurrentStaminaMasterRequestFromJson(data string) UpdateCurrentStaminaMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStaminaMasterRequestFromDict(dict)
}

func NewUpdateCurrentStaminaMasterRequestFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterRequest {
	return UpdateCurrentStaminaMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentStaminaMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentStaminaMasterRequest) Pointer() *UpdateCurrentStaminaMasterRequest {
	return &p
}

type UpdateCurrentStaminaMasterFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentStaminaMasterFromGitHubRequestFromJson(data string) UpdateCurrentStaminaMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStaminaMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentStaminaMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterFromGitHubRequest {
	return UpdateCurrentStaminaMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubRequest) Pointer() *UpdateCurrentStaminaMasterFromGitHubRequest {
	return &p
}

type DescribeStaminaModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeStaminaModelsRequestFromJson(data string) DescribeStaminaModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaModelsRequestFromDict(dict)
}

func NewDescribeStaminaModelsRequestFromDict(data map[string]interface{}) DescribeStaminaModelsRequest {
	return DescribeStaminaModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeStaminaModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStaminaModelsRequest) Pointer() *DescribeStaminaModelsRequest {
	return &p
}

type GetStaminaModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
}

func NewGetStaminaModelRequestFromJson(data string) GetStaminaModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaModelRequestFromDict(dict)
}

func NewGetStaminaModelRequestFromDict(data map[string]interface{}) GetStaminaModelRequest {
	return GetStaminaModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p GetStaminaModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p GetStaminaModelRequest) Pointer() *GetStaminaModelRequest {
	return &p
}

type DescribeStaminasRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeStaminasRequestFromJson(data string) DescribeStaminasRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminasRequestFromDict(dict)
}

func NewDescribeStaminasRequestFromDict(data map[string]interface{}) DescribeStaminasRequest {
	return DescribeStaminasRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStaminasRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStaminasRequest) Pointer() *DescribeStaminasRequest {
	return &p
}

type DescribeStaminasByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeStaminasByUserIdRequestFromJson(data string) DescribeStaminasByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminasByUserIdRequestFromDict(dict)
}

func NewDescribeStaminasByUserIdRequestFromDict(data map[string]interface{}) DescribeStaminasByUserIdRequest {
	return DescribeStaminasByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStaminasByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStaminasByUserIdRequest) Pointer() *DescribeStaminasByUserIdRequest {
	return &p
}

type GetStaminaRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	AccessToken        *string `json:"accessToken"`
}

func NewGetStaminaRequestFromJson(data string) GetStaminaRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaRequestFromDict(dict)
}

func NewGetStaminaRequestFromDict(data map[string]interface{}) GetStaminaRequest {
	return GetStaminaRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetStaminaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetStaminaRequest) Pointer() *GetStaminaRequest {
	return &p
}

type GetStaminaByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
}

func NewGetStaminaByUserIdRequestFromJson(data string) GetStaminaByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaByUserIdRequestFromDict(dict)
}

func NewGetStaminaByUserIdRequestFromDict(data map[string]interface{}) GetStaminaByUserIdRequest {
	return GetStaminaByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
	}
}

func (p GetStaminaByUserIdRequest) Pointer() *GetStaminaByUserIdRequest {
	return &p
}

type UpdateStaminaByUserIdRequest struct {
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	DuplicationAvoider     *string `json:"duplicationAvoider"`
	NamespaceName          *string `json:"namespaceName"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	Value                  *int32  `json:"value"`
	MaxValue               *int32  `json:"maxValue"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
	RecoverValue           *int32  `json:"recoverValue"`
}

func NewUpdateStaminaByUserIdRequestFromJson(data string) UpdateStaminaByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStaminaByUserIdRequestFromDict(dict)
}

func NewUpdateStaminaByUserIdRequestFromDict(data map[string]interface{}) UpdateStaminaByUserIdRequest {
	return UpdateStaminaByUserIdRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		StaminaName:            core.CastString(data["staminaName"]),
		UserId:                 core.CastString(data["userId"]),
		Value:                  core.CastInt32(data["value"]),
		MaxValue:               core.CastInt32(data["maxValue"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:           core.CastInt32(data["recoverValue"]),
	}
}

func (p UpdateStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"staminaName":            p.StaminaName,
		"userId":                 p.UserId,
		"value":                  p.Value,
		"maxValue":               p.MaxValue,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
		"recoverValue":           p.RecoverValue,
	}
}

func (p UpdateStaminaByUserIdRequest) Pointer() *UpdateStaminaByUserIdRequest {
	return &p
}

type ConsumeStaminaRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	AccessToken        *string `json:"accessToken"`
	ConsumeValue       *int32  `json:"consumeValue"`
}

func NewConsumeStaminaRequestFromJson(data string) ConsumeStaminaRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaRequestFromDict(dict)
}

func NewConsumeStaminaRequestFromDict(data map[string]interface{}) ConsumeStaminaRequest {
	return ConsumeStaminaRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ConsumeValue:  core.CastInt32(data["consumeValue"]),
	}
}

func (p ConsumeStaminaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"accessToken":   p.AccessToken,
		"consumeValue":  p.ConsumeValue,
	}
}

func (p ConsumeStaminaRequest) Pointer() *ConsumeStaminaRequest {
	return &p
}

type ConsumeStaminaByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	ConsumeValue       *int32  `json:"consumeValue"`
}

func NewConsumeStaminaByUserIdRequestFromJson(data string) ConsumeStaminaByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaByUserIdRequestFromDict(dict)
}

func NewConsumeStaminaByUserIdRequestFromDict(data map[string]interface{}) ConsumeStaminaByUserIdRequest {
	return ConsumeStaminaByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
		ConsumeValue:  core.CastInt32(data["consumeValue"]),
	}
}

func (p ConsumeStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
		"consumeValue":  p.ConsumeValue,
	}
}

func (p ConsumeStaminaByUserIdRequest) Pointer() *ConsumeStaminaByUserIdRequest {
	return &p
}

type RecoverStaminaByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RecoverValue       *int32  `json:"recoverValue"`
}

func NewRecoverStaminaByUserIdRequestFromJson(data string) RecoverStaminaByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecoverStaminaByUserIdRequestFromDict(dict)
}

func NewRecoverStaminaByUserIdRequestFromDict(data map[string]interface{}) RecoverStaminaByUserIdRequest {
	return RecoverStaminaByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
		RecoverValue:  core.CastInt32(data["recoverValue"]),
	}
}

func (p RecoverStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
		"recoverValue":  p.RecoverValue,
	}
}

func (p RecoverStaminaByUserIdRequest) Pointer() *RecoverStaminaByUserIdRequest {
	return &p
}

type RaiseMaxValueByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RaiseValue         *int32  `json:"raiseValue"`
}

func NewRaiseMaxValueByUserIdRequestFromJson(data string) RaiseMaxValueByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRaiseMaxValueByUserIdRequestFromDict(dict)
}

func NewRaiseMaxValueByUserIdRequestFromDict(data map[string]interface{}) RaiseMaxValueByUserIdRequest {
	return RaiseMaxValueByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
		RaiseValue:    core.CastInt32(data["raiseValue"]),
	}
}

func (p RaiseMaxValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
		"raiseValue":    p.RaiseValue,
	}
}

func (p RaiseMaxValueByUserIdRequest) Pointer() *RaiseMaxValueByUserIdRequest {
	return &p
}

type SetMaxValueByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	MaxValue           *int32  `json:"maxValue"`
}

func NewSetMaxValueByUserIdRequestFromJson(data string) SetMaxValueByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByUserIdRequestFromDict(dict)
}

func NewSetMaxValueByUserIdRequestFromDict(data map[string]interface{}) SetMaxValueByUserIdRequest {
	return SetMaxValueByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
		MaxValue:      core.CastInt32(data["maxValue"]),
	}
}

func (p SetMaxValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
		"maxValue":      p.MaxValue,
	}
}

func (p SetMaxValueByUserIdRequest) Pointer() *SetMaxValueByUserIdRequest {
	return &p
}

type SetRecoverIntervalByUserIdRequest struct {
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	DuplicationAvoider     *string `json:"duplicationAvoider"`
	NamespaceName          *string `json:"namespaceName"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
}

func NewSetRecoverIntervalByUserIdRequestFromJson(data string) SetRecoverIntervalByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByUserIdRequestFromDict(dict)
}

func NewSetRecoverIntervalByUserIdRequestFromDict(data map[string]interface{}) SetRecoverIntervalByUserIdRequest {
	return SetRecoverIntervalByUserIdRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		StaminaName:            core.CastString(data["staminaName"]),
		UserId:                 core.CastString(data["userId"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
	}
}

func (p SetRecoverIntervalByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"staminaName":            p.StaminaName,
		"userId":                 p.UserId,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
	}
}

func (p SetRecoverIntervalByUserIdRequest) Pointer() *SetRecoverIntervalByUserIdRequest {
	return &p
}

type SetRecoverValueByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RecoverValue       *int32  `json:"recoverValue"`
}

func NewSetRecoverValueByUserIdRequestFromJson(data string) SetRecoverValueByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByUserIdRequestFromDict(dict)
}

func NewSetRecoverValueByUserIdRequestFromDict(data map[string]interface{}) SetRecoverValueByUserIdRequest {
	return SetRecoverValueByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
		RecoverValue:  core.CastInt32(data["recoverValue"]),
	}
}

func (p SetRecoverValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
		"recoverValue":  p.RecoverValue,
	}
}

func (p SetRecoverValueByUserIdRequest) Pointer() *SetRecoverValueByUserIdRequest {
	return &p
}

type SetMaxValueByStatusRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func NewSetMaxValueByStatusRequestFromJson(data string) SetMaxValueByStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByStatusRequestFromDict(dict)
}

func NewSetMaxValueByStatusRequestFromDict(data map[string]interface{}) SetMaxValueByStatusRequest {
	return SetMaxValueByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetMaxValueByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetMaxValueByStatusRequest) Pointer() *SetMaxValueByStatusRequest {
	return &p
}

type SetRecoverIntervalByStatusRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func NewSetRecoverIntervalByStatusRequestFromJson(data string) SetRecoverIntervalByStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByStatusRequestFromDict(dict)
}

func NewSetRecoverIntervalByStatusRequestFromDict(data map[string]interface{}) SetRecoverIntervalByStatusRequest {
	return SetRecoverIntervalByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetRecoverIntervalByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetRecoverIntervalByStatusRequest) Pointer() *SetRecoverIntervalByStatusRequest {
	return &p
}

type SetRecoverValueByStatusRequest struct {
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func NewSetRecoverValueByStatusRequestFromJson(data string) SetRecoverValueByStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByStatusRequestFromDict(dict)
}

func NewSetRecoverValueByStatusRequestFromDict(data map[string]interface{}) SetRecoverValueByStatusRequest {
	return SetRecoverValueByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetRecoverValueByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetRecoverValueByStatusRequest) Pointer() *SetRecoverValueByStatusRequest {
	return &p
}

type DeleteStaminaByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
}

func NewDeleteStaminaByUserIdRequestFromJson(data string) DeleteStaminaByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStaminaByUserIdRequestFromDict(dict)
}

func NewDeleteStaminaByUserIdRequestFromDict(data map[string]interface{}) DeleteStaminaByUserIdRequest {
	return DeleteStaminaByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"userId":        p.UserId,
	}
}

func (p DeleteStaminaByUserIdRequest) Pointer() *DeleteStaminaByUserIdRequest {
	return &p
}

type RecoverStaminaByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewRecoverStaminaByStampSheetRequestFromJson(data string) RecoverStaminaByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecoverStaminaByStampSheetRequestFromDict(dict)
}

func NewRecoverStaminaByStampSheetRequestFromDict(data map[string]interface{}) RecoverStaminaByStampSheetRequest {
	return RecoverStaminaByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RecoverStaminaByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RecoverStaminaByStampSheetRequest) Pointer() *RecoverStaminaByStampSheetRequest {
	return &p
}

type RaiseMaxValueByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewRaiseMaxValueByStampSheetRequestFromJson(data string) RaiseMaxValueByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRaiseMaxValueByStampSheetRequestFromDict(dict)
}

func NewRaiseMaxValueByStampSheetRequestFromDict(data map[string]interface{}) RaiseMaxValueByStampSheetRequest {
	return RaiseMaxValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RaiseMaxValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RaiseMaxValueByStampSheetRequest) Pointer() *RaiseMaxValueByStampSheetRequest {
	return &p
}

type SetMaxValueByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewSetMaxValueByStampSheetRequestFromJson(data string) SetMaxValueByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByStampSheetRequestFromDict(dict)
}

func NewSetMaxValueByStampSheetRequestFromDict(data map[string]interface{}) SetMaxValueByStampSheetRequest {
	return SetMaxValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetMaxValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetMaxValueByStampSheetRequest) Pointer() *SetMaxValueByStampSheetRequest {
	return &p
}

type SetRecoverIntervalByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewSetRecoverIntervalByStampSheetRequestFromJson(data string) SetRecoverIntervalByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByStampSheetRequestFromDict(dict)
}

func NewSetRecoverIntervalByStampSheetRequestFromDict(data map[string]interface{}) SetRecoverIntervalByStampSheetRequest {
	return SetRecoverIntervalByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRecoverIntervalByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRecoverIntervalByStampSheetRequest) Pointer() *SetRecoverIntervalByStampSheetRequest {
	return &p
}

type SetRecoverValueByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewSetRecoverValueByStampSheetRequestFromJson(data string) SetRecoverValueByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByStampSheetRequestFromDict(dict)
}

func NewSetRecoverValueByStampSheetRequestFromDict(data map[string]interface{}) SetRecoverValueByStampSheetRequest {
	return SetRecoverValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRecoverValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRecoverValueByStampSheetRequest) Pointer() *SetRecoverValueByStampSheetRequest {
	return &p
}

type ConsumeStaminaByStampTaskRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampTask          *string `json:"stampTask"`
	KeyId              *string `json:"keyId"`
}

func NewConsumeStaminaByStampTaskRequestFromJson(data string) ConsumeStaminaByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaByStampTaskRequestFromDict(dict)
}

func NewConsumeStaminaByStampTaskRequestFromDict(data map[string]interface{}) ConsumeStaminaByStampTaskRequest {
	return ConsumeStaminaByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeStaminaByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeStaminaByStampTaskRequest) Pointer() *ConsumeStaminaByStampTaskRequest {
	return &p
}
