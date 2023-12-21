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

package enhance

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
	RequestId           *string             `json:"requestId"`
	ContextStack        *string             `json:"contextStack"`
	Name                *string             `json:"name"`
	Description         *string             `json:"description"`
	EnableDirectEnhance *bool               `json:"enableDirectEnhance"`
	TransactionSetting  *TransactionSetting `json:"transactionSetting"`
	EnhanceScript       *ScriptSetting      `json:"enhanceScript"`
	LogSetting          *LogSetting         `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
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
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		EnableDirectEnhance: core.CastBool(data["enableDirectEnhance"]),
		TransactionSetting:  NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		EnhanceScript:       NewScriptSettingFromDict(core.CastMap(data["enhanceScript"])).Pointer(),
		LogSetting:          NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:    core.CastString(data["queueNamespaceId"]),
		KeyId:               core.CastString(data["keyId"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                p.Name,
		"description":         p.Description,
		"enableDirectEnhance": p.EnableDirectEnhance,
		"transactionSetting":  p.TransactionSetting.ToDict(),
		"enhanceScript":       p.EnhanceScript.ToDict(),
		"logSetting":          p.LogSetting.ToDict(),
		"queueNamespaceId":    p.QueueNamespaceId,
		"keyId":               p.KeyId,
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
	RequestId           *string             `json:"requestId"`
	ContextStack        *string             `json:"contextStack"`
	NamespaceName       *string             `json:"namespaceName"`
	Description         *string             `json:"description"`
	EnableDirectEnhance *bool               `json:"enableDirectEnhance"`
	TransactionSetting  *TransactionSetting `json:"transactionSetting"`
	EnhanceScript       *ScriptSetting      `json:"enhanceScript"`
	LogSetting          *LogSetting         `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
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
		NamespaceName:       core.CastString(data["namespaceName"]),
		Description:         core.CastString(data["description"]),
		EnableDirectEnhance: core.CastBool(data["enableDirectEnhance"]),
		TransactionSetting:  NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		EnhanceScript:       NewScriptSettingFromDict(core.CastMap(data["enhanceScript"])).Pointer(),
		LogSetting:          NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:    core.CastString(data["queueNamespaceId"]),
		KeyId:               core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"description":         p.Description,
		"enableDirectEnhance": p.EnableDirectEnhance,
		"transactionSetting":  p.TransactionSetting.ToDict(),
		"enhanceScript":       p.EnhanceScript.ToDict(),
		"logSetting":          p.LogSetting.ToDict(),
		"queueNamespaceId":    p.QueueNamespaceId,
		"keyId":               p.KeyId,
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

type DescribeRateModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeRateModelsRequestFromJson(data string) DescribeRateModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeRateModelsRequestFromDict(dict2)
}

func NewDescribeRateModelsRequestFromDict(data map[string]interface{}) DescribeRateModelsRequest {
	return DescribeRateModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeRateModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRateModelsRequest) Pointer() *DescribeRateModelsRequest {
	return &p
}

type GetRateModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewGetRateModelRequestFromJson(data string) GetRateModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetRateModelRequestFromDict(dict2)
}

func NewGetRateModelRequestFromDict(data map[string]interface{}) GetRateModelRequest {
	return GetRateModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p GetRateModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p GetRateModelRequest) Pointer() *GetRateModelRequest {
	return &p
}

type DescribeRateModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeRateModelMastersRequestFromJson(data string) DescribeRateModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeRateModelMastersRequestFromDict(dict2)
}

func NewDescribeRateModelMastersRequestFromDict(data map[string]interface{}) DescribeRateModelMastersRequest {
	return DescribeRateModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRateModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRateModelMastersRequest) Pointer() *DescribeRateModelMastersRequest {
	return &p
}

type CreateRateModelMasterRequest struct {
	RequestId                  *string     `json:"requestId"`
	ContextStack               *string     `json:"contextStack"`
	NamespaceName              *string     `json:"namespaceName"`
	Name                       *string     `json:"name"`
	Description                *string     `json:"description"`
	Metadata                   *string     `json:"metadata"`
	TargetInventoryModelId     *string     `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string     `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string     `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy []*string   `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
}

func NewCreateRateModelMasterRequestFromJson(data string) CreateRateModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateRateModelMasterRequestFromDict(dict2)
}

func NewCreateRateModelMasterRequestFromDict(data map[string]interface{}) CreateRateModelMasterRequest {
	return CreateRateModelMasterRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetInventoryModelId:     core.CastString(data["targetInventoryModelId"]),
		AcquireExperienceSuffix:    core.CastString(data["acquireExperienceSuffix"]),
		MaterialInventoryModelId:   core.CastString(data["materialInventoryModelId"]),
		AcquireExperienceHierarchy: core.CastStrings(core.CastArray(data["acquireExperienceHierarchy"])),
		ExperienceModelId:          core.CastString(data["experienceModelId"]),
		BonusRates:                 CastBonusRates(core.CastArray(data["bonusRates"])),
	}
}

func (p CreateRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"targetInventoryModelId":   p.TargetInventoryModelId,
		"acquireExperienceSuffix":  p.AcquireExperienceSuffix,
		"materialInventoryModelId": p.MaterialInventoryModelId,
		"acquireExperienceHierarchy": core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		),
		"experienceModelId": p.ExperienceModelId,
		"bonusRates": CastBonusRatesFromDict(
			p.BonusRates,
		),
	}
}

func (p CreateRateModelMasterRequest) Pointer() *CreateRateModelMasterRequest {
	return &p
}

type GetRateModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewGetRateModelMasterRequestFromJson(data string) GetRateModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetRateModelMasterRequestFromDict(dict2)
}

func NewGetRateModelMasterRequestFromDict(data map[string]interface{}) GetRateModelMasterRequest {
	return GetRateModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p GetRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p GetRateModelMasterRequest) Pointer() *GetRateModelMasterRequest {
	return &p
}

type UpdateRateModelMasterRequest struct {
	RequestId                  *string     `json:"requestId"`
	ContextStack               *string     `json:"contextStack"`
	NamespaceName              *string     `json:"namespaceName"`
	RateName                   *string     `json:"rateName"`
	Description                *string     `json:"description"`
	Metadata                   *string     `json:"metadata"`
	TargetInventoryModelId     *string     `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string     `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string     `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy []*string   `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
}

func NewUpdateRateModelMasterRequestFromJson(data string) UpdateRateModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateRateModelMasterRequestFromDict(dict2)
}

func NewUpdateRateModelMasterRequestFromDict(data map[string]interface{}) UpdateRateModelMasterRequest {
	return UpdateRateModelMasterRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		RateName:                   core.CastString(data["rateName"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetInventoryModelId:     core.CastString(data["targetInventoryModelId"]),
		AcquireExperienceSuffix:    core.CastString(data["acquireExperienceSuffix"]),
		MaterialInventoryModelId:   core.CastString(data["materialInventoryModelId"]),
		AcquireExperienceHierarchy: core.CastStrings(core.CastArray(data["acquireExperienceHierarchy"])),
		ExperienceModelId:          core.CastString(data["experienceModelId"]),
		BonusRates:                 CastBonusRates(core.CastArray(data["bonusRates"])),
	}
}

func (p UpdateRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"rateName":                 p.RateName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"targetInventoryModelId":   p.TargetInventoryModelId,
		"acquireExperienceSuffix":  p.AcquireExperienceSuffix,
		"materialInventoryModelId": p.MaterialInventoryModelId,
		"acquireExperienceHierarchy": core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		),
		"experienceModelId": p.ExperienceModelId,
		"bonusRates": CastBonusRatesFromDict(
			p.BonusRates,
		),
	}
}

func (p UpdateRateModelMasterRequest) Pointer() *UpdateRateModelMasterRequest {
	return &p
}

type DeleteRateModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewDeleteRateModelMasterRequestFromJson(data string) DeleteRateModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteRateModelMasterRequestFromDict(dict2)
}

func NewDeleteRateModelMasterRequestFromDict(data map[string]interface{}) DeleteRateModelMasterRequest {
	return DeleteRateModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p DeleteRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p DeleteRateModelMasterRequest) Pointer() *DeleteRateModelMasterRequest {
	return &p
}

type DirectEnhanceRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	RateName           *string    `json:"rateName"`
	AccessToken        *string    `json:"accessToken"`
	TargetItemSetId    *string    `json:"targetItemSetId"`
	Materials          []Material `json:"materials"`
	Config             []Config   `json:"config"`
}

func NewDirectEnhanceRequestFromJson(data string) DirectEnhanceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDirectEnhanceRequestFromDict(dict2)
}

func NewDirectEnhanceRequestFromDict(data map[string]interface{}) DirectEnhanceRequest {
	return DirectEnhanceRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		RateName:        core.CastString(data["rateName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		TargetItemSetId: core.CastString(data["targetItemSetId"]),
		Materials:       CastMaterials(core.CastArray(data["materials"])),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p DirectEnhanceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rateName":        p.RateName,
		"accessToken":     p.AccessToken,
		"targetItemSetId": p.TargetItemSetId,
		"materials": CastMaterialsFromDict(
			p.Materials,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p DirectEnhanceRequest) Pointer() *DirectEnhanceRequest {
	return &p
}

type DirectEnhanceByUserIdRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	RateName           *string    `json:"rateName"`
	UserId             *string    `json:"userId"`
	TargetItemSetId    *string    `json:"targetItemSetId"`
	Materials          []Material `json:"materials"`
	Config             []Config   `json:"config"`
}

func NewDirectEnhanceByUserIdRequestFromJson(data string) DirectEnhanceByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDirectEnhanceByUserIdRequestFromDict(dict2)
}

func NewDirectEnhanceByUserIdRequestFromDict(data map[string]interface{}) DirectEnhanceByUserIdRequest {
	return DirectEnhanceByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		RateName:        core.CastString(data["rateName"]),
		UserId:          core.CastString(data["userId"]),
		TargetItemSetId: core.CastString(data["targetItemSetId"]),
		Materials:       CastMaterials(core.CastArray(data["materials"])),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p DirectEnhanceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rateName":        p.RateName,
		"userId":          p.UserId,
		"targetItemSetId": p.TargetItemSetId,
		"materials": CastMaterialsFromDict(
			p.Materials,
		),
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p DirectEnhanceByUserIdRequest) Pointer() *DirectEnhanceByUserIdRequest {
	return &p
}

type DirectEnhanceByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewDirectEnhanceByStampSheetRequestFromJson(data string) DirectEnhanceByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDirectEnhanceByStampSheetRequestFromDict(dict2)
}

func NewDirectEnhanceByStampSheetRequestFromDict(data map[string]interface{}) DirectEnhanceByStampSheetRequest {
	return DirectEnhanceByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DirectEnhanceByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DirectEnhanceByStampSheetRequest) Pointer() *DirectEnhanceByStampSheetRequest {
	return &p
}

type CreateProgressByUserIdRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	UserId             *string    `json:"userId"`
	RateName           *string    `json:"rateName"`
	TargetItemSetId    *string    `json:"targetItemSetId"`
	Materials          []Material `json:"materials"`
	Force              *bool      `json:"force"`
}

func NewCreateProgressByUserIdRequestFromJson(data string) CreateProgressByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateProgressByUserIdRequestFromDict(dict2)
}

func NewCreateProgressByUserIdRequestFromDict(data map[string]interface{}) CreateProgressByUserIdRequest {
	return CreateProgressByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		RateName:        core.CastString(data["rateName"]),
		TargetItemSetId: core.CastString(data["targetItemSetId"]),
		Materials:       CastMaterials(core.CastArray(data["materials"])),
		Force:           core.CastBool(data["force"]),
	}
}

func (p CreateProgressByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rateName":        p.RateName,
		"targetItemSetId": p.TargetItemSetId,
		"materials": CastMaterialsFromDict(
			p.Materials,
		),
		"force": p.Force,
	}
}

func (p CreateProgressByUserIdRequest) Pointer() *CreateProgressByUserIdRequest {
	return &p
}

type GetProgressRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
}

func NewGetProgressRequestFromJson(data string) GetProgressRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetProgressRequestFromDict(dict2)
}

func NewGetProgressRequestFromDict(data map[string]interface{}) GetProgressRequest {
	return GetProgressRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetProgressRequest) Pointer() *GetProgressRequest {
	return &p
}

type GetProgressByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
}

func NewGetProgressByUserIdRequestFromJson(data string) GetProgressByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetProgressByUserIdRequestFromDict(dict2)
}

func NewGetProgressByUserIdRequestFromDict(data map[string]interface{}) GetProgressByUserIdRequest {
	return GetProgressByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetProgressByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p GetProgressByUserIdRequest) Pointer() *GetProgressByUserIdRequest {
	return &p
}

type StartRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	RateName           *string    `json:"rateName"`
	TargetItemSetId    *string    `json:"targetItemSetId"`
	Materials          []Material `json:"materials"`
	AccessToken        *string    `json:"accessToken"`
	Force              *bool      `json:"force"`
	Config             []Config   `json:"config"`
}

func NewStartRequestFromJson(data string) StartRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewStartRequestFromDict(dict2)
}

func NewStartRequestFromDict(data map[string]interface{}) StartRequest {
	return StartRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		RateName:        core.CastString(data["rateName"]),
		TargetItemSetId: core.CastString(data["targetItemSetId"]),
		Materials:       CastMaterials(core.CastArray(data["materials"])),
		AccessToken:     core.CastString(data["accessToken"]),
		Force:           core.CastBool(data["force"]),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p StartRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rateName":        p.RateName,
		"targetItemSetId": p.TargetItemSetId,
		"materials": CastMaterialsFromDict(
			p.Materials,
		),
		"accessToken": p.AccessToken,
		"force":       p.Force,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p StartRequest) Pointer() *StartRequest {
	return &p
}

type StartByUserIdRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	RateName           *string    `json:"rateName"`
	TargetItemSetId    *string    `json:"targetItemSetId"`
	Materials          []Material `json:"materials"`
	UserId             *string    `json:"userId"`
	Force              *bool      `json:"force"`
	Config             []Config   `json:"config"`
}

func NewStartByUserIdRequestFromJson(data string) StartByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewStartByUserIdRequestFromDict(dict2)
}

func NewStartByUserIdRequestFromDict(data map[string]interface{}) StartByUserIdRequest {
	return StartByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		RateName:        core.CastString(data["rateName"]),
		TargetItemSetId: core.CastString(data["targetItemSetId"]),
		Materials:       CastMaterials(core.CastArray(data["materials"])),
		UserId:          core.CastString(data["userId"]),
		Force:           core.CastBool(data["force"]),
		Config:          CastConfigs(core.CastArray(data["config"])),
	}
}

func (p StartByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rateName":        p.RateName,
		"targetItemSetId": p.TargetItemSetId,
		"materials": CastMaterialsFromDict(
			p.Materials,
		),
		"userId": p.UserId,
		"force":  p.Force,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p StartByUserIdRequest) Pointer() *StartByUserIdRequest {
	return &p
}

type EndRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	Config             []Config `json:"config"`
}

func NewEndRequestFromJson(data string) EndRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewEndRequestFromDict(dict2)
}

func NewEndRequestFromDict(data map[string]interface{}) EndRequest {
	return EndRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p EndRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p EndRequest) Pointer() *EndRequest {
	return &p
}

type EndByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	Config             []Config `json:"config"`
}

func NewEndByUserIdRequestFromJson(data string) EndByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewEndByUserIdRequestFromDict(dict2)
}

func NewEndByUserIdRequestFromDict(data map[string]interface{}) EndByUserIdRequest {
	return EndByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p EndByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p EndByUserIdRequest) Pointer() *EndByUserIdRequest {
	return &p
}

type DeleteProgressRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
}

func NewDeleteProgressRequestFromJson(data string) DeleteProgressRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteProgressRequestFromDict(dict2)
}

func NewDeleteProgressRequestFromDict(data map[string]interface{}) DeleteProgressRequest {
	return DeleteProgressRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DeleteProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p DeleteProgressRequest) Pointer() *DeleteProgressRequest {
	return &p
}

type DeleteProgressByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewDeleteProgressByUserIdRequestFromJson(data string) DeleteProgressByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteProgressByUserIdRequestFromDict(dict2)
}

func NewDeleteProgressByUserIdRequestFromDict(data map[string]interface{}) DeleteProgressByUserIdRequest {
	return DeleteProgressByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteProgressByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p DeleteProgressByUserIdRequest) Pointer() *DeleteProgressByUserIdRequest {
	return &p
}

type CreateProgressByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewCreateProgressByStampSheetRequestFromJson(data string) CreateProgressByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateProgressByStampSheetRequestFromDict(dict2)
}

func NewCreateProgressByStampSheetRequestFromDict(data map[string]interface{}) CreateProgressByStampSheetRequest {
	return CreateProgressByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p CreateProgressByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p CreateProgressByStampSheetRequest) Pointer() *CreateProgressByStampSheetRequest {
	return &p
}

type DeleteProgressByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewDeleteProgressByStampTaskRequestFromJson(data string) DeleteProgressByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteProgressByStampTaskRequestFromDict(dict2)
}

func NewDeleteProgressByStampTaskRequestFromDict(data map[string]interface{}) DeleteProgressByStampTaskRequest {
	return DeleteProgressByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DeleteProgressByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DeleteProgressByStampTaskRequest) Pointer() *DeleteProgressByStampTaskRequest {
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

type GetCurrentRateMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentRateMasterRequestFromJson(data string) GetCurrentRateMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentRateMasterRequestFromDict(dict2)
}

func NewGetCurrentRateMasterRequestFromDict(data map[string]interface{}) GetCurrentRateMasterRequest {
	return GetCurrentRateMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentRateMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentRateMasterRequest) Pointer() *GetCurrentRateMasterRequest {
	return &p
}

type UpdateCurrentRateMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentRateMasterRequestFromJson(data string) UpdateCurrentRateMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentRateMasterRequestFromDict(dict2)
}

func NewUpdateCurrentRateMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRateMasterRequest {
	return UpdateCurrentRateMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentRateMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentRateMasterRequest) Pointer() *UpdateCurrentRateMasterRequest {
	return &p
}

type UpdateCurrentRateMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRateMasterFromGitHubRequestFromJson(data string) UpdateCurrentRateMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentRateMasterFromGitHubRequestFromDict(dict2)
}

func NewUpdateCurrentRateMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRateMasterFromGitHubRequest {
	return UpdateCurrentRateMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubRequest) Pointer() *UpdateCurrentRateMasterFromGitHubRequest {
	return &p
}
