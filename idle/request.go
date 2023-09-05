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

package idle

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
	ReceiveScript      *ScriptSetting      `json:"receiveScript"`
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
		ReceiveScript:      NewScriptSettingFromDict(core.CastMap(data["receiveScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"receiveScript":      p.ReceiveScript.ToDict(),
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
	ReceiveScript      *ScriptSetting      `json:"receiveScript"`
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
		ReceiveScript:      NewScriptSettingFromDict(core.CastMap(data["receiveScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"receiveScript":      p.ReceiveScript.ToDict(),
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

type DescribeCategoryModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeCategoryModelMastersRequestFromJson(data string) DescribeCategoryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelMastersRequestFromDict(dict)
}

func NewDescribeCategoryModelMastersRequestFromDict(data map[string]interface{}) DescribeCategoryModelMastersRequest {
	return DescribeCategoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCategoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCategoryModelMastersRequest) Pointer() *DescribeCategoryModelMastersRequest {
	return &p
}

type CreateCategoryModelMasterRequest struct {
	RequestId                 *string             `json:"requestId"`
	ContextStack              *string             `json:"contextStack"`
	NamespaceName             *string             `json:"namespaceName"`
	Name                      *string             `json:"name"`
	Description               *string             `json:"description"`
	Metadata                  *string             `json:"metadata"`
	RewardIntervalMinutes     *int32              `json:"rewardIntervalMinutes"`
	DefaultMaximumIdleMinutes *int32              `json:"defaultMaximumIdleMinutes"`
	AcquireActions            []AcquireActionList `json:"acquireActions"`
	IdlePeriodScheduleId      *string             `json:"idlePeriodScheduleId"`
	ReceivePeriodScheduleId   *string             `json:"receivePeriodScheduleId"`
}

func NewCreateCategoryModelMasterRequestFromJson(data string) CreateCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCategoryModelMasterRequestFromDict(dict)
}

func NewCreateCategoryModelMasterRequestFromDict(data map[string]interface{}) CreateCategoryModelMasterRequest {
	return CreateCategoryModelMasterRequest{
		NamespaceName:             core.CastString(data["namespaceName"]),
		Name:                      core.CastString(data["name"]),
		Description:               core.CastString(data["description"]),
		Metadata:                  core.CastString(data["metadata"]),
		RewardIntervalMinutes:     core.CastInt32(data["rewardIntervalMinutes"]),
		DefaultMaximumIdleMinutes: core.CastInt32(data["defaultMaximumIdleMinutes"]),
		AcquireActions:            CastAcquireActionList(core.CastArray(data["acquireActions"])),
		IdlePeriodScheduleId:      core.CastString(data["idlePeriodScheduleId"]),
		ReceivePeriodScheduleId:   core.CastString(data["receivePeriodScheduleId"]),
	}
}

func (p CreateCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"name":                      p.Name,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"rewardIntervalMinutes":     p.RewardIntervalMinutes,
		"defaultMaximumIdleMinutes": p.DefaultMaximumIdleMinutes,
		"acquireActions": CastAcquireActionListFromDict(
			p.AcquireActions,
		),
		"idlePeriodScheduleId":    p.IdlePeriodScheduleId,
		"receivePeriodScheduleId": p.ReceivePeriodScheduleId,
	}
}

func (p CreateCategoryModelMasterRequest) Pointer() *CreateCategoryModelMasterRequest {
	return &p
}

type GetCategoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	CategoryName  *string `json:"categoryName"`
}

func NewGetCategoryModelMasterRequestFromJson(data string) GetCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelMasterRequestFromDict(dict)
}

func NewGetCategoryModelMasterRequestFromDict(data map[string]interface{}) GetCategoryModelMasterRequest {
	return GetCategoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p GetCategoryModelMasterRequest) Pointer() *GetCategoryModelMasterRequest {
	return &p
}

type UpdateCategoryModelMasterRequest struct {
	RequestId                 *string             `json:"requestId"`
	ContextStack              *string             `json:"contextStack"`
	NamespaceName             *string             `json:"namespaceName"`
	CategoryName              *string             `json:"categoryName"`
	Description               *string             `json:"description"`
	Metadata                  *string             `json:"metadata"`
	RewardIntervalMinutes     *int32              `json:"rewardIntervalMinutes"`
	DefaultMaximumIdleMinutes *int32              `json:"defaultMaximumIdleMinutes"`
	AcquireActions            []AcquireActionList `json:"acquireActions"`
	IdlePeriodScheduleId      *string             `json:"idlePeriodScheduleId"`
	ReceivePeriodScheduleId   *string             `json:"receivePeriodScheduleId"`
}

func NewUpdateCategoryModelMasterRequestFromJson(data string) UpdateCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCategoryModelMasterRequestFromDict(dict)
}

func NewUpdateCategoryModelMasterRequestFromDict(data map[string]interface{}) UpdateCategoryModelMasterRequest {
	return UpdateCategoryModelMasterRequest{
		NamespaceName:             core.CastString(data["namespaceName"]),
		CategoryName:              core.CastString(data["categoryName"]),
		Description:               core.CastString(data["description"]),
		Metadata:                  core.CastString(data["metadata"]),
		RewardIntervalMinutes:     core.CastInt32(data["rewardIntervalMinutes"]),
		DefaultMaximumIdleMinutes: core.CastInt32(data["defaultMaximumIdleMinutes"]),
		AcquireActions:            CastAcquireActionList(core.CastArray(data["acquireActions"])),
		IdlePeriodScheduleId:      core.CastString(data["idlePeriodScheduleId"]),
		ReceivePeriodScheduleId:   core.CastString(data["receivePeriodScheduleId"]),
	}
}

func (p UpdateCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"categoryName":              p.CategoryName,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"rewardIntervalMinutes":     p.RewardIntervalMinutes,
		"defaultMaximumIdleMinutes": p.DefaultMaximumIdleMinutes,
		"acquireActions": CastAcquireActionListFromDict(
			p.AcquireActions,
		),
		"idlePeriodScheduleId":    p.IdlePeriodScheduleId,
		"receivePeriodScheduleId": p.ReceivePeriodScheduleId,
	}
}

func (p UpdateCategoryModelMasterRequest) Pointer() *UpdateCategoryModelMasterRequest {
	return &p
}

type DeleteCategoryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	CategoryName  *string `json:"categoryName"`
}

func NewDeleteCategoryModelMasterRequestFromJson(data string) DeleteCategoryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCategoryModelMasterRequestFromDict(dict)
}

func NewDeleteCategoryModelMasterRequestFromDict(data map[string]interface{}) DeleteCategoryModelMasterRequest {
	return DeleteCategoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p DeleteCategoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p DeleteCategoryModelMasterRequest) Pointer() *DeleteCategoryModelMasterRequest {
	return &p
}

type DescribeCategoryModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeCategoryModelsRequestFromJson(data string) DescribeCategoryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelsRequestFromDict(dict)
}

func NewDescribeCategoryModelsRequestFromDict(data map[string]interface{}) DescribeCategoryModelsRequest {
	return DescribeCategoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeCategoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeCategoryModelsRequest) Pointer() *DescribeCategoryModelsRequest {
	return &p
}

type GetCategoryModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	CategoryName  *string `json:"categoryName"`
}

func NewGetCategoryModelRequestFromJson(data string) GetCategoryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelRequestFromDict(dict)
}

func NewGetCategoryModelRequestFromDict(data map[string]interface{}) GetCategoryModelRequest {
	return GetCategoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetCategoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p GetCategoryModelRequest) Pointer() *GetCategoryModelRequest {
	return &p
}

type DescribeStatusesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStatusesRequestFromJson(data string) DescribeStatusesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesRequestFromDict(dict)
}

func NewDescribeStatusesRequestFromDict(data map[string]interface{}) DescribeStatusesRequest {
	return DescribeStatusesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
	return &p
}

type DescribeStatusesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeStatusesByUserIdRequestFromJson(data string) DescribeStatusesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeStatusesByUserIdRequest {
	return DescribeStatusesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusesByUserIdRequest) Pointer() *DescribeStatusesByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	CategoryName  *string `json:"categoryName"`
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
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"categoryName":  p.CategoryName,
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
	CategoryName  *string `json:"categoryName"`
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
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"categoryName":  p.CategoryName,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type PredictionRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	CategoryName       *string `json:"categoryName"`
}

func NewPredictionRequestFromJson(data string) PredictionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionRequestFromDict(dict)
}

func NewPredictionRequestFromDict(data map[string]interface{}) PredictionRequest {
	return PredictionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p PredictionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"categoryName":  p.CategoryName,
	}
}

func (p PredictionRequest) Pointer() *PredictionRequest {
	return &p
}

type PredictionByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	CategoryName       *string `json:"categoryName"`
}

func NewPredictionByUserIdRequestFromJson(data string) PredictionByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionByUserIdRequestFromDict(dict)
}

func NewPredictionByUserIdRequestFromDict(data map[string]interface{}) PredictionByUserIdRequest {
	return PredictionByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p PredictionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"categoryName":  p.CategoryName,
	}
}

func (p PredictionByUserIdRequest) Pointer() *PredictionByUserIdRequest {
	return &p
}

type ReceiveRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	CategoryName       *string  `json:"categoryName"`
	Config             []Config `json:"config"`
}

func NewReceiveRequestFromJson(data string) ReceiveRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveRequestFromDict(dict)
}

func NewReceiveRequestFromDict(data map[string]interface{}) ReceiveRequest {
	return ReceiveRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		CategoryName:  core.CastString(data["categoryName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReceiveRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"categoryName":  p.CategoryName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReceiveRequest) Pointer() *ReceiveRequest {
	return &p
}

type ReceiveByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	CategoryName       *string  `json:"categoryName"`
	Config             []Config `json:"config"`
}

func NewReceiveByUserIdRequestFromJson(data string) ReceiveByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByUserIdRequestFromDict(dict)
}

func NewReceiveByUserIdRequestFromDict(data map[string]interface{}) ReceiveByUserIdRequest {
	return ReceiveByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		CategoryName:  core.CastString(data["categoryName"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p ReceiveByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"categoryName":  p.CategoryName,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReceiveByUserIdRequest) Pointer() *ReceiveByUserIdRequest {
	return &p
}

type IncreaseMaximumIdleMinutesByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	CategoryName       *string `json:"categoryName"`
	IncreaseMinutes    *int32  `json:"increaseMinutes"`
}

func NewIncreaseMaximumIdleMinutesByUserIdRequestFromJson(data string) IncreaseMaximumIdleMinutesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumIdleMinutesByUserIdRequestFromDict(dict)
}

func NewIncreaseMaximumIdleMinutesByUserIdRequestFromDict(data map[string]interface{}) IncreaseMaximumIdleMinutesByUserIdRequest {
	return IncreaseMaximumIdleMinutesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		CategoryName:    core.CastString(data["categoryName"]),
		IncreaseMinutes: core.CastInt32(data["increaseMinutes"]),
	}
}

func (p IncreaseMaximumIdleMinutesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"categoryName":    p.CategoryName,
		"increaseMinutes": p.IncreaseMinutes,
	}
}

func (p IncreaseMaximumIdleMinutesByUserIdRequest) Pointer() *IncreaseMaximumIdleMinutesByUserIdRequest {
	return &p
}

type DecreaseMaximumIdleMinutesByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	CategoryName       *string `json:"categoryName"`
	DecreaseMinutes    *int32  `json:"decreaseMinutes"`
}

func NewDecreaseMaximumIdleMinutesByUserIdRequestFromJson(data string) DecreaseMaximumIdleMinutesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumIdleMinutesByUserIdRequestFromDict(dict)
}

func NewDecreaseMaximumIdleMinutesByUserIdRequestFromDict(data map[string]interface{}) DecreaseMaximumIdleMinutesByUserIdRequest {
	return DecreaseMaximumIdleMinutesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		CategoryName:    core.CastString(data["categoryName"]),
		DecreaseMinutes: core.CastInt32(data["decreaseMinutes"]),
	}
}

func (p DecreaseMaximumIdleMinutesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"categoryName":    p.CategoryName,
		"decreaseMinutes": p.DecreaseMinutes,
	}
}

func (p DecreaseMaximumIdleMinutesByUserIdRequest) Pointer() *DecreaseMaximumIdleMinutesByUserIdRequest {
	return &p
}

type IncreaseMaximumIdleMinutesByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewIncreaseMaximumIdleMinutesByStampSheetRequestFromJson(data string) IncreaseMaximumIdleMinutesByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumIdleMinutesByStampSheetRequestFromDict(dict)
}

func NewIncreaseMaximumIdleMinutesByStampSheetRequestFromDict(data map[string]interface{}) IncreaseMaximumIdleMinutesByStampSheetRequest {
	return IncreaseMaximumIdleMinutesByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p IncreaseMaximumIdleMinutesByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p IncreaseMaximumIdleMinutesByStampSheetRequest) Pointer() *IncreaseMaximumIdleMinutesByStampSheetRequest {
	return &p
}

type DecreaseMaximumIdleMinutesByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewDecreaseMaximumIdleMinutesByStampTaskRequestFromJson(data string) DecreaseMaximumIdleMinutesByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumIdleMinutesByStampTaskRequestFromDict(dict)
}

func NewDecreaseMaximumIdleMinutesByStampTaskRequestFromDict(data map[string]interface{}) DecreaseMaximumIdleMinutesByStampTaskRequest {
	return DecreaseMaximumIdleMinutesByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DecreaseMaximumIdleMinutesByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DecreaseMaximumIdleMinutesByStampTaskRequest) Pointer() *DecreaseMaximumIdleMinutesByStampTaskRequest {
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

type GetCurrentCategoryMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentCategoryMasterRequestFromJson(data string) GetCurrentCategoryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentCategoryMasterRequestFromDict(dict)
}

func NewGetCurrentCategoryMasterRequestFromDict(data map[string]interface{}) GetCurrentCategoryMasterRequest {
	return GetCurrentCategoryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentCategoryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentCategoryMasterRequest) Pointer() *GetCurrentCategoryMasterRequest {
	return &p
}

type UpdateCurrentCategoryMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentCategoryMasterRequestFromJson(data string) UpdateCurrentCategoryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCategoryMasterRequestFromDict(dict)
}

func NewUpdateCurrentCategoryMasterRequestFromDict(data map[string]interface{}) UpdateCurrentCategoryMasterRequest {
	return UpdateCurrentCategoryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentCategoryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentCategoryMasterRequest) Pointer() *UpdateCurrentCategoryMasterRequest {
	return &p
}

type UpdateCurrentCategoryMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentCategoryMasterFromGitHubRequestFromJson(data string) UpdateCurrentCategoryMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCategoryMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentCategoryMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentCategoryMasterFromGitHubRequest {
	return UpdateCurrentCategoryMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentCategoryMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentCategoryMasterFromGitHubRequest) Pointer() *UpdateCurrentCategoryMasterFromGitHubRequest {
	return &p
}
