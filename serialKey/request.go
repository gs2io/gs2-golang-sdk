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

package serialKey

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
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	Name            *string     `json:"name"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
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
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Description:   core.CastString(data["description"]),
		LogSetting:    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"logSetting":    p.LogSetting.ToDict(),
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

type DescribeIssueJobsRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	PageToken         *string `json:"pageToken"`
	Limit             *int32  `json:"limit"`
}

func NewDescribeIssueJobsRequestFromJson(data string) DescribeIssueJobsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIssueJobsRequestFromDict(dict)
}

func NewDescribeIssueJobsRequestFromDict(data map[string]interface{}) DescribeIssueJobsRequest {
	return DescribeIssueJobsRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		PageToken:         core.CastString(data["pageToken"]),
		Limit:             core.CastInt32(data["limit"]),
	}
}

func (p DescribeIssueJobsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
		"pageToken":         p.PageToken,
		"limit":             p.Limit,
	}
}

func (p DescribeIssueJobsRequest) Pointer() *DescribeIssueJobsRequest {
	return &p
}

type GetIssueJobRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
}

func NewGetIssueJobRequestFromJson(data string) GetIssueJobRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIssueJobRequestFromDict(dict)
}

func NewGetIssueJobRequestFromDict(data map[string]interface{}) GetIssueJobRequest {
	return GetIssueJobRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		IssueJobName:      core.CastString(data["issueJobName"]),
	}
}

func (p GetIssueJobRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
		"issueJobName":      p.IssueJobName,
	}
}

func (p GetIssueJobRequest) Pointer() *GetIssueJobRequest {
	return &p
}

type IssueRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	Metadata          *string `json:"metadata"`
	IssueRequestCount *int32  `json:"issueRequestCount"`
}

func NewIssueRequestFromJson(data string) IssueRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueRequestFromDict(dict)
}

func NewIssueRequestFromDict(data map[string]interface{}) IssueRequest {
	return IssueRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		Metadata:          core.CastString(data["metadata"]),
		IssueRequestCount: core.CastInt32(data["issueRequestCount"]),
	}
}

func (p IssueRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
		"metadata":          p.Metadata,
		"issueRequestCount": p.IssueRequestCount,
	}
}

func (p IssueRequest) Pointer() *IssueRequest {
	return &p
}

type DescribeSerialKeysRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
	PageToken         *string `json:"pageToken"`
	Limit             *int32  `json:"limit"`
}

func NewDescribeSerialKeysRequestFromJson(data string) DescribeSerialKeysRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSerialKeysRequestFromDict(dict)
}

func NewDescribeSerialKeysRequestFromDict(data map[string]interface{}) DescribeSerialKeysRequest {
	return DescribeSerialKeysRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		IssueJobName:      core.CastString(data["issueJobName"]),
		PageToken:         core.CastString(data["pageToken"]),
		Limit:             core.CastInt32(data["limit"]),
	}
}

func (p DescribeSerialKeysRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
		"issueJobName":      p.IssueJobName,
		"pageToken":         p.PageToken,
		"limit":             p.Limit,
	}
}

func (p DescribeSerialKeysRequest) Pointer() *DescribeSerialKeysRequest {
	return &p
}

type DownloadSerialCodesRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
}

func NewDownloadSerialCodesRequestFromJson(data string) DownloadSerialCodesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDownloadSerialCodesRequestFromDict(dict)
}

func NewDownloadSerialCodesRequestFromDict(data map[string]interface{}) DownloadSerialCodesRequest {
	return DownloadSerialCodesRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		IssueJobName:      core.CastString(data["issueJobName"]),
	}
}

func (p DownloadSerialCodesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
		"issueJobName":      p.IssueJobName,
	}
}

func (p DownloadSerialCodesRequest) Pointer() *DownloadSerialCodesRequest {
	return &p
}

type GetSerialKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Code            *string `json:"code"`
}

func NewGetSerialKeyRequestFromJson(data string) GetSerialKeyRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSerialKeyRequestFromDict(dict)
}

func NewGetSerialKeyRequestFromDict(data map[string]interface{}) GetSerialKeyRequest {
	return GetSerialKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Code:          core.CastString(data["code"]),
	}
}

func (p GetSerialKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"code":          p.Code,
	}
}

func (p GetSerialKeyRequest) Pointer() *GetSerialKeyRequest {
	return &p
}

type UseRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Code               *string `json:"code"`
}

func NewUseRequestFromJson(data string) UseRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseRequestFromDict(dict)
}

func NewUseRequestFromDict(data map[string]interface{}) UseRequest {
	return UseRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Code:          core.CastString(data["code"]),
	}
}

func (p UseRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"code":          p.Code,
	}
}

func (p UseRequest) Pointer() *UseRequest {
	return &p
}

type UseByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Code               *string `json:"code"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewUseByUserIdRequestFromJson(data string) UseByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseByUserIdRequestFromDict(dict)
}

func NewUseByUserIdRequestFromDict(data map[string]interface{}) UseByUserIdRequest {
	return UseByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Code:            core.CastString(data["code"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"code":            p.Code,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UseByUserIdRequest) Pointer() *UseByUserIdRequest {
	return &p
}

type RevertUseByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Code               *string `json:"code"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRevertUseByUserIdRequestFromJson(data string) RevertUseByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertUseByUserIdRequestFromDict(dict)
}

func NewRevertUseByUserIdRequestFromDict(data map[string]interface{}) RevertUseByUserIdRequest {
	return RevertUseByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Code:            core.CastString(data["code"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RevertUseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"code":            p.Code,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RevertUseByUserIdRequest) Pointer() *RevertUseByUserIdRequest {
	return &p
}

type UseByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewUseByStampTaskRequestFromJson(data string) UseByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseByStampTaskRequestFromDict(dict)
}

func NewUseByStampTaskRequestFromDict(data map[string]interface{}) UseByStampTaskRequest {
	return UseByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p UseByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p UseByStampTaskRequest) Pointer() *UseByStampTaskRequest {
	return &p
}

type RevertUseByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewRevertUseByStampSheetRequestFromJson(data string) RevertUseByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertUseByStampSheetRequestFromDict(dict)
}

func NewRevertUseByStampSheetRequestFromDict(data map[string]interface{}) RevertUseByStampSheetRequest {
	return RevertUseByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RevertUseByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RevertUseByStampSheetRequest) Pointer() *RevertUseByStampSheetRequest {
	return &p
}

type DescribeCampaignModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeCampaignModelsRequestFromJson(data string) DescribeCampaignModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCampaignModelsRequestFromDict(dict)
}

func NewDescribeCampaignModelsRequestFromDict(data map[string]interface{}) DescribeCampaignModelsRequest {
	return DescribeCampaignModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeCampaignModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeCampaignModelsRequest) Pointer() *DescribeCampaignModelsRequest {
	return &p
}

type GetCampaignModelRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewGetCampaignModelRequestFromJson(data string) GetCampaignModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCampaignModelRequestFromDict(dict)
}

func NewGetCampaignModelRequestFromDict(data map[string]interface{}) GetCampaignModelRequest {
	return GetCampaignModelRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
	}
}

func (p GetCampaignModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
	}
}

func (p GetCampaignModelRequest) Pointer() *GetCampaignModelRequest {
	return &p
}

type DescribeCampaignModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCampaignModelMastersRequestFromJson(data string) DescribeCampaignModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCampaignModelMastersRequestFromDict(dict)
}

func NewDescribeCampaignModelMastersRequestFromDict(data map[string]interface{}) DescribeCampaignModelMastersRequest {
	return DescribeCampaignModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCampaignModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCampaignModelMastersRequest) Pointer() *DescribeCampaignModelMastersRequest {
	return &p
}

type CreateCampaignModelMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
}

func NewCreateCampaignModelMasterRequestFromJson(data string) CreateCampaignModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCampaignModelMasterRequestFromDict(dict)
}

func NewCreateCampaignModelMasterRequestFromDict(data map[string]interface{}) CreateCampaignModelMasterRequest {
	return CreateCampaignModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		EnableCampaignCode: core.CastBool(data["enableCampaignCode"]),
	}
}

func (p CreateCampaignModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"name":               p.Name,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"enableCampaignCode": p.EnableCampaignCode,
	}
}

func (p CreateCampaignModelMasterRequest) Pointer() *CreateCampaignModelMasterRequest {
	return &p
}

type GetCampaignModelMasterRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewGetCampaignModelMasterRequestFromJson(data string) GetCampaignModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCampaignModelMasterRequestFromDict(dict)
}

func NewGetCampaignModelMasterRequestFromDict(data map[string]interface{}) GetCampaignModelMasterRequest {
	return GetCampaignModelMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
	}
}

func (p GetCampaignModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
	}
}

func (p GetCampaignModelMasterRequest) Pointer() *GetCampaignModelMasterRequest {
	return &p
}

type UpdateCampaignModelMasterRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	CampaignModelName  *string `json:"campaignModelName"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
}

func NewUpdateCampaignModelMasterRequestFromJson(data string) UpdateCampaignModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCampaignModelMasterRequestFromDict(dict)
}

func NewUpdateCampaignModelMasterRequestFromDict(data map[string]interface{}) UpdateCampaignModelMasterRequest {
	return UpdateCampaignModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		CampaignModelName:  core.CastString(data["campaignModelName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		EnableCampaignCode: core.CastBool(data["enableCampaignCode"]),
	}
}

func (p UpdateCampaignModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"campaignModelName":  p.CampaignModelName,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"enableCampaignCode": p.EnableCampaignCode,
	}
}

func (p UpdateCampaignModelMasterRequest) Pointer() *UpdateCampaignModelMasterRequest {
	return &p
}

type DeleteCampaignModelMasterRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewDeleteCampaignModelMasterRequestFromJson(data string) DeleteCampaignModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCampaignModelMasterRequestFromDict(dict)
}

func NewDeleteCampaignModelMasterRequestFromDict(data map[string]interface{}) DeleteCampaignModelMasterRequest {
	return DeleteCampaignModelMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
	}
}

func (p DeleteCampaignModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"campaignModelName": p.CampaignModelName,
	}
}

func (p DeleteCampaignModelMasterRequest) Pointer() *DeleteCampaignModelMasterRequest {
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

type GetCurrentCampaignMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentCampaignMasterRequestFromJson(data string) GetCurrentCampaignMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentCampaignMasterRequestFromDict(dict)
}

func NewGetCurrentCampaignMasterRequestFromDict(data map[string]interface{}) GetCurrentCampaignMasterRequest {
	return GetCurrentCampaignMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentCampaignMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentCampaignMasterRequest) Pointer() *GetCurrentCampaignMasterRequest {
	return &p
}

type UpdateCurrentCampaignMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentCampaignMasterRequestFromJson(data string) UpdateCurrentCampaignMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCampaignMasterRequestFromDict(dict)
}

func NewUpdateCurrentCampaignMasterRequestFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterRequest {
	return UpdateCurrentCampaignMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentCampaignMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentCampaignMasterRequest) Pointer() *UpdateCurrentCampaignMasterRequest {
	return &p
}

type UpdateCurrentCampaignMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentCampaignMasterFromGitHubRequestFromJson(data string) UpdateCurrentCampaignMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCampaignMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentCampaignMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterFromGitHubRequest {
	return UpdateCurrentCampaignMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentCampaignMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentCampaignMasterFromGitHubRequest) Pointer() *UpdateCurrentCampaignMasterFromGitHubRequest {
	return &p
}
