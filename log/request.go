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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Type *string	`json:"type"`
	GcpCredentialJson *string	`json:"gcpCredentialJson"`
	BigQueryDatasetName *string	`json:"bigQueryDatasetName"`
	LogExpireDays *int32	`json:"logExpireDays"`
	AwsRegion *string	`json:"awsRegion"`
	AwsAccessKeyId *string	`json:"awsAccessKeyId"`
	AwsSecretAccessKey *string	`json:"awsSecretAccessKey"`
	FirehoseStreamName *string	`json:"firehoseStreamName"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Description *string	`json:"description"`
	Type *string	`json:"type"`
	GcpCredentialJson *string	`json:"gcpCredentialJson"`
	BigQueryDatasetName *string	`json:"bigQueryDatasetName"`
	LogExpireDays *int32	`json:"logExpireDays"`
	AwsRegion *string	`json:"awsRegion"`
	AwsAccessKeyId *string	`json:"awsAccessKeyId"`
	AwsSecretAccessKey *string	`json:"awsSecretAccessKey"`
	FirehoseStreamName *string	`json:"firehoseStreamName"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type QueryAccessLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountAccessLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type QueryIssueStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountIssueStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type QueryExecuteStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountExecuteStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type QueryExecuteStampTaskLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type CountExecuteStampTaskLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Service *string	`json:"service"`
	Method *string	`json:"method"`
	UserId *string	`json:"userId"`
	Action *string	`json:"action"`
	Begin *int64	`json:"begin"`
	End *int64	`json:"end"`
	LongTerm *bool	`json:"longTerm"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
