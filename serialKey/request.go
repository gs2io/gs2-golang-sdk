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
	RequestId    *string     `json:"requestId"`
	ContextStack *string     `json:"contextStack"`
	Name         *string     `json:"name"`
	Description  *string     `json:"description"`
	LogSetting   *LogSetting `json:"logSetting"`
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Description   *string     `json:"description"`
	LogSetting    *LogSetting `json:"logSetting"`
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

type DescribeIssueJobsRequest struct {
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	PageToken         *string `json:"pageToken"`
	Limit             *int32  `json:"limit"`
}

func NewDescribeIssueJobsRequestFromJson(data string) DescribeIssueJobsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeIssueJobsRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
}

func NewGetIssueJobRequestFromJson(data string) GetIssueJobRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetIssueJobRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	Metadata          *string `json:"metadata"`
	IssueRequestCount *int32  `json:"issueRequestCount"`
}

func NewIssueRequestFromJson(data string) IssueRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewIssueRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
	PageToken         *string `json:"pageToken"`
	Limit             *int32  `json:"limit"`
}

func NewDescribeSerialKeysRequestFromJson(data string) DescribeSerialKeysRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeSerialKeysRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
	IssueJobName      *string `json:"issueJobName"`
}

func NewDownloadSerialCodesRequestFromJson(data string) DownloadSerialCodesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDownloadSerialCodesRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Code          *string `json:"code"`
}

func NewGetSerialKeyRequestFromJson(data string) GetSerialKeyRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetSerialKeyRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Code               *string `json:"code"`
}

func NewUseRequestFromJson(data string) UseRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUseRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Code               *string `json:"code"`
}

func NewUseByUserIdRequestFromJson(data string) UseByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUseByUserIdRequestFromDict(dict2)
}

func NewUseByUserIdRequestFromDict(data map[string]interface{}) UseByUserIdRequest {
	return UseByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Code:          core.CastString(data["code"]),
	}
}

func (p UseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"code":          p.Code,
	}
}

func (p UseByUserIdRequest) Pointer() *UseByUserIdRequest {
	return &p
}

type RevertUseByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Code               *string `json:"code"`
}

func NewRevertUseByUserIdRequestFromJson(data string) RevertUseByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewRevertUseByUserIdRequestFromDict(dict2)
}

func NewRevertUseByUserIdRequestFromDict(data map[string]interface{}) RevertUseByUserIdRequest {
	return RevertUseByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Code:          core.CastString(data["code"]),
	}
}

func (p RevertUseByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"code":          p.Code,
	}
}

func (p RevertUseByUserIdRequest) Pointer() *RevertUseByUserIdRequest {
	return &p
}

type UseByStampTaskRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
}

func NewUseByStampTaskRequestFromJson(data string) UseByStampTaskRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUseByStampTaskRequestFromDict(dict2)
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
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
}

func NewRevertUseByStampSheetRequestFromJson(data string) RevertUseByStampSheetRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewRevertUseByStampSheetRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeCampaignModelsRequestFromJson(data string) DescribeCampaignModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeCampaignModelsRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewGetCampaignModelRequestFromJson(data string) GetCampaignModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCampaignModelRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeCampaignModelMastersRequestFromJson(data string) DescribeCampaignModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeCampaignModelMastersRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
}

func NewCreateCampaignModelMasterRequestFromJson(data string) CreateCampaignModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateCampaignModelMasterRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewGetCampaignModelMasterRequestFromJson(data string) GetCampaignModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCampaignModelMasterRequestFromDict(dict2)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	NamespaceName      *string `json:"namespaceName"`
	CampaignModelName  *string `json:"campaignModelName"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
}

func NewUpdateCampaignModelMasterRequestFromJson(data string) UpdateCampaignModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCampaignModelMasterRequestFromDict(dict2)
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
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	CampaignModelName *string `json:"campaignModelName"`
}

func NewDeleteCampaignModelMasterRequestFromJson(data string) DeleteCampaignModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteCampaignModelMasterRequestFromDict(dict2)
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

type GetCurrentCampaignMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentCampaignMasterRequestFromJson(data string) GetCurrentCampaignMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentCampaignMasterRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentCampaignMasterRequestFromJson(data string) UpdateCurrentCampaignMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentCampaignMasterRequestFromDict(dict2)
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
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentCampaignMasterFromGitHubRequestFromJson(data string) UpdateCurrentCampaignMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentCampaignMasterFromGitHubRequestFromDict(dict2)
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
