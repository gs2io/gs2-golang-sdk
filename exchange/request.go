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

package exchange

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
	RequestId            *string             `json:"requestId"`
	ContextStack         *string             `json:"contextStack"`
	Name                 *string             `json:"name"`
	Description          *string             `json:"description"`
	EnableAwaitExchange  *bool               `json:"enableAwaitExchange"`
	EnableDirectExchange *bool               `json:"enableDirectExchange"`
	TransactionSetting   *TransactionSetting `json:"transactionSetting"`
	ExchangeScript       *ScriptSetting      `json:"exchangeScript"`
	LogSetting           *LogSetting         `json:"logSetting"`
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
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		EnableAwaitExchange:  core.CastBool(data["enableAwaitExchange"]),
		EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
		TransactionSetting:   NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ExchangeScript:       NewScriptSettingFromDict(core.CastMap(data["exchangeScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:     core.CastString(data["queueNamespaceId"]),
		KeyId:                core.CastString(data["keyId"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                 p.Name,
		"description":          p.Description,
		"enableAwaitExchange":  p.EnableAwaitExchange,
		"enableDirectExchange": p.EnableDirectExchange,
		"transactionSetting":   p.TransactionSetting.ToDict(),
		"exchangeScript":       p.ExchangeScript.ToDict(),
		"logSetting":           p.LogSetting.ToDict(),
		"queueNamespaceId":     p.QueueNamespaceId,
		"keyId":                p.KeyId,
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
	RequestId            *string             `json:"requestId"`
	ContextStack         *string             `json:"contextStack"`
	NamespaceName        *string             `json:"namespaceName"`
	Description          *string             `json:"description"`
	EnableAwaitExchange  *bool               `json:"enableAwaitExchange"`
	EnableDirectExchange *bool               `json:"enableDirectExchange"`
	TransactionSetting   *TransactionSetting `json:"transactionSetting"`
	ExchangeScript       *ScriptSetting      `json:"exchangeScript"`
	LogSetting           *LogSetting         `json:"logSetting"`
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
		NamespaceName:        core.CastString(data["namespaceName"]),
		Description:          core.CastString(data["description"]),
		EnableAwaitExchange:  core.CastBool(data["enableAwaitExchange"]),
		EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
		TransactionSetting:   NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ExchangeScript:       NewScriptSettingFromDict(core.CastMap(data["exchangeScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:     core.CastString(data["queueNamespaceId"]),
		KeyId:                core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"description":          p.Description,
		"enableAwaitExchange":  p.EnableAwaitExchange,
		"enableDirectExchange": p.EnableDirectExchange,
		"transactionSetting":   p.TransactionSetting.ToDict(),
		"exchangeScript":       p.ExchangeScript.ToDict(),
		"logSetting":           p.LogSetting.ToDict(),
		"queueNamespaceId":     p.QueueNamespaceId,
		"keyId":                p.KeyId,
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

type DumpUserDataByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelsRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelMastersRequestFromDict(dict)
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
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	NamespaceName      *string         `json:"namespaceName"`
	Name               *string         `json:"name"`
	Description        *string         `json:"description"`
	Metadata           *string         `json:"metadata"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
}

func NewCreateRateModelMasterRequestFromJson(data string) CreateRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRateModelMasterRequestFromDict(dict)
}

func NewCreateRateModelMasterRequestFromDict(data map[string]interface{}) CreateRateModelMasterRequest {
	return CreateRateModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
	}
}

func (p CreateRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"timingType":    p.TimingType,
		"lockTime":      p.LockTime,
		"enableSkip":    p.EnableSkip,
		"skipConsumeActions": CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelMasterRequestFromDict(dict)
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
	RequestId          *string         `json:"requestId"`
	ContextStack       *string         `json:"contextStack"`
	NamespaceName      *string         `json:"namespaceName"`
	RateName           *string         `json:"rateName"`
	Description        *string         `json:"description"`
	Metadata           *string         `json:"metadata"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
}

func NewUpdateRateModelMasterRequestFromJson(data string) UpdateRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRateModelMasterRequestFromDict(dict)
}

func NewUpdateRateModelMasterRequestFromDict(data map[string]interface{}) UpdateRateModelMasterRequest {
	return UpdateRateModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		RateName:           core.CastString(data["rateName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
	}
}

func (p UpdateRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"timingType":    p.TimingType,
		"lockTime":      p.LockTime,
		"enableSkip":    p.EnableSkip,
		"skipConsumeActions": CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRateModelMasterRequestFromDict(dict)
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

type DescribeIncrementalRateModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeIncrementalRateModelsRequestFromJson(data string) DescribeIncrementalRateModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIncrementalRateModelsRequestFromDict(dict)
}

func NewDescribeIncrementalRateModelsRequestFromDict(data map[string]interface{}) DescribeIncrementalRateModelsRequest {
	return DescribeIncrementalRateModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeIncrementalRateModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeIncrementalRateModelsRequest) Pointer() *DescribeIncrementalRateModelsRequest {
	return &p
}

type GetIncrementalRateModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewGetIncrementalRateModelRequestFromJson(data string) GetIncrementalRateModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIncrementalRateModelRequestFromDict(dict)
}

func NewGetIncrementalRateModelRequestFromDict(data map[string]interface{}) GetIncrementalRateModelRequest {
	return GetIncrementalRateModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p GetIncrementalRateModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p GetIncrementalRateModelRequest) Pointer() *GetIncrementalRateModelRequest {
	return &p
}

type DescribeIncrementalRateModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeIncrementalRateModelMastersRequestFromJson(data string) DescribeIncrementalRateModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIncrementalRateModelMastersRequestFromDict(dict)
}

func NewDescribeIncrementalRateModelMastersRequestFromDict(data map[string]interface{}) DescribeIncrementalRateModelMastersRequest {
	return DescribeIncrementalRateModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeIncrementalRateModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeIncrementalRateModelMastersRequest) Pointer() *DescribeIncrementalRateModelMastersRequest {
	return &p
}

type CreateIncrementalRateModelMasterRequest struct {
	RequestId            *string         `json:"requestId"`
	ContextStack         *string         `json:"contextStack"`
	NamespaceName        *string         `json:"namespaceName"`
	Name                 *string         `json:"name"`
	Description          *string         `json:"description"`
	Metadata             *string         `json:"metadata"`
	ConsumeAction        *ConsumeAction  `json:"consumeAction"`
	CalculateType        *string         `json:"calculateType"`
	BaseValue            *int64          `json:"baseValue"`
	CoefficientValue     *int64          `json:"coefficientValue"`
	CalculateScriptId    *string         `json:"calculateScriptId"`
	ExchangeCountId      *string         `json:"exchangeCountId"`
	MaximumExchangeCount *int32          `json:"maximumExchangeCount"`
	AcquireActions       []AcquireAction `json:"acquireActions"`
}

func NewCreateIncrementalRateModelMasterRequestFromJson(data string) CreateIncrementalRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateIncrementalRateModelMasterRequestFromDict(dict)
}

func NewCreateIncrementalRateModelMasterRequestFromDict(data map[string]interface{}) CreateIncrementalRateModelMasterRequest {
	return CreateIncrementalRateModelMasterRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		Metadata:             core.CastString(data["metadata"]),
		ConsumeAction:        NewConsumeActionFromDict(core.CastMap(data["consumeAction"])).Pointer(),
		CalculateType:        core.CastString(data["calculateType"]),
		BaseValue:            core.CastInt64(data["baseValue"]),
		CoefficientValue:     core.CastInt64(data["coefficientValue"]),
		CalculateScriptId:    core.CastString(data["calculateScriptId"]),
		ExchangeCountId:      core.CastString(data["exchangeCountId"]),
		MaximumExchangeCount: core.CastInt32(data["maximumExchangeCount"]),
		AcquireActions:       CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p CreateIncrementalRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"name":                 p.Name,
		"description":          p.Description,
		"metadata":             p.Metadata,
		"consumeAction":        p.ConsumeAction.ToDict(),
		"calculateType":        p.CalculateType,
		"baseValue":            p.BaseValue,
		"coefficientValue":     p.CoefficientValue,
		"calculateScriptId":    p.CalculateScriptId,
		"exchangeCountId":      p.ExchangeCountId,
		"maximumExchangeCount": p.MaximumExchangeCount,
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p CreateIncrementalRateModelMasterRequest) Pointer() *CreateIncrementalRateModelMasterRequest {
	return &p
}

type GetIncrementalRateModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewGetIncrementalRateModelMasterRequestFromJson(data string) GetIncrementalRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIncrementalRateModelMasterRequestFromDict(dict)
}

func NewGetIncrementalRateModelMasterRequestFromDict(data map[string]interface{}) GetIncrementalRateModelMasterRequest {
	return GetIncrementalRateModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p GetIncrementalRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p GetIncrementalRateModelMasterRequest) Pointer() *GetIncrementalRateModelMasterRequest {
	return &p
}

type UpdateIncrementalRateModelMasterRequest struct {
	RequestId            *string         `json:"requestId"`
	ContextStack         *string         `json:"contextStack"`
	NamespaceName        *string         `json:"namespaceName"`
	RateName             *string         `json:"rateName"`
	Description          *string         `json:"description"`
	Metadata             *string         `json:"metadata"`
	ConsumeAction        *ConsumeAction  `json:"consumeAction"`
	CalculateType        *string         `json:"calculateType"`
	BaseValue            *int64          `json:"baseValue"`
	CoefficientValue     *int64          `json:"coefficientValue"`
	CalculateScriptId    *string         `json:"calculateScriptId"`
	ExchangeCountId      *string         `json:"exchangeCountId"`
	MaximumExchangeCount *int32          `json:"maximumExchangeCount"`
	AcquireActions       []AcquireAction `json:"acquireActions"`
}

func NewUpdateIncrementalRateModelMasterRequestFromJson(data string) UpdateIncrementalRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateIncrementalRateModelMasterRequestFromDict(dict)
}

func NewUpdateIncrementalRateModelMasterRequestFromDict(data map[string]interface{}) UpdateIncrementalRateModelMasterRequest {
	return UpdateIncrementalRateModelMasterRequest{
		NamespaceName:        core.CastString(data["namespaceName"]),
		RateName:             core.CastString(data["rateName"]),
		Description:          core.CastString(data["description"]),
		Metadata:             core.CastString(data["metadata"]),
		ConsumeAction:        NewConsumeActionFromDict(core.CastMap(data["consumeAction"])).Pointer(),
		CalculateType:        core.CastString(data["calculateType"]),
		BaseValue:            core.CastInt64(data["baseValue"]),
		CoefficientValue:     core.CastInt64(data["coefficientValue"]),
		CalculateScriptId:    core.CastString(data["calculateScriptId"]),
		ExchangeCountId:      core.CastString(data["exchangeCountId"]),
		MaximumExchangeCount: core.CastInt32(data["maximumExchangeCount"]),
		AcquireActions:       CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p UpdateIncrementalRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":        p.NamespaceName,
		"rateName":             p.RateName,
		"description":          p.Description,
		"metadata":             p.Metadata,
		"consumeAction":        p.ConsumeAction.ToDict(),
		"calculateType":        p.CalculateType,
		"baseValue":            p.BaseValue,
		"coefficientValue":     p.CoefficientValue,
		"calculateScriptId":    p.CalculateScriptId,
		"exchangeCountId":      p.ExchangeCountId,
		"maximumExchangeCount": p.MaximumExchangeCount,
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p UpdateIncrementalRateModelMasterRequest) Pointer() *UpdateIncrementalRateModelMasterRequest {
	return &p
}

type DeleteIncrementalRateModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RateName      *string `json:"rateName"`
}

func NewDeleteIncrementalRateModelMasterRequestFromJson(data string) DeleteIncrementalRateModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteIncrementalRateModelMasterRequestFromDict(dict)
}

func NewDeleteIncrementalRateModelMasterRequestFromDict(data map[string]interface{}) DeleteIncrementalRateModelMasterRequest {
	return DeleteIncrementalRateModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p DeleteIncrementalRateModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p DeleteIncrementalRateModelMasterRequest) Pointer() *DeleteIncrementalRateModelMasterRequest {
	return &p
}

type ExchangeRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	RateName           *string  `json:"rateName"`
	AccessToken        *string  `json:"accessToken"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewExchangeRequestFromJson(data string) ExchangeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRequestFromDict(dict)
}

func NewExchangeRequestFromDict(data map[string]interface{}) ExchangeRequest {
	return ExchangeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ExchangeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
		"accessToken":   p.AccessToken,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ExchangeRequest) Pointer() *ExchangeRequest {
	return &p
}

type ExchangeByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	RateName           *string  `json:"rateName"`
	UserId             *string  `json:"userId"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewExchangeByUserIdRequestFromJson(data string) ExchangeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeByUserIdRequestFromDict(dict)
}

func NewExchangeByUserIdRequestFromDict(data map[string]interface{}) ExchangeByUserIdRequest {
	return ExchangeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
		UserId:        core.CastString(data["userId"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ExchangeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
		"userId":        p.UserId,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ExchangeByUserIdRequest) Pointer() *ExchangeByUserIdRequest {
	return &p
}

type ExchangeByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewExchangeByStampSheetRequestFromJson(data string) ExchangeByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeByStampSheetRequestFromDict(dict)
}

func NewExchangeByStampSheetRequestFromDict(data map[string]interface{}) ExchangeByStampSheetRequest {
	return ExchangeByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p ExchangeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ExchangeByStampSheetRequest) Pointer() *ExchangeByStampSheetRequest {
	return &p
}

type IncrementalExchangeRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	RateName           *string  `json:"rateName"`
	AccessToken        *string  `json:"accessToken"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewIncrementalExchangeRequestFromJson(data string) IncrementalExchangeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeRequestFromDict(dict)
}

func NewIncrementalExchangeRequestFromDict(data map[string]interface{}) IncrementalExchangeRequest {
	return IncrementalExchangeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p IncrementalExchangeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
		"accessToken":   p.AccessToken,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p IncrementalExchangeRequest) Pointer() *IncrementalExchangeRequest {
	return &p
}

type IncrementalExchangeByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	RateName           *string  `json:"rateName"`
	UserId             *string  `json:"userId"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewIncrementalExchangeByUserIdRequestFromJson(data string) IncrementalExchangeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeByUserIdRequestFromDict(dict)
}

func NewIncrementalExchangeByUserIdRequestFromDict(data map[string]interface{}) IncrementalExchangeByUserIdRequest {
	return IncrementalExchangeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
		UserId:        core.CastString(data["userId"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p IncrementalExchangeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
		"userId":        p.UserId,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p IncrementalExchangeByUserIdRequest) Pointer() *IncrementalExchangeByUserIdRequest {
	return &p
}

type IncrementalExchangeByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewIncrementalExchangeByStampSheetRequestFromJson(data string) IncrementalExchangeByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeByStampSheetRequestFromDict(dict)
}

func NewIncrementalExchangeByStampSheetRequestFromDict(data map[string]interface{}) IncrementalExchangeByStampSheetRequest {
	return IncrementalExchangeByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p IncrementalExchangeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p IncrementalExchangeByStampSheetRequest) Pointer() *IncrementalExchangeByStampSheetRequest {
	return &p
}

type UnlockIncrementalExchangeByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RateName           *string `json:"rateName"`
	UserId             *string `json:"userId"`
	LockTransactionId  *string `json:"lockTransactionId"`
}

func NewUnlockIncrementalExchangeByUserIdRequestFromJson(data string) UnlockIncrementalExchangeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockIncrementalExchangeByUserIdRequestFromDict(dict)
}

func NewUnlockIncrementalExchangeByUserIdRequestFromDict(data map[string]interface{}) UnlockIncrementalExchangeByUserIdRequest {
	return UnlockIncrementalExchangeByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		RateName:          core.CastString(data["rateName"]),
		UserId:            core.CastString(data["userId"]),
		LockTransactionId: core.CastString(data["lockTransactionId"]),
	}
}

func (p UnlockIncrementalExchangeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"rateName":          p.RateName,
		"userId":            p.UserId,
		"lockTransactionId": p.LockTransactionId,
	}
}

func (p UnlockIncrementalExchangeByUserIdRequest) Pointer() *UnlockIncrementalExchangeByUserIdRequest {
	return &p
}

type UnlockIncrementalExchangeByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewUnlockIncrementalExchangeByStampSheetRequestFromJson(data string) UnlockIncrementalExchangeByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockIncrementalExchangeByStampSheetRequestFromDict(dict)
}

func NewUnlockIncrementalExchangeByStampSheetRequestFromDict(data map[string]interface{}) UnlockIncrementalExchangeByStampSheetRequest {
	return UnlockIncrementalExchangeByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p UnlockIncrementalExchangeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p UnlockIncrementalExchangeByStampSheetRequest) Pointer() *UnlockIncrementalExchangeByStampSheetRequest {
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

type GetCurrentRateMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentRateMasterRequestFromJson(data string) GetCurrentRateMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRateMasterRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterRequestFromDict(dict)
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
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterFromGitHubRequestFromDict(dict)
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

type CreateAwaitByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	RateName           *string  `json:"rateName"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewCreateAwaitByUserIdRequestFromJson(data string) CreateAwaitByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAwaitByUserIdRequestFromDict(dict)
}

func NewCreateAwaitByUserIdRequestFromDict(data map[string]interface{}) CreateAwaitByUserIdRequest {
	return CreateAwaitByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		RateName:      core.CastString(data["rateName"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p CreateAwaitByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"rateName":      p.RateName,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p CreateAwaitByUserIdRequest) Pointer() *CreateAwaitByUserIdRequest {
	return &p
}

type DescribeAwaitsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RateName      *string `json:"rateName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeAwaitsRequestFromJson(data string) DescribeAwaitsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAwaitsRequestFromDict(dict)
}

func NewDescribeAwaitsRequestFromDict(data map[string]interface{}) DescribeAwaitsRequest {
	return DescribeAwaitsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		RateName:      core.CastString(data["rateName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAwaitsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rateName":      p.RateName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAwaitsRequest) Pointer() *DescribeAwaitsRequest {
	return &p
}

type DescribeAwaitsByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	RateName      *string `json:"rateName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeAwaitsByUserIdRequestFromJson(data string) DescribeAwaitsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAwaitsByUserIdRequestFromDict(dict)
}

func NewDescribeAwaitsByUserIdRequestFromDict(data map[string]interface{}) DescribeAwaitsByUserIdRequest {
	return DescribeAwaitsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		RateName:      core.CastString(data["rateName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAwaitsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"rateName":      p.RateName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAwaitsByUserIdRequest) Pointer() *DescribeAwaitsByUserIdRequest {
	return &p
}

type GetAwaitRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	AwaitName     *string `json:"awaitName"`
}

func NewGetAwaitRequestFromJson(data string) GetAwaitRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAwaitRequestFromDict(dict)
}

func NewGetAwaitRequestFromDict(data map[string]interface{}) GetAwaitRequest {
	return GetAwaitRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		AwaitName:     core.CastString(data["awaitName"]),
	}
}

func (p GetAwaitRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"awaitName":     p.AwaitName,
	}
}

func (p GetAwaitRequest) Pointer() *GetAwaitRequest {
	return &p
}

type GetAwaitByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	AwaitName     *string `json:"awaitName"`
}

func NewGetAwaitByUserIdRequestFromJson(data string) GetAwaitByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAwaitByUserIdRequestFromDict(dict)
}

func NewGetAwaitByUserIdRequestFromDict(data map[string]interface{}) GetAwaitByUserIdRequest {
	return GetAwaitByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		AwaitName:     core.CastString(data["awaitName"]),
	}
}

func (p GetAwaitByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"awaitName":     p.AwaitName,
	}
}

func (p GetAwaitByUserIdRequest) Pointer() *GetAwaitByUserIdRequest {
	return &p
}

type AcquireRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	AwaitName          *string  `json:"awaitName"`
	Config             []Config `json:"config"`
}

func NewAcquireRequestFromJson(data string) AcquireRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireRequestFromDict(dict)
}

func NewAcquireRequestFromDict(data map[string]interface{}) AcquireRequest {
	return AcquireRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		AwaitName:     core.CastString(data["awaitName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"awaitName":     p.AwaitName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p AcquireRequest) Pointer() *AcquireRequest {
	return &p
}

type AcquireByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	AwaitName          *string  `json:"awaitName"`
	Config             []Config `json:"config"`
}

func NewAcquireByUserIdRequestFromJson(data string) AcquireByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireByUserIdRequestFromDict(dict)
}

func NewAcquireByUserIdRequestFromDict(data map[string]interface{}) AcquireByUserIdRequest {
	return AcquireByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		AwaitName:     core.CastString(data["awaitName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"awaitName":     p.AwaitName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p AcquireByUserIdRequest) Pointer() *AcquireByUserIdRequest {
	return &p
}

type AcquireForceByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	AwaitName          *string  `json:"awaitName"`
	Config             []Config `json:"config"`
}

func NewAcquireForceByUserIdRequestFromJson(data string) AcquireForceByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireForceByUserIdRequestFromDict(dict)
}

func NewAcquireForceByUserIdRequestFromDict(data map[string]interface{}) AcquireForceByUserIdRequest {
	return AcquireForceByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		AwaitName:     core.CastString(data["awaitName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireForceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"awaitName":     p.AwaitName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p AcquireForceByUserIdRequest) Pointer() *AcquireForceByUserIdRequest {
	return &p
}

type SkipRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	AwaitName          *string  `json:"awaitName"`
	Config             []Config `json:"config"`
}

func NewSkipRequestFromJson(data string) SkipRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSkipRequestFromDict(dict)
}

func NewSkipRequestFromDict(data map[string]interface{}) SkipRequest {
	return SkipRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		AwaitName:     core.CastString(data["awaitName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p SkipRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"awaitName":     p.AwaitName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p SkipRequest) Pointer() *SkipRequest {
	return &p
}

type SkipByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	AwaitName          *string  `json:"awaitName"`
	Config             []Config `json:"config"`
}

func NewSkipByUserIdRequestFromJson(data string) SkipByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSkipByUserIdRequestFromDict(dict)
}

func NewSkipByUserIdRequestFromDict(data map[string]interface{}) SkipByUserIdRequest {
	return SkipByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		AwaitName:     core.CastString(data["awaitName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p SkipByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"awaitName":     p.AwaitName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p SkipByUserIdRequest) Pointer() *SkipByUserIdRequest {
	return &p
}

type DeleteAwaitRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	AwaitName          *string `json:"awaitName"`
}

func NewDeleteAwaitRequestFromJson(data string) DeleteAwaitRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitRequestFromDict(dict)
}

func NewDeleteAwaitRequestFromDict(data map[string]interface{}) DeleteAwaitRequest {
	return DeleteAwaitRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		AwaitName:     core.CastString(data["awaitName"]),
	}
}

func (p DeleteAwaitRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"awaitName":     p.AwaitName,
	}
}

func (p DeleteAwaitRequest) Pointer() *DeleteAwaitRequest {
	return &p
}

type DeleteAwaitByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	AwaitName          *string `json:"awaitName"`
}

func NewDeleteAwaitByUserIdRequestFromJson(data string) DeleteAwaitByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitByUserIdRequestFromDict(dict)
}

func NewDeleteAwaitByUserIdRequestFromDict(data map[string]interface{}) DeleteAwaitByUserIdRequest {
	return DeleteAwaitByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		AwaitName:     core.CastString(data["awaitName"]),
	}
}

func (p DeleteAwaitByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"awaitName":     p.AwaitName,
	}
}

func (p DeleteAwaitByUserIdRequest) Pointer() *DeleteAwaitByUserIdRequest {
	return &p
}

type CreateAwaitByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewCreateAwaitByStampSheetRequestFromJson(data string) CreateAwaitByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAwaitByStampSheetRequestFromDict(dict)
}

func NewCreateAwaitByStampSheetRequestFromDict(data map[string]interface{}) CreateAwaitByStampSheetRequest {
	return CreateAwaitByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p CreateAwaitByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p CreateAwaitByStampSheetRequest) Pointer() *CreateAwaitByStampSheetRequest {
	return &p
}

type DeleteAwaitByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewDeleteAwaitByStampTaskRequestFromJson(data string) DeleteAwaitByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitByStampTaskRequestFromDict(dict)
}

func NewDeleteAwaitByStampTaskRequestFromDict(data map[string]interface{}) DeleteAwaitByStampTaskRequest {
	return DeleteAwaitByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DeleteAwaitByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DeleteAwaitByStampTaskRequest) Pointer() *DeleteAwaitByStampTaskRequest {
	return &p
}
