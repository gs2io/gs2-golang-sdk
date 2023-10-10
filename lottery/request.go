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

package lottery

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
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	Name                     *string             `json:"name"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	LotteryTriggerScriptId   *string             `json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *string             `json:"choicePrizeTableScriptId"`
	LogSetting               *LogSetting         `json:"logSetting"`
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
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		LotteryTriggerScriptId:   core.CastString(data["lotteryTriggerScriptId"]),
		ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:         core.CastString(data["queueNamespaceId"]),
		KeyId:                    core.CastString(data["keyId"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                     p.Name,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"lotteryTriggerScriptId":   p.LotteryTriggerScriptId,
		"choicePrizeTableScriptId": p.ChoicePrizeTableScriptId,
		"logSetting":               p.LogSetting.ToDict(),
		"queueNamespaceId":         p.QueueNamespaceId,
		"keyId":                    p.KeyId,
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
	RequestId                *string             `json:"requestId"`
	ContextStack             *string             `json:"contextStack"`
	NamespaceName            *string             `json:"namespaceName"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	LotteryTriggerScriptId   *string             `json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *string             `json:"choicePrizeTableScriptId"`
	LogSetting               *LogSetting         `json:"logSetting"`
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
		NamespaceName:            core.CastString(data["namespaceName"]),
		Description:              core.CastString(data["description"]),
		TransactionSetting:       NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		LotteryTriggerScriptId:   core.CastString(data["lotteryTriggerScriptId"]),
		ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:         core.CastString(data["queueNamespaceId"]),
		KeyId:                    core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"description":              p.Description,
		"transactionSetting":       p.TransactionSetting.ToDict(),
		"lotteryTriggerScriptId":   p.LotteryTriggerScriptId,
		"choicePrizeTableScriptId": p.ChoicePrizeTableScriptId,
		"logSetting":               p.LogSetting.ToDict(),
		"queueNamespaceId":         p.QueueNamespaceId,
		"keyId":                    p.KeyId,
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

type DescribeLotteryModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeLotteryModelMastersRequestFromJson(data string) DescribeLotteryModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryModelMastersRequestFromDict(dict)
}

func NewDescribeLotteryModelMastersRequestFromDict(data map[string]interface{}) DescribeLotteryModelMastersRequest {
	return DescribeLotteryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeLotteryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeLotteryModelMastersRequest) Pointer() *DescribeLotteryModelMastersRequest {
	return &p
}

type CreateLotteryModelMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	Name                     *string `json:"name"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	Mode                     *string `json:"mode"`
	Method                   *string `json:"method"`
	PrizeTableName           *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
}

func NewCreateLotteryModelMasterRequestFromJson(data string) CreateLotteryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateLotteryModelMasterRequestFromDict(dict)
}

func NewCreateLotteryModelMasterRequestFromDict(data map[string]interface{}) CreateLotteryModelMasterRequest {
	return CreateLotteryModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		Mode:                     core.CastString(data["mode"]),
		Method:                   core.CastString(data["method"]),
		PrizeTableName:           core.CastString(data["prizeTableName"]),
		ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
	}
}

func (p CreateLotteryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"mode":                     p.Mode,
		"method":                   p.Method,
		"prizeTableName":           p.PrizeTableName,
		"choicePrizeTableScriptId": p.ChoicePrizeTableScriptId,
	}
}

func (p CreateLotteryModelMasterRequest) Pointer() *CreateLotteryModelMasterRequest {
	return &p
}

type GetLotteryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LotteryName   *string `json:"lotteryName"`
}

func NewGetLotteryModelMasterRequestFromJson(data string) GetLotteryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryModelMasterRequestFromDict(dict)
}

func NewGetLotteryModelMasterRequestFromDict(data map[string]interface{}) GetLotteryModelMasterRequest {
	return GetLotteryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
	}
}

func (p GetLotteryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
	}
}

func (p GetLotteryModelMasterRequest) Pointer() *GetLotteryModelMasterRequest {
	return &p
}

type UpdateLotteryModelMasterRequest struct {
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	LotteryName              *string `json:"lotteryName"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	Mode                     *string `json:"mode"`
	Method                   *string `json:"method"`
	PrizeTableName           *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
}

func NewUpdateLotteryModelMasterRequestFromJson(data string) UpdateLotteryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateLotteryModelMasterRequestFromDict(dict)
}

func NewUpdateLotteryModelMasterRequestFromDict(data map[string]interface{}) UpdateLotteryModelMasterRequest {
	return UpdateLotteryModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		LotteryName:              core.CastString(data["lotteryName"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		Mode:                     core.CastString(data["mode"]),
		Method:                   core.CastString(data["method"]),
		PrizeTableName:           core.CastString(data["prizeTableName"]),
		ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
	}
}

func (p UpdateLotteryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"lotteryName":              p.LotteryName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"mode":                     p.Mode,
		"method":                   p.Method,
		"prizeTableName":           p.PrizeTableName,
		"choicePrizeTableScriptId": p.ChoicePrizeTableScriptId,
	}
}

func (p UpdateLotteryModelMasterRequest) Pointer() *UpdateLotteryModelMasterRequest {
	return &p
}

type DeleteLotteryModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LotteryName   *string `json:"lotteryName"`
}

func NewDeleteLotteryModelMasterRequestFromJson(data string) DeleteLotteryModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLotteryModelMasterRequestFromDict(dict)
}

func NewDeleteLotteryModelMasterRequestFromDict(data map[string]interface{}) DeleteLotteryModelMasterRequest {
	return DeleteLotteryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
	}
}

func (p DeleteLotteryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
	}
}

func (p DeleteLotteryModelMasterRequest) Pointer() *DeleteLotteryModelMasterRequest {
	return &p
}

type DescribePrizeTableMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribePrizeTableMastersRequestFromJson(data string) DescribePrizeTableMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeTableMastersRequestFromDict(dict)
}

func NewDescribePrizeTableMastersRequestFromDict(data map[string]interface{}) DescribePrizeTableMastersRequest {
	return DescribePrizeTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribePrizeTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribePrizeTableMastersRequest) Pointer() *DescribePrizeTableMastersRequest {
	return &p
}

type CreatePrizeTableMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
	Prizes        []Prize `json:"prizes"`
}

func NewCreatePrizeTableMasterRequestFromJson(data string) CreatePrizeTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePrizeTableMasterRequestFromDict(dict)
}

func NewCreatePrizeTableMasterRequestFromDict(data map[string]interface{}) CreatePrizeTableMasterRequest {
	return CreatePrizeTableMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		Prizes:        CastPrizes(core.CastArray(data["prizes"])),
	}
}

func (p CreatePrizeTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"prizes": CastPrizesFromDict(
			p.Prizes,
		),
	}
}

func (p CreatePrizeTableMasterRequest) Pointer() *CreatePrizeTableMasterRequest {
	return &p
}

type GetPrizeTableMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
}

func NewGetPrizeTableMasterRequestFromJson(data string) GetPrizeTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeTableMasterRequestFromDict(dict)
}

func NewGetPrizeTableMasterRequestFromDict(data map[string]interface{}) GetPrizeTableMasterRequest {
	return GetPrizeTableMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
	}
}

func (p GetPrizeTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
	}
}

func (p GetPrizeTableMasterRequest) Pointer() *GetPrizeTableMasterRequest {
	return &p
}

type UpdatePrizeTableMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	Description    *string `json:"description"`
	Metadata       *string `json:"metadata"`
	Prizes         []Prize `json:"prizes"`
}

func NewUpdatePrizeTableMasterRequestFromJson(data string) UpdatePrizeTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdatePrizeTableMasterRequestFromDict(dict)
}

func NewUpdatePrizeTableMasterRequestFromDict(data map[string]interface{}) UpdatePrizeTableMasterRequest {
	return UpdatePrizeTableMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		Prizes:         CastPrizes(core.CastArray(data["prizes"])),
	}
}

func (p UpdatePrizeTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"prizes": CastPrizesFromDict(
			p.Prizes,
		),
	}
}

func (p UpdatePrizeTableMasterRequest) Pointer() *UpdatePrizeTableMasterRequest {
	return &p
}

type DeletePrizeTableMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
}

func NewDeletePrizeTableMasterRequestFromJson(data string) DeletePrizeTableMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePrizeTableMasterRequestFromDict(dict)
}

func NewDeletePrizeTableMasterRequestFromDict(data map[string]interface{}) DeletePrizeTableMasterRequest {
	return DeletePrizeTableMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
	}
}

func (p DeletePrizeTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
	}
}

func (p DeletePrizeTableMasterRequest) Pointer() *DeletePrizeTableMasterRequest {
	return &p
}

type DescribeLotteryModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeLotteryModelsRequestFromJson(data string) DescribeLotteryModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryModelsRequestFromDict(dict)
}

func NewDescribeLotteryModelsRequestFromDict(data map[string]interface{}) DescribeLotteryModelsRequest {
	return DescribeLotteryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeLotteryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeLotteryModelsRequest) Pointer() *DescribeLotteryModelsRequest {
	return &p
}

type GetLotteryModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LotteryName   *string `json:"lotteryName"`
}

func NewGetLotteryModelRequestFromJson(data string) GetLotteryModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryModelRequestFromDict(dict)
}

func NewGetLotteryModelRequestFromDict(data map[string]interface{}) GetLotteryModelRequest {
	return GetLotteryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
	}
}

func (p GetLotteryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
	}
}

func (p GetLotteryModelRequest) Pointer() *GetLotteryModelRequest {
	return &p
}

type DescribePrizeTablesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribePrizeTablesRequestFromJson(data string) DescribePrizeTablesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeTablesRequestFromDict(dict)
}

func NewDescribePrizeTablesRequestFromDict(data map[string]interface{}) DescribePrizeTablesRequest {
	return DescribePrizeTablesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribePrizeTablesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribePrizeTablesRequest) Pointer() *DescribePrizeTablesRequest {
	return &p
}

type GetPrizeTableRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
}

func NewGetPrizeTableRequestFromJson(data string) GetPrizeTableRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeTableRequestFromDict(dict)
}

func NewGetPrizeTableRequestFromDict(data map[string]interface{}) GetPrizeTableRequest {
	return GetPrizeTableRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
	}
}

func (p GetPrizeTableRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
	}
}

func (p GetPrizeTableRequest) Pointer() *GetPrizeTableRequest {
	return &p
}

type DrawByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	LotteryName        *string  `json:"lotteryName"`
	UserId             *string  `json:"userId"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewDrawByUserIdRequestFromJson(data string) DrawByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawByUserIdRequestFromDict(dict)
}

func NewDrawByUserIdRequestFromDict(data map[string]interface{}) DrawByUserIdRequest {
	return DrawByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		UserId:        core.CastString(data["userId"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p DrawByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"userId":        p.UserId,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p DrawByUserIdRequest) Pointer() *DrawByUserIdRequest {
	return &p
}

type PredictionRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	LotteryName        *string `json:"lotteryName"`
	AccessToken        *string `json:"accessToken"`
	RandomSeed         *int64  `json:"randomSeed"`
	Count              *int32  `json:"count"`
}

func NewPredictionRequestFromJson(data string) PredictionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionRequestFromDict(dict)
}

func NewPredictionRequestFromDict(data map[string]interface{}) PredictionRequest {
	return PredictionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		RandomSeed:    core.CastInt64(data["randomSeed"]),
		Count:         core.CastInt32(data["count"]),
	}
}

func (p PredictionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"accessToken":   p.AccessToken,
		"randomSeed":    p.RandomSeed,
		"count":         p.Count,
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
	LotteryName        *string `json:"lotteryName"`
	UserId             *string `json:"userId"`
	RandomSeed         *int64  `json:"randomSeed"`
	Count              *int32  `json:"count"`
}

func NewPredictionByUserIdRequestFromJson(data string) PredictionByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionByUserIdRequestFromDict(dict)
}

func NewPredictionByUserIdRequestFromDict(data map[string]interface{}) PredictionByUserIdRequest {
	return PredictionByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		UserId:        core.CastString(data["userId"]),
		RandomSeed:    core.CastInt64(data["randomSeed"]),
		Count:         core.CastInt32(data["count"]),
	}
}

func (p PredictionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"userId":        p.UserId,
		"randomSeed":    p.RandomSeed,
		"count":         p.Count,
	}
}

func (p PredictionByUserIdRequest) Pointer() *PredictionByUserIdRequest {
	return &p
}

type DrawWithRandomSeedByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	LotteryName        *string  `json:"lotteryName"`
	UserId             *string  `json:"userId"`
	RandomSeed         *int64   `json:"randomSeed"`
	Count              *int32   `json:"count"`
	Config             []Config `json:"config"`
}

func NewDrawWithRandomSeedByUserIdRequestFromJson(data string) DrawWithRandomSeedByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawWithRandomSeedByUserIdRequestFromDict(dict)
}

func NewDrawWithRandomSeedByUserIdRequestFromDict(data map[string]interface{}) DrawWithRandomSeedByUserIdRequest {
	return DrawWithRandomSeedByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		UserId:        core.CastString(data["userId"]),
		RandomSeed:    core.CastInt64(data["randomSeed"]),
		Count:         core.CastInt32(data["count"]),
		Config:        CastConfigs(core.CastArray(data["config"])),
	}
}

func (p DrawWithRandomSeedByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"userId":        p.UserId,
		"randomSeed":    p.RandomSeed,
		"count":         p.Count,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p DrawWithRandomSeedByUserIdRequest) Pointer() *DrawWithRandomSeedByUserIdRequest {
	return &p
}

type DrawByStampSheetRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewDrawByStampSheetRequestFromJson(data string) DrawByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawByStampSheetRequestFromDict(dict)
}

func NewDrawByStampSheetRequestFromDict(data map[string]interface{}) DrawByStampSheetRequest {
	return DrawByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DrawByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DrawByStampSheetRequest) Pointer() *DrawByStampSheetRequest {
	return &p
}

type DescribeProbabilitiesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LotteryName   *string `json:"lotteryName"`
	AccessToken   *string `json:"accessToken"`
}

func NewDescribeProbabilitiesRequestFromJson(data string) DescribeProbabilitiesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProbabilitiesRequestFromDict(dict)
}

func NewDescribeProbabilitiesRequestFromDict(data map[string]interface{}) DescribeProbabilitiesRequest {
	return DescribeProbabilitiesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DescribeProbabilitiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeProbabilitiesRequest) Pointer() *DescribeProbabilitiesRequest {
	return &p
}

type DescribeProbabilitiesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	LotteryName   *string `json:"lotteryName"`
	UserId        *string `json:"userId"`
}

func NewDescribeProbabilitiesByUserIdRequestFromJson(data string) DescribeProbabilitiesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProbabilitiesByUserIdRequestFromDict(dict)
}

func NewDescribeProbabilitiesByUserIdRequestFromDict(data map[string]interface{}) DescribeProbabilitiesByUserIdRequest {
	return DescribeProbabilitiesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DescribeProbabilitiesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
		"userId":        p.UserId,
	}
}

func (p DescribeProbabilitiesByUserIdRequest) Pointer() *DescribeProbabilitiesByUserIdRequest {
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

type GetCurrentLotteryMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentLotteryMasterRequestFromJson(data string) GetCurrentLotteryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentLotteryMasterRequestFromDict(dict)
}

func NewGetCurrentLotteryMasterRequestFromDict(data map[string]interface{}) GetCurrentLotteryMasterRequest {
	return GetCurrentLotteryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentLotteryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentLotteryMasterRequest) Pointer() *GetCurrentLotteryMasterRequest {
	return &p
}

type UpdateCurrentLotteryMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentLotteryMasterRequestFromJson(data string) UpdateCurrentLotteryMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLotteryMasterRequestFromDict(dict)
}

func NewUpdateCurrentLotteryMasterRequestFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterRequest {
	return UpdateCurrentLotteryMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentLotteryMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentLotteryMasterRequest) Pointer() *UpdateCurrentLotteryMasterRequest {
	return &p
}

type UpdateCurrentLotteryMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentLotteryMasterFromGitHubRequestFromJson(data string) UpdateCurrentLotteryMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLotteryMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentLotteryMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterFromGitHubRequest {
	return UpdateCurrentLotteryMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentLotteryMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentLotteryMasterFromGitHubRequest) Pointer() *UpdateCurrentLotteryMasterFromGitHubRequest {
	return &p
}

type DescribePrizeLimitsRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	PageToken      *string `json:"pageToken"`
	Limit          *int32  `json:"limit"`
}

func NewDescribePrizeLimitsRequestFromJson(data string) DescribePrizeLimitsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeLimitsRequestFromDict(dict)
}

func NewDescribePrizeLimitsRequestFromDict(data map[string]interface{}) DescribePrizeLimitsRequest {
	return DescribePrizeLimitsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribePrizeLimitsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribePrizeLimitsRequest) Pointer() *DescribePrizeLimitsRequest {
	return &p
}

type GetPrizeLimitRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	PrizeId        *string `json:"prizeId"`
}

func NewGetPrizeLimitRequestFromJson(data string) GetPrizeLimitRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeLimitRequestFromDict(dict)
}

func NewGetPrizeLimitRequestFromDict(data map[string]interface{}) GetPrizeLimitRequest {
	return GetPrizeLimitRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		PrizeId:        core.CastString(data["prizeId"]),
	}
}

func (p GetPrizeLimitRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"prizeId":        p.PrizeId,
	}
}

func (p GetPrizeLimitRequest) Pointer() *GetPrizeLimitRequest {
	return &p
}

type ResetPrizeLimitRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	PrizeId        *string `json:"prizeId"`
}

func NewResetPrizeLimitRequestFromJson(data string) ResetPrizeLimitRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetPrizeLimitRequestFromDict(dict)
}

func NewResetPrizeLimitRequestFromDict(data map[string]interface{}) ResetPrizeLimitRequest {
	return ResetPrizeLimitRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		PrizeId:        core.CastString(data["prizeId"]),
	}
}

func (p ResetPrizeLimitRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"prizeId":        p.PrizeId,
	}
}

func (p ResetPrizeLimitRequest) Pointer() *ResetPrizeLimitRequest {
	return &p
}

type DescribeBoxesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBoxesRequestFromJson(data string) DescribeBoxesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBoxesRequestFromDict(dict)
}

func NewDescribeBoxesRequestFromDict(data map[string]interface{}) DescribeBoxesRequest {
	return DescribeBoxesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBoxesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBoxesRequest) Pointer() *DescribeBoxesRequest {
	return &p
}

type DescribeBoxesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeBoxesByUserIdRequestFromJson(data string) DescribeBoxesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBoxesByUserIdRequestFromDict(dict)
}

func NewDescribeBoxesByUserIdRequestFromDict(data map[string]interface{}) DescribeBoxesByUserIdRequest {
	return DescribeBoxesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBoxesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBoxesByUserIdRequest) Pointer() *DescribeBoxesByUserIdRequest {
	return &p
}

type GetBoxRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	AccessToken    *string `json:"accessToken"`
}

func NewGetBoxRequestFromJson(data string) GetBoxRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBoxRequestFromDict(dict)
}

func NewGetBoxRequestFromDict(data map[string]interface{}) GetBoxRequest {
	return GetBoxRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		AccessToken:    core.CastString(data["accessToken"]),
	}
}

func (p GetBoxRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"accessToken":    p.AccessToken,
	}
}

func (p GetBoxRequest) Pointer() *GetBoxRequest {
	return &p
}

type GetBoxByUserIdRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	PrizeTableName *string `json:"prizeTableName"`
	UserId         *string `json:"userId"`
}

func NewGetBoxByUserIdRequestFromJson(data string) GetBoxByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBoxByUserIdRequestFromDict(dict)
}

func NewGetBoxByUserIdRequestFromDict(data map[string]interface{}) GetBoxByUserIdRequest {
	return GetBoxByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		UserId:         core.CastString(data["userId"]),
	}
}

func (p GetBoxByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"userId":         p.UserId,
	}
}

func (p GetBoxByUserIdRequest) Pointer() *GetBoxByUserIdRequest {
	return &p
}

type ResetBoxRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PrizeTableName     *string `json:"prizeTableName"`
	AccessToken        *string `json:"accessToken"`
}

func NewResetBoxRequestFromJson(data string) ResetBoxRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetBoxRequestFromDict(dict)
}

func NewResetBoxRequestFromDict(data map[string]interface{}) ResetBoxRequest {
	return ResetBoxRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		AccessToken:    core.CastString(data["accessToken"]),
	}
}

func (p ResetBoxRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"accessToken":    p.AccessToken,
	}
}

func (p ResetBoxRequest) Pointer() *ResetBoxRequest {
	return &p
}

type ResetBoxByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PrizeTableName     *string `json:"prizeTableName"`
	UserId             *string `json:"userId"`
}

func NewResetBoxByUserIdRequestFromJson(data string) ResetBoxByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetBoxByUserIdRequestFromDict(dict)
}

func NewResetBoxByUserIdRequestFromDict(data map[string]interface{}) ResetBoxByUserIdRequest {
	return ResetBoxByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		PrizeTableName: core.CastString(data["prizeTableName"]),
		UserId:         core.CastString(data["userId"]),
	}
}

func (p ResetBoxByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"prizeTableName": p.PrizeTableName,
		"userId":         p.UserId,
	}
}

func (p ResetBoxByUserIdRequest) Pointer() *ResetBoxByUserIdRequest {
	return &p
}
