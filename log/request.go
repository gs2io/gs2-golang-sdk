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

package log

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
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Type                *string `json:"type"`
	GcpCredentialJson   *string `json:"gcpCredentialJson"`
	BigQueryDatasetName *string `json:"bigQueryDatasetName"`
	LogExpireDays       *int32  `json:"logExpireDays"`
	AwsRegion           *string `json:"awsRegion"`
	AwsAccessKeyId      *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey  *string `json:"awsSecretAccessKey"`
	FirehoseStreamName  *string `json:"firehoseStreamName"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Type:                core.CastString(data["type"]),
		GcpCredentialJson:   core.CastString(data["gcpCredentialJson"]),
		BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
		LogExpireDays:       core.CastInt32(data["logExpireDays"]),
		AwsRegion:           core.CastString(data["awsRegion"]),
		AwsAccessKeyId:      core.CastString(data["awsAccessKeyId"]),
		AwsSecretAccessKey:  core.CastString(data["awsSecretAccessKey"]),
		FirehoseStreamName:  core.CastString(data["firehoseStreamName"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                p.Name,
		"description":         p.Description,
		"type":                p.Type,
		"gcpCredentialJson":   p.GcpCredentialJson,
		"bigQueryDatasetName": p.BigQueryDatasetName,
		"logExpireDays":       p.LogExpireDays,
		"awsRegion":           p.AwsRegion,
		"awsAccessKeyId":      p.AwsAccessKeyId,
		"awsSecretAccessKey":  p.AwsSecretAccessKey,
		"firehoseStreamName":  p.FirehoseStreamName,
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
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	Description         *string `json:"description"`
	Type                *string `json:"type"`
	GcpCredentialJson   *string `json:"gcpCredentialJson"`
	BigQueryDatasetName *string `json:"bigQueryDatasetName"`
	LogExpireDays       *int32  `json:"logExpireDays"`
	AwsRegion           *string `json:"awsRegion"`
	AwsAccessKeyId      *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey  *string `json:"awsSecretAccessKey"`
	FirehoseStreamName  *string `json:"firehoseStreamName"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		Description:         core.CastString(data["description"]),
		Type:                core.CastString(data["type"]),
		GcpCredentialJson:   core.CastString(data["gcpCredentialJson"]),
		BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
		LogExpireDays:       core.CastInt32(data["logExpireDays"]),
		AwsRegion:           core.CastString(data["awsRegion"]),
		AwsAccessKeyId:      core.CastString(data["awsAccessKeyId"]),
		AwsSecretAccessKey:  core.CastString(data["awsSecretAccessKey"]),
		FirehoseStreamName:  core.CastString(data["firehoseStreamName"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"description":         p.Description,
		"type":                p.Type,
		"gcpCredentialJson":   p.GcpCredentialJson,
		"bigQueryDatasetName": p.BigQueryDatasetName,
		"logExpireDays":       p.LogExpireDays,
		"awsRegion":           p.AwsRegion,
		"awsAccessKeyId":      p.AwsAccessKeyId,
		"awsSecretAccessKey":  p.AwsSecretAccessKey,
		"firehoseStreamName":  p.FirehoseStreamName,
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

type QueryAccessLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewQueryAccessLogRequestFromJson(data string) QueryAccessLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryAccessLogRequestFromDict(dict)
}

func NewQueryAccessLogRequestFromDict(data map[string]interface{}) QueryAccessLogRequest {
	return QueryAccessLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryAccessLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryAccessLogRequest) Pointer() *QueryAccessLogRequest {
	return &p
}

type CountAccessLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewCountAccessLogRequestFromJson(data string) CountAccessLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountAccessLogRequestFromDict(dict)
}

func NewCountAccessLogRequestFromDict(data map[string]interface{}) CountAccessLogRequest {
	return CountAccessLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountAccessLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountAccessLogRequest) Pointer() *CountAccessLogRequest {
	return &p
}

type QueryIssueStampSheetLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewQueryIssueStampSheetLogRequestFromJson(data string) QueryIssueStampSheetLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryIssueStampSheetLogRequestFromDict(dict)
}

func NewQueryIssueStampSheetLogRequestFromDict(data map[string]interface{}) QueryIssueStampSheetLogRequest {
	return QueryIssueStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryIssueStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryIssueStampSheetLogRequest) Pointer() *QueryIssueStampSheetLogRequest {
	return &p
}

type CountIssueStampSheetLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewCountIssueStampSheetLogRequestFromJson(data string) CountIssueStampSheetLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountIssueStampSheetLogRequestFromDict(dict)
}

func NewCountIssueStampSheetLogRequestFromDict(data map[string]interface{}) CountIssueStampSheetLogRequest {
	return CountIssueStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountIssueStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountIssueStampSheetLogRequest) Pointer() *CountIssueStampSheetLogRequest {
	return &p
}

type QueryExecuteStampSheetLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewQueryExecuteStampSheetLogRequestFromJson(data string) QueryExecuteStampSheetLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryExecuteStampSheetLogRequestFromDict(dict)
}

func NewQueryExecuteStampSheetLogRequestFromDict(data map[string]interface{}) QueryExecuteStampSheetLogRequest {
	return QueryExecuteStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryExecuteStampSheetLogRequest) Pointer() *QueryExecuteStampSheetLogRequest {
	return &p
}

type CountExecuteStampSheetLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewCountExecuteStampSheetLogRequestFromJson(data string) CountExecuteStampSheetLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountExecuteStampSheetLogRequestFromDict(dict)
}

func NewCountExecuteStampSheetLogRequestFromDict(data map[string]interface{}) CountExecuteStampSheetLogRequest {
	return CountExecuteStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountExecuteStampSheetLogRequest) Pointer() *CountExecuteStampSheetLogRequest {
	return &p
}

type QueryExecuteStampTaskLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewQueryExecuteStampTaskLogRequestFromJson(data string) QueryExecuteStampTaskLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryExecuteStampTaskLogRequestFromDict(dict)
}

func NewQueryExecuteStampTaskLogRequestFromDict(data map[string]interface{}) QueryExecuteStampTaskLogRequest {
	return QueryExecuteStampTaskLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryExecuteStampTaskLogRequest) Pointer() *QueryExecuteStampTaskLogRequest {
	return &p
}

type CountExecuteStampTaskLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewCountExecuteStampTaskLogRequestFromJson(data string) CountExecuteStampTaskLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountExecuteStampTaskLogRequestFromDict(dict)
}

func NewCountExecuteStampTaskLogRequestFromDict(data map[string]interface{}) CountExecuteStampTaskLogRequest {
	return CountExecuteStampTaskLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountExecuteStampTaskLogRequest) Pointer() *CountExecuteStampTaskLogRequest {
	return &p
}

type QueryAccessLogWithTelemetryRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewQueryAccessLogWithTelemetryRequestFromJson(data string) QueryAccessLogWithTelemetryRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryAccessLogWithTelemetryRequestFromDict(dict)
}

func NewQueryAccessLogWithTelemetryRequestFromDict(data map[string]interface{}) QueryAccessLogWithTelemetryRequest {
	return QueryAccessLogWithTelemetryRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryAccessLogWithTelemetryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryAccessLogWithTelemetryRequest) Pointer() *QueryAccessLogWithTelemetryRequest {
	return &p
}

type PutLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
	LogCategory        *string `json:"logCategory"`
	Payload            *string `json:"payload"`
}

func NewPutLogRequestFromJson(data string) PutLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutLogRequestFromDict(dict)
}

func NewPutLogRequestFromDict(data map[string]interface{}) PutLogRequest {
	return PutLogRequest{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
		LogCategory:        core.CastString(data["logCategory"]),
		Payload:            core.CastString(data["payload"]),
	}
}

func (p PutLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
		"logCategory":        p.LogCategory,
		"payload":            p.Payload,
	}
}

func (p PutLogRequest) Pointer() *PutLogRequest {
	return &p
}

type DescribeInsightsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeInsightsRequestFromJson(data string) DescribeInsightsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInsightsRequestFromDict(dict)
}

func NewDescribeInsightsRequestFromDict(data map[string]interface{}) DescribeInsightsRequest {
	return DescribeInsightsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInsightsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInsightsRequest) Pointer() *DescribeInsightsRequest {
	return &p
}

type CreateInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewCreateInsightRequestFromJson(data string) CreateInsightRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateInsightRequestFromDict(dict)
}

func NewCreateInsightRequestFromDict(data map[string]interface{}) CreateInsightRequest {
	return CreateInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p CreateInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p CreateInsightRequest) Pointer() *CreateInsightRequest {
	return &p
}

type GetInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InsightName     *string `json:"insightName"`
}

func NewGetInsightRequestFromJson(data string) GetInsightRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInsightRequestFromDict(dict)
}

func NewGetInsightRequestFromDict(data map[string]interface{}) GetInsightRequest {
	return GetInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InsightName:   core.CastString(data["insightName"]),
	}
}

func (p GetInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"insightName":   p.InsightName,
	}
}

func (p GetInsightRequest) Pointer() *GetInsightRequest {
	return &p
}

type DeleteInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InsightName     *string `json:"insightName"`
}

func NewDeleteInsightRequestFromJson(data string) DeleteInsightRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInsightRequestFromDict(dict)
}

func NewDeleteInsightRequestFromDict(data map[string]interface{}) DeleteInsightRequest {
	return DeleteInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InsightName:   core.CastString(data["insightName"]),
	}
}

func (p DeleteInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"insightName":   p.InsightName,
	}
}

func (p DeleteInsightRequest) Pointer() *DeleteInsightRequest {
	return &p
}
