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
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	Type *core.String	`json:"type"`
	GcpCredentialJson *core.String	`json:"gcpCredentialJson"`
	BigQueryDatasetName *core.String	`json:"bigQueryDatasetName"`
	LogExpireDays *int32	`json:"logExpireDays"`
	AwsRegion *core.String	`json:"awsRegion"`
	AwsAccessKeyId *core.String	`json:"awsAccessKeyId"`
	AwsSecretAccessKey *core.String	`json:"awsSecretAccessKey"`
	FirehoseStreamName *core.String	`json:"firehoseStreamName"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Description *core.String	`json:"description"`
	Type *core.String	`json:"type"`
	GcpCredentialJson *core.String	`json:"gcpCredentialJson"`
	BigQueryDatasetName *core.String	`json:"bigQueryDatasetName"`
	LogExpireDays *int32	`json:"logExpireDays"`
	AwsRegion *core.String	`json:"awsRegion"`
	AwsAccessKeyId *core.String	`json:"awsAccessKeyId"`
	AwsSecretAccessKey *core.String	`json:"awsSecretAccessKey"`
	FirehoseStreamName *core.String	`json:"firehoseStreamName"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type QueryAccessLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *core.String	`json:"service"`
	Method *core.String	`json:"method"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type CountAccessLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *bool	`json:"service"`
	Method *bool	`json:"method"`
	UserId *bool	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type QueryIssueStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *core.String	`json:"service"`
	Method *core.String	`json:"method"`
	UserId *core.String	`json:"userId"`
	Action *core.String	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type CountIssueStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *bool	`json:"service"`
	Method *bool	`json:"method"`
	UserId *bool	`json:"userId"`
	Action *bool	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type QueryExecuteStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *core.String	`json:"service"`
	Method *core.String	`json:"method"`
	UserId *core.String	`json:"userId"`
	Action *core.String	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type CountExecuteStampSheetLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *bool	`json:"service"`
	Method *bool	`json:"method"`
	UserId *bool	`json:"userId"`
	Action *bool	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type QueryExecuteStampTaskLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *core.String	`json:"service"`
	Method *core.String	`json:"method"`
	UserId *core.String	`json:"userId"`
	Action *core.String	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type CountExecuteStampTaskLogRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Service *bool	`json:"service"`
	Method *bool	`json:"method"`
	UserId *bool	`json:"userId"`
	Action *bool	`json:"action"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}
