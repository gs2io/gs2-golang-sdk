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

package seasonRating

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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
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
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
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
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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

type DumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeMatchSessionsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMatchSessionsRequestFromJson(data string) DescribeMatchSessionsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMatchSessionsRequestFromDict(dict)
}

func NewDescribeMatchSessionsRequestFromDict(data map[string]interface{}) DescribeMatchSessionsRequest {
	return DescribeMatchSessionsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMatchSessionsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMatchSessionsRequest) Pointer() *DescribeMatchSessionsRequest {
	return &p
}

type CreateMatchSessionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SessionName     *string `json:"sessionName"`
	TtlSeconds      *int32  `json:"ttlSeconds"`
}

func NewCreateMatchSessionRequestFromJson(data string) CreateMatchSessionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMatchSessionRequestFromDict(dict)
}

func NewCreateMatchSessionRequestFromDict(data map[string]interface{}) CreateMatchSessionRequest {
	return CreateMatchSessionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SessionName:   core.CastString(data["sessionName"]),
		TtlSeconds:    core.CastInt32(data["ttlSeconds"]),
	}
}

func (p CreateMatchSessionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"sessionName":   p.SessionName,
		"ttlSeconds":    p.TtlSeconds,
	}
}

func (p CreateMatchSessionRequest) Pointer() *CreateMatchSessionRequest {
	return &p
}

type GetMatchSessionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SessionName     *string `json:"sessionName"`
}

func NewGetMatchSessionRequestFromJson(data string) GetMatchSessionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMatchSessionRequestFromDict(dict)
}

func NewGetMatchSessionRequestFromDict(data map[string]interface{}) GetMatchSessionRequest {
	return GetMatchSessionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SessionName:   core.CastString(data["sessionName"]),
	}
}

func (p GetMatchSessionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"sessionName":   p.SessionName,
	}
}

func (p GetMatchSessionRequest) Pointer() *GetMatchSessionRequest {
	return &p
}

type DeleteMatchSessionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SessionName     *string `json:"sessionName"`
}

func NewDeleteMatchSessionRequestFromJson(data string) DeleteMatchSessionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMatchSessionRequestFromDict(dict)
}

func NewDeleteMatchSessionRequestFromDict(data map[string]interface{}) DeleteMatchSessionRequest {
	return DeleteMatchSessionRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SessionName:   core.CastString(data["sessionName"]),
	}
}

func (p DeleteMatchSessionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"sessionName":   p.SessionName,
	}
}

func (p DeleteMatchSessionRequest) Pointer() *DeleteMatchSessionRequest {
	return &p
}

type DescribeSeasonModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSeasonModelMastersRequestFromJson(data string) DescribeSeasonModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelMastersRequestFromDict(dict)
}

func NewDescribeSeasonModelMastersRequestFromDict(data map[string]interface{}) DescribeSeasonModelMastersRequest {
	return DescribeSeasonModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSeasonModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSeasonModelMastersRequest) Pointer() *DescribeSeasonModelMastersRequest {
	return &p
}

type CreateSeasonModelMasterRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	NamespaceName          *string     `json:"namespaceName"`
	Name                   *string     `json:"name"`
	Description            *string     `json:"description"`
	Metadata               *string     `json:"metadata"`
	Tiers                  []TierModel `json:"tiers"`
	ExperienceModelId      *string     `json:"experienceModelId"`
	ChallengePeriodEventId *string     `json:"challengePeriodEventId"`
}

func NewCreateSeasonModelMasterRequestFromJson(data string) CreateSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSeasonModelMasterRequestFromDict(dict)
}

func NewCreateSeasonModelMasterRequestFromDict(data map[string]interface{}) CreateSeasonModelMasterRequest {
	return CreateSeasonModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		Tiers:                  CastTierModels(core.CastArray(data["tiers"])),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p CreateSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"tiers": CastTierModelsFromDict(
			p.Tiers,
		),
		"experienceModelId":      p.ExperienceModelId,
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p CreateSeasonModelMasterRequest) Pointer() *CreateSeasonModelMasterRequest {
	return &p
}

type GetSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func NewGetSeasonModelMasterRequestFromJson(data string) GetSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelMasterRequestFromDict(dict)
}

func NewGetSeasonModelMasterRequestFromDict(data map[string]interface{}) GetSeasonModelMasterRequest {
	return GetSeasonModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SeasonName:    core.CastString(data["seasonName"]),
	}
}

func (p GetSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p GetSeasonModelMasterRequest) Pointer() *GetSeasonModelMasterRequest {
	return &p
}

type UpdateSeasonModelMasterRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	NamespaceName          *string     `json:"namespaceName"`
	SeasonName             *string     `json:"seasonName"`
	Description            *string     `json:"description"`
	Metadata               *string     `json:"metadata"`
	Tiers                  []TierModel `json:"tiers"`
	ExperienceModelId      *string     `json:"experienceModelId"`
	ChallengePeriodEventId *string     `json:"challengePeriodEventId"`
}

func NewUpdateSeasonModelMasterRequestFromJson(data string) UpdateSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSeasonModelMasterRequestFromDict(dict)
}

func NewUpdateSeasonModelMasterRequestFromDict(data map[string]interface{}) UpdateSeasonModelMasterRequest {
	return UpdateSeasonModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		SeasonName:             core.CastString(data["seasonName"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		Tiers:                  CastTierModels(core.CastArray(data["tiers"])),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p UpdateSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"tiers": CastTierModelsFromDict(
			p.Tiers,
		),
		"experienceModelId":      p.ExperienceModelId,
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p UpdateSeasonModelMasterRequest) Pointer() *UpdateSeasonModelMasterRequest {
	return &p
}

type DeleteSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func NewDeleteSeasonModelMasterRequestFromJson(data string) DeleteSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSeasonModelMasterRequestFromDict(dict)
}

func NewDeleteSeasonModelMasterRequestFromDict(data map[string]interface{}) DeleteSeasonModelMasterRequest {
	return DeleteSeasonModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SeasonName:    core.CastString(data["seasonName"]),
	}
}

func (p DeleteSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p DeleteSeasonModelMasterRequest) Pointer() *DeleteSeasonModelMasterRequest {
	return &p
}

type DescribeSeasonModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeSeasonModelsRequestFromJson(data string) DescribeSeasonModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelsRequestFromDict(dict)
}

func NewDescribeSeasonModelsRequestFromDict(data map[string]interface{}) DescribeSeasonModelsRequest {
	return DescribeSeasonModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeSeasonModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeSeasonModelsRequest) Pointer() *DescribeSeasonModelsRequest {
	return &p
}

type GetSeasonModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func NewGetSeasonModelRequestFromJson(data string) GetSeasonModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelRequestFromDict(dict)
}

func NewGetSeasonModelRequestFromDict(data map[string]interface{}) GetSeasonModelRequest {
	return GetSeasonModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SeasonName:    core.CastString(data["seasonName"]),
	}
}

func (p GetSeasonModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p GetSeasonModelRequest) Pointer() *GetSeasonModelRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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

type GetCurrentSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentSeasonModelMasterRequestFromJson(data string) GetCurrentSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentSeasonModelMasterRequestFromDict(dict)
}

func NewGetCurrentSeasonModelMasterRequestFromDict(data map[string]interface{}) GetCurrentSeasonModelMasterRequest {
	return GetCurrentSeasonModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentSeasonModelMasterRequest) Pointer() *GetCurrentSeasonModelMasterRequest {
	return &p
}

type UpdateCurrentSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentSeasonModelMasterRequestFromJson(data string) UpdateCurrentSeasonModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentSeasonModelMasterRequestFromDict(dict)
}

func NewUpdateCurrentSeasonModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentSeasonModelMasterRequest {
	return UpdateCurrentSeasonModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentSeasonModelMasterRequest) Pointer() *UpdateCurrentSeasonModelMasterRequest {
	return &p
}

type UpdateCurrentSeasonModelMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentSeasonModelMasterFromGitHubRequestFromJson(data string) UpdateCurrentSeasonModelMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentSeasonModelMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentSeasonModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentSeasonModelMasterFromGitHubRequest {
	return UpdateCurrentSeasonModelMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentSeasonModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentSeasonModelMasterFromGitHubRequest) Pointer() *UpdateCurrentSeasonModelMasterFromGitHubRequest {
	return &p
}

type GetBallotRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
	SessionName     *string `json:"sessionName"`
	AccessToken     *string `json:"accessToken"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
}

func NewGetBallotRequestFromJson(data string) GetBallotRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotRequestFromDict(dict)
}

func NewGetBallotRequestFromDict(data map[string]interface{}) GetBallotRequest {
	return GetBallotRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		SeasonName:     core.CastString(data["seasonName"]),
		SessionName:    core.CastString(data["sessionName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetBallotRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"seasonName":     p.SeasonName,
		"sessionName":    p.SessionName,
		"accessToken":    p.AccessToken,
		"numberOfPlayer": p.NumberOfPlayer,
		"keyId":          p.KeyId,
	}
}

func (p GetBallotRequest) Pointer() *GetBallotRequest {
	return &p
}

type GetBallotByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
	SessionName     *string `json:"sessionName"`
	UserId          *string `json:"userId"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetBallotByUserIdRequestFromJson(data string) GetBallotByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotByUserIdRequestFromDict(dict)
}

func NewGetBallotByUserIdRequestFromDict(data map[string]interface{}) GetBallotByUserIdRequest {
	return GetBallotByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		SeasonName:      core.CastString(data["seasonName"]),
		SessionName:     core.CastString(data["sessionName"]),
		UserId:          core.CastString(data["userId"]),
		NumberOfPlayer:  core.CastInt32(data["numberOfPlayer"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetBallotByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"seasonName":      p.SeasonName,
		"sessionName":     p.SessionName,
		"userId":          p.UserId,
		"numberOfPlayer":  p.NumberOfPlayer,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetBallotByUserIdRequest) Pointer() *GetBallotByUserIdRequest {
	return &p
}

type VoteRequest struct {
	SourceRequestId *string      `json:"sourceRequestId"`
	RequestId       *string      `json:"requestId"`
	ContextStack    *string      `json:"contextStack"`
	NamespaceName   *string      `json:"namespaceName"`
	BallotBody      *string      `json:"ballotBody"`
	BallotSignature *string      `json:"ballotSignature"`
	GameResults     []GameResult `json:"gameResults"`
	KeyId           *string      `json:"keyId"`
}

func NewVoteRequestFromJson(data string) VoteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteRequestFromDict(dict)
}

func NewVoteRequestFromDict(data map[string]interface{}) VoteRequest {
	return VoteRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		BallotBody:      core.CastString(data["ballotBody"]),
		BallotSignature: core.CastString(data["ballotSignature"]),
		GameResults:     CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:           core.CastString(data["keyId"]),
	}
}

func (p VoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"ballotBody":      p.BallotBody,
		"ballotSignature": p.BallotSignature,
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteRequest) Pointer() *VoteRequest {
	return &p
}

type VoteMultipleRequest struct {
	SourceRequestId *string        `json:"sourceRequestId"`
	RequestId       *string        `json:"requestId"`
	ContextStack    *string        `json:"contextStack"`
	NamespaceName   *string        `json:"namespaceName"`
	SignedBallots   []SignedBallot `json:"signedBallots"`
	GameResults     []GameResult   `json:"gameResults"`
	KeyId           *string        `json:"keyId"`
}

func NewVoteMultipleRequestFromJson(data string) VoteMultipleRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteMultipleRequestFromDict(dict)
}

func NewVoteMultipleRequestFromDict(data map[string]interface{}) VoteMultipleRequest {
	return VoteMultipleRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SignedBallots: CastSignedBallots(core.CastArray(data["signedBallots"])),
		GameResults:   CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p VoteMultipleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"signedBallots": CastSignedBallotsFromDict(
			p.SignedBallots,
		),
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteMultipleRequest) Pointer() *VoteMultipleRequest {
	return &p
}

type CommitVoteRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
	SessionName     *string `json:"sessionName"`
}

func NewCommitVoteRequestFromJson(data string) CommitVoteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCommitVoteRequestFromDict(dict)
}

func NewCommitVoteRequestFromDict(data map[string]interface{}) CommitVoteRequest {
	return CommitVoteRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SeasonName:    core.CastString(data["seasonName"]),
		SessionName:   core.CastString(data["sessionName"]),
	}
}

func (p CommitVoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
		"sessionName":   p.SessionName,
	}
}

func (p CommitVoteRequest) Pointer() *CommitVoteRequest {
	return &p
}
